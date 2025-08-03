#!/bin/bash
set -e

# Kill any running Delve processes
kill_delve() {
    echo "[start-dlv] Killing old Delve if running..."

    pkill -f "dlv exec" || true

    # Additional check to ensure port is not in use
    local pid=$(lsof -ti:40000 2>/dev/null)
    if [ ! -z "$pid" ]; then
        echo "[start-dlv] Forcefully killing process using port 40000..."
        kill -9 $pid || true
        sleep 1
    fi
}

# Start Delve debugger
start_delve() {
    echo "[start-dlv] Starting Delve..."
    dlv debug /app/cmd/api/main.go --headless --listen=0.0.0.0:40000 --api-version=2 --accept-multiclient &
    dlv_pid=$!
}

# Wait for Delve to become available
wait_for_delve() {
    echo "[start-dlv] Waiting for port 40000 to become available..."
    local max_attempts=30
    local attempt=1

    while ! nc -z localhost 40000 && [ $attempt -le $max_attempts ]; do
        echo "[start-dlv] Attempt $attempt/$max_attempts - Port not available yet, waiting..."
        sleep 0.5
        ((attempt++))
    done

    if nc -z localhost 40000; then
        echo "[start-dlv] Delve is now listening!"
        return 0
    else
        echo "[start-dlv] Delve failed to start after $max_attempts attempts"
        return 1
    fi
}

main() {
    kill_delve
    start_delve
    wait_for_delve

    # Return the status of wait_for_delve
    return $?
}


# Execute the main function
main
exit $?
