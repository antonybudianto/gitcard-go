package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// try load from .env file
	err := godotenv.Load()
	if err == nil {
		log.Println("Read environment variables from .env file")
	}
}

// GithubAccessToken get GH_ACCESS_TOKEN from os env
// if not found it will panic
func GithubAccessToken() string {
	return readEnvOrPanic("GH_ACCESS_TOKEN")
}

func WebAddress() string {
	return readEnv("WEB_ADDRESS", ":8080")
}

func GrpcServerAddress() string {
	return readEnv("GRPC_SERVER_ADDRESS", ":50051")
}

func GrpcClientAddress() string {
	return readEnv("GRPC_CLIENT_ADDRESS", "localhost:50051")
}

func readEnvOrPanic(envName string) string {
	val := readEnv(envName, "")
	if val == "" {
		log.Fatal(fmt.Sprintf("%s not set", envName))
	}
	return val
}

func readEnv(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}
