package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	HttpPort string `yaml:"httpPort"`
	Mysql    struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"ip"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
		Table    struct {
			Users   string `yaml:"users"`
			Favlist string `yaml:"favlist"`
			Data    string `yaml:"data"`
		}
	}
	OSS struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"accessKeyID"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		Bucket          string `yaml:"bucket"`
		MainDir         string `yaml:"maindir"`
	}
	Admin struct {
		Username string `yaml:"username"`
		Secret   string `yaml:"secret"` //jwt算法秘钥
		Name     string `yaml:"name"`   //jwt签名颁发者
	}
}

var Config *conf

func Init() {
	tmp := new(conf)
	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		fmt.Println("getConf err", err)
	}
	err = yaml.Unmarshal(yamlFile, tmp)
	if err != nil {
		fmt.Println("getConf err", err)
	}
	Config = tmp
	fmt.Println("获取配置文件成功")
}
