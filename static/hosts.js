"use strict";

class HostNames {
  constructor() {
    this.nameMap = {};
    return this;
  }
  get(name) {
    if (typeof this.nameMap[name] !== "undefined") {
      return this.nameMap[name];
    }
    return name;
  }
  set(name, host) {
    this.nameMap[name] = host;
    this.updateDeviceDivs();
  }
  updateDeviceDivs() {
    // Go through all the devices
    for (var div in deviceDivs) {
      deviceDivs[div].host = this.get(deviceDivs[div].host);
    }
  }
}
