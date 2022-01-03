package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.Handle("/", middleware(http.HandlerFunc(helloWorldHandler)))
	http.ListenAndServe(":8080", nil)

}

func helloWorldHandler(w http.ResponseWriter, request *http.Request) {

	reqId := request.Context().Value("reqId")
	if reqId != nil {
		w.Write([]byte(fmt.Sprintf("hello world with request id! %s", reqId)))
		return
	}

	w.Write([]byte(fmt.Sprintf("hello world!")))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), "reqId", id)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
