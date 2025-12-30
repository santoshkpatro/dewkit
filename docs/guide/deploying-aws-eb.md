# Deploying on AWS Elastic Beanstalk

DewKit can be deployed on AWS Elastic Beanstalk using a Go environment.

## Steps

1. Create a Go Elastic Beanstalk environment
2. Configure PostgreSQL (RDS) and Redis (Elasticache)
3. Set environment variables in EB console
4. Upload the built binary and assets
5. Run the install command during initial deployment

Elastic Beanstalk handles scaling, health checks, and restarts.
