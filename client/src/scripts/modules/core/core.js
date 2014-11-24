'use strict';

// Use Applicaion configuration module to register a new module
angular.module('core', []);

angular.module('dockerboardApp').requires.push('core');