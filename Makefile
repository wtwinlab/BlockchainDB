binaries := cmd/bcdbnode/bcdbnode benchmark/ycsb/ycsbtest storage/ethereum/contracts/deploy/deyploycontract
nodes := 4
clients := 4
shards := 1
workload := a
distribution := ycsb_data

.PHONY: all build clean download $(binaries) ethnet install verify test

all: download build ethnet install verify

fast: build redisup ethnet install verify

clean:
	@rm -fv $(binaries)

build: clean
	@go build -o ./.bin/bcdbnode $(GCFLAGS) ./cmd/bcdbnode
	@go build -o ./.bin/benchmark_bcdb $(GCFLAGS) ./benchmark/ycsb
	@go build -o ./.bin/deploy_contract $(GCFLAGS) ./storage/ethereum/contracts/deploy

download:
	@/bin/bash scripts/ycsb/gen_ycsb_data.sh
	@/bin/bash scripts/libs/get_docker_images.sh
	@go mod download

redisup:
	@/bin/bash scripts/start_redis_db.sh ${shards}
	
ethnet:
	@/bin/bash scripts/start_eth_network.sh ${shards} $(nodes)
	
install:
	@/bin/bash scripts/stop_nodes.sh
	@/bin/bash scripts/gen_config.sh ${shards} $(nodes)
	@/bin/bash scripts/start_nodes.sh ${shards} $(nodes)
verify:
	@go run cmd/tests/main.go

test:
	@echo "Test start with node size: $(nodes), client size: $(clients)"
	@/bin/bash scripts/ycsb/start_ycsb_test.sh $(nodes) $(clients) ${workload} ${distribution} >> test.$(nodes).${clients}.log 2>&1 

