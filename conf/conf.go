package conf

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	Conf     = &Config{}
	confPath string
)

// config
type Config struct {
	HttpServer *Httpserver
	MySql      *Mysql
	Mongo      *Mongo
}

type Httpserver struct {
	Addr string
}

type Mysql struct {
	Dsn string
}

type Mongo struct {
	Addrs    []string
	Username string
	Password string
	MaxPool  uint64
}

type Duration time.Duration

func Init() {
	confPath = "conf/local.toml"
	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
		log.Fatal(err)
	}
}
