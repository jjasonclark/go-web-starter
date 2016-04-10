angular.module("app").controller("LoginController", function($scope, sessionService, authTokenService, $location) {
  "use strict";

  $scope.allowSubmit = true;
  $scope.newUser = false;
  $scope.username = null;
  $scope.password = null;
  $scope.passwordConfirmation = null;

  function callLoginApi(params) {
    $scope.allowSubmit = false;
    return sessionService.
      login($scope.username, $scope.password)
      ['finally'](reenableSubmitting);
  }

  function callRegisterApi(params) {
    $scope.allowSubmit = false;
    return sessionService.
      register($scope.username, $scope.password, $scope.passwordConfirmation)
      ['finally'](reenableSubmitting);
  }

  function reenableSubmitting() {
    $scope.allowSubmit = true;
  }

  function clearPasswords() {
    $scope.password = null;
    $scope.passwordConfirmation = null;
  }

  function success(response) {
    clearPasswords();
    $location.path("/");
  }

  function failure() {
    clearPasswords();
  }

  $scope.login = function() {
    callLoginApi().then(success, failure);
  };

  $scope.register = function () {
    callRegisterApi().then(success, failure);
  };
});
