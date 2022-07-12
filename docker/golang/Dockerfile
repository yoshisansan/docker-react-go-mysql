
FROM golang:1.15.7-alpine as dev

ENV CGO_ENABLED 0
RUN apk update
RUN apk add alpine-sdk
RUN apk add git
RUN go install -u github.com/cosmtrek/air
ENV PATH $AIR:~/go/bin/air
COPY ./go-app /app/go
WORKDIR /app/go
RUN go build -o /go/bin/air github.com/cosmtrek/air

RUN CGO_ENABLED=0 go build -o main .
CMD ["~/go/bin/air", "-c", ".air.toml"]