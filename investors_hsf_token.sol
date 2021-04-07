pragma solidity >=0.8.0 <0.9.0;
import "./investor.sol";

// @title 投资者合约
// @author zhc
contract HillStoneFinance {
    investor public ivToken;
    string  _name;
    string  _symbol;
    uint8  _decimals = 2;
    uint256  _totalSupply;
    address public owner;
    mapping (address => uint256)  balances;
    mapping (address => mapping (address => uint256))  allowed;
     event Transfer(address owner,address spender,uint256 value);
     event Approval(address owner,address spender,uint256 value);
     /* Initializes contract with initial supply tokens to the creator of the contract */
    function initCoin(
        uint256 initialSupply,
        address holder)  public{
        _totalSupply = initialSupply * 10 ** uint256(_decimals); // Update total supply
        balances[holder] += _totalSupply;                       // Give the creator all initial tokens
        _name = "hillstonefinance";                                      // Set the name for display purposes
        _symbol = "HSF";                                  // Set the symbol for display purposes
        owner = holder;
       
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
    function totalSupply() public view returns (uint256){
        return _totalSupply;
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
 
    //@dev 仅允许消息发送人扩大授权
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
    //@notice 转币交易同时影响质押合约
    function placeAndTransfer(address manager,uint256 _class,uint256 _value) payable public returns(bool success){//@param manager 投资选择的合约地址;_class 1:高优先级，2:低优先级
        require(balances[msg.sender] >= _value && balances[manager] + _value > balances[manager]);
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value
        balances[manager] += _value;//往接收账户增加token数量_value
        ivToken = investor(manager);
        ivToken.placeBet(msg.sender,_class,_value);
        emit Transfer(msg.sender, manager, _value);//触发转币交易事件
        return true;
    }
    //@notice 转币交易同时影响质押合约
    function removeAndTransfer(address manager, uint256 _class, uint256 _value) payable public returns(bool success){//@param manager 投资选择的合约地址;_class 1:高优先级，2:低优先级
        require(balances[manager] >= _value && allowed[manager][msg.sender] >= _value);
        balances[msg.sender] += _value;//接收账户增加token数量_value
        balances[manager] -= _value; //支出账户_from减去token数量_value
        allowed[manager][msg.sender] -= _value;//消息发送者可以从账户_from中转出的数量减少_value
        ivToken = investor(manager);
        ivToken.removeBet(msg.sender,_class,_value);
        emit Transfer(manager, msg.sender, _value);//触发转币交易事件
        return true;
    }
    //@notice 向符合条件投资人增发
    function seo(address _to, uint256 _value) public {
        require(msg.sender == owner);
        _totalSupply += _value;
        balances[_to] += _value;
    }
    //@notice 销毁
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
        _totalSupply -= amount;
        balances[account] -= amount;
        emit Transfer(account, address(0), amount);
   }

   function burnFrom(address account, uint256 amount) external {
        require(amount <= allowed[account][msg.sender]);
        allowed[account][msg.sender] -= amount;
        _burn(account, amount);
   }
}
