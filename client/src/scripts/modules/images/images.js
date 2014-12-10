(function () {
  'use strict';

  dockerboardApp.registerModule('images.ctrl');

  angular.module('images.ctrl')
    .controller('ImagesCtrl', ImagesController)
    .config(['$stateProvider',
      function ($stateProvider) {
        $stateProvider.
          state('images', {
            url: '/images',
            templateUrl: '/scripts/modules/images/images.html'
          });
      }
    ]);

  ImagesController.$inject = ['$scope', '$http'];
  function ImagesController($scope, $http) {
    $http.get('/api/images')
      .success(function (data) {
        $scope.images = data;
      });

    $scope.getRepo = function (tags) {
      var repo = '';
      if (tags.length) {
        repo = tags[0].split(':')[0];
      }
      return repo;
    };

    $scope.getTags = function (repos) {
      var tags = [];
      angular.forEach(repos, function (value) {
        var tag = value.split(':')[1];
        if (tag) this.push(tag);
      }, tags);
      return tags.join(', ');
    };
  }

})();
