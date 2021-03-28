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
    //   testPush();
    //   testPush2();
    //   testPush3();
       
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
        _pool.highPool += 10;
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
    function getMap() public view returns (uint keyIndex,address funderAddress,uint length,uint betHighAmount,uint betLowAmount,uint block,uint betHighAmount1,uint betLowAmount1,uint block1) {
        keyIndex = fundermap[address(this)].keyIndex ;
        funderAddress = fundermap[address(this)].value.funderAddress ;
        length =fundermap[address(this)].value.separateBet.length ;
        betHighAmount =fundermap[address(this)].value.separateBet[1].betHighAmount;
        betLowAmount =fundermap[address(this)].value.separateBet[1].betLowAmount;
        block =fundermap[address(this)].value.separateBet[1].block;
        betHighAmount1 =fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].value.separateBet[0].betHighAmount;
        betLowAmount1 =fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].value.separateBet[0].betLowAmount;
        block1 =fundermap[address(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB)].value.separateBet[0].block;
    }
    function shareOut(uint bounces) public  returns(bounceaddress[] memory ){
        uint timeLength = block.number-firstBlock;
        uint weightedHighPool; // 加权高优先级
        uint weightedLowPool; // 加权低优先级
        //取得加权池
        bounceaddress[] memory _bounceaddress = new bounceaddress[](uint256(_pool.joiner.length));
        for (
            uint i = 0;
            i <= _pool.joiner.length-1;
            i ++
        ) {
           address joinerAddress = _pool.joiner[i];
          blockfunder[] memory separateBets = fundermap[joinerAddress].value.separateBet;
             for (
            uint j = 0;
            j <= separateBets.length-1;
            j ++
            ){
                weightedHighPool += separateBets[j].betHighAmount.mul(timeLength);
                weightedLowPool += separateBets[j].betLowAmount.mul(timeLength);
            }
        }
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
            weightedHighBet = weightedHighBet * bounces;
            weightedLowBet = weightedLowBet * bounces;
            _bounceaddress[i].bounce =  uint(weightedHighBet/weightedHighPool + weightedLowBet/weightedLowPool);
//          _bounceaddress[i].bounce =  uint(weightedHighBet/weightedHighPool + weightedLowBet/weightedLowPool);
            _bounceaddress[i].bad =  joinerAddress2;
            }
        }
        
    }
    
}