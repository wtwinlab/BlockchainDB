binaries := cmd/bcdbnode/bcdbnode benchmark/ycsb/ycsbtest
nodes := 1
clients := 1
shards := 1

.PHONY: all build clean download $(binaries) init generate install test

all: $(binaries) generate init install

build: $(binaries)

$(binaries):
	@go build -o ./$@ $(GCFLAGS) ./$(dir $@)

download:
	@/bin/bash scripts/gen_ycsb_data.sh

clean:
	@rm -fv $(binaries)

generate:
	@/bin/bash scripts/eth/gen_eth_config.sh ${shards} $(nodes)
	@/bin/bash scripts/gen_config.sh ${shards} $(nodes)

init:
	@/bin/bash scripts/start_redis_db.sh ${shards}
	@/bin/bash scripts/eth/start_eth_node.sh $(nodes)
	
install:
	@/bin/bash scripts/stop_nodes.sh
	@/bin/bash scripts/start_nodes.sh $(nodes) > server.${shards}.$(nodes).log 2>&1 && cat server.${shards}.$(nodes).log

test:
	@echo "Test start with node size: $(nodes), client size: $(clients)"
	@/bin/bash scripts/start_ycsb_test.sh $(nodes) $(clients)

