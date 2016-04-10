angular.module('app').config(function($routeProvider, $locationProvider, $httpProvider) {
  $httpProvider.interceptors.push('apiInterceptor');
  $locationProvider.html5Mode(false).hashPrefix("!");
  $routeProvider.
    when("/", {
      controller: "RootController",
      templateUrl: "templates/hello_world.html"
    }).
    when("/login", {
      controller: 'LoginController',
      templateUrl: "templates/login.html"
    }).
    otherwise("/");
});
