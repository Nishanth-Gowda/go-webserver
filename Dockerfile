# Use the official Go image to build the binary
FROM golang:alpine

# Set the working directory within the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o myapp

# Expose the port your application will listen on
EXPOSE 5000

# Specify the command to run when the container starts
CMD [ "./myapp" ]
