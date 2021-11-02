// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract KVStore {
    mapping(bytes => bytes) public data;

    function get(bytes memory key) public view returns (bytes memory val) {
        return data[key];
    }

    // write method in contract
    function set(bytes memory key, bytes memory val)
        public
        returns (bool success)
    {
        data[key] = val;
        return true;
    }
}
