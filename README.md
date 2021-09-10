Microsoft Windows [版本 10.0.22000.132]
(c) Microsoft Corporation。保留所有权利。

D:\blockchain-master\blockchain-master>testProject
00000c629224cd017d3b94b88a3d3491c186e1a79295be19bdef2f87df689c4f

使用信息：
addblock -data 交易信息-添加区块
printchain -遍历区块并打印

D:\blockchain-master\blockchain-master>testProject addblock -data 测试
0000091c969bd86e334c33d2f573ef35af6c2cde41fa2a84682f2c3935e79a4b

成功添加新区块

D:\blockchain-master\blockchain-master>testProject addblock -data 测试2
00000cc085c49853e4096a774b474201ff86e95ddc92f7522fbf121ad713f5ef

成功添加新区块

D:\blockchain-master\blockchain-master>testProject printchain
上一个区块的Hash：
区块信息：第一个区块
当前区块的Hash：00000c629224cd017d3b94b88a3d3491c186e1a79295be19bdef2f87df689c4f
Pow: true
