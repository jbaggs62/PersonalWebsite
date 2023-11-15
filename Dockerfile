# Build.
FROM golang:1.21 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY ./main.go /app/ 
COPY ./templates /app/templates



RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint


# Create a smaller final image
FROM alpine:3.18.4
WORKDIR /app

# Copy only the necessary files from the builder image
COPY --from=build-stage /entrypoint /entrypoint

# Set the entry point for the container
ENTRYPOINT ["./PersonalWebsite"]
