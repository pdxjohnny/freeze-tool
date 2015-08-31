var Status = function Status(name) {
  var div = document.createElement("div");
  var column_one = document.createElement("div");
  var column_two = document.createElement("div");
  div.id = name;
  div.className = "row";
  column_one.className = "col-md-4";
  column_two.className = "col-md-4";
  div.appendChild(column_one);
  div.appendChild(column_two);
  div.status_div = column_one;
  div.host_div = column_two;

  for (var i in this) {
    if (typeof this[i] === "function") {
      div[i] = this[i].bind(div);
    } else {
      div[i] = this[i];
    }
  }

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
  div.host = name;
  div.status = "OK";
  return div;
};

Status.prototype.update = function (object) {
  for (var prop in object) {
    if (prop.toLowerCase() === "host") {
      object[prop] = hosts.get(object[prop]);
    }
    this[prop.toLowerCase()] = object[prop];
  }
}
