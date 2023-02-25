package server

import (
	_ "embed"
	"net/http"
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

func (s *Server) Run(path string, c chan string) error {
	//Refresh the mux in every new Run
	mux := http.NewServeMux()
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			c <- r.URL.Query().Get("token")

			w.WriteHeader(http.StatusOK)
		}
	})

	s.Handler = mux

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
