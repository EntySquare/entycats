pragma solidity ^0.4.17;
 
contract testStorageandMemory{
    uint[] public array;
    function initTest() public{
        array.push(1);
        array.push(2);
        uint[] temp = array; //这里相当于unit [] storage temp = array;
        temp.push(3);
        temp[0] = 99;
        
    }
    function getArray(uint num) public view returns (uint result){
        result = array[num];
    }
    function getLow(uint a,uint b) public view returns (uint result){
        result = uint(a*10000/b+b/a);
    }
}