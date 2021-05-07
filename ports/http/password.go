package http

import "net/http"

func (hs *HttpServer) Forget(w http.ResponseWriter, r *http.Request) {}

func (hs *HttpServer) Reset(w http.ResponseWriter, r *http.Request) {}

func (hs *HttpServer) ValidateToken(w http.ResponseWriter, r *http.Request) {}
