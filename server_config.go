package main

type ServerConfig struct {
	ModelBase
	PrivateKey string `gorm:"not null;type:varchar(100)"`
	PublicKey  string `gorm:"not null;type:varchar(100)"`
	EndPoint   string `gorm:"not null;type:varchar(100)"`
	Port       int    `gorm:"not null;type:int;default:27"`
	IpCidr     string `gorm:"not null;type:varchar(100)"`
}
