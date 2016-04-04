angular.module('app', ['ngRoute']).
  config(function($routeProvider, $locationProvider) {
    $locationProvider.html5Mode(false).hashPrefix("!");
    $routeProvider.
      when("/", {
        templateUrl: "templates/hello_world.html"
      }).
      otherwise("/");
  });
