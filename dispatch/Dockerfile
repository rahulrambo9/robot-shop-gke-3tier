FROM golang:1.22

WORKDIR /go/src/app

COPY *.go .

# Initialize the module
RUN go mod init dispatch || { echo 'go mod init failed'; exit 1; }

# Install dependencies
RUN go get || { echo 'go get failed'; exit 1; }

# Build the application
RUN go build -o dispatch || { echo 'go build failed'; exit 1; }

# Command to run the application
CMD ["./dispatch"]
