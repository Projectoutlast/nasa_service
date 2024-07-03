#!/bin/sh

make run-nasa-grpc-service &
PID1=$!

make run-space-web-app &
PID2=$!

shutdown() {
    echo "Stopping..."
    kill -s SIGTERM $PID1
    kill -s SIGTERM $PID2
    wait $PID1
    wait $PID2
    echo "Bye!"
    exit 0
}

trap 'shutdown' SIGTERM SIGINT

wait $PID1
wait $PID2