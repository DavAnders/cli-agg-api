package handler

import (
	"database/sql"
	"log"
	"net/http"
)

func (cfg *ApiConfig) HandlerListArticleByQuery(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("query")
	if queryParam == "" {
		respondWithError(w, http.StatusBadRequest, "Query parameter is required")
		return
	}
	sqlQueryParam := sql.NullString{String: queryParam, Valid: true}

	articles, err := cfg.DB.ListArticlesByQuery(r.Context(), sqlQueryParam)
	if err != nil {
		log.Printf("Error listing articles by query: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error listing articles by query")
		return
	}

	respondWithJSON(w, http.StatusOK, articles)
}
