FROM golang:1.20

# Set the working directory
WORKDIR /go/src/app

# Copy all Go files to the working directory
COPY *.go ./

# Initialize the Go module with a proper module name
RUN go mod init github.com/rahulrambo9/dispatch

# Get dependencies
RUN go mod tidy

# Build the application
RUN go build -o dispatch

# Command to run the application
CMD ["./dispatch"]
