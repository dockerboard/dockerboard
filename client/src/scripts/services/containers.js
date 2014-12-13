(function () {
  'use strict';

  angular.module('dockerboard.services')
    .factory('Containers', ['$resource', function ($resource) {
      var res = $resource('/api/containers/:id');
      res.basicAttributes = [
        'Id',
        'Name',
        'Created',
        'Image'
      ];
      return res;
    }]);

})();
