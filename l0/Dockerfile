# first stage
# using golang apline builder with version from go.mod
FROM golang:1.24-alpine AS builder

# installing git for go mod download
RUN apk update && apk add --no-cache git

# set the container working directory
WORKDIR /app

# copying only go mod and go sum
COPY go.mod go.sum ./

# and downloading dependencies
RUN go mod download
# it allows to not download them again when rebuilding
# if no chages in go.mod were made

# copying all source code
COPY . .

# building application with flags
# CGO_ENABLED=0 - makes binary fully static without c libraries
# GOOS=linux - building for linux
# GOARCH=amd64 - building for amd64 architecture
# -ldflags - linking flags:
#   -w - remove debug data
#   -s - remove symbol table
# -a - rebuild all packages
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o main ./cmd/main.go

# second stage
# using alpine image to reduce final image size
FROM alpine:latest

# create a group and user to run the application (without root rights)
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

# set the container working directory
WORKDIR /app

# switching to new appuser
USER appuser

# copuing binary from builder stage
COPY --from=builder --chown=appuser:appgroup /app/main .

# set permission to execute file
RUN chmod +x /app/main

# declare that container uses port 8080
EXPOSE 8080

# starting application
CMD ["./main"]