pragma solidity >=0.8.0 <0.9.0;
import "./investor.sol";


contract HillStoneFinance {
    investor public ivToken;
    string public name;
    string public symbol;
    uint8 public decimals = 18;
    uint256 public totalSupply;
    address public owner;
    mapping (address => uint256)  balances;
    mapping (address => mapping (address => uint256))  allowed;
     event Transfer(address owner,address spender,uint256 value);
     event Approval(address owner,address spender,uint256 value);
     /* Initializes contract with initial supply tokens to the creator of the contract */
    function initCoin(
        uint256 initialSupply,
        address holder)  public{
        totalSupply = initialSupply * 10 ** uint256(decimals); // Update total supply
        balances[holder] += totalSupply;                       // Give the creator all initial tokens
        name = "hillstonefinance";                                      // Set the name for display purposes
        symbol = "HSF";                                  // Set the symbol for display purposes
        owner = holder;
       
    }
 
    function transfer(address _to, uint256 _value) public payable returns (bool success) {
 
 
        require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to]);
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value
        balances[_to] += _value;//往接收账户增加token数量_value
        emit Transfer(msg.sender, _to, _value);//触发转币交易事件
        return true;
    }
 
 
    function transferFrom(address _from, address _to, uint256 _value) payable public returns 
    (bool success) {
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value);
        balances[_to] += _value;//接收账户增加token数量_value
        balances[_from] -= _value; //支出账户_from减去token数量_value
        allowed[_from][msg.sender] -= _value;//消息发送者可以从账户_from中转出的数量减少_value
        emit Transfer(_from, _to, _value);//触发转币交易事件
        return true;
    }
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
 
 
    function approve(address _spender, uint256 _value) public returns (bool success)   
    { 
         require((_value == 0) || (_value >= allowed[msg.sender][_spender]));
        allowed[msg.sender][_spender] = _value;
       emit Approval(msg.sender, _spender, _value);
        return true;
    }
    
 
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];//允许_spender从_owner中转出的token数
    }
    
    function placeAndTransfer(address manager,uint256 _class,uint256 _value) payable public returns(bool success){
        require(balances[msg.sender] >= _value && balances[manager] + _value > balances[manager]);
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value
        balances[manager] += _value;//往接收账户增加token数量_value
        ivToken = investor(manager);
        ivToken.placeBet(msg.sender,_class,_value);
        emit Transfer(msg.sender, manager, _value);//触发转币交易事件
        return true;
    }
    function removeAndTransfer(address manager, uint256 _class, uint256 _value) payable public returns(bool success){
        require(balances[manager] >= _value && allowed[manager][msg.sender] >= _value);
        balances[msg.sender] += _value;//接收账户增加token数量_value
        balances[manager] -= _value; //支出账户_from减去token数量_value
        allowed[manager][msg.sender] -= _value;//消息发送者可以从账户_from中转出的数量减少_value
        ivToken = investor(manager);
        ivToken.removeBet(msg.sender,_class,_value);
        emit Transfer(manager, msg.sender, _value);//触发转币交易事件
        return true;
    }
    function seo(address _to, uint256 _value) public {
        require(msg.sender == owner);
        totalSupply += _value;
        balances[_to] += _value;
    }
    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
    }
  
    function burnAll() external {
        uint256 amount = balances[msg.sender];
        _burn(msg.sender, amount);
    }
  
   function _burn(address account, uint256 amount) internal {
        require(amount != 0);
        require(amount <= balances[account]);
        totalSupply -= amount;
        balances[account] -= amount;
        emit Transfer(account, address(0), amount);
   }

   function burnFrom(address account, uint256 amount) external {
        require(amount <= allowed[account][msg.sender]);
        allowed[account][msg.sender] -= amount;
        _burn(account, amount);
   }
}
