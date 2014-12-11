'use strict';

(function (angular, app) {

  var appName = app.name = 'dockerboard';
  app.registerModule = registerModule;
  app.dependencies = ['famous.angular', 'ui.router', 'ngResource', 'ngMaterial', 'ngAnimate', 'prettyBytes'];

  function registerModule(moduleName, dependencies) {
    angular.module(moduleName, dependencies || []);
    angular.module(appName).requires.push(moduleName);
  }

})(window.angular, window.dockerboardApp || (window.dockerboardApp = {}));
