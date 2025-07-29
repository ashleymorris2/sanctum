#!/bin/sh
set -e

# Only initialize if not already done
if [ ! -f "/app/svelte.config.js" ]; then
  echo "⚙️ Project hasn't been initialised - manually run initialisation"

   tail -f /dev/null
else
  echo "Svelte project found. Running debug..."
  # Just ensure dependencies are installed

#  NODE_OPTIONS="--inspect=0.0.0.0:9000" npm run debug
   npm run debug
fi

# Executeç the passed command (CMD)
exec "$@"