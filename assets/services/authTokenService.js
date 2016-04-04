angular.module('app').
  factory('authTokenService', function() {
    var tokens = [];

    function get() {
      if (tokens.lenth > 0) {
        return tokens[tokens.length - 1];
      }
      return "";
    }

    function push(token) {
      tokens.push(token);
    }

    return {
      get: get,
      push: push
    };
  });
