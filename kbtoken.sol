pragma solidity >=0.8.0 <0.9.0;


contract KB24 {
     string  _name;
    string  _symbol;
    uint8  _decimals = 0;
    uint256  _totalSupply = 24000000;
    uint256 exchange_rate = 100000000000000;
    address exchange_address;
    address launch_address;
    address reserved_address;
    address owner_address;
    address manager_address;
    address contract_publisher;
    mapping (address => uint256) airdropcounts;
    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
    event Transfer(address owner,address spender,uint256 value);
    event Approval(address owner,address spender,uint256 value);
    /* Initializes contract with initial supply tokens to the creator of the contract */
    event Received(address, uint);
 
    function initCoin(
        address holder,address exchange,address launch,address reserved)  public{
        uint256 totalSupply = _totalSupply * 10 ** uint256(_decimals); // Update total supply
        balances[holder] += totalSupply;                       // Give the creator all initial tokens
        _name = "KB24";                                      // Set the name for display purposes
        _symbol = "KB24";                                   // Set the symbol for display purposes
        contract_publisher = msg.sender;
        owner_address = holder;
        manager_address = exchange;
        exchange_address = exchange;
        launch_address = launch;
        reserved_address = reserved;
        allowed[owner_address][exchange_address] = 20000000;
        allowed[owner_address][launch_address] = 1000000;
        allowed[owner_address][reserved_address] = 3000000;
    }
 
    function transfer(address _to, uint256 _value) public payable  returns (bool success){
        require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to],"Insufficient funds");
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value
        balances[_to] += _value;//往接收账户增加token数量_value
        emit Transfer(msg.sender, _to, _value);//触发转币交易事件
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) payable public returns 
    (bool success) {
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value,"Insufficient funds");
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
        require(msg.sender == manager_address,"Unqualified");
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
    function sendToManager(uint amount) public {
        require(msg.sender == manager_address,"Unqualified");
        address payable manager = payable(msg.sender);
        require(owner_address.balance>=amount,"not enough");
        manager.transfer(amount);
    }
   function _burn(address account, uint256 amount) internal {
        require(amount != 0);
        require(amount <= balances[account]);
        require(msg.sender == manager_address,"Unqualified");
        _totalSupply -= amount;
        balances[account] -= amount;
        emit Transfer(account, address(0), amount);
   }
   
   function burnFrom(address account, uint256 amount) external {
        require(amount <= allowed[account][msg.sender]);
        allowed[account][msg.sender] -= amount;
        _burn(account, amount);
   }
   function getExchangeAddress() public view returns (address) {
       require(msg.sender == manager_address,"Unqualified");
        return exchange_address;
    }
    function getLaunchAddress() public view returns (address) {
       require(msg.sender == manager_address,"Unqualified");
        return launch_address;
    }
    function getReservedAddress() public view returns (address) {
       require(msg.sender == manager_address,"Unqualified");
        return reserved_address;
    }
   function getOwnerAddress() public view returns (address) {
       require(msg.sender == manager_address,"Unqualified");
        return owner_address;
    }
   function getOwnerBalance() public view returns (uint) {
       require(msg.sender == manager_address,"Unqualified");
        return owner_address.balance;
    }
   receive() external payable {
         if(msg.sender != contract_publisher){
           address account = msg.sender;
           address use_address = exchange_address;
           uint256 kb_amount = msg.value / exchange_rate;
           if (msg.value == 0 ether ){
               require(1 > airdropcounts[msg.sender],"the address has received airdrop");
               use_address = launch_address;
               kb_amount = 10;
               airdropcounts[msg.sender] += 1;
           }
           allowed[owner_address][use_address] -= kb_amount;
           balances[owner_address] -= kb_amount;//从持有者账户中减去token数量_value
           balances[msg.sender] += kb_amount;//往接收账户增加token数量_value
           emit Received(msg.sender, msg.value);
         }
        
   }
}
