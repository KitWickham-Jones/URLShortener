package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleShorten(w http.ResponseWriter, r *http.Request){
	var body struct{
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if body.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	
	ctx := r.Context()
	slug := generateSlug()

	err := s.store.InsertURL(ctx, slug, body.URL)
	if err != nil{
		http.Error(w, "Failed to write url to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"short_url" : s.config.BaseURL + "/"+ slug,
	})
}

func (s *Server) handleRedirect(w http.ResponseWriter, r *http.Request){
	
	slug := r.PathValue("slug"); if slug == ""{
		http.NotFound(w, r)
		return
	}

	longurl, err := s.store.GetURL(r.Context(), slug)
	if err != nil{
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longurl, http.StatusFound)

}