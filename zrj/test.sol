// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

contract simpletransfer {

    uint public balance;
    string public Nickname;
    address payable Address;
    mapping(address => uint256) public balances;

    constructor(
    ) {
       balance = 50;
    }

    event Sent(address from, address to, uint amount);

    function transfer(string sendername, address _to, int256 _value) payable public returns  (bool success) {
        require( balances[msg.sender] >= _value && balances[msg.sender] - _value >= 0 );
        Nickname = sendername;
        balances[_to] += _value;//接收账户增加token数量_value
        balances[msg.sender] -= _value; //支出账户_from减去token数量_value
        emit Sent(msg.sender, _to, _value);
        return true;
    }

    function observe(address inputaddr)  public view returns (uint accountbalance){
        accountbalance = balances[inputaddr];
    }


}