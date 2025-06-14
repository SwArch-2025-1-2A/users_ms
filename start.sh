#!/bin/sh
set -euo pipefail # strict mode

# Check if the DATABASE_URL environment variable is set
if [ -z "${DATABASE_URL}" ]; then
	echo "Error: DATABASE_URL environment variable for users microservice, is not set."
	exit 1
fi

./wait_for_db.sh
if [ $? -ne 0 ]; then
	echo "Error connecting to the database"
	exit 1
fi

# Run database migrations
echo "Running migrations..."
migrate -path db/migrations -database ${DATABASE_URL} -verbose up

echo "[[[USERS MICROSERVICE]]] Running the USERS MICROSERVICE now!"
./main
