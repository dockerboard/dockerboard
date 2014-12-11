(function () {
  'use strict';

  angular.module('dockerboard.services')
    .factory('Containers', ['$http', function ($http) {
      return {
        index: function () {
          return $http.get('/api/containers');
        },
        show: function (id) {
          return $http.get('/api/containers/' + id);
        }
      };
    }]);

})();
