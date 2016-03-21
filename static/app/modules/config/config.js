/**
 * Created by mihkelk on 16.02.2016.
 */
var foxApp = angular.module("fox", []);

foxApp.config(function ( $httpProvider) {
    $httpProvider.defaults.headers.common["Accept"] = "application/json";
    $httpProvider.defaults.headers.common["Content-Type"] = "application/json";
});

foxApp.constant("configConstant", {
   backendURL: 'http://localhost:9000/api'
});