api.DeviceStatus = function (data) {
  // Updates the status of a single device
  updateDeviceDiv(data);
}

api.Hostname = function (data) {
  // Updates the status of a single device
  hosts.set(data["Name"], data["Host"]);
}

api.Closed = function (data) {
  // Mark all of the devices for this host as disconnected
  hostDisconnected(data["Name"]);
}

// References to all of the divs that respresent devices
var deviceDivs = {};
// Map MicroSocket client names to hostnames
var hosts = new HostNames();

// When the page loads
$(function () {
  getDevices();
  getHostnames();
})

// Request updated device list
function getDevices() {
  api.send({
    "Method": "SendDevices"
  });
}

// Request hostnames of services
function getHostnames() {
  api.send({
    "Method": "SendHostname"
  });
}

function hostDisconnected(host) {
  // Go through all the devices
  for (var div in deviceDivs) {
    // If the divs host is the same as the host that
    // just sent the message
    if (deviceDivs[div].host === hosts.get(host)) {
      // Change the status to the disconnected message
      deviceDivs[div].status = "Disconnected";
    }
  }
}

function updateDeviceDiv(status) {
  // The div all the device divs are in
  var device_list_div = document.getElementById("devices");
  // The div respresenting the device
  var device_div = deviceDivs[status["Device"]];
  // Create a new div if it does not already exist
  if (typeof device_div === "undefined") {
    // Create a new device div
    device_div = new Device(status["Device"]);
    // Add it to the object that has all the divs
    deviceDivs[status["Device"]] = device_div;
    // Append it to the div that holds all the device divs
    device_list_div.appendChild(device_div);
  }
  // Update the div properties (host, status) based on the status object
  device_div.update(status);
}
