
# Dockerboard

**Work-In-Process**

Simple dashboards, visualizations, managements for your dockers.

This is a APIs service.

[Dockerboard][] and [Bluewhale][] are working together to make to the Docker awesome.

Lovingly created and maintained by [DockerPool][] Team.

[![badge](http://dockeri.co/image/dockerboard/dockerboard)](https://registry.hub.docker.com/u/dockerboard/dockerboard/)

## Quick Start

```
docker build -t dockerboard/bluewhale github.com/dockerboard/bluewhale
docker build -t dockerboard/dockerboard github.com/dockerboard/dockerboard
docker run -d -v /bluewhale/dist --name bluewhale dockerboard/bluewhale
docker run -d -p 8001:8001 -v /var/run/docker.sock:/var/run/docker.sock --volumes-from bluewhale --name dockerboard  dockerboard/dockerboard
open http://127.0.0.1:8001
```

### Connect vai a http/https Or a unix sock

  If using [boot2docker][], these are some ENV variables.

```
export DOCKER_HOST="tcp://0.0.0.0:2376"
export DOCKER_CERT_PATH="$HOME/.boot2docker/certs/boot2docker-vm"
export DOCKER_TLS_VERIFY="1"
```


## Screenshots

![Dockerboard Screenshot](https://github.com/dockerboard/bluewhale/blob/master/screenshots/dockerboard.gif?raw=true)

## Development

```
go get ./..
go run dockerboard.go
```

[Dockerboard]: https://github.com/dockerboard/dockerboard
[Bluewhale]: https://github.com/dockerboard/bluewhale
[DockerPool]: http://dockerpool.com/
[boot2docker]: http://boot2docker.io/
