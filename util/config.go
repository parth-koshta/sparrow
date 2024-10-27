package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	ServerUrl            string        `mapstructure:"SERVER_URL"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	LinkedInClientID     string        `mapstructure:"LINKEDIN_CLIENT_ID"`
	LinkedInClientSecret string        `mapstructure:"LINKEDIN_CLIENT_SECRET"`
	OpenAIApiKey         string        `mapstructure:"OPENAI_API_KEY"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetConfigName("app-local")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
