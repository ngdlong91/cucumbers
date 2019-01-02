package config

type GinConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewGinConfig() *GinConfig {
	return &GinConfig{
		Port: "3100",
	}
}
