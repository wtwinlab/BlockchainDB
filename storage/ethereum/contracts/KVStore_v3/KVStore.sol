// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

contract KVStore {

    event ItemSet(string key, string value);
    mapping (string => string) public items;
    
    function setItem(string memory key, string memory value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }
}