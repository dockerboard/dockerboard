(function () {
  'use strict';

  angular.module('dockerboard.services')
    .factory('Containers', ['$resource', function ($resource) {
      var res = $resource('/api/containers/:id');

      res.queryParams = {
        all: false,
        limit: 20,
        size: false,
        since: '',
        before: ''
      };

      res.basicAttributes = [
        'Id',
        'Name',
        'Created',
        'Image'
      ];
      return res;
    }]);

})();
