#!/usr/bin/env bash
#set -x
# trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT

dir=$(pwd)
echo $dir

bin="$dir/benchmark/ycsb/ycsbtest"
defaultAddrs="127.0.0.1:50001"
loadPath="$dir/temp/ycsb_data/workloada.dat"
runPath="$dir/temp/ycsb_data/run_workloada.dat"

size=${1:-4}
clients=${2:-4} 
ndrivers=${size}
nthreads=$(( ${clients} / ${ndrivers} ))

# nthreads= 1 2 4 16 32 64
# clients="4 8 16 32 64 128 192 256"

if [ ! -f ${bin} ]; then
    echo "Binary file ${bin} not found!"
    echo "Hint: "
    echo " Please build binaries by run command: make build "
    echo "exit 1 "
    exit 1
fi


for (( c=2; c<=${size}; c++ ))
do 
defaultAddrs="${defaultAddrs},127.0.0.1:$((50000 + ${c}))"
done
echo "start test with bcdbnode addrs: ${defaultAddrs}"


$bin --load-path=$loadPath --run-path=$runPath --ndrivers=$ndrivers --nthreads=$nthreads --server-addrs=${defaultAddrs} &
#2>&1 | tee test.log 