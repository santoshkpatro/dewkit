# Architecture

DewKit follows a simple two-tier architecture with a single backend service and a frontend web application.

## High-Level Architecture

- Backend service written in Go
- Frontend written in Vue.js
- PostgreSQL for persistent data storage
- Redis for caching and background operations

## Backend

The backend is a single Go application responsible for:

- Handling HTTP requests
- Managing authentication
- Ticket and conversation logic
- Database access

The backend produces a single executable binary (`dewkit`).

## Frontend

The frontend is built using Vue.js and bundled during the build process.

The compiled frontend assets are served by the Go backend.

## Database

PostgreSQL is used as the primary datastore.

The schema is defined in `schema.sql` and must be applied during initial installation.

## Cache

Redis is used for caching and transient data.
