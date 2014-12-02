'use strict';

// Setting up route
angular.module('site').config(['$stateProvider', '$urlRouterProvider',
  function ($stateProvider, $urlRouterProvider) {
    // Home state routing
    $stateProvider.
      state('site', {
        url: '/',
        templateUrl: '/scripts/modules/site/views/index.html'
      });
  }
]);