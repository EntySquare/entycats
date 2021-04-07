// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
import "./investors_hsf_token.sol";
import "./investors_usdt.sol";
import "./safemath.sol";
// @title 投资者合约
// @author zhc
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
        bytes32 password;
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
    // struct bounceaddress{
    //     address bad;
    //     uint bounce;
    // }
    mapping(address  => IndexValue)  fundermap;
    betPool  _pool ;
    uint index ;
    constructor(uint _startover,bool _institutionflag,address _institution) public{
       index = 0 ;
       firstBlock = block.number;
       startover = _startover;
       institutionflag = _institutionflag;
       maneger = msg.sender;
       if(institutionflag){
       institution = _institution;
       }
       
    }
    //@notice 实例化两种币的合约
    function initToken  (address hsfaddress,address tcaddress) public{
        hsfToken = HillStoneFinance(hsfaddress);
        usToken = USDT(tcaddress);
    }
    using SafeMath for uint;
    //@notice 质押方法 
    function placeBet(address fromAdress,uint _betClass,uint betAmount) external payable returns(bool success){ //@param _betClass 高低优先级 1:高优先级，2:低优先级
        //@notice 增加质押
        if(contains(fromAdress)) {  
            uint needapprove ;
            if(_betClass == 1){ 
              fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : betAmount,
                   betLowAmount : 0,
                   block : block.number
                   
               })); 
               _pool.highPool = _pool.highPool.add(betAmount);
               needapprove =  getAllValue(fromAdress);
              }
            if(_betClass ==2){ 
                fundermap[fromAdress].value.separateBet.push(blockfunder({
                   betHighAmount : 0,
                   betLowAmount : betAmount,
                   block : block.number
               })); 
               _pool.lowPool = _pool.lowPool.add(betAmount);
               needapprove =  getAllValue(fromAdress);
              }
              //@notice 调用hsf币的授权方法使用户可以随时取消
              hsfToken.approve(fromAdress,needapprove);
              return true;
        }
        //@notice 第一次加入用户新增
        //@dev  可考虑第一次授权是否同时设置密码
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
               _pool.joiner.push(fromAdress);
               //@notice 调用hsf币的授权方法使用户可以随时取消
               hsfToken.approve(fromAdress,betAmount);
               return true;
        }
    }
    //@notice 解除质押方法()
    function removeBet(address toAddress,uint _betClass,uint removeAmount) external payable returns(bool success){//@param _betClass 高低优先级 1:高优先级，2:低优先级
        //@notice 并没有地址
        if(!contains(toAddress)) { 
            return false;
        }
        else{
            uint count = fundermap[toAddress].value.separateBet.length;
            uint surplus = removeAmount;
             if(_betClass == 1){
               uint last =  getValue(toAddress,1);
               require(last >= removeAmount);
               _pool.highPool = _pool.highPool.sub(removeAmount);
                   //@notice 从后循环该地址的质押数列
                   for (
                         uint i = count-1;
                         i >= 0;
                         i --
                       ) { 
                        //@notice 如果本次下注不足以抵扣则删除本次
                        if(fundermap[toAddress].value.separateBet[i].betHighAmount < surplus){ 
                            surplus -= fundermap[toAddress].value.separateBet[i].betHighAmount;
                            fundermap[toAddress].value.separateBet[i].betHighAmount = 0;
                        }
                        //@notice 如果足够以抵扣则结束
                        else{ 
                           fundermap[toAddress].value.separateBet[i].betHighAmount -=surplus; 
                            return true;
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
                            return true;
                        }
                   }
               
              }
              //return hsfToken.transfer(toAddress,removeAmount);
        }
    }
     //@notice 分红算法
     //@dev onlyManager 仅本合约的创建者可调用且需保证合约账户上usdt值须大于bounces
     //@dev 仅盈利情况需考虑
    function shareOut(uint bounces) external onlyManager  payable{//@param bounces	投资期结束后所得盈余 
            if(bounces > startover){
            //@notice 保荐机构则将利润分与机构固定10%
            uint bou_sta = bounces - startover;
            if(institutionflag){
                usToken.transfer(institution,uint((bounces)/10));
            }
            // uint timeLength = block.number-firstBlock;
            // uint weightedHighPool = _pool.highPool * timeLength * 10; // 加权高优先级
            // uint weightedLowPool = _pool.lowPool * timeLength * 10; // 加权低优先级
            uint weightedHighPool ; // 加权高优先级
            uint weightedLowPool ; // 加权低优先级
           //@notice 取得加权后质押池
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
                    //@notice 当前地址的加权下注
                    weightedHighPool += separateBets[j].betHighAmount.mul(block.number-separateBets[j].block);
                    weightedLowPool += separateBets[j].betLowAmount.mul(block.number-separateBets[j].block);
                }
            }
            //@notice 循环遍历所有投资者账户
            for (
                uint x = 0;
                x <= _pool.joiner.length-1;
                x ++
            ) {
              address joinerAddress2 = _pool.joiner[x];
              blockfunder[] memory separateBets2 = fundermap[joinerAddress2].value.separateBet;
              uint weightedHighBounces;
              uint weightedLowBounces;
                for (
                uint y = 0;
                y <= separateBets2.length-1;
                y ++
                ){
                    //@notice 当前地址的加权下注
                    weightedHighBounces += separateBets2[y].betHighAmount.mul(block.number-separateBets2[y].block);
                    weightedLowBounces += separateBets2[y].betLowAmount.mul(block.number-separateBets2[y].block);
                }
                if(weightedHighBounces != 0 || weightedLowBounces != 0){
                    //@dev 考虑特殊情况
                    if(weightedHighPool == 0){
                        weightedHighPool = 1;
                    }
                    if(weightedLowPool == 0){
                        weightedLowPool = 1;
                    }
                        weightedHighBounces = weightedHighBounces * uint(bou_sta/10) * 3 ;
                        //@notice  本金分配
                        uint thisHigh ;
                        uint thisLow = uint((uint(startover/10) * 3 * getValue(joinerAddress2,2))/_pool.lowPool);
                    if(institutionflag){
                        weightedLowBounces = weightedLowBounces * uint(bou_sta/10) * 6;
                        thisHigh = uint((uint(startover/10) * 6 * getValue(joinerAddress2,1))/_pool.highPool);
                    }
                    else{
                        weightedLowBounces = weightedLowBounces * uint(bou_sta/10) * 7;
                         thisHigh = uint((uint(startover/10) * 7 * getValue(joinerAddress2,1))/_pool.highPool);
                    }
                   
                    
        //             _bounceaddress[i].bounce =  uint(weightedHighBet/weightedHighPool + weightedLowBet/weightedLowPool);
        //             _bounceaddress[i].bad =  joinerAddress2;
                    //@notice 调用转账方法
                    usToken.transfer(joinerAddress2,uint(weightedHighBounces/weightedHighPool + weightedLowBounces/weightedLowPool + thisHigh + thisLow));
                }
            }
            //@dev 向下取值后剩余值转入管理员账户
            uint256 lastCoin = usToken.balanceOf(address(this));
            if(lastCoin != 0){
                usToken.transfer(maneger,lastCoin);
            }
        }
        //@dev 本金小于利息
        else{
              uint sta_bou = startover - bounces;
              for (
                uint m = 0;
                m <= _pool.joiner.length-1;
                m ++
                ) {
                  if(_pool.highPool != 0 && _pool.lowPool != 0){
                  address joinerAddress3 = _pool.joiner[m];
                  uint lastBounces ;
                  if(institutionflag){
                    if (sta_bou * 10 <= startover){
                        lastBounces = uint(startover * 6  * getValue(joinerAddress3,1) /(10 * _pool.highPool)) + uint( startover * 3  * getValue(joinerAddress3,2) /(10 * _pool.lowPool));
                    } 
                    else if(sta_bou * 10 <= startover * 4){
                        lastBounces = uint(startover * 6  * getValue(joinerAddress3,1) /(10 * _pool.highPool)) + uint( (bounces - (startover * 6 / 10)) * (getValue(joinerAddress3,2) / _pool.lowPool));
                    }
                    else{
                        lastBounces = uint(bounces * getValue(joinerAddress3,1) / _pool.highPool); 
                    }
                  }
                  else{
                    if (sta_bou * 10 <= startover * 3){
                       lastBounces = uint(startover * 7  * getValue(joinerAddress3,1) /(10 * _pool.highPool)) + uint( (bounces - (startover * 7 / 10)) * (getValue(joinerAddress3,2) / _pool.lowPool)); 
                    }
                    else{
                       lastBounces = uint(bounces * getValue(joinerAddress3,1) / _pool.highPool); 
                    }
                  }
                   usToken.transfer(joinerAddress3,lastBounces);
                }
                }
                if(institutionflag){
                    if (sta_bou * 10 <= startover){
                     usToken.transfer(institution,uint(bounces - (startover * 9 / 10)));
                    } 
                }
        }
        //@notice 销毁本合约以使用的hsf货币
            hsfToken.burnAll();
    }
    //@notice 设置密码
    function setPassword(bytes32 _password) public returns (bool success){
        if(!contains(msg.sender)){
            index ++;
            fundermap[msg.sender].keyIndex = index;
            fundermap[msg.sender].password = _password;
            return true;
        }
        else{
            return false;
        }
    }
    //@notice 查看本人地址质押情况
    //@dev 尚无下注时间展示
    function getOwnersBetCondition(address _owner ,bytes32 _password) public view returns (uint highbet,uint lowbet){
        require(fundermap[_owner].password == _password,"wrong password");
        highbet = getValue(_owner,1);
        lowbet = getValue(_owner,2);
        
    }
    //@dev 内部方法方便合约内复用 
    function getHighLevelAmount() internal  view returns (uint) {
            return _pool.highPool;
    }
    function getLowLevelAmount() internal  view returns (uint) {
            return _pool.lowPool;
    }
    function contains(address adkey) internal view returns (bool) {
            return fundermap[adkey].keyIndex > 0;
    }
    function getMap(address adkey) internal view returns (uint keyIndex){
        keyIndex = fundermap[adkey].keyIndex;
        
    }
    function getValue(address adkey ,uint class) internal view returns (uint summary) {
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
    function getAllValue(address adkey) internal view returns (uint summary) {
            uint count = fundermap[adkey].value.separateBet.length;
            for (
                uint i = 0;
                i <= count-1;
                i ++ ) {
                  summary += fundermap[adkey].value.separateBet[i].betHighAmount;
                  summary += fundermap[adkey].value.separateBet[i].betLowAmount;
                }
            
    }
    // @notice 修改器
    modifier onlyManager() { 
        require(
            msg.sender == maneger,
            "Only maneger can call this."
        );
        _;
    }
}