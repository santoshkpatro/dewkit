# Configuration

DewKit requires only three environment variables.

## Required Environment Variables

```env
DB_URL=postgres://user:password@localhost:5432/dewkit
CACHE_URL=redis://localhost:6379
SECRET_KEY=your-secret-key
```

### DB_URL

PostgreSQL connection string.

### CACHE_URL

Redis connection string.

### SECRET_KEY

Used for signing sessions and security-sensitive operations.

## Configuration Philosophy

DewKit intentionally minimizes configuration to keep operations simple.
