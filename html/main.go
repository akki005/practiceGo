package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func sendBadRequest(rw http.ResponseWriter) {
	rw.WriteHeader(400)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(struct {
		Message string
	}{
		Message: "Bad request",
	})
}

func sendNotFoundRequest(rw http.ResponseWriter) {
	rw.WriteHeader(404)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(struct {
		Message string
	}{
		Message: "Not found",
	})
}

func sendInternalErrorRequest(rw http.ResponseWriter, err error) {
	rw.WriteHeader(500)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(struct {
		Message string
	}{
		Message: err.Error(),
	})
}

func sendSuccess(rw http.ResponseWriter) {
	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(struct {
		Message string
	}{
		Message: "success",
	})
}

func main() {
	os.Setenv("token", "abcd")
	token := os.Getenv("token")

	if token == "" {
		log.Fatal("please set token")
	}

	http.Handle("/upload", tokenValidatorMiddleWare(fileValidatorMiddleWare(http.HandlerFunc(fileUploadHandler))))
	http.Handle("/home", http.HandlerFunc(homeHandler))
	log.Println("started server on :8080")
	http.ListenAndServe(":8080", nil)
}

func tokenValidatorMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		tokenEnv := os.Getenv("token")
		if token != tokenEnv {
			sendBadRequest(rw)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func fileValidatorMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		_, header, err := r.FormFile("data")
		if header.Header["Content-Type"][0] != "image/png" {
			sendBadRequest(rw)
			return
		}
		if err != nil {
			sendInternalErrorRequest(rw, err)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func fileUploadHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		file, header, err := r.FormFile("data")
		if err != nil {
			sendInternalErrorRequest(rw, err)
			return
		}
		err = uploadFile(file, header.Filename)
		if err != nil {
			sendInternalErrorRequest(rw, err)
			return
		}
		sendSuccess(rw)
		return
	}

	sendNotFoundRequest(rw)

}

func uploadFile(srcFileContent io.Reader, fileName string) error {
	dst, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if _, err := io.Copy(dst, srcFileContent); err != nil {
		return err

	}

	return nil
}

type htmlData struct {
	Token string
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := htmlData{
			Token: "abcd",
		}
		tmpl.Execute(rw, data)
		return
	}

	sendNotFoundRequest(rw)
}
