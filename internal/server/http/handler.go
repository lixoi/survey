package internalhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/lixoi/survey/internal/app"
)

const defaultInterval = 2 * time.Second

type Service struct {
	sync.RWMutex
	Storage  app.App
	Interval time.Duration
}

type Response struct {
	Data  interface{} `json:"data"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func NewService(a app.App) *Service {
	return &Service{
		Storage:  a,
		Interval: defaultInterval,
	}
}

func (s *Service) GetHello(w http.ResponseWriter, r *http.Request) {
	resp := &Response{}
	resp.Data = "{\"hello\": \"hello word\"}"

	w.WriteHeader(http.StatusOK)
	WriteResponse(w, resp)
	return
}

func WriteResponse(w http.ResponseWriter, resp *Response) {
	resBuf, err := json.Marshal(resp)
	if err != nil {
		log.Printf("responce marshal error: %s", err)
	}
	_, err = w.Write(resBuf)
	if err != nil {
		log.Printf("responce marshal error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return
}
