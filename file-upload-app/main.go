package main

import (
	"encoding/base64"
	"log"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"file-upload-app/config"
	"file-upload-app/handlers"
)

func main() {
	// Initialize Viper configuration
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	uploadPath := viper.GetString("upload_path")
	if uploadPath == "" {
		log.Fatal("Upload path not specified")
	}

	// Ensure the upload path exists
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		log.Fatalf("Error creating upload directory, %s", err)
	}

	// Setup router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/uploads", uploadPath)

	// Add custom template functions
	router.SetFuncMap(template.FuncMap{
		"base64": func(data []byte) string {
			return base64.StdEncoding.EncodeToString(data)
		},
	})

	// Setup routes
	handlers.SetupRoutes(router, uploadPath)

	// Bind to all network interfaces
	router.Run("0.0.0.0:8080")
}
