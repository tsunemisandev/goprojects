package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"file-upload-app/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, uploadPath string) {
	// Render home page with uploaded files
	router.GET("/", func(c *gin.Context) {
		renderHomePage(c, uploadPath)
	})

	// Handle file upload
	router.POST("/upload", func(c *gin.Context) {
		handleFileUpload(c, uploadPath)
	})

	// Display QR code with local IP address
	router.GET("/qrcode", showQRCode)

	// Render upload form
	router.GET("/upload", renderForm)

	// Show single image
	router.GET("/image/:filename", func(c *gin.Context) {
		showImage(c, uploadPath)
	})
}

func renderHomePage(c *gin.Context, uploadPath string) {
	files, err := os.ReadDir(uploadPath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Unable to read upload directory: %v", err))
		return
	}

	var fileInfos []gin.H
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(uploadPath, file.Name())
			info, err := os.Stat(filePath)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Unable to read file info: %v", err))
				return
			}
			fileInfos = append(fileInfos, gin.H{
				"Name": file.Name(),
				"Size": float64(info.Size()) / (1024 * 1024), // size in MB
			})
		}
	}

	qrCodeData, err := utils.GenerateQRCodeWithLocalIP()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Files":     fileInfos,
		"QRCode":    qrCodeData,
		"UploadURL": "/upload",
		"Success":   c.Query("success"),
	})
}

func renderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func handleFileUpload(c *gin.Context, uploadPath string) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	files := form.File["files"]
	var uploadedFiles []string
	for _, file := range files {
		filePath := filepath.Join(uploadPath, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Unable to save file: %v", err))
			return
		}
		uploadedFiles = append(uploadedFiles, file.Filename)
	}

	c.Redirect(http.StatusSeeOther, "/?success=Files%20uploaded%20successfully")
}

func showQRCode(c *gin.Context) {
	qrCodeData, err := utils.GenerateQRCodeWithLocalIP()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, qrCodeData)
}

func showImage(c *gin.Context, uploadPath string) {
	filename := c.Param("filename")
	filePath := filepath.Join(uploadPath, filename)

	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("File not found: %v", err))
		return
	}
	defer file.Close()

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Unable to read file: %v", err))
		return
	}

	c.HTML(http.StatusOK, "image.html", gin.H{
		"ImageData": base64.StdEncoding.EncodeToString(fileBytes),
		"Filename":  filename,
	})
}
