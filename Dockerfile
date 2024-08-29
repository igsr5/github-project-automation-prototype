# Use Go 1.21 as the base image
FROM golang:1.23

# Install required tools
RUN apt-get update && apt-get install -y git \
    && rm -rf /var/lib/apt/lists/* \
    # Install the latest version (2.37.0) of the gh command
    && curl -sSL https://github.com/cli/cli/releases/download/v2.37.0/gh_2.37.0_linux_amd64.tar.gz | tar xz -C /tmp \
    && mv /tmp/gh_2.37.0_linux_amd64/bin/gh /usr/local/bin/ \
    && rm -rf /tmp/gh_2.37.0_linux_amd64

# Set the working directory
WORKDIR /app

ENV APP_PORT=80

# Copy the source code
COPY . .

# Execute the build command
RUN make build

# Command to run the server when the container starts
CMD ["./github-project-automation"]
