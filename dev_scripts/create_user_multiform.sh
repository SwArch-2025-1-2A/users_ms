#!/bin/bash

if [ $# -ne 3 ]; then
    echo "Error: expected three arguments: id (a valid uuid), username and a path to a profile pic"
    exit 1
fi

ID=$1
NAME=$2
PIC_PATH=$3

url="localhost:8008"
response=$(curl -X POST -s -i "$url/api/users" -F "id=$ID" -F "name=$NAME" -F "profilePic=@$PIC_PATH")
status=$?
if [ $status -ne 0 ]; then
    echo "Error when using the curl command"
    echo "Response: $response"
    echo "Status: $status"
    exit 1
else
    status_code=$(echo "$response" | head -n 1 | awk '{print $2}')
    total_lines=$(echo "$response" | wc -l)
    # I know that the first five lines don't hold the JSON
    json=$(echo "$response" | tail -n $((total_lines - 5)))
    # 201 is Created
    if [ $status_code -eq 201 ]; then 
        echo "Successfully created a user"
        
        echo "Response body:"
        echo "$json"
        exit 0
    elif [ $status_code -eq 400 ]; then
        echo "Bad request:"
        echo "$json"
        exit 1
    elif [ $status_code -eq 500 ]; then
        echo "Error from the server"
        exit 1
    else 
        echo "Strange error: neither 400 nor 500"
        exit 1
    fi
fi