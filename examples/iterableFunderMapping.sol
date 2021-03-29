// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;
 

// 1.质押
//  i.构造对象(质押人)
//   一.地址（质押人）
//   二.质押货币数量
//   三.质押类型
   
library IterableFunderMapping {
    struct funder {
        address funderAddress ;
        uint betHighAmount;
        uint betLowAmount;
    }
    
    
    

    //  ii.构造对象（质押池）
    //   一.地址（质押池）
    //   二.低优先级总量
    //   三.高优先级总量
    //   四.质押人数量
    //   五.质押人映射
    //  address betPoolAddress;
        
    struct IndexValue { 
        uint keyIndex; 
        funder value; 
    }
    struct keyFlag{
        address adkey;
        bool deleted;
    }
    struct itFundermap{
        uint funderAccount;
        keyFlag[] keys;
        mapping(address => IndexValue) fundermap;
      
    }
        //可遍历集合的insert方法
        function insert(itFundermap storage self, address adkey,funder memory value) public returns (bool success) {
            uint keyIndex = self.fundermap[adkey].keyIndex;
 //           self.fundermap[adkey].value = value;
            if (keyIndex > 0)
                return false;
            else {
                keyIndex = self.keys.length; //0
                self.keys.push(keyFlag({
                    adkey : adkey,
                    deleted : false
                })); //1
                self.fundermap[adkey].keyIndex = keyIndex + 1 ;  //1
//              self.keys[keyIndex].adkey = adkey;
                self.funderAccount++;
                return true;
            }
        }
        //可遍历集合的remove方法
        function remove(itFundermap storage self, address adkey) public returns (bool success) {
            uint keyIndex = self.fundermap[adkey].keyIndex;
            if (keyIndex == 0)
                return false;
            delete self.fundermap[adkey];
            self.keys[keyIndex].deleted = true;
            self.funderAccount --;
        }
        //可遍历集合的update方法 1.增持h 2. 减持h 3.增持l 4. 减持l
        function update(itFundermap storage self,address adkey,uint class,uint value)public returns (bool success) {
            uint keyIndex = self.fundermap[adkey].keyIndex;
              if (keyIndex == 0)
                return false;
              else {
                  if (class == 1){
                      self.fundermap[adkey].value.betHighAmount += value;
                  }
                  if (class == 2){
                      self.fundermap[adkey].value.betHighAmount -= value;
                  }
                  if (class == 3){
                      self.fundermap[adkey].value.betLowAmount += value;
                  }
                  if (class == 4){
                      self.fundermap[adkey].value.betLowAmount -= value;
                  }
                  return true;
              }
        }
        //可遍历集合的contains方法
        function contains(itFundermap storage self, address adkey) public view returns (bool) {
            return self.fundermap[adkey].keyIndex > 0;
        }
        //可遍历集合的get方法 1.h 2.l
        function getValue(itFundermap storage self, address adkey ,uint class) public view returns (uint success) {
            if(class == 1){
                return self.fundermap[adkey].value.betHighAmount;
            }
            if(class == 2){
                return self.fundermap[adkey].value.betLowAmount;
            }
        }
        function iterate_start(itFundermap storage self) public view returns (uint keyIndex) {
            return iterate_next(self, 0); //1
        }   // length 1 keyIndex 1
    
        function iterate_valid(itFundermap storage self, uint keyIndex) public view returns (bool) {
            return keyIndex <= self.keys.length;
        }
    
        function iterate_next(itFundermap storage self, uint keyIndex) public view returns (uint r_keyIndex) {
            keyIndex++;
            while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
                keyIndex++;
            return keyIndex;
        }


        function iterate_get(itFundermap storage self, uint keyIndex) public view returns (address adkey, funder storage value) {
            adkey = self.keys[keyIndex].adkey;
            value = self.fundermap[adkey].value;
           
            
        }

    }
