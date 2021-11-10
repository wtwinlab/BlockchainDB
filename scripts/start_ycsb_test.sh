#!/usr/bin/env bash
#set -x
# trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT

dir=$(pwd)
echo $dir

bin="$dir/benchmark/ycsb/ycsbtest"
defaultAddrs="127.0.0.1:50001"
loadPath="$dir/temp/ycsb_data/workloada.dat"
runPath="$dir/temp/ycsb_data/run_workloada.dat"

size=${1:-1}
ndrivers=${2:-4}
#THREADS="4 8 16 32 64 128 192 256"
nthreads=${3:-4}


for (( c=2; c<=${size}; c++ ))
do 
defaultAddrs="${defaultAddrs},127.0.0.1:$((50000 + ${c}))"
done
echo "start test with bcdbnode addrs: ${defaultAddrs}"


$bin --load-path=$loadPath --run-path=$runPath --ndrivers=$ndrivers --nthreads=$nthreads --server-addrs=${defaultAddrs} &
