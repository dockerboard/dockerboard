(function () {
  'use strict';

  angular.module('dockerboard.services')
    .factory('Containers', ['$resource', function ($resource) {
      return $resource('/api/containers/:id');
    }]);

})();
