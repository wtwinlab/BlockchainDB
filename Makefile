binaries := cmd/bcdbnode/bcdbnode benchmark/ycsb/ycsbtest storage/ethereum/contracts/deploy/deyploycontract
nodes := 4
clients := 4
shards := 1
workload := a

.PHONY: all build clean download $(binaries) ethnet install test

all: download build ethnet install

fast: build ethnet install

clean:
	@rm -fv $(binaries)

build: clean $(binaries)
	
$(binaries):
	@go build -o ./$@ $(GCFLAGS) ./$(dir $@)

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
	@/bin/bash scripts/start_nodes.sh ${shards} $(nodes) > server.${shards}.$(nodes).log 2>&1 && cat server.${shards}.$(nodes).log

test:
	@echo "Test start with node size: $(nodes), client size: $(clients)"
	@/bin/bash scripts/ycsb/start_ycsb_test.sh $(nodes) $(clients) ${workload} > test.$(nodes).${clients}.log 2>&1 && cat test.$(nodes).${clients}.log

