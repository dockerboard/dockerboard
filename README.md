
# Dockerboard

Simple dashboards, visualizations, managements for your dockers.


## Development

Currently it's simple.   
Yes, we need docker. :)


### Install nppm & bower packages

```
cd client/
npm install
bower install

cd ../
go get ./..
go run dockerboard.go
```


## Build With

- [negroni](https://github.com/codegangsta/negroni/) &mdash; Our back end API is negroni app. It responds to requests RESTfully in JSON.
- [Angular.js](https://www.angularjs.org/) &mdash; Our front end is an Angular.js app that communicates with the negroni API.
- [Famous.js](http://famo.us/) &mdash;  We use Famous.js to build beautiful experiences on any device.
- [D3.js](http://d3js.org/) &mdash; We use D3.js for drawing, bring data.

