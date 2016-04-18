/**
 * Created by mihkelk on 16.02.2016.
 */

var foxApp = angular.module("fox", [
    'ngRoute',
    'ngSanitize',
    'swaggerUi',
    'pascalprecht.translate'
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

foxApp.config(['$translateProvider', function ($translateProvider) {
    // add translation table
    var language = localStorage.language;
    if (!language) {
        language = 'en';
        localStorage.language = language;
    }

    $translateProvider.translations('en', translations_EN);
    $translateProvider.translations('et', translations_ET);
    $translateProvider.preferredLanguage(language);
}]);
