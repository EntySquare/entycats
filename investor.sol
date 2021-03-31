// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
import "./HillStoneFinance.sol";
import "./USDT.sol";
import "./safemath.sol";
 
contract investor{
    HillStoneFinance public hsfToken;
    USDT public usToken;
    address public maneger;
    uint public firstBlock ;
    uint public startover;
    address public institution;
    bool public institutionflag;
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
    constructor(uint _startover,bool institutionflag,address _institution) public{
       index = 0 ;
       firstBlock = block.number;
       startover = _startover;
       if(institutionflag){
       institution = _institution;
       }
       
    }
    function initToken  (address hsfaddress,address tcaddress) public{
        hsfToken = HillStoneFinance(hsfaddress);
        usToken = USDT(tcaddress);
        maneger = address(this);
    }
    using SafeMath for uint;
    function placeBet(address fromAdress,uint _betClass,uint betAmount) external payable returns(bool){
        if(contains(fromAdress)) {  //增加下注 
            uint needapprove ;
            if(_betClass == 1){ //高优先级
              fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
               _pool.highPool = _pool.highPool.add(betAmount);
               needapprove =  getAllValue(fromAdress);
              }
            if(_betClass ==2){ //低优先级
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
               needapprove =  getAllValue(fromAdress);
              }
              hsfToken.approve(fromAdress,needapprove);
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
            uint surplus = removeAmount;
             if(_betClass == 1){
               uint last =  getValue(toAddress,1);
               require(last >= removeAmount);
               _pool.highPool = _pool.highPool.sub(removeAmount);
                   for (
                         uint i = count-1;
                         i >= 0;
                         i --
                       ) { //循环该地址
                        if(fundermap[toAddress].value.separateBet[i].betHighAmount < surplus){ //如果本次下注不足以抵扣则删除本次
                            surplus -= fundermap[toAddress].value.separateBet[i].betHighAmount;
                            fundermap[toAddress].value.separateBet[i].betHighAmount = 0;
                        }
                        else{ //如果足够以抵扣则结束
                           fundermap[toAddress].value.separateBet[i].betHighAmount -=surplus; 
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
                        if(fundermap[toAddress].value.separateBet[i].betLowAmount < surplus){
                            surplus -= fundermap[toAddress].value.separateBet[i].betLowAmount;
                            fundermap[toAddress].value.separateBet[i].betLowAmount = 0; 
                        }
                        else{
                            fundermap[toAddress].value.separateBet[i].betLowAmount -=surplus;
                            return success;
                        }
                   }
               
              }
              //return hsfToken.transfer(toAddress,removeAmount);
        }
    }
    function shareOut(uint bounces) external payable{
         if(institutionflag){
            usToken.transfer(institution,uint(bounces/10));
        }
        uint timeLength = block.number-firstBlock;
        uint weightedHighPool = _pool.highPool * timeLength * 10; // 加权高优先级
        uint weightedLowPool = _pool.lowPool * timeLength * 10; // 加权低优先级
//         //取得加权池
// //        bounceaddress[] memory _bounceaddress = new bounceaddress[](uint256(_pool.joiner.length));
//         for (
//             uint i = 0;
//             i <= _pool.joiner.length-1;
//             i ++
//         ) {
//           address joinerAddress = _pool.joiner[i];
//           blockfunder[] memory separateBets = fundermap[joinerAddress].value.separateBet;
//              for (
//             uint j = 0;
//             j <= separateBets.length-1;
//             j ++
//             ){
//                 weightedHighPool += separateBets[j].betHighAmount.mul(timeLength * 10);
//                 weightedLowPool += separateBets[j].betLowAmount.mul(timeLength * 10);
//             }
//         }
        //循环遍历所有funder
        for (
            uint i = 0;
            i <= _pool.joiner.length-1;
            i ++
        ) {
          address joinerAddress2 = _pool.joiner[i];
          blockfunder[] memory separateBets2 = fundermap[joinerAddress2].value.separateBet;
          uint weightedHighBet;
          uint weightedLowBet;
            for (
            uint j = 0;
            j <= separateBets2.length-1;
            j ++
            ){
                //当前地址的加权下注
                weightedHighBet += separateBets2[j].betHighAmount.mul(block.number-separateBets2[j].block);
                weightedLowBet += separateBets2[j].betLowAmount.mul(block.number-separateBets2[j].block);
            }
            if(weightedHighBet != 0 || weightedLowBet != 0){
                if(weightedHighPool == 0){
                    weightedHighPool = 1;
                }
                if(weightedLowPool == 0){
                    weightedLowPool = 1;
                }
                    weightedHighBet = weightedHighBet * bounces * 3;
                if(institutionflag){
                     weightedLowBet = weightedLowBet * bounces * 6;
                }
                else{
                    weightedLowBet = weightedLowBet * bounces * 7;
                }
    //             _bounceaddress[i].bounce =  uint(weightedHighBet/weightedHighPool + weightedLowBet/weightedLowPool);
    //             _bounceaddress[i].bad =  joinerAddress2;
                usToken.transfer(joinerAddress2,uint(weightedHighBet/weightedHighPool + weightedLowBet/weightedLowPool));
            }
        }
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
    function getMap(address adkey) internal  returns (uint keyIndex){
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
    function getAllValue(address adkey) internal  returns (uint summary) {
            uint count = fundermap[adkey].value.separateBet.length;
            for (
                uint i = 0;
                i <= count-1;
                i ++ ) {
                  summary += fundermap[adkey].value.separateBet[i].betHighAmount;
                  summary += fundermap[adkey].value.separateBet[i].betLowAmount;
                }
            
    }
}