# Operations

## Database Migrations

Currently, DewKit relies on `schema.sql` for database initialization.

Future schema changes should be applied manually or via migration tooling.

## Backups

- Schedule PostgreSQL backups regularly
- Redis can be treated as ephemeral

## Monitoring

Monitor:

- Application logs
- Database connections
- Memory usage
