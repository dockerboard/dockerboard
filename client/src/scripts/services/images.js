(function () {
  'use strict';

  angular.module('dockerboard.services')
    .factory('Images', ['$http', function ($http) {
      return {
        index: function () {
          return $http.get('/api/images');
        },
        show: function (id) {
          return $http.get('/api/images/' + id);
        },
        basicAttributes: [
          'Id',
          'Author',
          'Comment',
          'DockerVersion',
          'Architecture',
          'Os',
          'Size',
          'VirtualSize',
          'Created',
          'Parent'
        ]
      };
    }]);

})();
