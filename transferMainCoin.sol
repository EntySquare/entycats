// Copyright 2021 EntySquare Software Studio
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

// @title transferMainCoin
// @author yueliyangzi
contract mainCoinTransfer{
event Received(address, uint);
  //一对一主网币转账事件
event Transfer(
    address indexed _from,
    uint indexed _value,
    address _to,
    uint _amount
  );

  //一对一主网币转账方法
  function transfer(
    address  _addresses,
    uint256 _amounts
  ) payable public{
    address payable xp = payable(_addresses);
    xp.transfer(_amounts);
    emit Transfer(msg.sender, msg.value, _addresses, _amounts);   
  }
receive() external payable {
           emit Received(msg.sender, msg.value);
   }
}