package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port        int    `env:"PORT" envDefault:"80"`
	AWSEndpoint string `env:"AWS_ENDPOINT" envDefault:""`
}

var config Config

// Load は、環境変数から設定を読み込みます。
// 1回のみ呼び出してください。
func Load() error {
	if err := env.Parse(&config); err != nil {
		return err
	}
	return nil
}

func Get() Config { return config }
