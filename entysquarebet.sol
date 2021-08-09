pragma solidity >=0.8.0 <0.9.0;

contract entySquareRollMoney{
    mapping (address => ManagerStatus) manager;
    struct ManagerStatus{
        address[] jackpot_joiners;
    }
    function join() external payable{
       manager[address(this)].jackpot_joiners.push(msg.sender);
       if(manager[address(this)].jackpot_joiners.length > 2){
          address _winner = _roll();
          address payable winnerP = payable(_winner);
          winnerP.transfer(1 ether);
       }
    }
    function _roll() internal view returns (address){
          uint256 a = 0;
          address winner = manager[address(this)].jackpot_joiners[a];
          return winner;
    }
    fallback() external payable  {
    }


}