// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Server string `config:"server"`
	Port int `config:"port"`
	User string `config:"user"`
	Password string `config:"password"`
	Database string `config:"database"`
	Query string `config:"query"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Server: "localhost",
	Port: 1433,
	User: "root",
	Password: "root",
	Database: "master",
	Query: "",
}
