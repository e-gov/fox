
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

    //TODO LISA FOX ADD IMPL
    $scope.add = function (foxName) {
        if (!foxName || foxName == '') {
            return;
        }
        FoxRegisterService.addFox(foxName, function(result) {
            initRegisterList();
        });
    };

    //TODO LISA FOX UPDATE IMPL
    $scope.update = function(fox) {
        console.log(fox);
    };

    //TODO LISA FOX DELETE IMPL
    $scope.delete = function(fox) {
        console.log(fox);
    };

    initRegisterList();

});