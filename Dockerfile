# ./Dockerfile
FROM golang:1.24-alpine AS builder
# Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .
# ADD .env .env

# Set necessary environment variables needed for our image 
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o warrantyapi .

# FROM alpine:latest
# # Move to /dist directory as the place for resulting binary folder
# WORKDIR /dist/
# # Copy binary from build to dist folder
# COPY --from=builder /build .
# # Export necessary port
# EXPOSE 9999
# # Command to run when starting the container.
# ENTRYPOINT ["./main"]

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/warrantyapi", "/"]
ENV TZ=Asia/Bangkok
# Command to run when starting the container.
ENTRYPOINT ["/warrantyapi"]
