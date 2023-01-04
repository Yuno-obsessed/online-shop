package interfaces

import (
	"fmt"
	"github.com/minio/minio-go"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

type fileUpload struct{}

func NewFileUpload() *fileUpload {
	return &fileUpload{}
}

type FileUploadInterface interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func (fu *fileUpload) UploadFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	size := file.Size
	if size > int64(10240000) {
		return "", fmt.Errorf("The file size is greater than 1024KB")
	}

	buffer := make([]byte, size)
	f.Read(buffer)
	fileType := http.DetectContentType(buffer)
	if !strings.HasPrefix(fileType, "image") {
		return "", fmt.Errorf("The file format is not valid")
	}
	filePath := path.Ext(file.Filename) + file.Filename

	accessKey := os.Getenv("minio_secret_key")
	secKey := os.Getenv("minio_acess_key")
	//endpoint := os.Getenv()
	client,err := minio.New("/image/",accessKey,secKey,true)
	if err != nil{
		return "", fmt.Errorf(err.Error())
	}
	return filePath, nil
}
