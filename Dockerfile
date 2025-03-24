FROM golang:1.13
WORKDIR /go/src/quickstart
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
# Install server application
CMD ["go", "run", "main.go"]