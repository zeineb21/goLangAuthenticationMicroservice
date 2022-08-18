FROM golang:1.17 
#AS builder
WORKDIR /go/src/securityMS
COPY . .
RUN go env -w GO111MODULE=on
RUN go build -o /securityMS


EXPOSE 8080
CMD ["/securityMS"]

FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/securityMS

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/securityMS/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 