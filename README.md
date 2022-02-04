# go-migration-example
Example of database migration for Go application


```bash
docker run  -e POSTGRES_USER=postgres \
 -e POSTGRES_PASSWORD=postgres \
 -e POSTGRES_DB=myapp \
 -p 5432:5432 -it postgres:14.1-alpine
```