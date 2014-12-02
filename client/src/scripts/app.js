'use strict';

angular.module(dockerboardApp.name, ['famous.angular', 'ui.router', 'ngMaterial'])
  .config(['$locationProvider', '$urlRouterProvider',

    function ($locationProvider, $urlRouterProvider) {
      // Redirect to home view when route not found
      $urlRouterProvider.otherwise('/');

      // use the HTML5 History API
      $locationProvider.html5Mode(true);
    }
  ]);

//Then define the init function for starting up the application
angular.element(document).ready(function() {
  //Fixing facebook bug with redirect
  if (window.location.hash === '#_=_') window.location.hash = '#!';

  //Then init the app
  //angular.bootstrap(document, [dockerboardApp.name]);
});
