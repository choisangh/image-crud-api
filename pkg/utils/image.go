package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IsFileExists(filename string, path string) bool {
	// 파일 존재 여부 확인
	_, err := os.Stat(filepath.Join(path, filename))
	return err == nil
}

func IsValidImageFormat(data string) bool {
	// 이미지 유효성 검사
	prefix := "data:image/"
	suffixes := []string{"jpeg", "png", "gif"}

	// 이미지 데이터 포함 여부
	if !strings.HasPrefix(data, prefix) {
		return false
	}

	// 올바른 이미지 형식인지
	for _, suffix := range suffixes {
		if strings.Contains(data, suffix) {
			return true
		}
	}
	return false
}

func CreateImageFile(filename string, imageBase64 string, path string) error {
	// 이미지 디코딩
	data := strings.Split(imageBase64, ";base64,")[1]
	imageBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// 이미지 생성
	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// 이미지 파일 저장
	_, err = file.Write(imageBytes)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}
