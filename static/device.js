var Device = function Device(data) {
  var div = document.createElement("div");
  var column_one = document.createElement("div");
  var column_two = document.createElement("div");
  div.className = "row";
  column_one.className = "col-md-8";
  column_two.className = "col-md-4";
  div.appendChild(column_one);
  div.appendChild(column_two);
  div.device_div = column_one;
  div.status_div = column_two;

  for(var i in this) {
    if(typeof this[i] === "function") {
      div[i] = this[i].bind(div);
    } else {
      div[i] = this[i];
    }
  }

  Object.defineProperty(div, "device", {
    set: function (name) {
      this.device_name = name;
      this.device_div.innerHTML = name;
      return this.device_div;
    },
    get: function (name) {
      return this.device_name;
    }
  });

  Object.defineProperty(div, "status", {
    set: function (name) {
      this.status_name = name;
      this.status_div.innerHTML = name;
      return this.status_div;
    },
    get: function (name) {
      return this.status_name;
    }
  });

  return div;
};

Device.prototype.open = function open() {
  alert(this.getAttribute("Id"));
};
