
# Dockerboard

**Work-In-Process**

Simple dashboards, visualizations, managements for your dockers.

[Dockerboard][] and [Bluewhale][] are working together to make to the Docker awesome.

Lovingly created and maintained by [DockerPool][] Team.

[![gitter][gitter-image]][gitter-url]
[![GoDoc][godoc-image]][godoc-url]
[![build status][travis-image]][travis-url]
[![wercker status][wercker-image]][wercker-url]   
[![DockerBadge][docker-image]][docker-url]


## Screenshots

![Dockerboard Hub Screenshot](https://raw.githubusercontent.com/dockerboard/bluewhale/master/screenshots/hub_version_ping.gif)

![Dockerboard Hosts Screenshot](https://raw.githubusercontent.com/dockerboard/bluewhale/master/screenshots/hosts.gif)

![Dockerboard Screenshot](https://raw.githubusercontent.com/dockerboard/bluewhale/master/screenshots/dockerboard.gif)


## Features

* Multi Hosts, switch, add, delete
* Search for an image on Docker Hub, and create image
* Run locally


## Quick Start

### Run locally

```
$(boot2docker shellinit)
go build
./dockerboard -h
./dockerboard server -h
./dockerboard server -s bluewhale/dist -p 8888
```

### Open Brower

```
open http://127.0.0.1:8001
# Or
open http://$(boot2docker ip 2>/dev/null):8001 
```

### Build & Run

```
docker build -t dockerboard/dockerboard github.com/dockerboard/dockerboard
docker run -d -p 8001:8001 -v /var/run/docker.sock:/var/run/docker.sock --name dockerboard  dockerboard/dockerboard
```

Or Pull From [Docker Hub][]

```
docker pull dockerboard/dockerboard
docker run -d -p 8001:8001 -v /var/run/docker.sock:/var/run/docker.sock --name dockerboard  dockerboard/dockerboard
```

Or Built with [Bluewhale][]

```
docker build -t dockerboard/bluewhale github.com/dockerboard/bluewhale
docker build -t dockerboard/dockerboard github.com/dockerboard/dockerboard
docker run -d -v /bluewhale/dist --name bluewhale dockerboard/bluewhale
docker run -d -p 8001:8001 -v /var/run/docker.sock:/var/run/docker.sock --volumes-from bluewhale --name dockerboard  dockerboard/dockerboard
```

### Connect via a http/https Or a unix sock

  If using [boot2docker][], these are some ENV variables.

```
export DOCKER_HOST="tcp://0.0.0.0:2376"
export DOCKER_CERT_PATH="$HOME/.boot2docker/certs/boot2docker-vm"
export DOCKER_TLS_VERIFY="1"
```


## Development

Dockerboard RESTful API Prefix: `http://localhost:8001/api`.

```
// Maybe you need.
$(boot2docker shellinit)
go get ./..
go run dockerboard.go
```

[Dockerboard]: https://github.com/dockerboard/dockerboard
[Bluewhale]: https://github.com/dockerboard/bluewhale
[DockerPool]: http://dockerpool.com/
[boot2docker]: http://boot2docker.io/
[docker hub]: https://hub.docker.com/

[gitter-image]: https://badges.gitter.im/Join%20Chat.svg
[gitter-url]: https://gitter.im/dockerboard/dockerboard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge
[godoc-image]: https://godoc.org/github.com/dockerboard/dockerboard?status.svg
[godoc-url]: http://godoc.org/github.com/dockerboard/dockerboard
[travis-image]: https://img.shields.io/travis/dockerboard/dockerboard/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dockerboard/dockerboard
[wercker-image]: https://app.wercker.com/status/6134bd7bdb90915629df6d86e569a284/s
[wercker-url]: https://app.wercker.com/project/bykey/6134bd7bdb90915629df6d86e569a284
[docker-image]: http://dockeri.co/image/dockerboard/dockerboard
[docker-url]: https://registry.hub.docker.com/u/dockerboard/dockerboard/
