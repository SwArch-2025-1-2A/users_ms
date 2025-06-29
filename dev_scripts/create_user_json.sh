#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Error: expected two arguments: id and username"
    exit 1
fi

ID=$1
USERNAME=$2
endpoint="localhost:8008/api/users"
data="{\"id\": \"$ID\", \"username\": \"$USERNAME\"}"
response=$(curl -X "POST" -s -i "$endpoint" -H "Content-Type: application/json" -d "$data")
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
    if [ $status_code -eq 201 ]; then
        echo "Successfully created user:"
        echo "$json"
        exit 0
    else
        echo "Error creating the user:"
        echo "Status code: $status_code"
        echo "Response: $response"
        exit 1
    fi
fi