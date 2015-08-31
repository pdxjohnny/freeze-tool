"use strict";

class HostNames {
  constructor() {
    this.nameMap = {};
    this.changeDivs = {};
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
    for (var div in this.changeDivs) {
      this.changeDivs[div].host = this.get(this.changeDivs[div].host);
    }
  }
}
