package config

import (
	"github.com/Hajymuhammet03/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"0.0.0.0"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage               StorageConfig `yaml:"storage"`
	JwtKey                string        `yaml:"jwt_key" env-required:"true"`
	JwtKeySupAdmin        string        `yaml:"jwt_key_1" env-required:"true"`
	PublicFilePath        string        `yaml:"public_path"`
	MaxFileSize           int64         `yaml:"max_file_size" env-required:"true"`
	MessageMaxFileSize    int64         `yaml:"message_max_file_size" env-required:"true"`
	MessageMimetypes      []string      `yaml:"message_mimetypes" env-required:"true"`
	VideoCollectionServer string        `yaml:"video_collection_server" env-required:"true"`

	GetPublicFilePath string `yaml:"get_public_path"`

	PublicFilePathPost  string `yaml:"public_file_path_post" env-required:"true"`
	PublicFilePathVideo string `yaml:"public_file_path_video" env-required:"true"`
}

type StorageConfig struct {
	PgPoolMaxConn int    `yaml:"pg_pool_max_conn" env-required:"true"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Database      string `json:"database"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type RedisConfig struct {
	Addr       int    `yaml:"addr" env-required:"true"`
	Password   string `json:"password"`
	MaxRetries string `json:"max_retries"`
	Db         string `json:"db"`
	PoolSize   string `json:"pool_size"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		// TODO path config
		pathConfig := "./../../config.yml"

		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
