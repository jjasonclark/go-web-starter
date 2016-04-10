angular.module('app').
  factory('authTokenService', function($window) {
    var storageKey = "authToken";
    var token = getStoredToken();

    function get() {
      return token;
    }

    function set(newValue) {
      token = newValue;
      setStoredToken();
    }

    function getStoredToken() {
      return $window.localStorage.getItem(storageKey) || "";
    }

    function setStoredToken() {
      $window.localStorage.setItem(storageKey, token);
    }

    function clear() {
      token = undefined;
      $window.localStorage.removeItem(storageKey);
    }

    return {
      get: get,
      set: set,
      clear: clear
    };
  });
