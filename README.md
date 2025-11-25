# Gamma Mobility Test

This code base contains source for my Gamma Mobility Go Lang coding test.

## Architectural Decisions

The endpoints expect 100k customer payment requests every minute. To handle requests at this scale, I integrated a message broker (using kafka) that asynchronously handles processing payment requests to the API while allowing the endpoint process request at the required scale effortlessly.

I ensure the worker is Idempotent to avoid duplicate processing of payments.

I adotped hexagonal architecture as my major driver for seperation of concerns, abstracting business logic and overall codebase maintainability.

## Code Setup
- clone repository
- install go mod ependencies

## Setup Insfrastructure

The source contains a `docker-compose.yml` that sets up a kafka cluster, redis and postgres db.

- Run `docker-compose up`

To setup project infrastructure.

## API Server
To start the API server

- cd in project root folder
- run `go run main.go`

## Worker
To start the worker process

- open a new terminal
- cd into project root directory
- run `go run worker/main.go`

## SQL
I have included a `db.sql` file that contains the sql script required to create tables for the database. Please run the script in your DBMS of choice after setting up project infrastructure and before calling the API endpoints.

## 
