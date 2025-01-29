#!/bin/bash

FOLDER1="test/input"
FOLDER2="test/response"

COUNT1=$(find "$FOLDER1" -type f | wc -l)
COUNT2=$(find "$FOLDER2" -type f | wc -l)

if [ "$COUNT1" -gt "$COUNT2" ]; then
    echo "Error: $FOLDER1 contains more files than $FOLDER2"
    exit 1
fi

for ((i=1; i<=COUNT1; i++)); do
    TESTFILE="test/input/${i}_input.txt"
    RESPONSEFILE="test/response/${i}_response.txt"
    
    # echo "Executing main.go with $TESTFILE and $RESPONSEFILE for i = $i"
    
    go run src/main.go "$TESTFILE" "$RESPONSEFILE" a
    EXIT_CODE=$?
    # echo "Exit code: $EXIT_CODE"

    if [ "$EXIT_CODE" -ne 0 ]; then
        echo "Error: $TESTFILE and $RESPONSEFILE Test failed"
    fi
    if [ "$EXIT_CODE" -eq 0 ]; then
        echo "$TESTFILE and $RESPONSEFILE Test are OK"
    fi

    sleep 0.2
done