# Hello World

A simple REST API built in `golang` using Labstack's [Echo](https://https://echo.labstack.com/]), to demonstrate how to build multi architecture container images using [Drone CI](https://drone.io) and [Buildx](https://docs.docker.com/build/buildx/install/)

## Pre-requisites

- [Docker Desktop](https://docs.docker.com/desktop/)
- [Drone CI CLI](https://docs.drone.io/cli/install/)

## Download Sources

Clone the sources and CD into it,

```shell
git clone https://github.com/kameshsampath/get-started-with-drone-ci.git && cd "$(basename "$_" .git)"
```

## Environment Setup

Start a local registry where the images will be pushed to,

```shell
docker run --name=registry -d -p 5001:5000 registry:2
```

## Build the Application

```shell
drone exec --trusted --env-file=.env
```

>**NOTE:** The drone pipeline uses .env file to load the config

## Run Application

```shell
docker-compose up
```

## Testing

```shell
curl http://localhost:8080/
```

The command should return `Hello World!`.

Add a `post`,

```shell
curl -X POST \
  --header "Content-Type: application/json" \
  --data @sample.json \
  http://localhost:8080/add
```

You should see a response like,

```json
{"id":1,"name":"Tom","text":"Searching for Jerry!"}
```

## Clean up

```shell
docker-compose down
```

Stop the local registry

```shell
docker stop registry
```
