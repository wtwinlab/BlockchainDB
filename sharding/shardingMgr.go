package sharding

import (
	"strings"

	BlockchainConnector "github.com/sbip-sg/BlockchainDB/connectors"
)

type ShardingMgr struct {
	ethConn *BlockchainConnector.EthereumConnector
	fabConn *BlockchainConnector.FabricConnector
}

func shard(key string) string {
	if strings.HasPrefix(key, "eth") {
		return "ethereum"
	} else if strings.HasPrefix(key, "fab") {
		return "fabric"
	} else {
		return "default"
	}
}

func (mgr *ShardingMgr) Read(key string) {

	switch shard(key) {
	case "ethereum":
		mgr.ethConn.Read(key)
		return
	case "fabric":
		mgr.fabConn.Read(key)
		return
	default:
		return
	}

}

func (mgr *ShardingMgr) Write(key string, value string) {
	switch shard(key) {
	case "ethereum":
		mgr.ethConn.Write(key, value)
		return
	case "fabric":
		mgr.fabConn.Write(key, value)
		return
	default:
		return
	}
}
