# gotstock

REST API providing CRUD for pharmaceutical products. It can be run as a Docker container alongside an instance of PostgreSQL, both created by the `docker-compose.yml`. The app also uses a Schema Migration tool to keep track of DB schema changes. 

# Usage

Use `docker compose up` to run the API. After the API is up and running, run in another Shell `go test -v ./it`.
