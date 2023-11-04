FROM golang:latest
ADD . /golangserver
WORKDIR /golangserver
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .
RUN rm -rf ./services ./utils ./go.mod ./go.sum ./app ./main.go ./data ./README.md
EXPOSE 4080
CMD ["/golangserver/main"]