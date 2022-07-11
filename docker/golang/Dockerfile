
FROM golang:1.15.7-alpine as dev

ENV CGO_ENABLED 0
RUN apk update
RUN apk add alpine-sdk
COPY ./go-app /app/go
WORKDIR /app/go
RUN CGO_ENABLED=0 go build -o main .
CMD ["go", "run", "main.go"]