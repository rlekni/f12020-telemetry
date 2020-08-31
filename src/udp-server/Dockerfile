FROM golang:latest AS builder

# Set Environment Variables
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
ADD . /app/

# Build app
RUN go get -d
RUN go build -a -installsuffix cgo -o main .

# Make the image smaller
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 20777/udp
CMD ["./main"]
