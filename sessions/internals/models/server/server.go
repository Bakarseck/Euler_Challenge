package server

import (
	"fmt"
	"net/http"
	"sessions/internals/handlers"
	"sessions/internals/models"
	"sessions/internals/utils"
)

type Server struct {
	Router *http.ServeMux
}

var (
	ENDPOINTS = []models.EndPoint{
		{Path: "/", Handler: handlers.Home, Method: http.MethodGet},
	}
)

func NewServer() Server {
	return Server{
		Router: http.NewServeMux(),
	}
}

func (s *Server) ConfigureRoutes() {
	for _, endpoint := range ENDPOINTS {
		s.Router.HandleFunc(endpoint.Path, s.handleRequest(endpoint.Path, endpoint.Handler, endpoint.Method))
	}
}

func (s *Server) handleRequest(path string, handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if path == r.URL.Path {
			if r.Method == method {
				handler(w, r)
				return
			}

			data := map[string]interface{}{"error": "Method Not allowed"}
			utils.RespondWithJSON(w, data, http.StatusMethodNotAllowed)
			return
		}

		data := map[string]interface{}{"error": "Resources not found"}
		utils.RespondWithJSON(w, data, http.StatusNotFound)
	}
}

func (s *Server) StartServer(port string) error {
	s.ConfigureRoutes()
	fmt.Printf("http://localhost:%v", port)
	fmt.Println()

	return http.ListenAndServe(":"+port, s.Router)
}
