package middleware

import (
	"log"
	"net/http"
	"time"
)

type WriterWithStatus struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *WriterWithStatus) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.StatusCode = statusCode
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapperWriter := &WriterWithStatus{
			w, http.StatusOK,
		}
		next.ServeHTTP(wrapperWriter, r)
		if wrapperWriter.StatusCode != http.StatusOK {
      log.Println(wrapperWriter.StatusCode, r.Method, r.URL, "", time.Since(start)," - Error : ", http.StatusText(wrapperWriter.StatusCode))
		} else {
			log.Println(wrapperWriter.StatusCode, r.Method, r.URL, "", time.Since(start))
		}
	})
}
