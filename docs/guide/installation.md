# Installation

This guide explains how to install DewKit for the first time.

## Prerequisites

- Go (latest stable)
- Node.js (18+)
- PostgreSQL
- Redis

## Clone Repository

```bash
git clone https://github.com/santoshkpatro/dewkit
cd dewkit
```

## Build Project

```bash
go run build
```

This will compile the backend binary and frontend assets.

## Initial Installation

Run the install command once to initialize the database:

```bash
./dewkit install
```

This command applies `schema.sql` and performs first-time setup.
