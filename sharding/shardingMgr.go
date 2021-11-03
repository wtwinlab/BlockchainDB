package sharding

import (
	"context"
	"fmt"
	"strings"

	EthConnector "github.com/sbip-sg/BlockchainDB/blockchainconnectors/ethereumconnector"
	FabConnector "github.com/sbip-sg/BlockchainDB/blockchainconnectors/fabricconnector"
)

type ShardingMgr struct {
	EthConn *EthConnector.EthereumConnector
	FabConn *FabConnector.FabricConnector
}

func shard(key string) string {
	if strings.HasPrefix(key, PARTITION_ETH().Key) {
		return PARTITION_ETH().Shard
	} else if strings.HasPrefix(key, PARTITION_FAB().Shard) {
		return PARTITION_FAB().Shard
	} else {
		return PARTITION_DEFAULT().Shard
	}
}

func (mgr *ShardingMgr) Read(ctx context.Context, key string) (string, error) {

	switch shard(key) {
	case PARTITION_ETH().Shard:
		return mgr.EthConn.Read(key)

	case PARTITION_FAB().Shard:
		return mgr.FabConn.Read(key)

	default:
		return "", fmt.Errorf("Error sharding key %s", key)
	}

}

func (mgr *ShardingMgr) Write(ctx context.Context, key string, value string) error {
	switch shard(key) {
	case PARTITION_ETH().Shard:
		return mgr.EthConn.Write(key, value)

	case PARTITION_FAB().Shard:
		return mgr.FabConn.Write(key, value)

	default:
		return fmt.Errorf("Error sharding key %s", key)
	}
}
