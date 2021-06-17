pragma solidity >=0.8.0 <0.9.0;


contract KB24 {
     string  _name;
    string  _symbol;
    uint8  _decimals = 2;
    uint256  _totalSupply = 24000000;
    uint256 exchange_rate = 10000;
    address exchange_address;
    address launch_address;
    address reserved_address;
    address  owner_address;
    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
    event Transfer(address owner,address spender,uint256 value);
    event Approval(address owner,address spender,uint256 value);
    /* Initializes contract with initial supply tokens to the creator of the contract */

    function initCoin(
        address holder,address exchange, address launch, address reserved)  public{
        uint256 totalSupply = _totalSupply * 10 ** uint256(_decimals); // Update total supply
        balances[holder] += totalSupply;                       // Give the creator all initial tokens
        _name = "KB24";                                      // Set the name for display purposes
        _symbol = "KB24";                                  // Set the symbol for display purposes
        owner_address = holder;
        exchange_address = exchange;
        launch_address = launch;
        reserved_address = reserved;
        allowed[msg.sender][exchange_address] = 20000000;
        emit Approval(msg.sender, exchange_address, 20000000);
        allowed[msg.sender][launch_address] = 1000000;
        emit Approval(msg.sender, launch_address, 1000000);
        allowed[msg.sender][reserved_address] = 3000000;
        emit Approval(msg.sender, reserved_address, 3000000);
    }
 
    function transfer(address _to, uint256 _value) public payable  returns (bool success){
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
    function totalSupply() public view returns (uint256){
        return _totalSupply;
    }
   //@notice 向符合条件投资人增发
    function seo(address _to, uint256 _value) public {
        require(msg.sender == owner_address);
        _totalSupply += _value;
        balances[msg.sender] += _value;
        allowed[msg.sender][_to] = _value;
        emit Approval(msg.sender, _to, _value);
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
   //@notice 空投
   function airdrop() payable public returns(bool success){
        require(balances[owner_address] >= 100 && allowed[owner_address][launch_address] >= 100);
        balances[owner_address] -= 100;//从持有者账户中减去token数量_value
        balances[msg.sender] += 100;//往接收账户增加token数量_value
        emit Transfer(owner_address, msg.sender, 100);//触发转币交易事件
        uint256 last;
        last=allowed[owner_address][launch_address]-100;
        allowed[owner_address][launch_address] = last;
        emit Approval(owner_address, launch_address, last);
        return true;
   }
   //@notice 兑换
   function exchange() payable public returns(bool success){
      address account = msg.sender;
      uint256 amount = msg.value;
      require(account.balance >= amount);
      bool result = payable(msg.sender).send(amount);
      if(result == false){
          return false;
      }
      uint256 kb_amount = amount * exchange_rate;
      balances[owner_address] -= kb_amount;//从持有者账户中减去token数量_value
      balances[msg.sender] += kb_amount;//往接收账户增加token数量_value
      emit Transfer(owner_address, msg.sender, kb_amount);//触发转币交易事件
      uint256 last;
      last=allowed[owner_address][exchange_address]-kb_amount;
      allowed[owner_address][exchange_address] = last;
      emit Approval(owner_address, exchange_address, last);
      return true;
      
   }


}
