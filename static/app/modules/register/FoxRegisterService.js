/**
 * Created by mihkelk on 17.02.2016.
 */
foxApp.service("FoxRegisterService", function ($http, $log, configConstant, $translate) {

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

    this.getFox = function (uuid, onSuccess, onError) {
        $http({
            method: "GET",
            cache: false,
            responseType: 'json',
            url: configConstant.backendURL + "/fox/foxes/" + uuid,
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            }
        }).then(function (result) {
            if (!result.data.parents) {
                result.data.parents = [];
            }
            return result;
        }).then(onSuccess, onError);
    };

    this.addFox = function (fox, onSuccess, onError) {
        $http({
            method: "POST",
            cache: false,
            responseType: 'json',
            data: fox,
            headers: { 'Content-Type': 'application/json; charset=utf-8' },
            url: configConstant.backendURL + "/fox/foxes"
        }).then(onSuccess, onError);
    };

    this.updateFox = function (fox, onSuccess, onError) {
        $http({
            method: "PUT",
            cache: false,
            responseType: 'json',
            data: fox,
            headers: { 'Content-Type': 'application/json; charset=utf-8' },
            url: configConstant.backendURL + "/fox/foxes/" + fox.uuid
        }).then(onSuccess, onError);
    };

    this.deleteFox = function (uuid, onSuccess, onError) {
        $http({
            method: "DELETE",
            cache: false,
            responseType: 'json',
            headers: { 'Content-Type': 'application/json; charset=utf-8' },
            url: configConstant.backendURL + "/fox/foxes/" + uuid
        }).then(onSuccess, onError);
    };

    this.changeLanguage = function(key) {
        $translate.use(key);
        localStorage.language = key;
    };

});
