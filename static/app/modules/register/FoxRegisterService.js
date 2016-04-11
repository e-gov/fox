/**
 * Created by mihkelk on 17.02.2016.
 */
foxApp.service("FoxRegisterService", function ($http, configConstant, $translate) {

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
            headers: {'Content-Type': 'application/json; charset=utf-8'},
            url: configConstant.backendURL + "/fox/foxes"
        }).then(onSuccess, onError);
    };

    this.deleteFox = function (uuid, onSuccess, onError) {
        $http({
            method: "DELETE",
            cache: false,
            responseType: 'json',
            headers: {'Content-Type': 'application/json; charset=utf-8'},
            url: configConstant.backendURL + "/fox/foxes/" + uuid + "/delete"
        }).then(onSuccess, onError);
    };

    this.changeLanguage = function(key) {
        $translate.use(key);
        localStorage.language = key;
    };

});