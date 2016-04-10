angular.module("app").
  factory('apiInterceptor', function() {
    var jsonContentType = 'application/json;charset=utf-8';
    var whiteListPrefixes = [
      '/api/'
    ];

    function isAPIPath(url) {
      for(var index in whiteListPrefixes) {
        var path = whiteListPrefixes[index];
        var result = url.split(path);
        if (result.length >= 2 && result[0] === "") {
          return true;
        }
      }
      return false;
    }

    function addAPIHeaders(config) {
      config.headers['X-Requested-With'] = "XMLHttpRequest";
      config.headers['Content-Type'] = jsonContentType;
      config.headers.accept = jsonContentType;
      return config;
    }

    return {
      request: function(current) {
        var config = angular.extend({
          url: "",
          headers: {}
        }, current);
        if (isAPIPath(config.url)) {
          addAPIHeaders(config);
        }
        return config;
      }
    };
  });
