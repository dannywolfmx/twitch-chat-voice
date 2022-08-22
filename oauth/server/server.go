package server

import (
	"context"
	_ "embed"
	"net/http"
	"time"
)

//go:embed index.html
var page []byte

type Server struct {
	*http.Server
}

func NewServer(port string) *Server {
	return &Server{
		Server: &http.Server{
			Addr: port,
		},
	}

}

func (s *Server) Run(path string) (string, error) {
	token := ""

	//Refresh the mux in every new Run
	mux := http.NewServeMux()
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token = r.URL.Query().Get("token")
		w.WriteHeader(http.StatusOK)
		go func() {
			time.Sleep(time.Second * 4)
			s.Shutdown(context.TODO())
		}()
	})

	s.Handler = mux

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return token, nil
}
