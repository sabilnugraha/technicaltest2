package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	dto "technicaltest/dto/result"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("image")

		if err != nil && r.Method == "PATCH" {
			ctx := context.WithValue(r.Context(), "dataFile", "false")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		defer file.Close()

		// setup file type filtering
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "The provided file format is not allowed. Please upload a JPEG or PNG image"}
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		const MAX_UPLOAD_SIZE = 8000 // 8MB
		// Parse our multipart form, 8 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Max size in 8mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()

		filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(context.Background(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}