#!/bin/bash
set -e

echo "[start-dlv] Killing old Delve if running..."
pkill -f "dlv exec" || true

echo "[start-dlv] Starting Delve..."
dlv debug ./cmd/api/main.go  ../main --headless --listen=0.0.0.0:40000 --api-version=2 --accept-multiclient & dlv_pid=$!

echo "[start-dlv] Waiting for port 40000 to become available..."
while ! nc -z localhost 40000; do
  sleep 0.2
done

echo "[start-dlv] Delve is now listening!"

# Signal success
exit 0