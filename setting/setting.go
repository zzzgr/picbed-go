package setting

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
)

var App = new(Application)

func SetUp() (err error) {

	err = ifConfig("config/config.yaml")
	if err != nil {
		fmt.Println("setting SetUp err")
		return err
	}

	// 加载配置
	loadConfig()

	return
}

func ifConfig(src string) (err error) {
	file, err := os.Stat(src)
	if err != nil {
		// config文件夹
		err = os.Mkdir("config", os.ModePerm)
		if err != nil {
			return errors.New(fmt.Sprintf("create config dir rrror: %s", err))
		}

		// config.yaml文件
		_, err = os.Create(src)
		if err != nil {
			return errors.New(fmt.Sprintf("create config file error: %s", err))
		}

		// 默认配置
		app := &Application{
			Server: Server{
				Name:   "picbed-go",
				Port:   8080,
				Notice: "欢迎光临picbed-go",
				Mode:   "debug",
			},
			Config: Config{
				BliCookie: "填写bilibili cookie, 格式: SESSDATA=5034xxxxxx2A21;bili_jct=6axxxxxx89",
			},
		}
		appYaml, err := yaml.Marshal(app)
		if err != nil {
			return err
		}

		err = os.WriteFile(src, appYaml, 0777)
		if err != nil {
			return errors.New(fmt.Sprintf("init default config error: %s", err))
		}

		return nil
	}

	if file.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a file, please check!", src))
	} else {
		return nil
	}
}

func loadConfig() {
	// 配置文件名称
	viper.SetConfigName("config")
	// 配置文件扩展名
	viper.SetConfigType("yaml")
	// 配置文件所在路径
	viper.AddConfigPath("./config")
	// 查找并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	// 配置信息绑定到结构体变量
	err = viper.Unmarshal(&App)
	if err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v \n", err)
	}

	// 热加载配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(c fsnotify.Event) {
		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		// 配置信息绑定到结构体变量
		err = viper.Unmarshal(&App)
		if err != nil {
			return
		}
	})
}
