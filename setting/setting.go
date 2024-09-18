package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify" //监控文件的库

	"github.com/spf13/viper"
)

// 全局变量，用来保存程序所有信息，不再用viper
var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() error {
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.SetConfigFile("config.yaml")
	viper.SetConfigFile("D:/bluebell/setting/config.yaml") // 查找配置文件所在的路径
	err := viper.ReadInConfig()                            // 查找并读取配置文件
	if err != nil {                                        // 处理读取配置文件的错误
		fmt.Println("viper.ReadInConfig()")
		return err
	}
	//把读取到的配置信息，反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%s\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了")
		//变化后继续反序列化
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%s\n", err)
		}
	})
	return err
}
