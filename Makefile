binaries := cmd/bcdbnode/bcdbnode benchmark/ycsb/ycsbtest
nodes := 1
clients := 1

.PHONY: all build clean download $(binaries) init generate install test

all: $(binaries) init generate install

build: $(binaries)

$(binaries):
	@go build -o ./$@ $(GCFLAGS) ./$(dir $@)

download:
	@/bin/bash scripts/gen_ycsb_data.sh

clean:
	@rm -fv $(binaries)

init:
	@/bin/bash scripts/start_eth_node.sh $(nodes)

generate:
	@/bin/bash scripts/gen_config.sh $(nodes)
	
install:
	@/bin/bash scripts/stop_nodes.sh
	@/bin/bash scripts/start_nodes.sh $(nodes) > server.log 2>&1 && cat server.log

test:
	@echo "Test start with node size: $(nodes), client size: $(clients)"
	@/bin/bash scripts/start_ycsb_test.sh $(nodes) $(clients)

