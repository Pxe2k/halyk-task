# Use official Golang image as the base
FROM golang:1.25

# Set the working directory inside the container
WORKDIR /app

# Copy all project files to the container
COPY . .

# Install dependencies
RUN go mod tidy

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "cmd/app/app.go"]
