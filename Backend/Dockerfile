# ---- Runtime stage ----------------------------------------------------------
FROM alpine:3.20

# CGO-linked binary requirements
RUN apk add --no-cache ca-certificates tzdata libstdc++

WORKDIR /app
COPY --from=build /out/studypal ./studypal

# Make the app listen on 8000
ENV PORT=8000
EXPOSE 8000

# Run the binary
ENTRYPOINT ["./studypal"]
