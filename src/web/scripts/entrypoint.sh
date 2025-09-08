#!/usr/bin/env sh
set -eu

cd /web

# Not initialised yet? Park the container nicely.
if [ ! -f "svelte.config.js" ]; then
  echo "⚙️  Project hasn't been initialised — run your init steps inside the container."
  exec tail -f /dev/null
fi

echo "🔧 Ensuring pnpm is available..."
if command -v corepack >/dev/null 2>&1; then
  corepack enable >/dev/null 2>&1 || true
  corepack prepare pnpm@9 --activate >/dev/null 2>&1 || true
fi

# Fallback if pnpm isn't on PATH after corepack
if ! command -v pnpm >/dev/null 2>&1; then
  echo "📦 corepack didn't activate pnpm; installing pnpm@9 globally..."
  npm i -g pnpm@9
fi

# Install deps if missing/empty
if [ ! -d "node_modules" ] || [ -z "$(ls -A node_modules 2>/dev/null || true)" ]; then
  echo "📚 Installing dependencies with pnpm..."
  pnpm install
fi

echo "🚀 Starting..."
pnpm run debug