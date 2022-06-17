pragma solidity ^0.8.6;

interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address recipient, uint256 amount)
        external
        returns (bool);
    function allowance(address owner, address spender)
        external
        view
        returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}

abstract contract Ownable {
    address private _owner;
    address private _previousOwner;
    uint256 private _lockTime;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    constructor ()  {
        address msgSender = msg.sender;
        _owner = msgSender;
        emit OwnershipTransferred(address(0), msgSender);
    }

    function owner() public view returns (address) {
        return _owner;
    }   
    
    modifier onlyOwner() {
        require(_owner == msg.sender, "Ownable: caller is not the owner");
        _;
    }
    
    function renounceOwnership() public virtual onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

library SafeMath {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return sub(a, b, "SafeMath: subtraction overflow");
    }
    function sub(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        uint256 c = a - b;

        return c;
    }
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }
    function div(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        uint256 c = a / b;
        return c;
    }
}

interface IUniswapV2Factory {
    event PairCreated(
        address indexed token0,
        address indexed token1,
        address pair,
        uint256
    );

    function feeTo() external view returns (address);

    function feeToSetter() external view returns (address);

    function getPair(address tokenA, address tokenB)
        external
        view
        returns (address pair);

    function allPairs(uint256) external view returns (address pair);

    function allPairsLength() external view returns (uint256);

    function createPair(address tokenA, address tokenB)
        external
        returns (address pair);

    function setFeeTo(address) external;

    function setFeeToSetter(address) external;
}

interface IUniswapV2Pair {
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
    event Transfer(address indexed from, address indexed to, uint256 value);

    function name() external pure returns (string memory);

    function symbol() external pure returns (string memory);

    function decimals() external pure returns (uint8);

    function totalSupply() external view returns (uint256);

    function balanceOf(address owner) external view returns (uint256);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 value) external returns (bool);

    function transfer(address to, uint256 value) external returns (bool);

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external returns (bool);

    function DOMAIN_SEPARATOR() external view returns (bytes32);

    function PERMIT_TYPEHASH() external pure returns (bytes32);

    function nonces(address owner) external view returns (uint256);

    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;

    event Mint(address indexed sender, uint256 amount0, uint256 amount1);
    event Burn(
        address indexed sender,
        uint256 amount0,
        uint256 amount1,
        address indexed to
    );
    event Swap(
        address indexed sender,
        uint256 amount0In,
        uint256 amount1In,
        uint256 amount0Out,
        uint256 amount1Out,
        address indexed to
    );
    event Sync(uint112 reserve0, uint112 reserve1);

    function MINIMUM_LIQUIDITY() external pure returns (uint256);

    function factory() external view returns (address);

    function token0() external view returns (address);

    function token1() external view returns (address);

    function getReserves()
        external
        view
        returns (
            uint112 reserve0,
            uint112 reserve1,
            uint32 blockTimestampLast
        );

    function price0CumulativeLast() external view returns (uint256);

    function price1CumulativeLast() external view returns (uint256);

    function kLast() external view returns (uint256);

    function mint(address to) external returns (uint256 liquidity);

    function burn(address to)
        external
        returns (uint256 amount0, uint256 amount1);

    function swap(
        uint256 amount0Out,
        uint256 amount1Out,
        address to,
        bytes calldata data
    ) external;

    function skim(address to) external;

    function sync() external;

    function initialize(address, address) external;
}

interface IUniswapV2Router01 {
    function factory() external pure returns (address);

