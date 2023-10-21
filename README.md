# docker-phasik.tv

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-green?logo=pre-commit&labelColor=1f2d23&color=brightgreen)](https://github.com/pre-commit/pre-commit)
[![pre-commit](https://github.com/LyraPhase/docker-phasik.tv/actions/workflows/pre-commit.yml/badge.svg)](https://github.com/LyraPhase/docker-phasik.tv/actions/workflows/pre-commit.yml)
![DigitalOcean](https://img.shields.io/badge/DigitalOcean-%230167ff.svg?style=for-the-badge&logo=digitalOcean&logoColor=white)

Docker container to host `watch.phasik.tv`, and possibly other endpoints in the future.

## Building

    docker build .  -t  lyraphase/phasik.tv:latest

## Running

    docker run -d -p 8080:80  lyraphase/phasik.tv:latest

## Environment Variables

The docker container currently supports the following environment variables:

- `PORT`: The port to listen on
