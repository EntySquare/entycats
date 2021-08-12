contract model{
    using uintArrayLimit for uint[180];
    mapping(address => UserIncome) details;
    struct UserIncome{
        uint8 time;
        uint[180] income_pool;
    }
}
// @title array limit library
// @author yueliyangzi
library uintArrayLimit{
    function pushLimit(uint[180] memory origin, uint input) internal pure returns (uint[180] memory result) {
        for(
          uint i = 0; i < 180; i++
            ){
             origin[i] = origin[i+1];   
        }
        origin[179] = input;
        return origin;
    }
}

// @title cell library
// @author yueliyangzi
library SafeMath {
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

}