    function WETH() external pure returns (address);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint256 amountADesired,
        uint256 amountBDesired,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    )
        external
        returns (
            uint256 amountA,
            uint256 amountB,
            uint256 liquidity
        );

    function addLiquidityETH(
        address token,
        uint256 amountTokenDesired,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    )
        external
        payable
        returns (
            uint256 amountToken,
            uint256 amountETH,
            uint256 liquidity
        );

    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint256 liquidity,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountA, uint256 amountB);

    function removeLiquidityETH(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountToken, uint256 amountETH);

    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint256 liquidity,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountA, uint256 amountB);

    function removeLiquidityETHWithPermit(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountToken, uint256 amountETH);

    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapTokensForExactTokens(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapExactETHForTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts);

    function swapTokensForExactETH(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapExactTokensForETH(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function swapETHForExactTokens(
        uint256 amountOut,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts);

    function quote(
        uint256 amountA,
        uint256 reserveA,
        uint256 reserveB
    ) external pure returns (uint256 amountB);

    function getAmountOut(
        uint256 amountIn,
        uint256 reserveIn,
        uint256 reserveOut
    ) external pure returns (uint256 amountOut);

    function getAmountIn(
        uint256 amountOut,
        uint256 reserveIn,
        uint256 reserveOut
    ) external pure returns (uint256 amountIn);

    function getAmountsOut(uint256 amountIn, address[] calldata path)
        external
        view
        returns (uint256[] memory amounts);

    function getAmountsIn(uint256 amountOut, address[] calldata path)
        external
        view
        returns (uint256[] memory amounts);
}

interface IUniswapV2Router02 is IUniswapV2Router01 {
    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountETH);

    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 amountETH);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external;

    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable;

    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external;
}

