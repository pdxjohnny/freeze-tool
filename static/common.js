api.Hostname = function (data) {
  // Updates the status of a single device
  hosts.set(data["Name"], data["Host"]);
}

// Map MicroSocket client names to hostnames
var hosts = new HostNames();

// When the page loads
$(function () {
  getHostnames();
})

// Request hostnames of services
function getHostnames() {
  api.send({
    "Method": "SendHostname"
  });
}
