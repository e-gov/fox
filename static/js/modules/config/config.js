/**
 * Created by mihkelk on 16.02.2016.
 */


var foxApp = angular.module("fox", []);


foxApp.controller('RegisterController', ['$scope', function ($scope) {
    $scope.testing = 'Hola!';

    $scope.foxList = [{id: 1, name: "Reinuvader"}, {id: 2, name: "Juhan"}];

    //TODO LISA FOX ADD IMPL
    $scope.add = function (fox) {
        console.log(fox);
    };

    //TODO LISA FOX UPDATE IMPL
    $scope.update = function(fox) {
        console.log(fox);
    };

    //TODO LISA FOX DELETE IMPL
    $scope.delete = function(fox) {
        console.log(fox);
    };

}]);