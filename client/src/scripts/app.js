'use strict';

angular.module('dockerboardApp',
  ['famous.angular', 'ui.router', 'ngAnimate'])
  .config(function ($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise("/");
  })
;
