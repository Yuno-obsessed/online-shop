package interfaces

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type fileUpload struct {
	client     *minio.Client
	bucketName string
}

func NewFileUpload(entity string) *fileUpload {
	ctx := context.Background()
	client, err := minio.New("minio:9000/", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalf("Failed to create a new minio client, %v", err)
		return nil
	}
	bucketName := "zusammen.storage." + entity
	location := "us-east-1"
	exists, _ := client.BucketExists(ctx, bucketName)
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatalf("Failed to create a bucket %s, %v", bucketName, err)
			return nil
		}
		fmt.Printf("Bucker %s was created successfully", bucketName)
	}
	return &fileUpload{client: client, bucketName: bucketName}
}

type FileUploadInterface interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	ReplaceFile(file string, newFile *multipart.FileHeader) (string, error)
	DeleteFile(file string) error
}

func (fu *fileUpload) UploadFile(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()
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
	filePath := file.Filename

	fileBytes := bytes.NewReader(buffer)
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = fu.client.PutObject(ctx, fu.bucketName, filePath, fileBytes, size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("Error putting object in bucket %v", err)
	}
	return filePath, nil
}

func (fu *fileUpload) ReplaceFile(file string, newFile *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	obj, err := fu.client.GetObject(ctx, fu.bucketName, file, minio.GetObjectOptions{})
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
	filePath := newFile.Filename
	fileBytes := bytes.NewReader(buffer)

	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = fu.client.PutObject(ctx, fu.bucketName, filePath, fileBytes, size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("Error putting object in bucket %v", err)
	}
	return filePath, nil
}

func (fu *fileUpload) DeleteFile(file string) error {
	ctx := context.Background()
	err := fu.client.RemoveObject(ctx, fu.bucketName, file, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		return err
	}
	return nil
}
