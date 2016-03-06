/**
 * Created by mihkelk on 17.02.2016.
 */

foxApp.service("FoxRegisterService", function ($http, configConstant) {

    this.getAll = function (onSuccess, onError) {
        $http({
            method: "GET",
            cache: false,
            responseType: 'json',
            url: configConstant.backendURL + "/fox/foxes",
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            }
        }).then(onSuccess, onError);
    };

    this.addFox = function (foxName, onSuccess, onError) {
        $http({
            method: "POST",
            cache: false,
            responseType: 'json',
            data: {name: foxName},
            headers: {'Content-Type': 'application/x-www-form-urlencoded'},
            url: configConstant.backendURL + "/fox/foxes"
        }).then(onSuccess, onError);
    };

    this.deleteFox = function (foxId, onSuccess, onError) {
        $http({
            method: "DELETE",
            cache: false,
            responseType: 'json',
            data: {foxId: foxId},
            headers: {'Content-Type': 'application/x-www-form-urlencoded'},
            url: configConstant.backendURL + "/fox/foxes/" + foxId
        }).then(onSuccess, onError);
    };

});