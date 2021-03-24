// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
 
contract example2{
    
    struct IndexValue { 
        uint keyIndex; 
        funder value; 
    }
    struct funder {
        address funderAddress ;
        uint betHighAmount;
        uint betLowAmount;
    }
    struct betPool {
        uint highPool;
        uint lowPool;
    }
    mapping(address  => IndexValue) fundermap;
    betPool  _pool ;
    uint index = 0;
    function placeBet(address fromAdress,uint _betClass,uint betAmount) public payable returns(bool success){
        if(this.contains(fromAdress)) {
            if(_betClass == 1){
               fundermap[fromAdress].value.betHighAmount += betAmount;
               _pool.highPool += betAmount;
              }
            if(_betClass ==2){
               fundermap[fromAdress].value.betLowAmount += betAmount;
               _pool.lowPool +=betAmount;
              }
              return true;
        }
        else{
            index ++;
            fundermap[fromAdress].keyIndex = index;
            if(_betClass == 1){
                fundermap[fromAdress].value = funder({
                   funderAddress : fromAdress,
                   betHighAmount : betAmount,
                   betLowAmount : 0
               }); 
               _pool.highPool += betAmount;
              }
            if(_betClass ==2){
                fundermap[fromAdress].value = funder({
                   funderAddress : fromAdress,
                   betHighAmount : 0,
                   betLowAmount : betAmount
               }); 
               _pool.lowPool +=betAmount;
              }
             return true;
        }
    }
    function removeBet(address toAddress,uint _betClass,uint removeAmount) public payable returns(bool success){
        if(!this.contains(toAddress)) {
            return false;
        }
        else{
             if(_betClass == 1){
               uint last =  this.getValue(toAddress,1);
               require(last >= removeAmount);
               fundermap[toAddress].value.betHighAmount -= removeAmount;
               _pool.highPool -= removeAmount;
              }
            if(_betClass ==2){
               uint last =  this.getValue(toAddress,2);
               require(last >= removeAmount);
               fundermap[toAddress].value.betLowAmount -= removeAmount;
               _pool.lowPool -= removeAmount;
              }
              return true;
        }
    }
    function getHighLevelAmount() public  view returns (uint) {
            return _pool.highPool;
    }
    function getLowLevelAmount() public  view returns (uint) {
            return _pool.lowPool;
    }
    function contains(address adkey) public view returns (bool) {
            return fundermap[adkey].keyIndex > 0;
    }
    function getValue(address adkey ,uint class) public view returns (uint last) {
            if(class == 1){
                return fundermap[adkey].value.betHighAmount;
            }
            if(class == 2){
                return fundermap[adkey].value.betLowAmount;
            }
        }
}