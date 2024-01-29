package internalhttp

import (
	"net/http"
	"strings"
	"time"

	"github.com/lixoi/survey/internal/app"
)

/*
IP клиента;
* дата и время запроса;
* метод, path и версия HTTP;
* код ответа;
* latency (время обработки запроса, посчитанное, например, с помощью middleware);
* user agent, если есть.
*/

func requestLog(r *http.Request) (strings.Builder, time.Time) {
	var log strings.Builder
	// IP:Port клиента
	if r.RemoteAddr != "" {
		log.WriteString(r.RemoteAddr)
		log.WriteString(" ")
	}
	// время запроса
	timeStamp := time.Now()
	log.WriteString(timeStamp.UTC().String())
	log.WriteString(" ")
	// метод
	log.WriteString(r.Method)
	log.WriteString(" ")
	// путь
	log.WriteString(r.URL.Path)
	log.WriteString(" ")
	// версия HTTP
	log.WriteString(r.Proto)
	log.WriteString(" ")
	// код ответа
	if r.Response != nil {
		log.WriteString(r.Response.Status)
		log.WriteString(" ")
	}
	// user agent
	if len(r.Header) > 0 && len(r.Header.Get("User-Agent")) > 0 {
		log.WriteString(string(r.Header.Get("User-Agent")[0]))
		log.WriteString(" ")
	}

	return log, timeStamp
}

func loggingMiddleware(next http.HandlerFunc, logger app.Logger) http.HandlerFunc { //nolint:unused
	return func(w http.ResponseWriter, r *http.Request) {
		// args := r.URL.Query()
		// arg := args.Get("method")
		if r.URL.Path == "" {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error("request")
			return
		}
		log, latency := requestLog(r)
		next(w, r)
		// всремя обработки запроса
		log.WriteString(time.Now().Sub(latency).String())
		logger.Info(log.String())
	}
}
