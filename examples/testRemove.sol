pragma solidity >=0.8.0 <0.9.0;
import "./safemath.sol";


contract example2{
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
    
    uint bounce;
    using SafeMath for uint;
     constructor() public{
       index = 0 ;
       bounce = 1000;
       firstBlock = block.number;
       testPush();
       testPush2();
       testPush2();
       
    }
    function getNowBlock() public view returns (uint _b){
        _b = block.number;
    }
    function testPush() public {
        fundermap[address(this)].keyIndex = 1 ;
        fundermap[address(this)].value.funderAddress = address(this);
        fundermap[address(this)].value.separateBet.push(blockfunder({
                   betHighAmount : 10,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
        _pool.joiner.push(address(this));
        _pool.highPool += 10;
    }
    function testPush2() public {
        fundermap[address(this)].value.separateBet.push(blockfunder({
                   betHighAmount : 10,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
        _pool.highPool += 10;
    }
   
    function getArray(uint i) public view returns (uint keyIndex,address funderAddress,uint length,
    uint betHighAmount,uint betLowAmount,uint block) {
        keyIndex = fundermap[address(this)].keyIndex ;
        funderAddress = fundermap[address(this)].value.funderAddress ;
        length =fundermap[address(this)].value.separateBet.length ;
        betHighAmount =fundermap[address(this)].value.separateBet[i].betHighAmount;
        betLowAmount =fundermap[address(this)].value.separateBet[i].betLowAmount;
        block =fundermap[address(this)].value.separateBet[i].block;

    }
    
function removeBet(address toAddress,uint removeAmount) public  returns(bool success){
        if(contains(toAddress)) { // 并没有地址
            uint count = fundermap[toAddress].value.separateBet.length;
            uint surplus = removeAmount;
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
                            removeFunderArray(toAddress,i,count);
                        }
                        else{ //如果足够以抵扣则结束
                            fundermap[toAddress].value.separateBet[i].betHighAmount -=surplus;
                            return success;
                        }
                   }
                    
             
              //return hsfToken.transfer(toAddress,removeAmount);
        }
    }    
 
 
 function removeFunderArray(address _address ,uint removei,uint length) internal{
     if(removei != length -1){
         for(uint i = 0;i < length-1;i ++){
                 if(i==removei){
                  fundermap[_address].value.separateBet[i] = fundermap[_address].value.separateBet[i+1];
                 }
             }
         
    }
    delete fundermap[_address].value.separateBet[length-1];
 }
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
    function contains(address adkey) internal  returns (bool) {
            return fundermap[adkey].keyIndex > 0;
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