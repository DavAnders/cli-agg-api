-- +goose Up
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL UNIQUE,
    content TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    polarity FLOAT,
    subjectivity FLOAT,
    query VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS articles;