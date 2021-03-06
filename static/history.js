api.DeviceHistory = function (data) {
  if (DumpKey === data["DumpKey"]) {
    // Adds a status to the history of the device
    addStatus(data);
  }
}

api.DeviceStatus = function (data) {
  // In case the host restarts
  data["Host"] = hosts.get(data["Host"]);
  // SO compare to the old
  var stringified = JSON.stringify(data);
  if (deviceName === data["Device"] &&
    stringified !== lastStatus) {
    // Update the lastStatus
    lastStatus = stringified;
    // Adds a status to the history of the device
    addStatus(data);
  }
}

api.SendDeviceHistoryConfirm = function (data) {
  if (DumpAccepted == false) {
    DumpAccepted = true;
    api.send({
      "Method": "SendDeviceHistoryConfirmed",
      "ClientId": data["ClientId"],
      "DumpKey": data["DumpKey"]
    });
  }
}

// The device were getting the history of
var deviceName = location.hash.slice(1);
// So we know what history dump to accept
var DumpKey = false;
// So we no only to accept one history dump
var DumpAccepted = false;
// The last status sent so there are no duplicates
var lastStatus = "";

// When the page loads
$(function () {
  var deviceNameDiv = document.getElementById("deviceName");
  deviceNameDiv.innerHTML = deviceName;
  getDeviceHistory();
})

// Request updated device list
function getDeviceHistory() {
  DumpKey = String(Math.random());
  api.send({
    "Method": "SendDeviceHistory",
    "DumpKey": DumpKey,
    "Device": deviceName
  });
}

function addStatus(status) {
  // The div all the status divs are in
  var history_list = document.getElementById("history");
  // The div respresenting the status
  var status_div = new Status(status["Device"]);
  // Make sure hosts can change the hostname if need
  hosts.changeDivs[Object.keys(hosts.changeDivs).length] = status_div;
  // Append it to the div that holds all the status divs
  history_list.appendChild(status_div);
  // Update the div properties (host, status) based on the status object
  status_div.update(status);
}
