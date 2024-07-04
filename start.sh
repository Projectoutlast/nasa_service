#!/bin/sh

make run-nasa-grpc-service &
PID1=$!

make run-space-web-app &
PID2=$!

make run-auth-service &
PID3=$!

shutdown() {
    echo "Stopping..."
    kill -s SIGTERM $PID1
    kill -s SIGTERM $PID2
    kill -s SIGTERM $PID3
    wait $PID1
    wait $PID2
    wait $PID3
    echo "Bye!"
    exit 0
}

trap 'shutdown' SIGTERM SIGINT

wait $PID1
wait $PID2
wait $PID3