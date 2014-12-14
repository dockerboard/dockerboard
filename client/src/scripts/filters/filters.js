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
      var reg = /[\-\:\._]/;

      return formatImageId;

      function formatImageId(id) {
        if (reg.exec(id)) {
          return id;
        }
        return limitToFilter(id, 12);
      }
    }]);

})();