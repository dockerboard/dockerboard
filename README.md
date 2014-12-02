
# Dockerboard

Simple dashboards, visualizations, managements for your dockers.

Lovingly created and maintained by [DockerPool][] Team.


## Development

Currently it's very simple.   
Yes, we need `Dockerfile`. :)


### Install npm & bower packages

```
cd client/
npm install
bower install

cd ../
go get ./..
go run dockerboard.go
```


## Build With

- [gohttp](https://github.com/gohttp) &mdash; Our back end API is gohttp app. It responds to requests RESTfully in JSON.
- [Angular.js](https://www.angularjs.org/) &mdash; Our front end is an Angular.js app that communicates with the negroni API.
- [Famous.js](http://famo.us/) &mdash;  We use Famous.js to build beautiful experiences on any device.
- [D3.js](http://d3js.org/) &mdash; We use D3.js for drawing, bring data.


[DockerPool]: http://dockerpool.com/
