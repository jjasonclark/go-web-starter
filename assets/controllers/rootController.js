angular.module("app").controller("RootController", function($scope, authTokenService, sessionService) {
  "use strict";

  function setup() {
    $scope.authToken = authTokenService.get();
    $scope.loggedIn = $scope.authToken && $scope.authToken !== "";
  }

  $scope.logout = function () {
    sessionService.logout()['finally'](function() {
      setup();
    });
  };

  setup();
});
