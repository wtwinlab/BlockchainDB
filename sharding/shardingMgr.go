package sharding

import (
	"context"
	"hash/adler32"
	"log"
	"strconv"
	"strings"

	Connectors "github.com/sbip-sg/BlockchainDB/blockchainconnectors"

	//EthConnector "github.com/sbip-sg/BlockchainDB/blockchainconnectors/ethereumconnector"
	//FabConnector "github.com/sbip-sg/BlockchainDB/blockchainconnectors/fabricconnector"
	"github.com/sbip-sg/BlockchainDB/node/config"
	EthClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
	FabClientSDK "github.com/sbip-sg/BlockchainDB/storage/fabric/clientSDK"
)

type ShardingMgr struct {
	Shards      map[string]Connectors.BlockchainConnector
	Conf        map[string]config.Shard
	ShardNumber int
	// BCConn Connectors.BlockchainConnector
	// EthConn *EthConnector.EthereumConnector
	// FabConn *FabConnector.FabricConnector
}

func NewShardingMgr(conf *config.Options) (*ShardingMgr, error) {
	var shards map[string]Connectors.BlockchainConnector
	var confs map[string]config.Shard
	for _, shard := range conf.Shards {
		switch shard.Type {
		case PARTITION_ETH().Shard:
			ethconn, err := EthClientSDK.NewEthereumKVStoreInstance(shard.EthNode, shard.EthHexAddr, shard.EthHexKey)
			if err != nil {
				log.Println("Failed to NewEthereumKVStoreInstance", err)
				break
			}
			shards[shard.ID] = ethconn
			confs[shard.ID] = shard
			log.Println("Sucess NewEthereumKVStoreInstance for shard ", shard.ID)
		case PARTITION_FAB().Shard:
			fabconn, err := FabClientSDK.NewFabricKVStoreInstance()
			if err != nil {
				log.Println("Failed to NewFabricKVStoreInstance", err)
				break
			}
			shards[shard.ID] = fabconn
			confs[shard.ID] = shard
			log.Println("Sucess NewFabricKVStoreInstance for shard ", shard.ID)
		default:
			log.Println("Error sharding key", shard.ID)
			break
		}
	}
	// bcConn, err := EthClientSDK.NewEthereumKVStoreInstance(conf.EthNode, conf.EthHexAddr, conf.EthHexKey)
	// if err != nil {
	// 	log.Println("Failed to NewEthereumKVStoreInstance", err)
	// 	return nil, err
	// }
	return &ShardingMgr{Shards: shards, Conf: confs, ShardNumber: conf.ShardNumber}, nil
}

func (mgr *ShardingMgr) partitionScheme(key string) string {

	partitionId := hash(key) % mgr.ShardNumber

	return strconv.Itoa(partitionId)
}

func hash(data string) int {
	// 1. adler32
	sum := int(adler32.Checksum([]byte(data)))

	// 2. fnv
	// algorithm := fnv.New32a()
	// algorithm.Write([]byte(data))
	// sum := int(algorithm.Sum32())
	return sum
}

func partition(key string) string {
	if strings.HasPrefix(key, PARTITION_ETH().Key) {
		return PARTITION_ETH().Shard
	} else if strings.HasPrefix(key, PARTITION_FAB().Shard) {
		return PARTITION_FAB().Shard
	} else {
		return PARTITION_DEFAULT().Shard
	}
}

func (mgr *ShardingMgr) Read(ctx context.Context, key string) (string, error) {

	// switch partition(key) {
	// case PARTITION_ETH().Shard:
	// 	return mgr.EthConn.Read(key)

	// case PARTITION_FAB().Shard:
	// 	return mgr.FabConn.Read(key)

	// default:
	// 	return "", fmt.Errorf("Error sharding key %s", key)
	// }
	partitionkey := mgr.partitionScheme(key)
	return mgr.Shards[partitionkey].Read(key)

}

func (mgr *ShardingMgr) Write(ctx context.Context, key string, value string) (string, error) {
	// switch partition(key) {
	// case PARTITION_ETH().Shard:
	// 	return mgr.EthConn.Write(key, value)

	// case PARTITION_FAB().Shard:
	// 	return mgr.FabConn.Write(key, value)

	// default:
	// 	return fmt.Errorf("Error sharding key %s", key)
	// }
	partitionkey := mgr.partitionScheme(key)
	return mgr.Shards[partitionkey].Write(key, value)
}
