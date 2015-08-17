var Albom = function Albom(data) {
  var div = document.createElement("div");

  for(var i in this) {
    if(typeof this[i] === "function") {
      div[i] = this[i].bind(div);
    } else {
      div[i] = this[i];
    }
  }

  return div;
};

Albom.prototype.open = function open() {
  alert(this.getAttribute("data-id"));
};
