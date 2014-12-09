'use strict';

dockerboardApp.registerModule('board');

angular.module('board')
  .config(['$stateProvider',
    function ($stateProvider) {
      $stateProvider.
        state('board', {
          url: '/board',
          templateUrl: '/scripts/modules/board/board.html'
        });
    }
  ]);

