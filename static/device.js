var Device = function Device(name) {
  var div = document.createElement("div");
  var column_one = document.createElement("div");
  var column_two = document.createElement("div");
  var column_three = document.createElement("div");
  div.id = name;
  div.className = "row";
  column_one.className = "col-md-4";
  column_two.className = "col-md-4";
  column_three.className = "col-md-4";
  div.appendChild(column_one);
  div.appendChild(column_two);
  div.appendChild(column_three);
  div.device_div = column_one;
  div.status_div = column_two;
  div.host_div = column_three;

  for (var i in this) {
    if (typeof this[i] === "function") {
      div[i] = this[i].bind(div);
    } else {
      div[i] = this[i];
    }
  }

  Object.defineProperty(div, "device", {
    set: function (name) {
      this.device_name = name;
      this.device_div.innerHTML = this.historyLink(name);
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

  Object.defineProperty(div, "host", {
    set: function (name) {
      this.host_name = name;
      this.host_div.innerHTML = name;
      return this.host_div;
    },
    get: function (name) {
      return this.host_name;
    }
  });

  // Set the name and status
  div.device = name;
  div.status = "OK";
  return div;
};

Device.prototype.update = function (object) {
  for (var prop in object) {
    if (prop.toLowerCase() === "host") {
      object[prop] = hosts.get(object[prop]);
    }
    this[prop.toLowerCase()] = object[prop];
  }
}

Device.prototype.historyLink = function (name) {
  return "<a href=\"history.html#" + name + "\">" + name + "</a>";
}
