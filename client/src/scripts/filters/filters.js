(function () {
  'use strict';

  dockerboardApp.registerModule('dockerboard.filters');

  angular.module('dockerboard.filters')
    .filter('sanitize', ['$sce', function ($sce) {
      return function(htmlCode) {
        return htmlCode ? $sce.trustAsHtml(htmlCode + '') : '';
      }
    }])

    .filter('formatImageId', ['limitToFilter', function (limitToFilter) {
      var reg = /[\-_:.]/g;
      function formatImageId(image) {
        if (reg.test(image)) {
          return image;
        }
        return limitToFilter(image, 12);
      }
      return formatImageId;
    }]);

})();