
/**
 * Created by mihkelk on 17.02.2016.
 */


foxApp.controller('RegisterController', function ($scope, FoxRegisterService) {

    function initRegisterList() {
        $scope.foxName = undefined;
        FoxRegisterService.getAll(function(result) {
            $scope.foxList = result.data;
        });
    }

    $scope.add = function (foxName) {
        if (!foxName || foxName == '') {
            return;
        }
        FoxRegisterService.addFox(foxName, function() {
            initRegisterList();
            $scope.newFoxName = undefined;
        });
    };

    //TODO LISA FOX UPDATE IMPL
    $scope.update = function(fox) {
        console.log(fox);
    };

    $scope.delete = function(fox) {
        FoxRegisterService.deleteFox(fox.uuid, function() {
            initRegisterList();
        });
    };

    $scope.changeLanguage = function(key) {
        console.log("Test");
        //FoxRegisterService.changeLanguage(key);
    };

    initRegisterList();

});