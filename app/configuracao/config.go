package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"
)

const projectDirName = "minotauro_backend"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func EnvCloudName() string {
	loadEnv()
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	loadEnv()

	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	loadEnv()

	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	loadEnv()

	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}

func EnvDev() bool {
	loadEnv()

	dev, _ := strconv.ParseBool(os.Getenv("DEV"))
	return dev

}

func EnvLocalHostDb() string {
	loadEnv()

	return os.Getenv("SQL_HOST")

}

func EnvUsrDb() string {
	loadEnv()

	return os.Getenv("SQL_USER")

}

func EnvSqlPassword() string {
	loadEnv()

	return os.Getenv("SQL_PASSWORD")

}

func EnvLocalDb() string {
	loadEnv()

	return os.Getenv("SQL_DATABASE")

}
