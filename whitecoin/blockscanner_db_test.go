/*
 * Copyright 2018 The openwallet Authors
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

//func TestSaveLocalBlockHead(t *testing.T) {
//	height := uint32(50775550)
//	hash := "0306c5fe30a97ad125c5edb3f8bf027f777f2f0b5cf34fba31974375f82ce08a"
//	wm := testNewWalletManager()
//
//	err := wm.Blockscanner.SaveLocalBlockHead(height, hash)
//	if err != nil {
//		t.Error(err)
//	}
//}
//
////GetLocalNewBlock 获取本地记录的区块高度和hash
//func TestGetLocalNewBlock(t *testing.T) {
//	height := uint32(50987586)
//	hash := "030a0242b4fe359d247dd531a026d44dba62fdba7e38c765d86377ae7f68fc18"
//	wm := testNewWalletManager()
//
//	wm.Blockscanner.SaveLocalBlockHead(height, hash)
//
//	_height, _hash, err := wm.Blockscanner.GetLocalBlockHead()
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	if _height != height || _hash != hash {
//		t.Errorf("Expecting local height: %v and hash: %v, but returned height: %v and hash: %v.", height, hash, _height,