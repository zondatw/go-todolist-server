FROM golang:1.13-alpine as builder

WORKDIR /go/app
ADD . .

# fix $GOPATH/go.mod exists but should not
ENV GOPATH=""

ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 5000


FROM debian

RUN apt update
RUN apt install -y curl

WORKDIR /usr/local/
COPY --from=builder /go/app/migrations /usr/local/migrations
COPY --from=builder /go/app/go-todolist-server /usr/local/go-todolist-server
COPY --from=builder /go/app/docker_entry.sh /usr/local/docker_entry.sh

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /usr/local/migrate

ENTRYPOINT ["sh", "-x", "docker_entry.sh"]
