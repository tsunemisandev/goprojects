package utils

import (
	"encoding/base64"
	"fmt"
	"net"

	"github.com/skip2/go-qrcode"
)

func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func GenerateQRCodeWithLocalIP() (string, error) {
	localIP, err := GetLocalIP()
	if err != nil {
		return "", fmt.Errorf("Unable to get local IP address: %v", err)
	}

	url := fmt.Sprintf("http://%s:8080", localIP)
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return "", fmt.Errorf("Unable to generate QR code: %v", err)
	}

	png, err := qrCode.PNG(256)
	if err != nil {
		return "", fmt.Errorf("Unable to generate QR code: %v", err)
	}

	return base64.StdEncoding.EncodeToString(png), nil
}
