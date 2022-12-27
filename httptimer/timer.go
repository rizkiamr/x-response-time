package httptimer

import (
	"net/http"
	"strconv"
	"time"
)

type responseWriterWithTimer struct {
	http.ResponseWriter

	isHeaderWritten bool
	start           time.Time
}

func Timed(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(&responseWriterWithTimer{w, false, time.Now()}, r)
	})
}

func (w *responseWriterWithTimer) WriteHeader(statusCode int) {
	duration := time.Now().Sub(w.start)
	us := int(duration.Truncate(1000*time.Nanosecond).Nanoseconds() / 1000)
	w.Header().Set("X-Response-Time", strconv.Itoa(us)+" us")

	w.ResponseWriter.WriteHeader(statusCode)
	w.isHeaderWritten = true
}

func (w *responseWriterWithTimer) Write(b []byte) (int, error) {
	if !w.isHeaderWritten {
		w.WriteHeader(200)
	}

	return w.ResponseWriter.Write(b)
}
