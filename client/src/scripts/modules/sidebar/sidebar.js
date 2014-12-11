(function () {
  'use strict';

  // Use Applicaion configuration module to register a new module
  dockerboardApp.registerModule('sidebar.component');

  angular.module('sidebar.component')
    .controller('SidebarCtrl', SidebarController);

  SidebarController.$inject = ['$scope', '$location', '$mdSidenav', '$timeout', '$rootScope', 'Menu'];
  function SidebarController($scope, $location, $mdSidenav, $timeout, $rootScope, Menu) {
    $scope.menu = Menu;

    var mainContentArea = document.querySelector("[role='main']");

    $rootScope.$on('$locationChangeSuccess', openPage);

    $scope.closeMenu = function() {
      $timeout(function() { $mdSidenav('left').close(); });
    };

    $scope.openMenu = function() {
      $timeout(function() { $mdSidenav('left').open(); });
    };

    $scope.path = function() {
      return $location.path();
    };

    $scope.goHome = function($event) {
      $scope.menu.selectSection(null);
      $scope.menu.selectPage(null, null);
      $location.path( '/' );
    };

    function openPage() {
      $scope.closeMenu();
      mainContentArea && mainContentArea.focus();
    }
  }

})();
