#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Error: expected two arguments: the user id (a uuid) and a new username"
    exit 1
fi

ID=$1
NAME=$2
url="localhost:8008"

response=$(curl -X PUT "$url/api/users/$ID" -s -i -d "{\"name\": \"$NAME\"}" -H "Content-Type: application/json")
status=$?
if [ $status -ne 0 ]; then
    echo "Error calling the curl command"
    echo "Response: $response"
    echo "Status: $status"
    exit 1
else
    status_code=$(echo "$response" | head -n 1 | awk '{print $2}')
    total_lines=$(echo "$response" | wc -l)
    json=$(echo "$response" | tail -n $((total_lines - 5)))
    if [ $status_code -eq 200 ]; then
        echo "Successfully changed the user's name"
        echo "Server response:"
        echo "$json"
        exit 0
    else
        echo "Error. Server response:"
        echo "$json"
        exit 1
    fi
fi