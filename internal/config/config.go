package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/valyala/fastjson"
)

func Read(fpath string) (c Config, err error) {
	//_, err = toml.DecodeFile(fpath, &c)
	config, _ := ioutil.ReadFile(fpath) // filename is the JSON file to read
	fastjson.ParseBytes(config)
	return
}

type Config struct {
	Logger LoggerConf
	PSQL   PSQLConfig
	Server ServerConf
	Certs  CertsConf
}

type LoggerConf struct {
	Level string
	// TODO
}

type PSQLConfig struct {
	DSN  string
	Port string
	User string
	Pass string
	DB   string
}

type ServerConf struct {
	HostName string
	HttpPort string
	GrpcPort string
	Swagger  bool
}

type CertsConf struct {
	SrvCert string
	SrvKey  string
}

func NewConfig(fpath string) (c Config, err error) {
	config, err := os.ReadFile(fpath) // filename is the JSON file to read
	if err != nil {
		return
	}

	v, err := fastjson.ParseBytes(config)
	if err != nil {
		return
	}

	if !v.Exists("Logger") {
		err = fmt.Errorf("not init Logger in %s", fpath)
		return
	}
	vv := v.Get("Logger")
	if !vv.Exists("Level") {
		err = fmt.Errorf("not init Level in %s", fpath)
		return
	}
	c.Logger.Level = string(vv.Get("Level").GetStringBytes())

	if !v.Exists("PSQL") {
		err = fmt.Errorf("not init PSQL in %s", fpath)
		return
	}
	vv = v.Get("PSQL")
	if !vv.Exists("DNS") || !vv.Exists("Port") || !vv.Exists("User") || !vv.Exists("Pass") {
		err = fmt.Errorf("not init parameters of PSQL in %s", fpath)
		return
	}
	c.PSQL.DSN = string(vv.Get("DNS").GetStringBytes())
	c.PSQL.Port = string(vv.Get("Port").GetStringBytes())
	c.PSQL.User = string(vv.Get("User").GetStringBytes())
	c.PSQL.Pass = string(vv.Get("Pass").GetStringBytes())
	c.PSQL.DB = string(vv.Get("DB").GetStringBytes())

	if !v.Exists("Server") {
		err = fmt.Errorf("not init Server in %s", fpath)
		return
	}
	vv = v.Get("Server")
	if !vv.Exists("HostName") || !vv.Exists("HttpPort") || !vv.Exists("GrpcPort") || !vv.Exists("Swagger") {
		err = fmt.Errorf("not init parameters of Server in %s", fpath)
		return
	}
	c.Server.HostName = string(vv.Get("HostName").GetStringBytes())
	c.Server.Swagger = vv.Get("Swagger").GetBool()
	c.Server.HttpPort = string(vv.Get("HttpPort").GetStringBytes())
	c.Server.GrpcPort = string(vv.Get("GrpcPort").GetStringBytes())

	if !v.Exists("Certs") {
		err = fmt.Errorf("not init certificates in %s", fpath)
		return
	}
	vv = v.Get("Certs")
	if !vv.Exists("SrvCert") || !vv.Exists("SrvKey") {
		err = fmt.Errorf("not init certificates for Server in %s", fpath)
		return
	}
	c.Certs.SrvCert = string(vv.Get("SrvCert").GetStringBytes())
	c.Certs.SrvKey = string(vv.Get("SrvKey").GetStringBytes())

	return
}
