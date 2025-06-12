FROM golang:1.24-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only
# redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

# Install migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Copy start script and make it executable
COPY start.sh .
RUN chmod +x start.sh

# Copy the rest of the application code.
# (Necessary for building and code generation)
COPY . .

# Generate repositories
RUN sqlc generate

# Build the application
RUN go build -v -o main .

# `MODE` can be one of: debug, release or test
ARG MODE=release
ENV GIN_MODE=${MODE}

EXPOSE ${PORT:-8008}

RUN echo "Built in ${MODE} mode with port set to:" ${PORT:-8008}

# Run migrations before starting the application
CMD ["./start.sh"]
