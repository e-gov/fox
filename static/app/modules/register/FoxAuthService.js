/**
 * Created by mihkelk on 19.04.2016.
 */
foxApp.service("FoxAuthService", function ($http, $log, configConstant) {

    //var User = "mihkel";
    //var Token = "gAAAAABXHnIMI2xOhmLYoW29-TZ-XTQ3vRJJCRCH0hueKSnIO59NGLJWKrpiF7ilI1d1GrvhhgFgDmxomNtXVK5e2A3MBdGs4AwV0EZ_c3pz4BIoItIkKUzZSjFZlteqKOGKGISGSs3ujWLEQr_lYolEsspv7a6u0C13hN-_Z9gaD4z_oFkQjRI=";

    //
    var User;
    var Token;

    this.login = function (username, password, onSuccess, onError) {
        $http({
            method: "GET",
            cache: false,
            responseType: 'json',
            params: {
                username: username,
                challenge: password,
                provider: "provider"
            },
            url: configConstant.loginURL,
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            }
        }).then(function(result) {
            Token = result.data.token;
            User = username;
            $http.defaults.headers.common.Authorization = 'Bearer ' + Token;
            console.log("Recived token", Token);
            onSuccess();
        }, onError);
    };

    this.isAuthenticated = function () {
        return angular.isDefined(Token);
    };
    this.logout = function() {
        Token = null;
        User = null;
        $http.defaults.headers.common.Authorization = null;
    };
    this.getUser = function () {
        return User;
    }
});