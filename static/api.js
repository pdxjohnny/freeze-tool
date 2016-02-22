var make_api = function () {
  this.connected = false;
  this.messages = [];
  this.sender = false;
  this.ws = false;
  this.name = false;
  this.debug = false;
  // Set the default server
  this.change_server(location.origin.split("//")[1]);
  // Handlers for messages
  return this;
}

make_api.prototype.change_server = function (url) {
  this.origin = "ws://" + location.origin.split("//")[1] + "/ws";
}

make_api.prototype.connect = function () {
  this.ws = new WebSocket(this.origin);
  this.ws.onopen = this.onopen.bind(this);
  this.ws.onclose = this.onclose.bind(this);
  this.ws.onmessage = this.onmessage.bind(this);
}

make_api.prototype.onopen = function (data) {
  this.connected = true;
  this.sendall();
}

make_api.prototype.onclose = function (data) {
  this.connected = false;
}

make_api.prototype.onmessage = function (data) {
  try {
    data = JSON.parse(data["data"])
    if (this.debug === true) {
      console.log(data);
    }
  } catch (err) {
    console.error("Could not decode", data["data"]);
  }
  if (typeof data["Method"] !== "undefined" &&
    typeof this[data["Method"]] === "function") {
    this[data["Method"]](data);
  }
}

make_api.prototype.send = function (data) {
  this.messages.push(data);
  this.sendall();
}

make_api.prototype.sendall = function (data) {
  if (this.connected) {
    var numToSend = this.messages.length;
    for (var message = 0; message < numToSend; message++) {
      // Allways send the first message
      if (this.sender) {
        this.async("send", this.messages[0])
      } else {
        this.messages[0] = JSON.stringify(this.messages[0]);
        this.ws.send(this.messages[0]);
      }
      this.messages.shift();
    }
  }
}

make_api.prototype.async = function (method, data) {
  if (this.sender) {
    this.sender.postMessage([method, data]);
  }
}

make_api.prototype.startsender = function () {
  if (!this.sender) {
    this.sender = new Worker("sender.js");
    this.sender.onmessage = this.onmessage.bind(this);
  }
}

make_api.prototype.WorkerConnected = function () {
  this.connected = true;
}

make_api.prototype.MicroSocketName = function (data) {
  this.name = data["Name"];
}

api = new make_api();
api.connect()
