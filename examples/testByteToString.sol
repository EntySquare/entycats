pragma solidity ^0.4.17;
 
contract testByteToString{
    string password ;
   /// string类型转化为bytes12型转
    function stringTobytes12(string memory source) constant public returns(bytes12 result){
        assembly{
            result := mload(add(source,32))
        }
    }
  /// bytes12类型转化为string型转
    function bytes12ToString(bytes12 x) constant public returns(string){
        bytes memory bytesString = new bytes(32);
        uint charCount = 0 ;
        for(uint j = 0 ; j<32;j++){
            byte char = byte(bytes12(uint(x) *2 **(8*j)));
            if(char !=0){
                bytesString[charCount] = char;
                charCount++;
            }
        }
        bytes memory bytesStringTrimmed = new bytes(charCount);
        for(j=0;j<charCount;j++){
            bytesStringTrimmed[j]=bytesString[j];
        }
        return string(bytesStringTrimmed);
    }
}