package service

import (
	"context"
	"log"

	"github.com/sbip-sg/BlockchainDB/node/config"
	pbv "github.com/sbip-sg/BlockchainDB/proto/blockchaindb"
	"github.com/sbip-sg/BlockchainDB/sharding"
)

var _ pbv.BCdbNodeServer = (*ServerNode)(nil)

type ServerNode struct {
	shardingMgr *sharding.ShardingMgr
}

func NewServerNode(conf *config.Options) (*ServerNode, error) {

	// ethereumconn, err := EthClientSDK.NewEthereumKVStoreInstance(conf.EthNode, conf.EthHexAddr, conf.EthHexKey)
	// if err != nil {
	// 	log.Println("Failed to NewEthereumKVStoreInstance", err)
	// 	return nil, err
	// }
	// fabricconn, err := FabClientSDK.NewFabricKVStoreInstance()
	// if err != nil {
	// 	fmt.Println("Failed to NewFabricKVStoreInstance", err)
	// 	return nil, err
	// }
	// shamgr := &sharding.ShardingMgr{EthConn: ethereumconn, FabConn: fabricconn}
	shamgr, err := sharding.NewShardingMgr(conf)
	if err != nil {
		log.Println("Failed to NewShardingMgr", err)
		return nil, err
	}
	return &ServerNode{shardingMgr: shamgr}, nil
}

func (sv *ServerNode) Get(ctx context.Context, req *pbv.GetRequest) (*pbv.GetResponse, error) {
	val, err := sv.shardingMgr.Read(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pbv.GetResponse{Value: string(val)}, nil
}

func (sv *ServerNode) Set(ctx context.Context, req *pbv.SetRequest) (*pbv.SetResponse, error) {
	// Use serverclient instance to set
	err := sv.shardingMgr.Write(ctx, req.GetKey(), req.GetValue())
	if err != nil {
		return nil, err
	}
	return &pbv.SetResponse{}, nil
}
