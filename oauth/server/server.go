package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

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
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		buf, err := os.ReadFile("./server/index.html")

		if err != nil {
			fmt.Println(err)
			s.Shutdown(context.TODO())
			return
		}

		w.Write(buf)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token = r.URL.Query().Get("token")
		s.Shutdown(context.TODO())
	})

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return token, nil
}
