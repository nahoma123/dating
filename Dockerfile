FROM golang:latest

# Set the working directory to the app root
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the app
RUN go build cmd/main.go

# Expose the app on port 8000
EXPOSE 8000

# Run the app
CMD ["./main"]
