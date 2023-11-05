FROM golang:latest
WORKDIR /golangserver
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o golangapi .
RUN rm -rf ./services ./utils ./go.mod ./go.sum ./app ./main.go ./data ./README.md
EXPOSE 4080
CMD ["/golangserver/golangapi"]