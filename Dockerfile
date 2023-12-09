# syntax=docker/dockerfile:1
FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY ./server/*.go ./server/
COPY ./trivia/*.go ./trivia/
COPY trivia.json ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /TriviaAPI
CMD ["/TriviaAPI"]
