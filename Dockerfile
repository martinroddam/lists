# pulling a lightweight version of golang (in Alpine Linux)
FROM golang:1.8

ADD . /go/src/lists
WORKDIR /go/src/lists
COPY . .

RUN go get -u -d -v ./...
RUN go install -v ./...

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/lists"]
# Document that the service listens on port 9000.
EXPOSE 9000