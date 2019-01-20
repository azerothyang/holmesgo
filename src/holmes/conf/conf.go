package conf

import (
	"errors"
	"holmes/yaml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//config struct
type ConfS struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Mysql struct {
			Master struct {
				Host     string
				Port     string
				User     string
				Password string
				Database string
				Charset  string
				Timeout  string
				MaxIdle  int
				MaxConn  int
			}
			Slave struct {
				Host     string
				Port     string
				User     string
				Password string
				Database string
				Charset  string
				Timeout  string
				MaxIdle  int
				MaxConn  int
			}
		}
		Redis struct {
			Host    string
			Port    string
			Auth    string
			Timeout string
			Db      int
		}
	}
}

var (
	Conf ConfS
	Env  = "debug"
)

//init config
func InitConf() {
	dir, err := GetCurrentPath()
	if err != nil {
		log.Fatal(err)
	}

	var envb []byte
	envb, err = ioutil.ReadFile(dir + FmtSlash("conf/.env"))
	if err != nil {
		log.Fatal(err)
	}

	var config []byte
	//check release/.env, if .env 's content == debug
	if string(envb) != "debug" {
		Env = "production"
		config, err = ioutil.ReadFile(dir + FmtSlash("conf/prod.yml"))
	} else {
		Env = "debug"
		config, err = ioutil.ReadFile(dir + FmtSlash("conf/dev.yml"))
	}
	err = yaml.Unmarshal(config, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}

//get current exec file path
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

// according to operating system to change path slash, default use linux path slash
func FmtSlash(path string) string {
	sys := runtime.GOOS
	if sys == `windows` {
		return strings.Replace(path, "/", "\\", -1)
	}
	return path
}
