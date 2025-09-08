# Development stage
FROM node:20-bookworm AS dev

WORKDIR /web

# Use Corepack to manage pnpm (much more reliable than global install)
RUN corepack enable && corepack prepare pnpm@9 --activate

# Copy only manifests first to leverage layer caching
COPY package.json pnpm-lock.yaml* ./

# Install deps early (dev container will mount code later)
RUN pnpm install

# Copy entrypoint
COPY ./scripts/entrypoint.sh ./scripts/entrypoint.sh
RUN chmod +x ./scripts/entrypoint.sh

ENV VITE_HOST=0.0.0.0 \
    CHOKIDAR_USEPOLLING=1 \
    WATCHPACK_POLLING=true

EXPOSE 5173 6006 9229

ENTRYPOINT ["./scripts/entrypoint.sh"]
