angular.module('app').config(function($routeProvider, $locationProvider, $httpProvider) {
  $httpProvider.interceptors.push('apiInterceptor');
  $locationProvider.html5Mode(false).hashPrefix("!");
  $routeProvider.
    when("/", {
      templateUrl: "templates/hello_world.html"
    }).
    otherwise("/");
});
