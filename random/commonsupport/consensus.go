// Copyright (c) 2018-2019 The MATRIX Authors
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php
package commonsupport

import (
	"fmt"
	"github.com/MatrixAINetwork/go-matrix"
	"github.com/MatrixAINetwork/go-matrix/baseinterface"
	"github.com/MatrixAINetwork/go-matrix/ca"
	"github.com/MatrixAINetwork/go-matrix/common"
	"github.com/MatrixAINetwork/go-matrix/common/readstatedb"
	"github.com/MatrixAINetwork/go-matrix/core/vm"
	"github.com/MatrixAINetwork/go-matrix/log"
	"github.com/MatrixAINetwork/go-matrix/mc"
	"math/big"
)

const (
	ModeleRandomCommon = "随机数服务公共组件"
)

func GetElectGenTimes(support matrix.StateReader, height uint64) (*mc.ElectGenTimeStruct, error) {
	return readstatedb.GetElectGenTimes(support, height)
}

func getRandomInfo(hash common.Hash, support baseinterface.RandomChainSupport) (*mc.RandomInfoStruct, error) {
	height, err := GetNumberByHash(hash, support)
	if err != nil {
		log.Error(ModeleRandomCommon, "获取随机数信息失败", err, "hash", hash.String())
		return nil, fmt.Errorf("获取随机数信息失败 %v", hash)
	}
	randonInfo, err := readstatedb.GetRandomInfo(support.BlockChain(), height)
	if err != nil {
		log.Error(ModeleRandomCommon, "获取随机数信息失败,从状态树获取信息失败 err", err)
		return nil, fmt.Errorf("获取随机数信息失败,从状态树获取信息失败 %v", err)
	}
	if randonInfo == nil {
		log.Error(ModeleRandomCommon, "获取随机数信息失败", "获取到的信息为空")
		return nil, fmt.Errorf("获取随机数信息失败,获取到的信息为空")
	}
	return randonInfo, nil
}

func GetMinHash(hash common.Hash, support baseinterface.RandomChainSupport) (common.Hash, error) {
	randomInfo, err := getRandomInfo(hash, support)
	if err != nil {
		log.Error(ModeleRandomCommon, "获取最小hash阶段 err", err)
		return common.Hash{}, fmt.Errorf("获取最小hash阶段 %v", err)
	}
	return randomInfo.MinHash, nil
}
func GetMaxNonce(hash common.Hash, support baseinterface.RandomChainSupport) (uint64, error) {
	randomInfo, err := getRandomInfo(hash, support)
	if err != nil {
		log.Error(ModeleRandomCommon, "获取最大Nonce阶段 err", err)
		return 0, fmt.Errorf("获取最大Nonce阶段 %v", err)
	}
	return randomInfo.MaxNonce, nil
}

func GetDepositListByHeightAndRole(height *big.Int, role common.RoleType) ([]vm.DepositDetail, error) {
	return ca.GetElectedByHeightAndRole(height, role)
}

func GetSelfAddress() common.Address {
	return ca.GetDepositAddress()
}
