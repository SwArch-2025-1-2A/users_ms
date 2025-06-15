#!/bin/bash
if [ $# -ne 1 ]; then
    echo "Error: expected one argument: the user uuid"
    exit 1
fi

ID=$1
url="localhost:8008"
response=$(curl -s -i "$url/api/users/$ID")
status=$?
if [ $status -ne 0 ]; then
    echo "Error when calling the curl command"
    echo "Response: $response"
    exit 1
else 
    status_code=$(echo "$response" | head -n 1 | awk '{print $2}')
    total_lines=$(echo "$response" | wc -l )
    # I know that the first five lines don't hold the JSON
    json=$(echo "$response" | tail -n $((total_lines - 5)))
    if [ $status_code -eq 200 ]; then
        echo "Success. Got this answer from the server:"
        echo "$json"
        exit 0
    elif [ $status_code -eq 404 ]; then
        echo "No user with the uuid $ID"
        echo "Server response:"
        echo "$json"
        exit 1
    elif [ $status_code -eq 400 ]; then
        echo "It seems the uuid $ID is not valid"
        echo "Server response:"
        echo "$json"
        exit 1
    elif [ $status_code -eq 500 ]; then
        echo "Error at the server"
        echo "Server response:"
        echo "$json"
        exit 1
    else
        echo "Strange error"
        echo "$json"
        exit 1
    fi
fi