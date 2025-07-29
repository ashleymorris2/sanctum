# Development stage
FROM node:20-alpine as dev

WORKDIR /app

# Install pnpm globally
RUN npm install -g pnpm

COPY ./scripts/entrypoint.sh scripts/entrypoint.sh
RUN chmod +x scripts/entrypoint.sh

ENTRYPOINT ["./scripts/entrypoint.sh"]
