package config

import (
	"github.com/pkg/errors"

	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/config"

	"os"
)

type Config struct {
	ServicePort string `yaml:"service_port"`

	RepositoryCfg *pgxpool.Config `yaml:"repository_cfg"`
	DBConnConfig  DBConnConfig    `yaml:"db_conn_config"`
}

type DBConnConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DBName       string `yaml:"db_name"`
	PoolMaxConns string `yaml:"pool_max_conns"`
}

func InitConfig() (*Config, error) {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		return nil, errors.New(ErrConfigPathIsEmpty)
	}

	base, err := os.Open(cfgPath + "config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config")
	}

	mergeCfg, err := config.NewYAML(config.Source(base))
	if err != nil {
		return nil, errors.Wrap(err, "failed merge config")
	}

	var cfg Config
	if err := mergeCfg.Get("config").Populate(&cfg); err != nil {
		return nil, errors.Wrap(err, "marshal config")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?pool_max_conns=%s",
		cfg.DBConnConfig.User,
		cfg.DBConnConfig.Password,
		cfg.DBConnConfig.Host,
		cfg.DBConnConfig.DBName,
		cfg.DBConnConfig.PoolMaxConns)

	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse data base config")
	}

	cfg.RepositoryCfg = repCfg

	return &cfg, err

}
