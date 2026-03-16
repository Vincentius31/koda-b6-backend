FROM golang:1.25.0-alpine AS build

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o backend-coffee cmd/main.go

RUN chmod +x backend-coffee


FROM alpine:latest

WORKDIR /app

COPY --from=build /workspace/backend-coffee /app

EXPOSE 8888

ENTRYPOINT ["/app/backend-coffee"]