api.Devices = function (devices) {
  updateDeviceList(devices["Devices"], devices["Host"]);
}

$(function() {
  api.send({"Method": "SendDevices"});
})

function updateDeviceList(deviceList, host) {
  var device_list_div = hostDevicesDiv(host);
  for (var device_id in deviceList) {
    var name = deviceList[device_id];
    var device_div = document.getElementById(name);
    if (device_div == null) {
      device_div = new Device(name);
      device_list_div.appendChild(device_div);
    }
    device_div.device = name;
    device_div.status = "OK";
  }
  var deviceInList = $(device_list_div).find("div.row").toArray();
  for (var div in deviceInList) {
    var div = deviceInList[div];
    if (-1 == $.inArray(div.id, deviceList)) {
      div.status = "Offline";
    }
  }
}

function hostDevicesDiv(host) {
  if (typeof host === "undefined") {
    host = location.origin;
  }
  var host_device_list_div = document.getElementById(host + "devices");
  if (host_device_list_div == null) {
    var all_hosts = document.getElementById("hosts");
    host_device_list_div = document.createElement("div");
    host_device_list_div.id = host + "devices";
    host_device_list_div.className = "row";
    all_hosts.appendChild(host_device_list_div);
  }
  return host_device_list_div;
}
