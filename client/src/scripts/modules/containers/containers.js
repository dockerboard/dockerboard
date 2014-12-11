(function () {
  'use strict';

  dockerboardApp.registerModule('containers.ctrl');

  angular.module('containers.ctrl')
    .controller('ContainersCtrl', ContainersController)
    .config(['$stateProvider',
      function ($stateProvider) {
        $stateProvider.
          state('containers', {
            url: '/containers',
            templateUrl: '/scripts/modules/containers/containers.html'
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

})();
