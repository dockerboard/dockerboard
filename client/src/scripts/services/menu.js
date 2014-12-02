'use strict';
(function () {

  angular.module('dockerboard.services')
    .factory('Menu', ['$rootScope', '$location', function ($rootScope, $location) {
      var sections = [
        {
          name: 'Apps'
        },
        {
          name: 'Images'
        }
      ]

      var self;

      $rootScope.$on('$locationChangeSuccess', onLocationChange);

      return self = {
        sections: sections,

        selectSection: function(section) {
          self.openedSection = section;
        },
        toggleSelectSection: function(section) {
          self.openedSection = (self.openedSection === section ? null : section);
        },
        isSectionSelected: function(section) {
          return self.openedSection === section;
        },
        selectPage: function(section, page) {
          page && page.url && $location.path(page.url);
          self.currentSection = section;
          self.currentPage = page;
        },
        isPageSelected: function(section, page) {
          return self.currentPage === page;
        }
      };

      function onLocationChange() {
        var activated = false;
        var path = $location.path();
        sections.forEach(function(section) {
          if (section && section.pages) {
            section.pages.forEach(function(page) {
              if (path === page.url) {
                self.selectSection(section);
                self.selectPage(section, page);
                activated = true;
              }
            });
          }
        });
        if (!activated) {
          //self.selectSection(sections[0]);
        }
      }
    }])

})();
