function Cont() {
    var qtypes = [];
    var kvs = new Map();
    var currentQtype = "blank";

    this.setCurrentQtype = function (qtype) {
        currentQtype = qtype;
    };

    this.appendQtype = function (name) {
        qtypes.push(name);
    };

    this.getQtypes = function () {
        return qtypes;
    };

    this.getQuestions = function (qtype) {
        return kvs.get(qtype).data();
    };

    this.addQuestion = function (number) {
        var qS = kvs.get(currentQtype);
        qS.appendOrRemove(number);
    };

    this.initkvs = function () {
        for (var i = 0; i < qtypes.length; i++) {
            kvs.set(qtypes[i], new List());
        }
    };
}