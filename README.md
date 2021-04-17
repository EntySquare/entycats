# entycats
contracts of HAP (HSF collection and usdt delievery)

delivery of real USDT/Tether is represented by mockusdt , find more information at investors_usdt.sol


#### build
we build our abi and golang lib from listed command through solc

    solc --abi investors.sol -o ./abi
    abigen --abi=./abi/USDT.abi --pkg=investors_usdt --out=investors_usdt.go


#### test
see infos at ./lib and ./mockusdt