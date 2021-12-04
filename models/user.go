package models

type User struct {
	Email       string `mapstructure:"username"`
	SystemAdmin bool   `mapstructure:"systemadmin"`
}
