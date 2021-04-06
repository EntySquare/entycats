pragma solidity ^0.4.17;
 
contract testencrypt{
   string public word;
   bytes32 public ciphertext;
   constructor() public{
       word = "hello world";
   }
    function initword(string _word) public{
        word = _word;
        
    }
    function encode() public view returns (bytes32 result){
        ciphertext = sha256(word);
        result = ciphertext;
    }
    function encrypt() public view returns (string result){
        
    }
    function decrypt() public view returns (string result){
        
    }
}