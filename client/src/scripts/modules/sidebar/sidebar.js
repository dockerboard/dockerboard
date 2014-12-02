'use strict';
(function () {

  // Use Applicaion configuration module to register a new module
  dockerboardApp.registerModule('sidebar.component');

  angular.module('sidebar.component')
    .controller('SidebarCtrl', SidebarController);

  SidebarController.$inject = ['$scope', 'Menu'];
  function SidebarController($scope, Menu) {
    $scope.menu = Menu;
  }

})();
