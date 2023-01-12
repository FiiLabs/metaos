# Environmental requirements：

1. go

2. openssl

```
git clone -b openssl-3.0.0-alpha4 https://github.com/openssl/openssl.git

cd openssl && ./config

sudo make install
```

3.intall jq(https://stedolan.github.io/jq/download/), sed



# compile and run

######  1.compile

```
make build
```

2. ###### run 

   local single node

```
bash ./scripts/single-node.sh
```

​	  local four node

```
terminal 1,cmd:
	bash ./scripts/four-node.sh
	
terminal 2,cmd:
  ./build/metaosd start  --pruning=nothing --home=./testnet/node1 --rpc.laddr=tcp://0.0.0.0:36657 --p2p.laddr=tcp://0.0.0.0:36656
  
terminal 3,cmd:
  ./build/metaosd start  --pruning=nothing --home=./testnet/node2 --rpc.laddr=tcp://0.0.0.0:46657 --p2p.laddr=tcp://0.0.0.0:46656

terminal 4,cmd:
./build/metaosd start  --pruning=nothing --home=./testnet/node3 --rpc.laddr=tcp://0.0.0.0:56657 --p2p.laddr=tcp://0.0.0.0:56656
```

run by docker

```
./scripts/four_validators_local.sh docker
```

