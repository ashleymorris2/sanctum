
# Sanctum (name TBD) - Work in early progress

Sanctum is a personal data and metric aggregator designed to help users collect, track, and visualize a wide range of information both from connected external sources and manual input.

Track habits, health, and productivity in one place. Some examples include:
	- Sleep duration and quality
	- Workout frequency or exercise data
	- Time spent on tasks or focus sessions
	- Custom, user-defined metrics

Sanctum aims to provide insight into daily patterns and long-term trends, the mission is to take back ownership of personal data that is often siloed into many different walled gardens. 

## Features (Planned & In Progress)
 - Unified dashboard for personal metrics
 - API ingestion from external providers
 - Manual metric entry via web UI
 - Custom metric definitions and visualizations
 - Aggregations and insights (e.g. daily, 7-day, monthly trends)
 - Self-hosted deployment via Docker Compose
 - Privacy-focused local storage (no cloud dependencies)

## Tech Stack
### Backend
 - Go | _API, Auth_
	- Echo (routing, middleware)
### Frontend
- Sveltekit | _Frontend_
	- Tailwind, DaisyUI (Styling)
### Data
- Postgres | _Relational Storage_
	- SQLC (code generate from Schema)
	- Goode (migrations)
