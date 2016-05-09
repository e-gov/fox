/**
 * Created by mihkelk on 17.02.2016.
 */
foxApp.service("FoxAlertService", function () {

    this.alerts = [];
    this.infos = [];

    this.addError = function(label) {
        this.alerts.push(label);
    };

    this.addInfo = function(label) {
        this.infos.push(label);
    };

    this.clearErrors = function() {
        this.alerts = []
    };

    this.clearInfos = function() {
        this.alerts = []
    }
});
