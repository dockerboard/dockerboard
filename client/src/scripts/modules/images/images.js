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

    $scope.queryParams = Images.queryParams;

    $scope.queryParamsFilters = '';

    $scope.fetch = function () {
      $scope.queryParams.filters = parseFilters($scope.queryParamsFilters);
      Images.query($scope.queryParams, function (data) {
        $scope.images = data;
      });
    };

    function parseFilters(text) {
      if (!text) return '';
      var filters = {};
      var arr = text.split(/\s+/g);
      for (var i = 0, l = arr.length; i < l; ++i) {
        var f = arr[i].split('=');
        if (f.length !== 2) {
          continue;
        }
        var name = f[0];
        var value = f[1];
        if (name && value) {
          filters[name] = filters[name] || [];
          filters[name].push(value);
        }
      }
      return JSON.stringify(filters);
    }

    $scope.fetch();

    $scope.search = function () {
      $scope.fetch();
    }

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

  ImageController.$inject = ['$scope', '$stateParams', 'limitToFilter', 'dateFilter', 'prettyBytesFilter', 'Images'];
  function ImageController($scope, $stateParams, limitToFilter, dateFilter, prettyBytesFilter, Images) {
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
          v = limitToFilter(v, 12);
          var href = '#/images/' + v;
          v = '<a ng-href="' + href + '" href="' + href + '">' + v + '</a>';
        } else if (k === 'Size' || k === 'VirtualSize') {
          v = prettyBytesFilter(v);
        } else if (k === 'Created') {
          v = dateFilter(v, 'yyyy-MM-dd HH:mm:ss Z');
        }

        this.push({
          key: k,
          value: v
        });
      }, $scope.basicAttributes);
    }

    Images.get({id: $stateParams.id}, function (data) {
      formatBasicAttributes(data);
      $scope.image = data;
    });

  }
})();
