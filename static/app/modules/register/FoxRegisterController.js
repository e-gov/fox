/**
 * Created by mihkelk on 17.02.2016.
 */

foxApp.controller("RegisterController", function ($scope, $log, FoxRegisterService) {

    function initRegisterList() {
        $scope.foxName = undefined;
        $scope.selectedFox = {};
        FoxRegisterService.getAll(function(result) {
            $scope.foxList = result.data;
        });
    }

    $scope.add = function (fox) {
        if (!fox.name || fox.name == '') {
            return;
        }
        FoxRegisterService.addFox(fox, function() {
            initRegisterList();
            // $scope.newFoxName = undefined;
        });
    };

    $scope.update = function(fox) {
        FoxRegisterService.updateFox(fox, function() {
            initRegisterList();
        });
    };

    $scope.delete = function(fox) {
        FoxRegisterService.deleteFox(fox.uuid, function() {
            initRegisterList();
        });
    };

    $scope.edit = function(fox) {
        FoxRegisterService.getFox(fox.uuid, function(result) {
            $scope.selectedFox = result.data;
        });
    };

    $scope.getAvailableParents = function(parentUuid) {
        if (!$scope.selectedFox) {
            return $scope.foxList;
        } else {
            return $scope.foxList.filter(function(fox) {
                return fox.uuid !== $scope.selectedFox.uuid
                    && $scope.selectedFox.parents
                        .filter(function(uuid) {
                            return uuid !== parentUuid; // exclude currently selected parent's uuid
                        }).indexOf(fox.uuid) === -1; // check if this fox is not already selected fox's parent
            });
        }
    };

    $scope.addParent = function() {
        if (!$scope.selectedFox.parents) {
            $scope.selectedFox.parents = [];
        }
        if ($scope.selectedFox.parents.indexOf("") < 0){
            $scope.selectedFox.parents.push("");
        }
    };

    $scope.removeParent = function(index) {
        $scope.selectedFox.parents.splice(index, 1);
    };

    $scope.changeLanguage = function(key) {
        FoxRegisterService.changeLanguage(key);
    };

    initRegisterList();

});
