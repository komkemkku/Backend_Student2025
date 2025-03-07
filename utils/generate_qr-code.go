package utils

import (
	"encoding/base64"
	"log"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCodeBase64 - ฟังก์ชันสร้าง QR Code และคืนค่าเป็น Base64
func GenerateQRCodeBase64(data string) (string, error) {
	// สร้าง QR Code เป็น PNG
	qrCode, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		log.Println("Error generating QR Code:", err)
		return "", err
	}

	// แปลง QR Code เป็น Base64
	qrBase64 := base64.StdEncoding.EncodeToString(qrCode)

	return qrBase64, nil
}
