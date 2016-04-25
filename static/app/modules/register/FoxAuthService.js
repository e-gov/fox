/**
 * Created by mihkelk on 19.04.2016.
 */
foxApp.service("FoxAuthService", function ($http, $log, configConstant) {

    var Token;
    var User;

    this.login = function (username, password, onSuccess, onError) {
        $http({
            method: "GET",
            cache: false,
            responseType: 'json',
            url: configConstant.loginURL,
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            }
        }).then(function(result) {
            Token = result.token;
            User = username;
            onSuccess();
        }, onError);
    };

    this.isAuthenticated = function () {
        return angular.isDefined(Token);
    };

    this.getUser = function () {
        return User;
    }
});
