package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}

func main() {
	loadConfig()
	port := viper.GetInt("server.port")
	env := viper.GetString("app.env")
	fmt.Printf("Server running on port %d in %s environment\n", port, env)
}
