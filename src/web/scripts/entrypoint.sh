#!/bin/sh
set -e

# Only initialize if not already done
if [ ! -f "/app/svelte.config.js" ]; then
  echo "⚙️ Project hasn't been initialised - manually run initialisation"

   tail -f /dev/null
else
  echo "Svelte project found. Checking for new dependencies..."
  # Just ensure dependencies are installed
  pnpm install
fi

# Executeç the passed command (CMD)
exec "$@"