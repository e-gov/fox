/**
 * Created by mihkelk on 17.02.2016.
 */
foxApp.service("FoxAlertService", function () {

    this.alerts = [];

    this.addError = function(label) {
        this.alerts.push(label);
    };

});
