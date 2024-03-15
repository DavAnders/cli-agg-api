package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DavAnders/go-blog-api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type ArticlePayload struct {
	Title        string     `json:"title"`
	Url          string     `json:"url"`
	Content      *string    `json:"content,omitempty"`
	PublishedAt  *time.Time `json:"published_at,omitempty"`
	Polarity     *float64   `json:"polarity,omitempty"`
	Subjectivity *float64   `json:"subjectivity,omitempty"`
	Query        *string    `json:"query,omitempty"`
}

func (p *ArticlePayload) ConvertToCreateArticleParams() database.CreateArticleParams {
	var content sql.NullString
	if p.Content != nil {
		content = sql.NullString{String: *p.Content, Valid: true}
	} else {
		content = sql.NullString{}
	}

	var publishedAt sql.NullTime
	if p.PublishedAt != nil {
		publishedAt = sql.NullTime{Time: *p.PublishedAt, Valid: true}
	} else {
		publishedAt = sql.NullTime{}
	}

	var polarity sql.NullFloat64
	if p.Polarity != nil {
		polarity = sql.NullFloat64{Float64: *p.Polarity, Valid: true}
	} else {
		polarity = sql.NullFloat64{}
	}

	var subjectivity sql.NullFloat64
	if p.Subjectivity != nil {
		subjectivity = sql.NullFloat64{Float64: *p.Subjectivity, Valid: true}
	} else {
		subjectivity = sql.NullFloat64{}
	}

	var query sql.NullString
	if p.Query != nil {
		query = sql.NullString{String: *p.Query, Valid: true}
	} else {
		query = sql.NullString{}
	}

	return database.CreateArticleParams{
		Title:        p.Title,
		Url:          p.Url,
		Content:      content,
		PublishedAt:  publishedAt,
		Polarity:     polarity,
		Subjectivity: subjectivity,
		Query:        query,
	}
}

func (cfg *ApiConfig) HandlerCreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var payload ArticlePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	arg := payload.ConvertToCreateArticleParams()

	id, err := cfg.DB.CreateArticle(r.Context(), arg)
	if err != nil {
		log.Printf("Error creating article: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error creating article")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]int32{"id": id})
}
