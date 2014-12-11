(function (angular, app) {
  'use strict';

  var appName = app.name = 'dockerboard';
  app.dependencies = ['ngResource', 'ngMaterial', 'ngAnimate', 'ui.router', 'prettyBytes'];
  app.registerModule = registerModule;

  function registerModule(moduleName, dependencies) {
    angular.module(moduleName, dependencies || []);
    angular.module(appName).requires.push(moduleName);
  }

})(window.angular, window.dockerboardApp || (window.dockerboardApp = {}));
