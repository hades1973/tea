function List() {
    var Node = function (element) {
        this.element = element;
        this.next = null;
    };
    var length = 0;
    var head = null;

    this.appendOrRemove = function (element) { // if element not in list, append; otherwise remove it
        var newnode = new Node(element);
        if (!head) {
            head = newnode;
            return;
        }
        var curNode = head;
        var prevNode = head;
        while (curNode) {
            if (curNode.element == element) {
                if (curNode == head) {
                    head = curNode.next;
                    return;
                } else {
                    prevNode.next = curNode.next;
                    return;
                }
            }
            prevNode = curNode;
            curNode = curNode.next;
        }
        prevNode.next = newnode;
    };

    this.data = function () {
        var ar = [];
        for (var node = head; node; node = node.next) {
            ar.push(node.element);
        }
        return ar;
    };
};