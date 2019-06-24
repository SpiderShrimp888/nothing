package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
)

func main() {
	privateKey := ReadFile("./private")
	publicKey := ReadFile("./public")
	endpoint := ReadFile("./endpoint")
	db, err := gorm.Open("sqlite3", "./jaynest.db")
	if err != nil {
		panic(err)
	}

	db.Create(&ServerConfig{
		ModelBase:  ModelBase{"1"},
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		EndPoint:   endpoint,
		Port:       27,
		IpCidr:     "192.168.243.",
	})

	fmt.Println("初始服务器配置已完成")
}

func ReadFile(path string) string {
	bytes, _ := ioutil.ReadFile(path)

	return string(bytes)
}
