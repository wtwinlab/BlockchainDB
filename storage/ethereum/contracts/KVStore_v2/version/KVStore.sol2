// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

contract KVStore {

    mapping(address => mapping(string => string)) store;
    uint MAX_STR_LENGTH = 1000;

    function get(address _acct, string memory _key) public view returns(string memory) {
        return store[_acct][_key];
    }
    function set(string memory _key, string memory _value) public {
        require(bytes(_value).length <= MAX_STR_LENGTH);
        store[msg.sender][_key] = _value;
    }
}
