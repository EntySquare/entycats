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
    struct bounceaddress{
        address bad;
        uint bounce;
    }
    mapping(address  => IndexValue)  fundermap;
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
    function placeBet(address fromAdress,uint _betClass,uint betAmount) external payable returns(bool){
        if(contains(fromAdress)) {  //增加下注 
            if(_betClass == 1){ //高优先级
              fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
               _pool.highPool = _pool.highPool.add(betAmount);
                betAmount += getValue(fromAdress,1);
              }
            if(_betClass ==2){ //低优先级
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
               betAmount += getValue(fromAdress,2);
              }
              
              hsfToken.approve(fromAdress,betAmount);
         //   return hsfToken.transferFrom(fromAdress,maneger,betAmount);
        }
        else{//新增
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
        if(!contains(toAddress)) { // 并没有地址
            return false;
        }
        else{
            uint count = fundermap[toAddress].value.separateBet.length;
             if(_betClass == 1){
               uint last =  getValue(toAddress,1);
               require(last >= removeAmount);
               _pool.highPool = _pool.highPool.sub(removeAmount);
                   for (
                         uint i = count-1;
                         i >= 0;
                         i --
                       ) { //循环该地址
                        if(fundermap[toAddress].value.separateBet[i].betHighAmount <= removeAmount){ //如果本次下注不足以抵扣则删除本次
                            removeAmount -= fundermap[toAddress].value.separateBet[i].betHighAmount;
                            fundermap[toAddress].value.separateBet[i].betHighAmount = 0; 
                        }
                        else{ //如果足够以抵扣则结束
                            fundermap[toAddress].value.separateBet[i].betHighAmount -=removeAmount;
                            return success;
                        }
                   }
                    
             }
            if(_betClass ==2){
              uint last =  getValue(toAddress,2);
              require(last >= removeAmount);
               _pool.lowPool = _pool.lowPool.sub(removeAmount);
              
                for (
                         uint i = count-1;
                         i >= 0;
                         i --
                       ) {
                        if(fundermap[toAddress].value.separateBet[i].betLowAmount <= removeAmount){
                            removeAmount -= fundermap[toAddress].value.separateBet[i].betLowAmount;
                            fundermap[toAddress].value.separateBet[i].betLowAmount = 0; 
                        }
                        else{
                            fundermap[toAddress].value.separateBet[i].betLowAmount -=removeAmount;
                            return success;
                        }
                   }
               
              }
              //return hsfToken.transfer(toAddress,removeAmount);
        }
    }
    function shareOut(uint bounces) external payable returns(bounceaddress[] memory){
        uint lastBlock = block.number; // 结算时区块高度
        uint timeLength = lastBlock-firstBlock;
        uint weightedHighPool; // 加权高优先级
        uint weightedLowPool; // 加权低优先级
        bounceaddress[] memory x;
        //取得加权池
        for (
            uint i = 0;
            i <= _pool.joiner.length-1;
            i ++
        ) {
           address joinerAddress = _pool.joiner[i];
          fundermap[joinerAddress].value.separateBet;
             for (
            uint j = 0;
            j <= fundermap[joinerAddress].value.separateBet.length-1;
            j ++
            ){
                weightedHighPool += fundermap[joinerAddress].value.separateBet[j].betHighAmount.mul(lastBlock.sub(fundermap[joinerAddress].value.separateBet[j].block));
                weightedLowPool += fundermap[joinerAddress].value.separateBet[j].betLowAmount.mul(lastBlock.sub(fundermap[joinerAddress].value.separateBet[j].block));
            }
        }
        //循环遍历所有funder
        for (
            uint i = 0;
            i <= _pool.joiner.length-1;
            i ++
        ) {
           address joinerAddress = _pool.joiner[i];
           uint weightedHighBet;
           uint weightedLowBet;
            for (
            uint j = 0;
            j <= fundermap[joinerAddress].value.separateBet.length-1;
            j ++
            ){
                //当前地址的加权下注
                weightedHighBet += fundermap[joinerAddress].value.separateBet[j].betHighAmount.mul(timeLength-fundermap[joinerAddress].value.separateBet[j].block);
                weightedLowBet += fundermap[joinerAddress].value.separateBet[j].betLowAmount.mul(timeLength-fundermap[joinerAddress].value.separateBet[j].block);
            }
            if(weightedHighBet != 0 || weightedLowBet != 0){
            uint _bounce = bounces.mul(weightedHighBet.div(weightedHighPool)) + bounces.mul(weightedLowBet.div(weightedLowPool));
            x[i] =bounceaddress({
                   bad : joinerAddress,
                   bounce : _bounce
               }); 
            }
        }
          return x;
    }
    function getHighLevelAmount() public  view returns (uint) {
            return _pool.highPool;
    }
    function getLowLevelAmount() public  view returns (uint) {
            return _pool.lowPool;
    }
    function contains(address adkey) internal  returns (bool) {
            return fundermap[adkey].keyIndex > 0;
    }
    function getMap(address adkey) internal   returns (uint keyIndex){
        keyIndex = fundermap[adkey].keyIndex;
        
    }
    function getValue(address adkey ,uint class) internal  returns (uint summary) {
        uint count = fundermap[adkey].value.separateBet.length;
        if(class == 1){
               for (
            uint i = 0;
            i <= count-1;
            i ++ ) {
              summary += fundermap[adkey].value.separateBet[i].betHighAmount;
            }
        }
        if(class == 2){
                for (
            uint i = 0;
            i <= count-1;
            i ++ ) {
              summary += fundermap[adkey].value.separateBet[i].betLowAmount;
            }
        }
        }
}