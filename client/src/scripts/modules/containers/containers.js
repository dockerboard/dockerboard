(function () {
  'use strict';

  dockerboardApp.registerModule('containers.ctrl');

  angular.module('containers.ctrl')
    .controller('ContainersCtrl', ContainersController)
    .controller('ContainerCtrl', ContainerController)
    .config(['$stateProvider',
      function ($stateProvider) {
        $stateProvider.
          state('containers', {
            url: '/containers',
            templateUrl: '/scripts/modules/containers/containers.html'
          })
          .state('containeritem', {
            url: '/containers/:id',
            templateUrl: '/scripts/modules/containers/container.html'
          });
      }
    ]);

  ContainersController.$inject = ['$scope', 'Containers'];
  function ContainersController($scope, Containers) {
    Containers
      .index()
      .success(function (data) {
        $scope.containers = data;
      });
  }

  ContainerController.$inject = ['$scope', '$stateParams', 'Containers'];
  function ContainerController($scope, $stateParams, Containers) {
    Containers
      .show($stateParams.id)
      .success(function (data) {
        $scope.container = data;
      });
  }
})();
