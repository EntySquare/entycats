pragma solidity >=0.8.0 <0.9.0;
import "./example2.sol";

contract TestCoin {
     example2 public exToken;
     string public name;
    string public symbol;
    uint8 public decimals = 18;
    uint256 public totalSupply;
    address public owner;
    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
     event Transfer(address owner,address spender,uint256 value);
     event Approval(address owner,address spender,uint256 value);
     /* Initializes contract with initial supply tokens to the creator of the contract */
    function initCoin(
        uint256 initialSupply,
        address holder)  public{
        totalSupply = initialSupply * 10 ** uint256(decimals); // Update total supply
        balances[holder] += totalSupply;                       // Give the creator all initial tokens
        name = "testcoin";                                      // Set the name for display purposes
        symbol = "tc";                                  // Set the symbol for display purposes
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
         require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
       emit Approval(msg.sender, _spender, _value);
        return true;
    }
 
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];//允许_spender从_owner中转出的token数
    }
    
    function shareOutAndTransfer(address manager,  uint256 bounces) payable public returns(bool success){
        exToken = example2(manager);
        example2.bounceaddress[] memory array =exToken.shareOut(bounces);
        for (uint i =0;i<array.length-1;i++){
            balances[msg.sender] -= array[i].bounce;//从消息发送者账户中减去token数量_value
            balances[array[i].bad] += array[i].bounce;//往接收账户增加token数量_value 
            emit Transfer(manager, array[i].bad, array[i].bounce);//触发转币交易事件
        }
        return true;
    }

}
