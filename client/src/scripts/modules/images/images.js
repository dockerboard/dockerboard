(function () {
  'use strict';

  dockerboardApp.registerModule('images.ctrl');

  angular.module('images.ctrl')
    .controller('ImagesCtrl', ImagesController)
    .controller('ImageCtrl', ImageController)
    .config(['$stateProvider',
      function ($stateProvider) {
        $stateProvider.
          state('images', {
            url: '/images',
            templateUrl: '/scripts/modules/images/images.html'
          })
          .state('imageitem', {
            url: '/images/:id',
            templateUrl: '/scripts/modules/images/image.html'
          });
      }
    ]);

  ImagesController.$inject = ['$scope', 'Images'];
  function ImagesController($scope, Images) {
    Images
      .index()
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

  ImageController.$inject = ['$scope', '$stateParams', '$filter', 'Images'];
  function ImageController($scope, $stateParams, $filter, Images) {
    $scope.tabs = [
      {
        title: 'Normal'
      },
      {
        title: 'Base'
      }
    ];

    $scope.basicAttributes = [];

    function formatBasicAttributes(image) {
      angular.forEach(Images.basicAttributes, function (k) {
        var v = image[k];
        if (k === 'Id' || k === 'Parent') {
          v = $filter('limitTo')(v, 8);
          var href = '#/images/' + v;
          v = '<a ng-href="' + href + '" href="' + href + '">' + v + '</a>';
        } else if (k === 'Size' || k === 'VirtualSize') {
          v = $filter('prettyBytes')(v);
        } else if (k === 'Created') {
          v = $filter('date')(v, 'yyyy-MM-dd HH:mm:ss Z');
        }

        this.push({
          key: k,
          value: v
        });
      }, $scope.basicAttributes);
    }

    Images
      .show($stateParams.id)
      .success(function (data) {
        $scope.image = data;
        formatBasicAttributes(data);
      });

  }
})();
