# syntax=docker/dockerfile:1.6
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/worker ./

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/worker /app/worker
USER nonroot:nonroot
ENTRYPOINT ["/app/worker"]
