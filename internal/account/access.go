package account

import (
	"bytes"
	"encoding/json"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Access interface {
	Error() error
	ChangeSecret(id, current, next string)
	CreateAccount(id, secret string)
	GetAccount(id string) *Account
	GetAccounts() []Account
	VerifyAccount(id, secret string) bool
}

type fileAccess struct {
	err      error
	path     string
	accounts []Account
	mutex    sync.Mutex
}

func (a *fileAccess) Error() error {
	return a.err
}

func (a *fileAccess) ChangeSecret(id, current, next string) {
	if a.err != nil {
		return
	}
	a.loadAccounts()
	for _, account := range a.accounts {
		if account.Id == id {
			if err := bcrypt.CompareHashAndPassword(account.Hash, []byte(current)); err == nil {
				hash, _ := bcrypt.GenerateFromPassword([]byte(next), 14)
				account.Hash = hash
			}
		}
	}
	a.saveAccounts()
}

func (a *fileAccess) CreateAccount(id, secret string) {
	if a.err != nil {
		return
	}
	a.loadAccounts()
	exists := false
	for _, account := range a.accounts {
		if account.Id == id {
			exists = true
			break
		}
	}
	if !exists {
		hash, _ := bcrypt.GenerateFromPassword([]byte(secret), 14)
		a.accounts = append(a.accounts, Account{Id: id, Hash: hash})
	}
	a.saveAccounts()
}

func (a *fileAccess) GetAccount(id string) *Account {
	a.loadAccounts()
	for _, account := range a.accounts {
		if account.Id == id {
			return &account
		}
	}
	return nil
}

func (a *fileAccess) GetAccounts() []Account {
	a.loadAccounts()
	return a.accounts
}

func (a *fileAccess) VerifyAccount(id, secret string) bool {
	a.loadAccounts()
	for _, account := range a.accounts {
		if account.Id == id {
			if err := bcrypt.CompareHashAndPassword(account.Hash, []byte(secret)); err == nil {
				return true
			}
		}
	}
	return false
}

func NewFileAccess(path string) Access {
	return &fileAccess{path: path}
}

func (a *fileAccess) loadAccounts() {
	if a.err != nil {
		return
	}
	a.mutex.Lock()
	defer a.mutex.Unlock()
	var accounts []Account
	file, err := os.ReadFile(a.path)
	// Don't overwrite the cache
	if err != nil {
		return
	}
	if err := json.NewDecoder(bytes.NewReader(file)).Decode(&accounts); err != nil {
		a.err = err
		return
	}
	a.accounts = accounts
}

func (a *fileAccess) saveAccounts() {
	if a.err != nil {
		return
	}
	a.mutex.Lock()
	defer a.mutex.Unlock()
	file, err := os.OpenFile(a.path, os.O_CREATE|os.O_SYNC|os.O_WRONLY, 0644)
	if err != nil {
		a.err = err
		return
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(a.accounts); err != nil {
		a.err = err
	}
}
