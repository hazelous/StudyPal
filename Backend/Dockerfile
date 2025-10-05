# Build stage
FROM golang:1.24-alpine AS build
ENV GOTOOLCHAIN=auto
RUN apk add --no-cache build-base
WORKDIR /src

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
# CGO is required by mattn/go-sqlite3 via gorm's sqlite driver
RUN CGO_ENABLED=1 GOOS=linux go build -o /out/studypal

# Runtime stage
FROM alpine:3.20
# libgcc is required for CGO runtime on Alpine
RUN apk add --no-cache ca-certificates libgcc
WORKDIR /app

COPY --from=build /out/studypal ./studypal

# Defaults (can be overridden by Compose)
ENV PORT=8000 \
    DB_PATH=/app/data/study_pal.db

# ensure the data dir exists (Compose volume will mount on top)
RUN mkdir -p /app/data

EXPOSE 8000

# lightweight healthcheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 \
  CMD wget -qO- "http://127.0.0.1:${PORT}/health" >/dev/null || exit 1

CMD ["./studypal"]
