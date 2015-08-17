var device_list_div = document.getElementById("devices");

api.Devices = function (devices) {
  var deviceList = devices["Devices"];
  updateDeviceList(deviceList);
}

$(function() {
  api.send({"Method": "SendDevices"});
})

function updateDeviceList(deviceList) {
  for (var device_id in deviceList) {
    var name = deviceList[device_id];
    var device_div = document.getElementById(name);
    if (device_div == null) {
      device_div = new Device(name);
      device_list_div.appendChild(device_div);
    }
    device_div.device = name;
  }
}
