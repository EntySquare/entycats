// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
import "./HillStoneFinance.sol";
import "./TestCoin.sol";
import "./safemath.sol";
 
contract example2{
    HillStoneFinance public hsfToken;
    TestCoin public tcToken;
    address public maneger;
    uint public firstBlock ;
    uint public startover;
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
        address[] joiner;
        uint highPool;
        uint lowPool;
    }
    mapping(address  => IndexValue) fundermap;
    betPool  _pool ;
    uint index ;
    constructor(uint _startover) public{
       index = 0 ;
       firstBlock = block.number;
       startover = _startover;
       
    }
    function initToken  (address hsfaddress,address tcaddress) public{
        hsfToken = HillStoneFinance(hsfaddress);
        tcToken = TestCoin(tcaddress);
        maneger = address(this);
    }
    using SafeMath for uint;
    function placeBet(address fromAdress,uint _betClass,uint betAmount) external payable returns(bool success){
        if(this.contains(fromAdress)) { //新增
            if(_betClass == 1){ //高优先级
               fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number 
                   
               })); 
               _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){ //低优先级
              fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
              
              }
              hsfToken.approve(fromAdress,betAmount);
         //   return hsfToken.transferFrom(fromAdress,maneger,betAmount);
        }
        else{ //增加下注
            index ++;
            fundermap[fromAdress].keyIndex = index;
            if(_betClass == 1){ //高优先级
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
                fundermap[fromAdress].value.funderAddress = fromAdress;
                _pool.highPool = _pool.highPool.add(betAmount);
              }
            if(_betClass ==2){ //低优先级
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               fundermap[fromAdress].value.funderAddress = fromAdress;
               _pool.lowPool = _pool.lowPool.add(betAmount);
              }
               _pool.joiner.push(fromAdress);
               hsfToken.approve(fromAdress,betAmount);
               //return hsfToken.transferFrom(fromAdress,maneger,betAmount);
        }
    }
    function removeBet(address toAddress,uint _betClass,uint removeAmount) external payable returns(bool success){
        if(!this.contains(toAddress)) { // 并没有地址
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
                       ) { //循环该地址
                        if(separateBets[i].betHighAmount <= removeAmount){ //如果本次下注不足以抵扣则删除本次
                            removeAmount -= separateBets[i].betHighAmount;
                            delete separateBets[i]; 
                        }
                        else{ //如果足够以抵扣则结束
                            separateBets[i].betHighAmount -=removeAmount;
                            return success;
                        }
                   }
                    _pool.highPool = _pool.highPool.sub(removeAmount);
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
                            removeAmount -= separateBets[i].betLowAmount;
                            delete separateBets[i];
                        }
                        else{
                            separateBets[i].betLowAmount -=removeAmount;
                            return success;
                        }
                   }
                _pool.lowPool = _pool.lowPool.sub(removeAmount);
              }
              //return hsfToken.transfer(toAddress,removeAmount);
        }
    }
    function shareOut(uint bounces) external payable returns(bool success){
        uint lastBlock = block.number; // 结算时区块高度
        uint timeLength = lastBlock-firstBlock;
        address[] memory joiners= _pool.joiner; //参与者地址数列
        uint weightedHighPool; // 加权高优先级
        uint weightedLowPool; // 加权低优先级
       
        //取得加权池
        for (
            uint i = 1;
            i <= joiners.length;
            i ++
        ) {
           address joinerAddress = joiners[i];
           blockfunder[] memory separateBets = fundermap[joinerAddress].value.separateBet;
             for (
            uint j = 1;
            j <= separateBets.length;
            j ++
            ){
                weightedHighPool += separateBets[j].betHighAmount.mul(lastBlock.sub(separateBets[j].block));
                weightedLowPool += separateBets[j].betLowAmount.mul(lastBlock.sub(separateBets[j].block));
            }
        }
        //循环遍历所有funder
        for (
            uint i = 1;
            i <= joiners.length;
            i ++
        ) {
           address joinerAddress = joiners[i];
           blockfunder[] memory separateBets = fundermap[joinerAddress].value.separateBet;
           uint weightedHighBet;
           uint weightedLowBet;
            for (
            uint j = 1;
            j <= separateBets.length;
            j ++
            ){
                //当前地址的加权下注
                weightedHighBet += separateBets[j].betHighAmount.mul(timeLength-separateBets[j].block);
                weightedLowBet += separateBets[j].betLowAmount.mul(timeLength-separateBets[j].block);
            }
            if(weightedHighBet != 0 || weightedLowBet != 0){
            uint bounce = bounces.mul(weightedHighBet.div(weightedHighPool)) + bounces.mul(weightedLowBet.div(weightedLowPool));
            tcToken.transferFrom(maneger,joinerAddress,bounce);
            }
        }
        return success;
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