FROM golang:1.20

WORKDIR /

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

COPY cmd /api/cmd
COPY pkg /api/pkg
COPY go.mod /api/go.mod
COPY go.sum /api/go.sum

WORKDIR /api/cmd
RUN go version
RUN go mod download
RUN go build -o main

EXPOSE 8080
RUN chmod +x /api/cmd/main
RUN chmod +x /api/cmd/main

