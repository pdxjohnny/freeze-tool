api.DeviceStatus = function (data) {
  // Updates the status of a single device
  updateDeviceDiv(data);
}

api.Closed = function (data) {
  // Mark all of the devices for this host as disconnected
  hostDisconnected(data["Name"]);
}

// References to all of the divs that respresent devices
// are stored in hosts.changeDivs

// When the page loads
$(function () {
  getDevices();
})

// Request updated device list
function getDevices() {
  api.send({
    "Method": "SendDevices"
  });
}

function hostDisconnected(host) {
  // Go through all the devices
  for (var div in hosts.changeDivs) {
    // If the divs host is the same as the host that
    // just sent the message
    if (hosts.changeDivs[div].host === hosts.get(host)) {
      // Change the status to the disconnected message
      hosts.changeDivs[div].status = "Disconnected";
    }
  }
}

function updateDeviceDiv(status) {
  // The div all the device divs are in
  var device_list_div = document.getElementById("devices");
  // The div respresenting the device
  var device_div = hosts.changeDivs[status["Device"]];
  // Create a new div if it does not already exist
  if (typeof device_div === "undefined") {
    // Create a new device div
    device_div = new Device(status["Device"]);
    // Add it to the object that has all the divs
    hosts.changeDivs[status["Device"]] = device_div;
    // Append it to the div that holds all the device divs
    device_list_div.appendChild(device_div);
  }
  // Update the div properties (host, status) based on the status object
  device_div.update(status);
}
