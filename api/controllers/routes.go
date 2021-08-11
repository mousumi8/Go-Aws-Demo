package controllers

import "DemoProject/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/generateToken", middlewares.SetMiddlewareJSON(s.Authenticate)).Methods("POST")
	s.Router.HandleFunc("/resources", middlewares.SetMiddlewareAuthentication(s.CreateResource)).Methods("POST")
	s.Router.HandleFunc("/resources", middlewares.SetMiddlewareAuthentication(s.GetResources)).Methods("GET")
}
