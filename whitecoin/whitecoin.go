
/*
 * Copyright 2019 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package whitecoin

import (
	"fmt"
	"math/big"

	xwc "github.com/Assetsadapter/whitecoin-adapter/libs/config"
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/shopspring/decimal"
)

const DECIMAL_COUNT = 8

//CurveType 曲线类型
func (wm *WalletManager) CurveType() uint32 {
	return wm.Config.CurveType
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "whitecoin"
}

//Symbol 币种标识
func (wm *WalletManager) Symbol() string {
	return wm.Config.Symbol
}

//Decimal 小数位精度
func (wm *WalletManager) Decimal() int32 {
	return DECIMAL_COUNT
}

//BalanceModelType 余额模型类型
func (wm *WalletManager) BalanceModelType() openwallet.BalanceModelType {
	return openwallet.BalanceModelTypeAddress
}

//GetAddressDecode 地址解析器
func (wm *WalletManager) GetAddressDecode() openwallet.AddressDecoder {
	return wm.Decoder
}

//GetTransactionDecoder 交易单解析器
func (wm *WalletManager) GetTransactionDecoder() openwallet.TransactionDecoder {
	return wm.TxDecoder
}

//GetBlockScanner 获取区块链
func (wm *WalletManager) GetBlockScanner() openwallet.BlockScanner {
	return wm.Blockscanner
}

//LoadAssetsConfig 加载外部配置
func (wm *WalletManager) LoadAssetsConfig(c config.Configer) error {

	wm.Config.ServerAPI = c.String("serverAPI")
	wm.Config.ChainID = c.String("ChainID")
	wm.Config.FixFees, _ = c.Int64("FixFees")
	wm.Api = NewWalletClient(wm.Config.ServerAPI, wm.Config.ServerAPI, false)
	wm.Config.DataDir = c.String("dataDir")

	//数据文件夹
	wm.Config.makeDataDir()

	//添加XWC的chain环境
	xwc.Add(xwc.ChainConfig{
		Name:      wm.FullName(),
		CoreAsset: wm.Symbol(),
		Prefix:    wm.Symbol(),
		ID:        wm.Config.ChainID,
	})

	xwc.SetCurrent(wm.Config.ChainID)

	return nil
}

//InitAssetsConfig 初始化默认配置
func (wm *WalletManager) InitAssetsConfig() (config.Configer, error) {
	return config.NewConfigData("ini", []byte(wm.Config.DefaultConfig))
}

//GetAssetsLogger 获取资产账户日志工具
func (wm *WalletManager) GetAssetsLogger() *log.OWLogger {
	return wm.Log
}

//GetSmartContractDecoder 获取智能合约解析器
func (wm *WalletManager) GetSmartContractDecoder() openwallet.SmartContractDecoder {
	return wm.ContractDecoder
}

//GetAddressDecode 地址解析器
//如果实现了AddressDecoderV2，就无需实现AddressDecoder
func (wm *WalletManager) GetAddressDecoderV2() openwallet.AddressDecoderV2 {
	return wm.DecoderV2
}

// 从区块链中的值转换为系统内显示的小数值
func ConvertToFloat(amount string) (decimal.Decimal, error) {
	return ConvertAmountToFloat(amount, DECIMAL_COUNT)
}

func ConvertAmountToFloat(amount string, decimals int) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(amount)
	if err != nil {
		log.Errorf("convert string to decimal failed, err = %v", err)
		return d, err
	}

	dAmount := d.Shift(-int32(decimals))
	return dAmount, nil
}

// 从系统内的小数值转换为区块链中的整数值
func ConvertToDecimal(amount string) (*big.Int, error) {
	return ConvertFloatStringToBigInt(amount, DECIMAL_COUNT)
}

func ConvertFloatStringToBigInt(amount string, decimals int) (*big.Int, error) {
	vDecimal, _ := decimal.NewFromString(amount)
	if decimals < 0 || decimals > 30 {
		return nil, fmt.Errorf("wrong decimal input through")
	}

	decimalInt := big.NewInt(1)
	for i := 0; i < decimals; i ++ {
		decimalInt.Mul(decimalInt,big.NewInt(10))
	}
	d, _ := decimal.NewFromString(decimalInt.String())
	vDecimal = vDecimal.Mul(d)
	rst:= new(big.Int)
	if _, valid := rst.SetString(vDecimal.String(), 10); !valid {
		log.Error("convert to big.int failed")
		return nil, fmt.Errorf("convert to big.int failed")
	}
	return rst, nil
}