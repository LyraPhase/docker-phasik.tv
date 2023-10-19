# docker-phasik.tv

Docker container to host `watch.phasik.tv`, and possibly other endpoints in the future.

## Building

    docker build .  -t  lyraphase/phasik.tv:latest

## Running

    docker run -d -p 8080:80  lyraphase/phasik.tv:latest

## Environment Variables

The docker container currently supports the following environment variables:

- `PORT`: The port to listen on
