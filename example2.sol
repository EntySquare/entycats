// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
import "./HillStoneFinance.sol";
import "./safemath.sol";
 
contract example2{
    HillStoneFinance public hsfToken;
    address public maneger;
    uint public firstBlock ;
    struct IndexValue { 
        uint keyIndex; 
        funder value; 
    }
    struct blockfunder{
        uint betHighAmount;
        uint betLowAmount;
        uint block;
    }
    struct funder {
        address funderAddress ;
        blockfunder[] separateBet;
      
    }
    struct betPool {
        address[] joiners;
        uint highPool;
        uint lowPool;
    }
    mapping(address  => IndexValue) fundermap;
    betPool  _pool ;
    uint index ;
     function contractexample2() public{
       hsfToken = HillStoneFinance(0xaE036c65C649172b43ef7156b009c6221B596B8b); //实例化一个token
       index = 0 ;
       maneger = 0xaE036c65C649172b43ef7156b009c6221B596B8b;
       firstBlock = block.number;
    }
    using SafeMath for uint;
    function placeBet(address fromAdress,uint _betClass,uint betAmount) public payable returns(bool success){
        if(this.contains(fromAdress)) {
            if(_betClass == 1){
               fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number 
                   
               })); 
               _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){
              fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
              
              }
              return hsfToken.transferFrom(fromAdress,maneger,betAmount);
        }
        else{
            index ++;
            fundermap[fromAdress].keyIndex = index;
            if(_betClass == 1){
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
                fundermap[fromAdress].value.funderAddress = fromAdress;
                _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               fundermap[fromAdress].value.funderAddress = fromAdress;
               _pool.lowPool = _pool.lowPool.add(betAmount);
              }
               _pool.joiners.push(fromAdress);
             return hsfToken.transferFrom(fromAdress,maneger,betAmount);
        }
    }
    function removeBet(address toAddress,uint _betClass,uint removeAmount) public payable returns(bool success){
        if(!this.contains(toAddress)) {
            return false;
        }
        else{
            blockfunder[] storage separateBets = fundermap[toAddress].value.separateBet;
            uint count = separateBets.length;
             if(_betClass == 1){
               uint last =  this.getValue(toAddress,1);
               require(last >= removeAmount);
                   for (
                         uint i = count;
                         i >= 1;
                         i --
                       ) {
                        if(separateBets[i].betHighAmount <= removeAmount){
                            delete separateBets[i];
                            removeAmount -= separateBets[i].betHighAmount;
                        }
                        else{
                            separateBets[i].betHighAmount -=removeAmount;
                        }
                   }
             }
            if(_betClass ==2){
              uint last =  this.getValue(toAddress,2);
              require(last >= removeAmount);
                for (
                         uint i = count;
                         i >= 1;
                         i --
                       ) {
                        if(separateBets[i].betLowAmount <= removeAmount){
                            delete separateBets[i];
                            removeAmount -= separateBets[i].betLowAmount;
                        }
                        else{
                            separateBets[i].betLowAmount -=removeAmount;
                        }
                   }
              }
              return hsfToken.transferFrom(maneger,toAddress,removeAmount);
        }
    }
    function shareOut() public payable returns(bool success){
        uint count = _pool.joiners.length;
        uint lastBlock = block.number;
        uint timeLength = lastBlock-firstBlock;
        for (
            uint i = 1;
            i <= count;
            i ++
        ) {
        //   address _address = _pool.joiners[i];
        // //   if(fundermap[_address].betHighAmount != 0 || fundermap[_address].betLowAmount != 0){
               
        // //   }
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
            uint summary;
            blockfunder[] memory separateBets = fundermap[adkey].value.separateBet;
            uint count = separateBets.length;
        if(class == 1){
               for (
            uint i = 1;
            i <= count;
            i ++ ) {
              summary += separateBets[i].betHighAmount;
            }
        }
        if(class == 2){
                for (
            uint i = 1;
            i <= count;
            i ++ ) {
              summary += separateBets[i].betLowAmount;
            }
        }
            return summary;
        }
}