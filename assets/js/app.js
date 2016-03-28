angular.module('app', ['ngRoute']).
  config(function($routeProvider, $locationProvider) {
    $locationProvider.html5Mode(false).hashPrefix("!");
    $routeProvider.
      when("/hello_world", {
        templateUrl: "templates/hello_world.html"
      }).
      otherwise("/hello_world");
  });
