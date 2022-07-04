package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var once sync.Once
var configInfo = new(Specification)

type Specification struct {
	Port int
	DB   struct {
		Typ   string `yaml:"type"`
		DSN   string `yaml:"dsn"`
		Debug bool
	}
	Cache struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
	}
	JWT struct {
		Timeout int64  `yaml:"timeout"`
		SignKey string `yaml:"signKey"`
	}
	SMS struct {
		APIKey   string `yaml:"apiKey"`
		Template string `yaml:"template"`
	}
	Zap struct {
		Level        string `yaml:"level"`          // 级别
		Format       string `yaml:"format"`         // 输出
		Prefix       string `yaml:"prefix"`         // 日志前缀
		Director     string `yaml:"director"`       // 日志文件夹
		LogInConsole bool   `yaml:"log-in-console"` // 输出控制台
	}
	Storage struct {
		AuthFile   string `yaml:"authFile"`   // 认证文件地址
		Bucket     string `yaml:"bucket"`     // bucket
		FileFormat string `yaml:"fileFormat"` // 文件名称模板
		GcsFormat  string `yaml:"gcsFormat"`  // zip模板
	} `yaml:"storage"`
	DdmcUpload struct {
		Appkey  string `yaml:"appkey"`
		Secret  string `yaml:"secret"`
		Host    string `yaml:"host"`
		FunName string `yaml:"funName"`
	} `yaml:"ddmcUpload"`
	FzUpload struct {
		AuthFile   string `yaml:"authFile"`   // 认证文件地址
		Bucket     string `yaml:"bucket"`     // bucket
		FileFormat string `yaml:"fileFormat"` // 文件名称模板
		GcsFormat  string `yaml:"gcsFormat"`  // zip模板
	} `yaml:"fzUpload"`
	GD struct {
		Key string `yaml:"key"`
		URL string `yaml:"url"`
	} `yaml:"gd"`
	EPPTask struct {
		Appkey         string `yaml:"appkey"`
		Secret         string `yaml:"secret"`
		ProductTaskURL string `yaml:"productTaskURL"`
		ConsumeTaskURL string `yaml:"consumeTaskURL"`
		UploadTaskURL  string `yaml:"uploadTaskURL"`
	} `yaml:"eppTask"`
	UploadPath      string `yaml:"uploadPath"`
	AliAppCode      string `yaml:"aliAppCode"`
	DataControlFlag bool   `yaml:"dataControlFlag"` // 是否识别风控开关;
	DataParseWork   int    `yaml:"dataParseWork"`   // 上传数据解析工作线程
	UserMaxRisk     int    `yaml:"userMaxRisk"`
}

func Get(fileName string) (*Specification, error) {
	once.Do(func() {
		if fileName == "" {
			fileName = "./conf/conf.yaml"
		}
		err := Refresh(fileName)
		if err != nil {
			log.Fatalln("read conf file: ", err)
		}
	})
	return configInfo, nil
}

func Refresh(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, configInfo); err != nil {
		return err
	}
	return nil
}
