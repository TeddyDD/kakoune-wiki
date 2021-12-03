package kakoune

import "github.com/caarlos0/env/v6"

type Config struct {
	Session  string `env:"kak_session"`
	Client   string `env:"kak_client"`
	WikiPath string `env:"kak_opt_wiki_path"`
	Buffile  string `env:"kak_buffile"`

	CommandFifo  string `env:"kak_command_fifo"`
	ResponseFifo string `env:"kak_response_fifo"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
