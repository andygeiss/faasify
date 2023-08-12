package config

import "github.com/andygeiss/faasify/internal/account"

type Config struct {
	AccountAccess account.Access `json:"account_access"`
	AppName       string         `json:"app_name"`
	Domain        string         `json:"domain"`
	Mode          string         `json:"mode"`
	Token         string         `json:"token"`
	Url           string         `json:"url"`
}
