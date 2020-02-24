FROM golang:1.13-alpine

ADD . .

# fix $GOPATH/go.mod exists but should not
ENV GOPATH=""

ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 5000
CMD ["./go-todolist-server"]