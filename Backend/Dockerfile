# Build stage
FROM golang:1.24-alpine AS build
# GOTOOLCHAIN auto to automatically fetch new versions of golang
ENV GOTOOLCHAIN=auto
RUN apk add --no-cache build-base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO may be required for me to use SQLite
RUN go run .

# Runtime stage
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/studypal ./studypal
EXPOSE 8000
CMD ["./studypal"]
