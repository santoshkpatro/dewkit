# Deploying on Heroku

DewKit can be deployed on Heroku using a custom buildpack or container-based deployment.

## Steps

1. Create a new Heroku app
2. Provision PostgreSQL and Redis add-ons
3. Set environment variables:

```bash
heroku config:set DB_URL=...
heroku config:set CACHE_URL=...
heroku config:set SECRET_KEY=...
```

4. Build and deploy the application
5. Run the install command once:

```bash
heroku run ./dewkit install
```

6. Start the server:

```bash
./dewkit runserver
```
