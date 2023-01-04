package middleware

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"net/http"
	"os"
)

func ImageUploadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a POST method with the "image" form field
		if r.Method == http.MethodPost && r.FormValue("image") != "" {
			// Parse the multipart form
			err := r.ParseMultipartForm(10 << 20)
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}
			// Get the image file from the form
			file, _, err := r.FormFile("image")
			if err != nil {
				http.Error(w, "Error getting image file", http.StatusBadRequest)
				return
			}
			defer file.Close()
			// Create a buffer to hold the image data
			buf := new(bytes.Buffer)
			io.Copy(buf, file)

			// Connect to Minio
			minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), false)
			if err != nil {
				http.Error(w, "Error connecting to Minio", http.StatusInternalServerError)
				return
			}

			// Save the image to Minio
			_, err = minioClient.PutObject("images", "image.jpg", buf, int64(buf.Len()), minio.PutObjectOptions{ContentType: "image/jpeg"})
			if err != nil {
				http.Error(w, "Error saving image to Minio", http.StatusInternalServerError)
				return
			}

			// Set the image URL in the request context
			url := fmt.Sprintf("%s/images/image.jpg", os.Getenv("MINIO_ENDPOINT"))
			//ctx := context.WithValue(r.Context(), "imageURL", url)
			ctx := context.WithValue(r.Context(), "imageURL", url)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	imageURL, ok := r.Context().Value("imageURL").(string)
	if !ok {
		http.Error(w, "Error getting image URL", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Image uploaded to %s", imageURL)))
}
