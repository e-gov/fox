/**
 * Created by mihkelk on 16.02.2016.
 */


var foxApp = angular.module("fox", []);

foxApp.config(function ( $httpProvider) {
    $httpProvider.defaults.useXDomain = true;
    //$httpProvider.defaults.withCredentials = true;
    delete $httpProvider.defaults.headers.common["X-Requested-With"];
    $httpProvider.defaults.headers.common["Accept"] = "application/json";
    $httpProvider.defaults.headers.common["Content-Type"] = "application/json";
});

foxApp.constant("configConstant", {
   backendURL: 'http://localhost:8090'
});