// SPDX-License-Identifier: MIT

const EmployeeVest = artifacts.require("EmployeeVest");
const GrouceCredit = artifacts.require("GrouceCredit");

module.exports = function (deployer) {
  deployer.deploy(EmployeeVest,10000000,0);
  deployer.deploy(GrouceCredit);
};

