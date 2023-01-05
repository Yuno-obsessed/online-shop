package interfaces

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

type fileUpload struct {
	client     *minio.Client
	bucketName string
}

func NewFileUpload() *fileUpload {
	client, err := minio.New("http://localhost:9000/",
		os.Getenv("minio_secret_key"),
		os.Getenv("minio_acess_key"),
		true)

	if err != nil {
		log.Fatalf("Failed to create a new minio client")
		return nil
	}
	bucketName := "zusammenStorage"
	location := "eu-south-1"
	exists, _ := client.BucketExists(bucketName)
	if !exists {
		err = client.MakeBucket(bucketName, location)
		if err != nil {
			log.Fatalf("Failed to create a bucket %s", bucketName)
			return nil
		}
	}
	return &fileUpload{client: client, bucketName: bucketName}
}

type FileUploadInterface interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	ReplaceFile(file string, newFile *multipart.FileHeader) (string, error)
	DeleteFile(file string) error
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

	fileBytes := bytes.NewReader(buffer)
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = fu.client.PutObject(fu.bucketName, filePath, fileBytes, size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("Error putting object in bucket %v", err)
	}
	return filePath, nil
}

func (fu *fileUpload) ReplaceFile(file string, newFile *multipart.FileHeader) (string, error) {
	obj, err := fu.client.GetObject(fu.bucketName, file, minio.GetObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("Error in getting object, %v", err)
	}
	if obj != nil {
		return "", fmt.Errorf("There is no such object in a bucket, %v", err)
	}

	// Not sure,, prolly client.RemoveObject is needed here
	f, err := newFile.Open()
	if err != nil {
		return "", fmt.Errorf("Error opening a file, %v", err)
	}
	defer f.Close()

	size := newFile.Size
	if size > int64(10240000) {
		return "", fmt.Errorf("The file size is greater than 1024KB")
	}

	buffer := make([]byte, size)
	f.Read(buffer)
	fileType := http.DetectContentType(buffer)
	if !strings.HasPrefix(fileType, "image") {
		return "", fmt.Errorf("The file format is not valid")
	}
	filePath := path.Ext(newFile.Filename) + newFile.Filename
	fileBytes := bytes.NewReader(buffer)

	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = fu.client.PutObject(fu.bucketName, filePath, fileBytes, size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("Error putting object in bucket %v", err)
	}
	return filePath, nil
}

func (fu *fileUpload) DeleteFile(file string) error {
	err := fu.client.RemoveObject(fu.bucketName, file)
	if err != nil {
		return err
	}
	return nil
}
