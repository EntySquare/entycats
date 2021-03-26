// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
import "./HillStoneFinance.sol";
import "./safemath.sol";
 
contract example2{
    HillStoneFinance public hsfToken;
    address public maneger;
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
    uint index ;
     function contractexample2() public{
       hsfToken = HillStoneFinance(0xDA0bab807633f07f013f94DD0E6A4F96F8742B53); //实例化一个token
       index = 0 ;
       maneger = 0xDA0bab807633f07f013f94DD0E6A4F96F8742B53;
    }
    using SafeMath for uint;
    function placeBet(address fromAdress,uint _betClass,uint betAmount) public payable returns(bool success){
        if(this.contains(fromAdress)) {
            if(_betClass == 1){
               fundermap[fromAdress].value.betHighAmount += betAmount;
               _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){
               fundermap[fromAdress].value.betLowAmount += betAmount;
               _pool.lowPool = _pool.lowPool.add(betAmount);
              }
              return hsfToken.transferFrom(fromAdress,maneger,betAmount);
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
                _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){
                fundermap[fromAdress].value = funder({
                   funderAddress : fromAdress,
                   betHighAmount : 0,
                   betLowAmount : betAmount
               }); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
              }
             return hsfToken.transferFrom(fromAdress,maneger,betAmount);
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
               _pool.highPool = _pool.highPool.sub(removeAmount);
              }
            if(_betClass ==2){
               uint last =  this.getValue(toAddress,2);
               require(last >= removeAmount);
               fundermap[toAddress].value.betLowAmount -= removeAmount;
                _pool.lowPool = _pool.lowPool.sub(removeAmount);
              }
              return hsfToken.transferFrom(maneger,toAddress,removeAmount);
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