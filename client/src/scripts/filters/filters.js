(function () {
  'use strict';

  dockerboardApp.registerModule('dockerboard.filters');

  angular.module('dockerboard.filters')
    .filter('sanitize', ['$sce', function ($sce) {
      return function(htmlCode) {
        return htmlCode ? $sce.trustAsHtml(htmlCode + '') : '';
      }
    }]);

})();