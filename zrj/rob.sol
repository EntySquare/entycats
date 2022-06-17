// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

contract rob {

    uint public balance;

    constructor(
    ) {
       balance = 0;
    }

    function store(uint inside) external  {
      balance += inside;
    }

    function seek()  public view returns (uint balance2){
        balance2 = balance;
    }

}