FROM golang:1.16-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app/cmd
RUN go mod download
RUN go mod download github.com/jackc/chunkreader 
RUN go mod download github.com/jackc/pgproto3
RUN go build main.go

EXPOSE 8880

ENTRYPOINT "./main"