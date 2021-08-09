// SPDX-License-Identifier: SimPL-2.0
pragma solidity >=0.8.0 <0.9.0;
import "./investors_usdt.sol";
interface IERC20 {
    function totalSupply()  external returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    event Transfer(address  owner, address  spender, uint256 value);
    event Approval(address  owner, address  spender, uint256 value);
}
// @title nucleus contract
// @author yueliyangzi
contract nucleus{
    string public  _name;
    string public _symbol;
    uint8  public _decimals = 12;
    uint256 public  _totalSupply = 1000000000000;
    uint8 constant ACTIVE = 0;
    uint8 constant SUSPENDED = 1;
    uint8 constant CLOSED = 2;
    uint8 public marketStatus;
    address manager_address;
    mapping (address => uint256)  balances;
    mapping (address => mapping (address => uint256)) allowed;
    event Transfer(address owner,address spender,uint256 value);
    event Approval(address owner,address spender,uint256 value);
    event ChangeMarketStatusEvent(uint8 status);
    
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
    function approve(address _spender, uint256 _value) public returns (bool success)   
    { 
         require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
       emit Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];//允许_spender从_owner中转出的token数
    }
    function name() public view returns (string memory) {
        return _name;
    }
    function symbol() public view returns (string memory){
        return _symbol;
    }
    function decimals() public view returns (uint8){
        return _decimals;
    }
    function viewtotalSupply() public view returns (uint256){
        return _totalSupply;
    }
    //@notice Interface reserved for change market status 
    function changeMarketStatus(uint8 _status) external {
        require (msg.sender == manager_address);
        if (marketStatus == CLOSED) revert();  // closed is forever
        marketStatus = _status;
        emit ChangeMarketStatusEvent(_status);
    }
}
// @title cell contract
// @author yueliyangzi
contract CELL is nucleus {
    using SafeMathCell for uint256;
    using AddressArrayLimitCell for address[31];
    using AddressArrayCell for address[];
    IERC20 public aToken;
    uint256 constant exchange_rate_usdt = 100000;
    mapping(uint8 => TokenInfo) tokens;
    address[] partners;
    mapping (address => ManagerStatus) manager;
    mapping (address => AddressStatus) details;
    struct TokenInfo{
        string token_name;
        address token_address;
        uint256 exchange_rate; 
    }
    struct AddressStatus{
        uint256 locked_balances;
        uint256 ido_balances;
        uint256 pledge_balances;
        uint256 avilable_balances;
        bool airdropflag;
        address recommender;
        uint256 pledge_power;
    }
    struct ManagerStatus{
        address owner_address;
        uint256 launch_pool;
        uint256 ido_pool;
        uint256 creator_pool;
        uint256 lp_pool;
        uint256 pledge_pool;
        uint256 primary_jackpot;
        uint256 secondary_jackpot;
        uint256 final_jackpot;
        address[31] jackpot_joiners;
        ComputingPowerPool pledge_period;
    }
    struct ComputingPowerPool{
        uint256 all_power;
        address[] miners;
        
    }
    //@notice 实例化测试的合约
    function initToken  (address tcaddress) public{
//         usToken = USDT(tcaddress);
      aToken = IERC20(tcaddress);
    }
     /* Initializes contract with initial supply tokens to the creator of the contract */
    //@notice Contract initial setting
    function initCoin(
        address holder)  public{
        uint256 totalSupply = _totalSupply * 10 ** uint256(_decimals); // Update total supply
        balances[holder] += totalSupply;                       // Give the creator all initial tokens
        _name = "CELL";                                      // Set the name for display purposes
        _symbol = "CELL";                                   // Set the symbol for display purposes
        manager_address = holder;
        marketStatus = ACTIVE;
        manager[holder].owner_address = holder;
        manager[holder].launch_pool = 20000000000;
        manager[holder].ido_pool = 105000000000;
        manager[holder].creator_pool = 45000000000;
        manager[holder].lp_pool = 180000000000;
        manager[holder].pledge_pool = 650000000000;
    }
    function transfer(address _to, uint256 _value) public payable  returns (bool success){
        require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to],"Insufficient funds");
        require(details[msg.sender].avilable_balances >= _value,"Insufficient funds");
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value
        if (details[msg.sender].ido_balances > _value){
            details[msg.sender].ido_balances -= _value;
        }
        else{
            details[msg.sender].ido_balances = 0;
        }
        details[msg.sender].avilable_balances -= _value;
        uint256 tax = (_value * 8).div(100);
        uint256 to_jackpot = tax.div(4);
        _burn(tax.div(4));
        uint256 to_partner = tax.div(80);
        manager[manager_address].primary_jackpot += to_jackpot;
        manager[manager_address].jackpot_joiners.pushLimit(msg.sender);
         for (
                uint i = 0;
                i <= partners.length-1;
                i ++
                ){
                address partner = partners[i];
                balances[partner] += to_partner;
        }
        balances[_to] += _value.sub(tax);//往接收账户增加token数量_value
        details[_to].avilable_balances += _value.sub(tax);
        emit Transfer(msg.sender, _to, _value);//触发转币交易事件
        if(manager[manager_address].primary_jackpot >= 50000){
            award();
        }
        return true;
    }
    function transferFrom(address _from, address _to, uint256 _value) payable public returns 
    (bool success) {
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value,"Insufficient funds");
        require(details[_from].avilable_balances >= _value,"Insufficient funds");
        balances[_from] -= _value; //支出账户_from减去token数量_value
        if (details[_from].ido_balances > _value){
            details[_from].ido_balances -= _value;
        }
        else{
            details[_from].ido_balances = 0;
        }
        details[_from].avilable_balances -= _value;
        uint256 tax = (_value * 8).div(100);
        uint256 to_jackpot = tax.div(4);
        _burn(tax.div(4));
        uint256 to_partner = tax.div(80);
        manager[manager_address].primary_jackpot += to_jackpot;
        manager[manager_address].jackpot_joiners.pushLimit(_from);
         for (
                uint i = 0;
                i <= partners.length-1;
                i ++
                ){
                address partner = partners[i];
                balances[partner] += to_partner;
        }
        balances[_to] += _value.sub(tax);//接收账户增加token数量_value
        details[_to].avilable_balances += _value.sub(tax);
        allowed[_from][msg.sender] -= _value;//消息发送者可以从账户_from中转出的数量减少_value
        emit Transfer(_from, _to, _value);//触发转币交易事件
        if(manager[manager_address].primary_jackpot >= 50000){
            award();
        }
        return true;
    }
    //@notice Logic of prize pool
    function award() internal {
        uint256 bonus = 0;
        uint256 pj = manager[manager_address].primary_jackpot;
        bonus += pj.mul(6).div(10);
        uint256 to_s = pj.mul(4).div(10);
        manager[manager_address].secondary_jackpot += to_s;
        manager[manager_address].primary_jackpot = 0;
        if(manager[manager_address].secondary_jackpot >= 100000){
            uint256 sj = manager[manager_address].secondary_jackpot;
            bonus += sj.mul(6).div(10);
            uint256 to_f = sj.mul(4).div(10);
            manager[manager_address].final_jackpot += to_f;
            manager[manager_address].secondary_jackpot = 0;
            if(manager[manager_address].final_jackpot >= 200000){
                uint256 fj = manager[manager_address].final_jackpot;
                bonus += fj.mul(9).div(10);
                uint256 to_p = fj.div(10);
                manager[manager_address].primary_jackpot += to_p;
                manager[manager_address].final_jackpot = 0;
            }
        }
         for (
                uint i = 30;
                i >= 0;
                i --
                ){
                address joiner = manager[manager_address].jackpot_joiners[i];
                if (i == 30){
                   balances[joiner] += bonus.div(2);
                   details[joiner].avilable_balances += bonus.div(2);
                }
                if (i <= 29 && i >= 20){
                   balances[joiner] += bonus.div(40);
                   details[joiner].avilable_balances += bonus.div(40);
                }
                if(i <= 19 && i >= 0){
                   balances[joiner] += bonus.div(80);    
                   details[joiner].avilable_balances += bonus.div(80);
                }
        }
        
    }    
  //@notice burn function internal
  function _burn(uint256 amount) internal {
        require(amount != 0);
        require(amount <= _totalSupply);
        _totalSupply -= amount;
  }     
  //@notice bind user with recommender (one can only have one recommender)
  function bindRecommender(address _recommender) external returns(bool){
      require(details[msg.sender].recommender == address(0),"this address has been bound");
      _bind(msg.sender,_recommender);
      return true;
  }
  //@notice bind function internal
  function _bind(address _self,address _to) internal{
      details[_self].recommender = _to;
  }
  //@notice Logic of airdrop
  function airdrop(address _recommender) external returns(bool){
      require(!details[msg.sender].airdropflag,"the address has received airdrop");
             if (_recommender == address(0)){
              require(manager[manager_address].launch_pool >= 1000000,"airdrop award has been used up");
              _airdrop(msg.sender,1000000);
             }
             else{
              require(manager[manager_address].launch_pool >= 1500000,"airdrop award has been used up");
              require(details[msg.sender].recommender == address(0),"this address has been bound");
              _airdrop(msg.sender,1000000);
              _airdrop(_recommender,500000);
              _bind(msg.sender,_recommender);
             }
             return true;
  }
  //@notice airdrop function internal
  function _airdrop(address _to,uint256 amount) internal{
       details[_to].locked_balances += amount;
       balances[_to] += amount;
       details[_to].airdropflag = true;
       manager[manager_address].launch_pool -= amount;
       balances[manager_address] -= amount;
  }
  
  
   //@notice Logic of join ido project
  function join_ido(uint256 _value) external payable returns(bool){
      uint256 cell_amount = _value * exchange_rate_usdt;
      require(manager[manager_address].ido_pool >= cell_amount,"not enough ido funds in pool");
      //TODO 
      aToken.transferFrom(msg.sender,manager_address,_value);
      balances[manager_address] -= cell_amount;
      details[msg.sender].ido_balances += cell_amount;
      details[msg.sender].avilable_balances += cell_amount;
      balances[msg.sender] += cell_amount;
      emit Transfer(msg.sender, manager_address, _value);//触发转币交易事件前端先
      return true;
  }
   //@notice user who join ido can recover 80% of investment
  function withdraw_ido(uint256 _value) external payable returns(bool){
      uint256 usdt_amount = _value * exchange_rate_usdt.mul(8).div(10);
      require(details[msg.sender].ido_balances >= _value && details[msg.sender].avilable_balances >= _value &&  balances[msg.sender] >= _value,"this address does not have enough ido funds");
      require(aToken.balanceOf(manager_address) >= usdt_amount);
      aToken.transfer(manager_address,usdt_amount);
      balances[manager_address] += _value;
      details[msg.sender].ido_balances -= _value;
      details[msg.sender].avilable_balances -= _value;
      balances[msg.sender] -= _value;
      emit Transfer(manager_address, msg.sender, _value);//触发转币交易事件
      return true;
  }   
  //@notice Participate in contract building community
  function join_partners() external payable returns(bool){
      require(partners.length <= 40,"There are already 40 people ");
      //TODO 
      aToken.transferFrom(msg.sender,manager_address,4000);
    //   balances[manager_address] -= cell_amount;
    //   details[msg.sender].ido_balances += cell_amount;
    //   details[msg.sender].avilable_balances += cell_amount;
    //   balances[msg.sender] += cell_amount;
      require(!partners.containAddress(msg.sender));
      partners.push(msg.sender);
      return true;
  } 
  //@notice LP liquidity
  function lp_miner() external{
      //TODO
  }   
  
  //@notice pledge miner
  function pledge_miner(uint256 _amount,bool pledgeflag) external returns(bool){
      if(!pledgeflag){
       return _quit(_amount);
      }
      return _pledge(_amount);
  }
  //@notice pledge function internal
  function _pledge(uint256 _amount) internal returns(bool){
      require (_amount <= details[msg.sender].avilable_balances,"not enough token in this address");
      bool exist = manager[manager_address].pledge_period.miners.containAddress(msg.sender);
       if (!exist){
           manager[manager_address].pledge_period.miners.push(msg.sender);
       }
       details[msg.sender].avilable_balances -= _amount;
       details[msg.sender].pledge_balances += _amount;
       return true;
  }
  //@notice _quit function internal
  function _quit(uint256 _amount) internal returns(bool){
      require (_amount <= details[msg.sender].pledge_balances,"not enough token in this pledge pool");
      details[msg.sender].pledge_balances -= _amount;
      details[msg.sender].avilable_balances += _amount;
      bool exist = manager[manager_address].pledge_period.miners.containAddress(msg.sender);
      if(exist&&details[msg.sender].pledge_balances==0){
          manager[manager_address].pledge_period.miners.deleteAddress(msg.sender);
      }
      return true;
  }
  //@notice Interface reserved for external timed execution
  function pledge_award() external returns(bool){
      require (msg.sender == manager_address);
      if (!_pledge_award_prepare()) revert();
      for (
                uint i = 0;
                i <= manager[manager_address].pledge_period.miners.length-1;
                i ++
                ){
                    address pointer = manager[manager_address].pledge_period.miners[i];
                    if (details[pointer].pledge_balances > 0){
                        
                    }
                }
      return true;
      
  }
  //@notice Logic of pledge prize pool
  function _pledge_award_prepare() internal returns(bool){
       for (
                uint i = 0;
                i <= manager[manager_address].pledge_period.miners.length-1;
                i ++
                ){
                   uint256 _pledge_power = _pledge_power_calculate(manager[manager_address].pledge_period.miners[i]);
                   details[manager[manager_address].pledge_period.miners[i]].pledge_power += _pledge_power;
                   manager[manager_address].pledge_period.all_power += _pledge_power;
                   bool success = _pledge_recommender_calculate(manager[manager_address].pledge_period.miners[i],manager[manager_address].pledge_period.miners[i],_pledge_power,1);
                   if (!success){
                       return false;
                   }
                }
                return true;
  }
  //@notice calc function internal
  function _pledge_recommender_calculate(address _self,address _now,uint256 _pledge_power,uint256 _generation) internal returns(bool){
            if (details[_now].recommender == address(0)){
                return true;
             }
             else{
               uint256 ratio = _generation.recommender_radio();
               if(ratio == 0){
                   return true;
               }
               details[_now].pledge_power += _pledge_power.mul(ratio).div(100);
               manager[manager_address].pledge_period.all_power += _pledge_power.mul(ratio).div(100);
               _generation ++ ;
               bool success = _pledge_recommender_calculate(_self,details[_now].recommender,_pledge_power,_generation);
               return success;
             }
             
  }
  //@notice calc function internal
  function _pledge_power_calculate(address _to) internal view returns(uint256){
       uint256 _pledge_balances = details[_to].pledge_balances;
       uint256 _pledge_weight = _pledge_balances.weight_coefficient();
       uint256 _pledge_power = _pledge_balances.mul(_pledge_weight).div(10); //TODO  Compute power
       return _pledge_power;
  }
  //@notice Interface reserved for add token can be used
  function addTokenExchange(uint8 _token_code,string memory _token_name,address _token_address,uint256 _exchange_rate) external{
       if (msg.sender != manager_address) revert();
       tokens[_token_code].token_name = _token_name;
       tokens[_token_code].token_address = _token_address;
       tokens[_token_code].exchange_rate = _exchange_rate;
  }
  //@notice Interface reserved for withdraw token from contract address (only for manager)
  function withdrawToken(uint8 _token_code,address _received,uint256 _amount) external{
      if (msg.sender != manager_address) revert();
         address _token_address = tokens[_token_code].token_address;
         aToken = IERC20(_token_address);
         require(aToken.balanceOf(manager_address) >= _amount);
         aToken.transferFrom(manager_address,_received,_amount);
  }
}
// @title array limit library
// @author yueliyangzi
library AddressArrayLimitCell{
    function pushLimit(address[31] memory origin, address input) internal pure returns (address[31] memory result) {
        for(
          uint i = 0; i < 30; i++
            ){
             origin[i] = origin[i+1];   
        }
        origin[30] = input;
        return origin;
    }
}
// @title array library
// @author yueliyangzi
library AddressArrayCell{
    function deleteAddress(address[] memory origin,address pointer) internal pure returns(address[] memory result){
        if(containAddress(origin,pointer)){
        uint index;
        for(
            uint i = 0;i < origin.length; i++
            ){
              if(origin[i] == pointer){
                  index = i;
              }    
            }
        delete origin[index];
        for(
            uint i = index;i < origin.length-1;i++
            ){
                origin[i]=origin[i+1];
            }
        }
        return origin;
    }
    function containAddress(address[] memory origin,address pointer) internal pure returns(bool){
          for(
            uint i = 0;i < origin.length; i++
            ){
              if(origin[i] == pointer){
                  return true;
              }    
            }
            return false;
    }
}
// @title cell library
// @author yueliyangzi
library SafeMathCell {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
 
        uint256 c = a * b;
        require(c / a == b, "SafeMath:multiplication overflow");
        return c;
    }
 

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath:division overflow");
        uint256 c = a / b;
        return c;
    }
 

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "SafeMath: subtraction overflow");
        uint256 c = a - b;
 
        return c;
    }
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");
 
        return c;
    }

    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0, "SafeMath: mod overflow");
        return a % b;
    }
    function recommender_radio(uint256 _generation) internal pure returns(uint256 ratio){
              if(_generation == 1){
                   return 20;
               }
               if(_generation >= 2 && _generation <= 4){
                   return 15;
               }
               if(_generation >= 5 && _generation <= 10){
                   return 10;
               }
               if(_generation >= 11 && _generation <= 15){
                   return 8;
               }
               if(_generation >= 16 && _generation <= 20){
                   return 6;
               }
               if(_generation >= 21){
                   return 0;
               }
    }
    function weight_coefficient(uint256 _pledge_balances) internal pure returns(uint256 ratio){
        if(_pledge_balances < 100000000){
           return 15;
       }
       if(_pledge_balances >= 100000000 && _pledge_balances < 500000000){
           return 12;
       }
       if(_pledge_balances >= 500000000 && _pledge_balances < 1000000000){
           return 10;
       }
       if(_pledge_balances >= 1000000000 && _pledge_balances < 2000000000){
           return 7;
       }
       if(_pledge_balances >= 2000000000){
           return 5;
       }
    }
}
