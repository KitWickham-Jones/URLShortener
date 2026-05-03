package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleShorten(w http.ResponseWriter, r *http.Request){
	var body struct{
		URL string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	

}