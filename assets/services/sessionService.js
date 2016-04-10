angular.module('app').
  factory('sessionService', function($http, authTokenService) {
    function callLoginApi(username, password) {
      return $http.post("/api/login", {
        username: username,
        password: password
      }).then(successLogin);
    }

    function callLogoutApi() {
      return $http.delete("/api/login").then(authTokenService.clear);
    }

    function callRegisterApi(username, password, passwordConfirmation) {
      return $http.post("/api/login/create", {
        username: username,
        password: password,
        passwordConfirmation: passwordConfirmation
      }).then(successLogin);
    }

    function successLogin(response) {
      authTokenService.set(response.data.token);
    }

    return {
      login: callLoginApi,
      logout: callLogoutApi,
      register: callRegisterApi
    };
  });
