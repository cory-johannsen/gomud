# Use a lightweight COS base image
FROM gcr.io/distroless/base:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go binary into the container
COPY bin/gunchete /app/

# Copy static assets into the container
COPY assets /app/assets
COPY server /app/server

# Specify the port your application listens on
EXPOSE 7000

# Set the entry point for the container
ENTRYPOINT ["/app/gunchete"]
