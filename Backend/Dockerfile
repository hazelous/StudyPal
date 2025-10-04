# Backend/Dockerfile
# syntax=docker/dockerfile:1

# ---- Build stage ------------------------------------------------------------
FROM golang:1.24-alpine AS build

# sqlite3 driver needs CGO enabled at build time
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOTOOLCHAIN=auto

RUN apk add --no-cache build-base
WORKDIR /src

# leverage module cache
COPY go.mod go.sum ./
RUN go mod download

# bring in the rest of the source
COPY . .

# compile a static-ish Linux binary
RUN go build -o /out/studypal .

# Runtime Stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# copy only the compiled binary
COPY --from=build /out/studypal ./studypal

# align with compose + app default
ENV PORT=8000

# document the port the app listens on (host publish is done by compose)
EXPOSE 8000

CMD ["./studypal"]
