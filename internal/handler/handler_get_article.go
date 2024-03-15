package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func (cfg *ApiConfig) HandlerGetArticleByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	article, err := cfg.DB.GetArticleByID(r.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Article not found")
			return
		}
		log.Printf("Error getting article by ID: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error getting article by ID")
		return
	}

	respondWithJSON(w, http.StatusOK, article)
}
