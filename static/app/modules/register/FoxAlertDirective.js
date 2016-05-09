/**
 * Created by mihkelk on 17.02.2016.
 */
foxApp.directive("foxAlert", function (FoxAlertService) {
    return {
        templateUrl: 'views/alert-template.html',
        link: function(scope) {
            scope.$watch(function() {
                return FoxAlertService.alerts;
            }, function(newVal) {
                scope.alerts = newVal;
            }, true);

            scope.$watch(function() {
                return FoxAlertService.infos;
            }, function(newVal) {
                scope.infos = newVal;
            }, true);
        }
    }
});