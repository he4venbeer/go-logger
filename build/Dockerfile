# Start from the official Go image.
FROM golang:1.21.1

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy the go mod and sum files, then download the dependencies.
# This is an optimization step, ensuring that dependencies are cached unless changes to go.mod or go.sum are made.
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from the current directory (on your machine) to the PWD (Present Working Directory) inside the container.
COPY . .

# Build the Go app. Adjust this if your main file is located elsewhere.
RUN go build -o main .

# Command to run the executable.
CMD ["./main"]
