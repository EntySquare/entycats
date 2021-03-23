// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
 

import "./iterableFunderMapping.sol";
//import "./brilliant.sol";
    // uint betPoolAmount;
    // mapping(uint => betPool) betPoolmap;
    // //  iii.生成质押池
    // function createNewBetPool(address _betPoolAddress){
    //     betPoolAmount++;
    //     betPoolmap[betPoolAmount] = betPool(_betPoolAddress,0,0,0);
    // }
    //  iv.实现质押功能的方法
    //   一.生成新的质押人
    //   二.改变目标质押池中的变量
    //   三.highLevel 1 lowLevel 2
contract example{
    IterableFunderMapping.itFundermap data;
    struct betPool{
          uint amount;
      }
     betPool[] betPools;
     address public manager;         //管理员地址
        //构造函数  指定管理员地址为合约部署者的地址
    function contractManager() public{
        manager =  msg.sender;
    }
    function placeBet(uint _betClass) public payable returns(bool success){
        //betPool storage _betPool = betPoolmap[_betPoolAmount];
        if (IterableFunderMapping.contains(data,msg.sender)){
            if (_betClass == 1) {
               return IterableFunderMapping.update(data,msg.sender,1,msg.value);
               betPools[1].amount +=  msg.value;
                
            }
            if (_betClass == 2) {
               return IterableFunderMapping.update(data,msg.sender,3,msg.value);
               betPools[2].amount +=  msg.value ;
            }
        }
        else{
             if (_betClass == 1) {
               return IterableFunderMapping.insert(data,msg.sender,IterableFunderMapping.funder({funderAddress:msg.sender,
               betHighAmount:msg.value,
               betLowAmount:0
               }));
               betPools[1].amount +=  msg.value ; 
            }
            if (_betClass == 2) {
               return IterableFunderMapping.insert(data,msg.sender,IterableFunderMapping.funder({funderAddress:msg.sender,
               betHighAmount:0,
               betLowAmount:msg.value
               }));
               betPools[2].amount +=  msg.value ;
            }
        }
           
    }
    // 2.解除质押
    function removeBet(uint _betClass,uint removeAmount) public payable returns(bool success){
         if(IterableFunderMapping.contains(data,msg.sender)){
             if (_betClass == 1) {
                   uint last =  IterableFunderMapping.getValue(data,msg.sender,1);
                   require(last >= removeAmount);
                   betPools[1].amount -=  msg.value ;
                   return IterableFunderMapping.update(data,msg.sender,2,removeAmount);
                }
             if (_betClass == 2) {
                   uint last =  IterableFunderMapping.getValue(data,msg.sender,1);
                   require(last >= removeAmount);
                    betPools[2].amount -=  msg.value ;
                   return IterableFunderMapping.update(data,msg.sender,4,removeAmount);
                }
         }
         else {
             return false;
         }
    }
    //3.查询全部资金池
    function getHighLevelAmount() public Manager view returns (uint) {
        uint s = betPools[1].amount;
        // for (
        //     uint i = IterableFunderMapping.iterate_start(data);
        //     IterableFunderMapping.iterate_valid(data,i);
        //     i = IterableFunderMapping.iterate_next(data,i)
        // ) {
        //      (,IterableFunderMapping.funder memory value) = IterableFunderMapping.iterate_get(data,i);
        //       s += value.betHighAmount;
        // }
        return s;
    }
    function getLowLevelAmount() public Manager view returns (uint) {
          uint s;
          for (
            uint i = IterableFunderMapping.iterate_start(data);
            IterableFunderMapping.iterate_valid(data,i);
            i = IterableFunderMapping.iterate_next(data,i)
        ) {
            (, IterableFunderMapping.funder storage value) = IterableFunderMapping.iterate_get(data,i);
             s += value.betLowAmount;
        }
        return s;
    }
     //只有管理员才能进行操作  关键字段modifier 
    modifier Manager(){
        require(msg.sender == manager);
        _;
    }
    
}
