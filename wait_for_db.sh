# We use sh instead of bash because this was made for Linux Alpine
# This script assumes that netcat (nc) is installed, so the Dockerfile should install it
# because it is not included in Linux Alpine

#!/bin/sh

TIMEOUT=${DB_TIMEOUT:-10}
DB_PORT=${DB_PORT:-5432}
DB_HOST=${DB_HOST:-mu_users_db}
echo "Checking connection to the database at $DB_HOST:$DB_PORT with timeout $TIMEOUT"

# Many bash commands don't work on Linux Alpine because it uses a different shell: Busybox Ash
# This explain the strange approach to this script
start_time=$(date +%s)

while true; do
  nc -z "$DB_HOST" "$DB_PORT" >/dev/null 2>&1
  if [ $? -eq 0 ]; then
    echo "Connected successfully to the database!"
    exit 0
  fi

  now=$(date +%s)
  elapsed=$((now - start_time))

  if [ "$elapsed" -ge "$TIMEOUT" ]; then
    echo "Timeout reached: Could not connect to database at $DB_HOST:$DB_PORT"
    exit 1
  fi

  echo "Waiting for database at $DB_HOST:$DB_PORT..."
  sleep 1
done