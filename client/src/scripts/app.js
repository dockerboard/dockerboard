'use strict';

angular.module(dockerboardApp.name, ['famous.angular', 'ui.router', 'ngMaterial'])
  .config(['$locationProvider', '$urlRouterProvider', function($locationProvider, $urlRouterProvider) {

      // Redirect to home view when route not found
      $urlRouterProvider.otherwise('/');

      // use the HTML5 History API
      $locationProvider.html5Mode(true);
    }
    ])
  .run(['$rootScope', function($rootScope) {
      $rootScope.$on('$stateChangeStart', function(event, toState, toParams, fromState, fromParams) {
        console.log("State Change: transition begins!");
      });

      $rootScope.$on('$stateChangeSuccess', function(event, toState, toParams, fromState, fromParams) {
        console.log("State Change: State change success!");
      });

      $rootScope.$on('$stateChangeError', function(event, toState, toParams, fromState, fromParams) {
        console.log("State Change: Error!");
      });

      $rootScope.$on('$stateNotFound', function(event, toState, toParams, fromState, fromParams) {
        console.log("State Change: State not found!");
      });

      $rootScope.$on('$viewContentLoading', function(event, viewConfig) {
        console.log("View Load: the view is loaded, and DOM rendered!");
      });

      $rootScope.$on('$viewcontentLoaded', function(event, viewConfig) {
        console.log("View Load: the view is loaded, and DOM rendered!");
      });

  }])

//Then define the init function for starting up the application
angular.element(document).ready(function() {
  //Fixing facebook bug with redirect
  if (window.location.hash === '#_=_') {
    window.location.hash = '#!';
  }

  //Then init the app
  //angular.bootstrap(document, [dockerboardApp.name]);
});
