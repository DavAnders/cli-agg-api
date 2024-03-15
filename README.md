# Go API for Content Aggregator and Analyzer

This Go API is part of a concept setup designed to work alongside a [Python CLI tool](https://github.com/DavAnders/cli-content-agg) for aggregating and analyzing content from NewsAPI. The API serves as a backend for storing, retrieving, and managing articles fetched by the CLI tool.

## Overview

The API is built with Go and uses PostgreSQL for data storage. It's designed to showcase how a Python CLI tool can interact with a Go-based backend API.

## Setup

### Prerequisites

- Go installed on your machine
- PostgreSQL database
- Configuration file setup for database connection details

### Configuration

Ensure your database connection details are correctly set up in the `config.go` file or whichever configuration method you prefer. The API expects the following environment variables for connecting to PostgreSQL:

- `DBHost`: Database host
- `DBPort`: Database port
- `DBUser`: Database user
- `DBPassword`: Database password
- `DBName`: Database name

### Running the API

To start the API server, either build main, or navigate to the project directory and run:

```bash
go run ./cmd/main.go
```
