/**
 * Created by mihkelk on 16.02.2016.
 */
var foxApp = angular.module("fox", [
    'ngRoute',
    'ngSanitize',
    'swaggerUi'
]);

foxApp.config(function ($httpProvider) {
    $httpProvider.defaults.headers.common["Accept"] = "application/json";
    $httpProvider.defaults.headers.common["Content-Type"] = "application/json";
});

foxApp.constant("configConstant", {
    backendURL: 'http://localhost:9000/api'
});

foxApp.config(function ($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'views/main.html',
            controller: 'RegisterController'
        })
        .when('/swagger', {
            templateUrl: 'views/swagger.html'
        })
        .when('/about', {
            templateUrl: 'views/about.html'
        })
        .otherwise({
            redirectTo: '/'
        });
});