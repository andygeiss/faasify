package config

import "github.com/andygeiss/faasify/internal/account"

type Config struct {
	AccountAccess account.Access `json:"account_access"`
	Domain        string         `json:"domain"`
	Mode          string         `json:"mode"`
	Token         string         `json:"token"`
	Url           string         `json:"url"`
}
