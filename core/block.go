package core

import (
	"PunkCoin/basetool"
	"PunkCoin/common"
	"PunkCoin/consensus"
	"bytes"

	"math/big"
)






type Block struct {

	//矿工自身地址
	MinerAddress *common.Address


	//自身哈希
	Hash *common.BlockHash

	//上一个主块的哈希值
	Mainblock *common.BlockHash

	//两笔验证其他交易的区块hash
	BlockOne *common.BlockHash
	BlockTwo *common.BlockHash

	//自身交易的输入hash
	TxInputs []TxInput

	//交易输出的地址方 第一笔是用来主块奖励给自己的 其他笔是用来发送给对方和给自己找零的
	SendTo []TxOutput

	//交易金额
	Amount int

	//难度值
	nonce *big.Int

	//目标值
	target *big.Int
}

type TxInput struct {

	//输入的区块地址
	InputAddress *common.BlockHash

	//解锁脚本以使用该区块
	decScript []byte

	//指向区块的哪笔交易
	Index int
}

type TxOutput struct {

	//输出的地址
	OutputAddress *common.BlockHash

	//加密的脚本
	encScript []byte

	//输出的金额
	Amount int
}

//全局的主块链

//var targetforTx = big.NewInt(math.MaxInt64)
//var targetforMb = big.NewInt(math.MaxInt64)
//
////全局的难度值 会有一个线程通过计算区块增长速度 来修改难度值 这里暂时设置为 前导零两个
//func Settarget(n1,n2 uint) {
//
//	targetforTx.Rsh(targetforTx, n1)
//	targetforMb.Rsh(targetforMb, n2)
//}


//初始化
//func init(){
//
//	genesis := CreateGenesisBlock()
//	mc.Settarget(2,5)
//	mc.ChainBlocks = []Block{}
//	mc.Add(genesis)
//
//
//}


func NewBlock(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []TxInput,SendTo []TxOutput,amount int,target *big.Int,nonce *big.Int,hash *common.BlockHash)*Block{
	return &Block{MinerAddress:miner,Hash:hash,Mainblock:Mainblock,BlockOne:BlockOne,BlockTwo:BlockTwo,TxInputs:TxInputs,SendTo:SendTo,Amount:amount,nonce:nonce,target:target}
}

//在调用该函数时 已经检验过交易的脚本是否满足，交易的输入输出已经构建好
func CreateBlock(miner *common.Address,Mainblock *common.BlockHash,BlockOne,BlockTwo *common.BlockHash,TxInputs []TxInput,SendTo []TxOutput,amount int,target *big.Int) *Block {

	//判断当前主块是不是最新

	latest := mc.Getlatest()
	if bytes.Equal(latest[:],Mainblock[:]){
		//如果是最新的话,验证两笔交易
		// 如果所有输入块有足够的余额用于输出 其实用户交易时已经经过检测
		if basetool.GetBalance(TxInputs) > amount && basetool.Calculate(SendTo) == amount {
				hash, nonce := consensus.Pow(miner,Mainblock, BlockOne, BlockTwo, TxInputs, SendTo, amount, target)

				block := NewBlock(miner,Mainblock, BlockOne, BlockTwo, TxInputs, SendTo, amount, target, nonce, &hash)
				return block
			}else {
				return nil
			}
	}else{
			return nil

	}

}
