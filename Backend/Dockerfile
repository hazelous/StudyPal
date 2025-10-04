FROM golang:1.22-alpine AS build
RUN apk add --no-cache build-base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /out/studypal
# I may need CG0 for the SQLite

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/studypal ./studypal
EXPOSE 8080
CMD ["./studypal"]
