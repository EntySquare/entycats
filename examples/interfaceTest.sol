pragma solidity >=0.8.0 <0.9.0;
interface IERC20 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function hellow() external  returns (uint256);
}
contract realIERC20 {
    function hellow()   public payable returns (uint256) {
        return 100;
    }
}
contract test{
    IERC20 public usToken;
      function initToken  (address _address) public{
        usToken = IERC20(_address);
      }
      function hellow() public  returns (uint256){
          return usToken.hellow();
      }
}