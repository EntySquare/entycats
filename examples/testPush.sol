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
       testPush3();
       
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
                   betHighAmount : 20,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
        _pool.highPool += 20;
    }
    function testPush3() public {
        fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].keyIndex = 1 ;
        fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].value.funderAddress = address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB);
        fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].value.separateBet.push(blockfunder({
                   betHighAmount : 10,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
        _pool.joiner.push(address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB));
        _pool.highPool += 10;
    }
    function getMap() public view returns (uint keyIndex,address funderAddress,uint length,
    uint betHighAmount,uint betLowAmount,uint block,
    uint betHighAmount1,uint betLowAmount1,uint block1,
    uint betHighAmount2,uint betLowAmount2,uint block2) {
        keyIndex = fundermap[address(this)].keyIndex ;
        funderAddress = fundermap[address(this)].value.funderAddress ;
        length =fundermap[address(this)].value.separateBet.length ;
        betHighAmount =fundermap[address(this)].value.separateBet[0].betHighAmount;
        betLowAmount =fundermap[address(this)].value.separateBet[0].betLowAmount;
        block =fundermap[address(this)].value.separateBet[0].block;
        betHighAmount1 =fundermap[address(this)].value.separateBet[1].betHighAmount;
        betLowAmount1 =fundermap[address(this)].value.separateBet[1].betLowAmount;
        block1 =fundermap[address(this)].value.separateBet[1].block;
        betHighAmount2 =fundermap[address(this)].value.separateBet[2].betHighAmount;
        betLowAmount2 =fundermap[address(this)].value.separateBet[2].betLowAmount;
        block2 =fundermap[address(this)].value.separateBet[2].block;
    }
    
    function testPushExist(address adkey,uint betAmount) public  returns (uint needapprove){
        if(contains(adkey)) {  //增加下注 
              fundermap[adkey].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
                _pool.highPool = _pool.highPool.add(betAmount);
                needapprove =  getValue(adkey,1);
              }
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