contract OceanToken is IERC20, Ownable {
    using SafeMath for uint256;

    mapping(address => uint256) private _tOwned;
    mapping(address => mapping(address => uint256)) private _allowances;
    mapping (address => bool) isDividendExempt;
    mapping(address => bool) private _isExcludedFromFee;
    mapping(address => bool) private _updated;
    
    uint256 private _tFeeTotal;
    string private _name = "Ocean";
    string private _symbol = "Ocean";
    uint8 private _decimals = 18;

    uint256 public _communityFee = 150;//社区分红3%
    uint256 public _fundFee = 50;//基金会1%
    uint256 public _marketingFee = 50;//营销1%
    uint256 public _lpFee = 150;//添加流动池分红3%
    uint256 public _fomoFee = 50;//FOMO奖励1%
    uint256 public _burnFee = 50;//销毁1%
    uint256 public _inviterFee = 200;//推广奖励4%

    uint256 currentIndex;  
    uint256 private _tTotal = 100 * 10**8 * 10**_decimals;
    uint256 distributorGas = 500000;
    uint256 public lpMinPeriod = 23 hours;
    uint256 public lpFeeShareTime;

    IUniswapV2Router02 public uniswapV2Router;
    address public uniswapV2Pair;
    address private fromAddress;
    address private toAddress;
    address public communityAddress;
    address public marketingAddress;
    address public fundAddress;
    address public fomoAddress;

    mapping(address => address) public inviter;
    address[] shareholders;
    mapping (address => uint256) shareholderIndexes;

    uint256 public fomoIndex = 0;
    uint256 public fomoMinPeriod = 4 hours;
    uint256 public fomoLastTime;
    mapping(uint => address) public fomoRewardAddress;
     
    constructor(address receiveAddress_, address communityAddress_, address marketingAddress_, address fundAddress_, address fomoAddress_, address routerAddress_, address usdtAddress_) {
        _tOwned[receiveAddress_] = _tTotal;
        communityAddress = communityAddress_;
        marketingAddress = marketingAddress_;
        fundAddress = fundAddress_;
        fomoAddress = fomoAddress_;
        fomoLastTime = block.timestamp + 30 days;
       
        uniswapV2Router = IUniswapV2Router02(routerAddress_);
        uniswapV2Pair = IUniswapV2Factory(uniswapV2Router.factory()).createPair(address(this), usdtAddress_);

        //exclude owner and this contract from fee
        _isExcludedFromFee[receiveAddress_] = true;
        _isExcludedFromFee[communityAddress] = true;
        _isExcludedFromFee[marketingAddress] = true;
        _isExcludedFromFee[fundAddress] = true;
        _isExcludedFromFee[fomoAddress] = true;
        _isExcludedFromFee[msg.sender] = true;
        _isExcludedFromFee[address(this)] = true;
        isDividendExempt[address(this)] = true;
        isDividendExempt[address(0)] = true;
        
        emit Transfer(address(0), receiveAddress_, _tTotal);
    }

    function name() public view returns (string memory) {
        return _name;
    }

    function symbol() public view returns (string memory) {
        return _symbol;
    }

    function decimals() public view returns (uint256) {
        return _decimals;
    }

    function totalSupply() public view override returns (uint256) {
        return _tTotal;
    }

    function balanceOf(address account) public view override returns (uint256) {
        return _tOwned[account];
    }

    function transfer(address recipient, uint256 amount)
        public
        override
        returns (bool)
    {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function allowance(address owner, address spender)
        public
        view
        override
        returns (uint256)
    {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 amount)
        public
        override
        returns (bool)
    {
        _approve(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public override returns (bool) {
        _transfer(sender, recipient, amount);
        _approve(
            sender,
            msg.sender,
            _allowances[sender][msg.sender].sub(
                amount,
                "ERC20: transfer amount exceeds allowance"
            )
        );
        return true;
    }

    function increaseAllowance(address spender, uint256 addedValue)
        public
        virtual
        returns (bool)
    {
        _approve(
            msg.sender,
            spender,
            _allowances[msg.sender][spender].add(addedValue)
        );
        return true;
    }

    function decreaseAllowance(address spender, uint256 subtractedValue)
        public
        virtual
        returns (bool)
    {
        _approve(
            msg.sender,
            spender,
            _allowances[msg.sender][spender].sub(
                subtractedValue,
                "ERC20: decreased allowance below zero"
            )
        );
        return true;
    }

    function totalFees() public view returns (uint256) {
        return _tFeeTotal;
    }

   function isExcludedFromFee(address account) public view returns (bool) {
        return _isExcludedFromFee[account];
    }
    function excludeFromFee(address account) public onlyOwner {
        _isExcludedFromFee[account] = true;
    }

    function includeInFee(address account) public onlyOwner {
        _isExcludedFromFee[account] = false;
    }

    function marketingAddres(address marketingAddress_) public onlyOwner {
        marketingAddress = marketingAddress_;
    }

    function fundAddres(address fundAddress_) public onlyOwner {
        fundAddress = fundAddress_;
    }

    function communityAddres(address communityAddress_) public onlyOwner {
        communityAddress = communityAddress_;
    }

    function setPairAddress(address uniswapV2Pair_) public onlyOwner {
        uniswapV2Pair = uniswapV2Pair_;
    }

    function setRouterAddress(address routerAddress_) public onlyOwner {
        uniswapV2Router = IUniswapV2Router02(routerAddress_);
    }

    function setLpMinPeriod(uint256 lpMinPeriod_) public onlyOwner {
        lpMinPeriod = lpMinPeriod_;
    }

    function setFomoMinPeriod(uint256 fomoMinPeriod_) public onlyOwner {
        fomoMinPeriod = fomoMinPeriod_;
    }

    //to recieve ETH from uniswapV2Router when swaping
    receive() external payable {}

    function chooseFeeMode(bool takeFee) private {
        if (takeFee) {
            _communityFee = 150;
            _fundFee = 50;
            _marketingFee = 50;
            _lpFee = 150;
            _fomoFee = 50;
            _burnFee = 50;
            _inviterFee = 200;
        } else {
            _communityFee = 0;
            _fundFee = 0;
            _marketingFee = 0;
            _lpFee = 0;
            _fomoFee = 0;
            _burnFee = 0;
            _inviterFee = 0;
        }
    }

    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) private {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    function _transfer(
        address from,
        address to,
        uint256 amount
    ) private {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        require(from != to, "ERC20: transfer from is the same as to ");
        require(amount > 0, "Transfer amount must be greater than zero");
        bool takeFee = true;
        if (_isExcludedFromFee[from] || _isExcludedFromFee[to]) {
            takeFee = false;
        }

        if(takeFee) {
            uint256 allowAmount = balanceOf(from) * 90 / 100;
            require(allowAmount >= amount, "ERC20: only 90% is allowed to be transferred each time");
        }
        
        bool shouldSetInviter = balanceOf(to) == 0 && inviter[to] == address(0) && from != uniswapV2Pair && amount >= 1 * 10**_decimals;
        _tokenTransfer(from, to, amount, takeFee);
        
        if (shouldSetInviter) {
            inviter[to] = from;
        }
        if(fromAddress == address(0) ) fromAddress = from;
        if(toAddress == address(0)) toAddress = to;
        if(!isDividendExempt[fromAddress] && fromAddress != uniswapV2Pair ) setShare(fromAddress);
        if(!isDividendExempt[toAddress] && toAddress != uniswapV2Pair ) setShare(toAddress);
        
        fromAddress = from;
        toAddress = to;  
         if(_tOwned[address(this)] >= 1 * 10**4 * 10**_decimals && from !=address(this) && lpFeeShareTime.add(lpMinPeriod) <= block.timestamp) {
             lpProcess(distributorGas) ;
             lpFeeShareTime = block.timestamp;
        }
        
        if(from == uniswapV2Pair) {
            fomoProcess(to);
        }

    }

    function fomoProcess(address account) private {
        fomoRewardAddress[fomoIndex] = account;

        fomoIndex += 1;
        if(fomoIndex > 2) {
            fomoIndex = 0;
        }

        if (block.timestamp > fomoLastTime) {
            uint256 fomoReward = balanceOf(fomoAddress) * 70 / 100;
            uint256 perFomoReward = fomoReward / 3;
            
            if(perFomoReward >= 10 * 10**_decimals) {
                for(uint i = 0; i < 3; i++) {
                    address from = fomoAddress;
                    address to = fomoRewardAddress[i];
                    _tOwned[from] = _tOwned[from].sub(perFomoReward);
                    _tOwned[to] = _tOwned[to].add(perFomoReward);
                    emit Transfer(from, to, perFomoReward);
                }
            }
        } 
        fomoLastTime = block.timestamp + fomoMinPeriod;

    }
    
    function lpProcess(uint256 gas) private {
        uint256 shareholderCount = shareholders.length;

        if(shareholderCount == 0)return;
        uint256 nowbanance = _tOwned[address(this)];
        uint256 gasUsed = 0;
        uint256 gasLeft = gasleft();

        uint256 iterations = 0;

        while(gasUsed < gas && iterations < shareholderCount) {
            if(currentIndex >= shareholderCount){
                currentIndex = 0;
            }

          uint256 amount = nowbanance.mul(IERC20(uniswapV2Pair).balanceOf(shareholders[currentIndex])).div(IERC20(uniswapV2Pair).totalSupply());
         if( amount < 1 * 10**_decimals) {
             currentIndex++;
             iterations++;
             return;
         }
         if(_tOwned[address(this)]  < amount )return;
            distributeDividend(shareholders[currentIndex],amount);
            
            gasUsed = gasUsed.add(gasLeft.sub(gasleft()));
            gasLeft = gasleft();
            currentIndex++;
            iterations++;
        }
    }
   
    function distributeDividend(address shareholder ,uint256 amount) internal {
            _tOwned[address(this)] = _tOwned[address(this)].sub(amount);
            _tOwned[shareholder] = _tOwned[shareholder].add(amount);
             emit Transfer(address(this), shareholder, amount);
    }

    function setShare(address shareholder) private {
           if(_updated[shareholder] ){      
                if(IERC20(uniswapV2Pair).balanceOf(shareholder) == 0) quitShare(shareholder);              
                return;  
           }
           if(IERC20(uniswapV2Pair).balanceOf(shareholder) == 0) return;  
            addShareholder(shareholder);
            _updated[shareholder] = true;   
    }

    function addShareholder(address shareholder) internal {
        shareholderIndexes[shareholder] = shareholders.length;
        shareholders.push(shareholder);
    }
    
    function quitShare(address shareholder) private {
           removeShareholder(shareholder);   
           _updated[shareholder] = false; 
    }

    function removeShareholder(address shareholder) internal {
        shareholders[shareholderIndexes[shareholder]] = shareholders[shareholders.length-1];
        shareholderIndexes[shareholders[shareholders.length-1]] = shareholderIndexes[shareholder];
        shareholders.pop();
    }

    function _tokenTransfer(
        address sender,
        address recipient,
        uint256 amount,
        bool takeFee
    ) private {
        chooseFeeMode(takeFee);
        _transferStandard(sender, recipient, amount);
    }

    //
    function _takeburnFee(
        address sender,
        uint256 tAmount
    ) private {
        if (_burnFee == 0) return;
        if(_tFeeTotal >= 999 * 10**7 * 10**_decimals)_burnFee = 0;
        _tOwned[address(0)] = _tOwned[address(0)].add(tAmount);
        _tFeeTotal = _tFeeTotal.add(tAmount);
        emit Transfer(sender, address(0), tAmount);
    }

    function _takeLPFee(address sender,uint256 tAmount) private {
        if (_lpFee == 0) return;
        _tOwned[address(this)] = _tOwned[address(this)].add(tAmount);
        emit Transfer(sender, address(this), tAmount);
    }

    function _takeMarketingFee(address sender,uint256 tAmount) private {
        if (_marketingFee == 0) return;
        _tOwned[marketingAddress] = _tOwned[marketingAddress].add(tAmount);
        emit Transfer(sender, marketingAddress, tAmount);
    }

    function _takeCommunityFee(address sender,uint256 tAmount) private {
        if (_communityFee == 0) return;
        _tOwned[communityAddress] = _tOwned[communityAddress].add(tAmount);
        emit Transfer(sender, communityAddress, tAmount);
    }

    function _takeFundFee(address sender,uint256 tAmount) private {
        if (_fundFee == 0) return;
        _tOwned[fundAddress] = _tOwned[fundAddress].add(tAmount);
        emit Transfer(sender, fundAddress, tAmount);
    }

    function _takeFomoFee(address sender,uint256 tAmount) private {
        if (_fomoFee == 0) return;
        _tOwned[fomoAddress] = _tOwned[fomoAddress].add(tAmount);
        emit Transfer(sender, fomoAddress, tAmount);
    }

    function _takeInviterFee(
        address sender,
        address recipient,
        uint256 tAmount
    ) private {
        if (_inviterFee == 0) return;
        address cur;
        if (sender == uniswapV2Pair) {
            cur = recipient;
        } else {
            cur = sender;
        }

        uint256 accurRate;
        for (int256 i = 0; i < 8; i++) {
            uint256 rate;
            if (i == 0) {
                rate = 50;
            } else if (i == 1){
                rate = 30;
            } else {
                rate = 20;
            }
            cur = inviter[cur];

            if (cur == address(0)) {
                break;
            }
            uint256 curBalance = balanceOf(cur);
            if(curBalance <= 100000 * 10**_decimals) {
                continue;
            }
            accurRate = accurRate.add(rate);
            uint256 curTAmount = tAmount.div(10000).mul(rate);
            _tOwned[cur] = _tOwned[cur].add(curTAmount);
            emit Transfer(sender, cur, curTAmount);
        }
        _tOwned[address(0)] = _tOwned[address(0)].add(tAmount.div(10000).mul(_inviterFee.sub(accurRate)));
        emit Transfer(sender, address(0), tAmount.div(10000).mul(_inviterFee.sub(accurRate)));
    }

    function _transferStandard(
        address sender,
        address recipient,
        uint256 tAmount
    ) private {
        _tOwned[sender] = _tOwned[sender].sub(tAmount);
        _takeCommunityFee(sender, tAmount.div(10000).mul(_communityFee));
        _takeFundFee(sender, tAmount.div(10000).mul(_fundFee));
        _takeMarketingFee(sender, tAmount.div(10000).mul(_marketingFee));
        _takeLPFee(sender, tAmount.div(10000).mul(_lpFee));
        _takeFomoFee(sender, tAmount.div(10000).mul(_fomoFee));
        _takeburnFee(sender, tAmount.div(10000).mul(_burnFee));
        _takeInviterFee(sender, recipient, tAmount);
       
        uint256 recipientRate = 10000 - _communityFee - _fundFee - _marketingFee - _lpFee - _fomoFee - _burnFee - _inviterFee;
        _tOwned[recipient] = _tOwned[recipient].add(tAmount.div(10000).mul(recipientRate));
        emit Transfer(sender, recipient, tAmount.div(10000).mul(recipientRate));
    }

}