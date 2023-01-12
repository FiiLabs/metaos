#!/bin/sh
Home=./testnet
ChainID=testnet # chain-id
ChainCMD=./build/metaosd
NodeName=metaos-node # node name
NodeIP=(tcp://127.0.0.1 tcp://127.0.0.1 tcp://127.0.0.1 tcp://127.0.0.1)
NodeNames=("node0" "node1" "node2" "node3")
NodeDic=("${Home}/node0" "${Home}/node1" "${Home}/node2" "${Home}/node3")
Mnemonics=("eagle marriage host height topple sorry exist nation screen affair bulk average medal flush candy alert amused alone hire clerk treat hybrid tip cake"
"width clap suspect squeeze rich exact lawn output play blanket join join measure charge they sword wheat light federal review true portion add rival"
"satisfy web truck wink canal use decrease glove glow skill always script differ speed eternal close today slow grass disorder robot match face consider"
"assist cute perfect during kiwi vacant marble happy smooth now isolate social birth maid just mixture federal pause ridge midnight picture cattle document inner"
"setup capital exact dad minimum pigeon blush claw cake find animal torch cry guide dirt settle parade host grief lunar indicate laptop bulk cherry"
)
Validators=("validator0" "validator1" "validator2" "validator3")
Stake=umeta
TotalStake=10000000000000000${Stake} # total stake in genesis
SendStake=10000000000000${Stake}
DataPath=/tmp

Point=upoint
PointOwner=metaos1g6gqr3s58dhw3jq5hm95qrng0sa9um7g2e2fcx # replace with actual address
PointToken=`echo {\"symbol\": \"point\", \"name\": \"Metaos point native token\", \"scale\": 6, \"min_unit\": \"upoint\", \"initial_supply\": \"1000000000\", \"max_supply\": \"1000000000000\", \"mintable\": true, \"owner\": \"${PointOwner}\"}`

rm -rf $Home
$ChainCMD keys delete admin -y

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do $ChainCMD keys delete ${Validators[$i]} -y; done

bash -c "echo -e \"${Mnemonics[4]}\n12345678\n12345678\" | ${ChainCMD} keys add admin --recover"

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do bash -c "echo -e \"${Mnemonics[$i]}\n12345678\n12345678\" | ${ChainCMD} keys add ${Validators[$i]} --recover --home=${NodeDic[$i]}"; done

#for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do mkdir ${NodeDic[$i]}; done
$ChainCMD init moniker --chain-id $ChainID --home=$Home
for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do $ChainCMD init moniker --chain-id $ChainID --home=${NodeDic[$i]}; done
for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do $ChainCMD genkey --out-file ${NodeDic[$i]}/priv_validator.pem --home=${NodeDic[$i]}; done
for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do $ChainCMD genkey --type node --out-file ${NodeDic[$i]}/priv_node.pem --home=${NodeDic[$i]}; done


sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $Home/config/config.toml

sed -i 's/addr_book_strict = true/addr_book_strict = false/' $Home/config/config.toml

sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/' $Home/config/config.toml

sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' $Home/config/config.toml

sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "1umeta"/' $Home/config/app.toml

#sed -i 's/filter_peers = false/filter_peers = true/' $Home/config/config.toml

sed -i "s/stake/$Stake/g" $Home/config/genesis.json

sed -i "s/\"point_token_denom\": \"$Stake\"/\"point_token_denom\": \"$Point\"/g" $Home/config/genesis.json

sed -i "s/node0token/$Point/g" $Home/config/genesis.json

sed -i "s/\"base_denom\": \"$Stake\"/\"base_denom\": \"$Point\"/g" $Home/config/genesis.json

sed -i "s/\"restricted_service_fee_denom\": false/\"restricted_service_fee_denom\": true/g" $Home/config/genesis.json

cat $Home/config/genesis.json | jq ".app_state.service.params.min_deposit[0].denom = \"$Point\"" > $Home/temp; cat $Home/temp; cp -f $Home/temp $Home/config/genesis.json

cat $Home/config/genesis.json | jq ".app_state.token.tokens |= . + [$PointToken]" > $Home/temp; cat $Home/temp; cp -f $Home/temp $Home/config/genesis.json

sed -i "s/\"base_token_manager\": \"\"/\"base_token_manager\": \"$(echo 12345678 | $ChainCMD keys show validator0 | grep address | cut -b 12-)\"/" $Home/config/genesis.json

sed -i "s/\"token_tax_rate\": \"0.400000000000000000\"/\"token_tax_rate\": \"1\"/g" $Home/config/genesis.json

sed -i "s/\"denom\": \"meta\"/\"denom\": \"point\"/g" $Home/config/genesis.json

sed -i "s/\"amount\": \"60000000000\"/\"amount\": \"60000\"/g" $Home/config/genesis.json

sed -i "s/\"amount\": \"1000000000\"/\"amount\": \"1000000000000000\"/g" $Home/config/genesis.json

sed -i "s/\"amount\": \"500000000\"/\"amount\": \"1000000000000000\"/g" $Home/config/genesis.json

sed -i "s/\"amount\": \"1000\"/\"amount\": \"1000000000\"/g" $Home/config/genesis.json

sed -i "s/\"amount\": \"5000\"/\"amount\": \"5000000000\"/g" $Home/config/genesis.json

sed -i "s/nodes\": \[/nodes\": \[{\"id\": \"$($ChainCMD tendermint show-node-id --home=$Home)\", \"name\": \"$NodeName\"}/" $Home/config/genesis.json

bash -c "$ChainCMD add-genesis-account \$(echo 12345678 | $ChainCMD keys show validator0 -a --home=$Home) ${TotalStake} --root-admin --home=$Home"

bash -c "$ChainCMD add-genesis-account ${PointOwner} 1000000000000000${Point} --home=$Home"

openssl ecparam -genkey -name SM2 -out $Home/root.key

echo -e "CN\nSH\nSH\nIT\nDEV\n'${NodeNames[0]}'\n\n" | openssl req -new -x509 -sm3 -sigopt "distid:1234567812345678" -key $Home/root.key -out $Home/root.crt -days 3650

$ChainCMD set-root-cert $Home/root.crt --home=$Home


for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do echo -e "CN\nSH\nSH\nIT\nDEV\n'${NodeNames[$i]}'\n\n\n\n" | openssl req -new -key ${NodeDic[$i]}/priv_validator.pem -out ${NodeDic[$i]}/validator_req.csr -sm3 -sigopt "distid:1234567812345678"; done

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do openssl x509 -req -in ${NodeDic[$i]}/validator_req.csr -out ${NodeDic[$i]}/validator.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA $Home/root.crt -CAkey $Home/root.key -CAcreateserial; done

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do echo -e "CN\nSH\nSH\nIT\nDEV\n'${NodeNames[$i]}'\n\n\n\n" | openssl req -new -key ${NodeDic[$i]}/priv_node.pem -out ${NodeDic[$i]}/node_req.csr -sm3 -sigopt "distid:1234567812345678"; done

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do openssl x509 -req -in ${NodeDic[$i]}/node_req.csr -out ${NodeDic[$i]}/node.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA $Home/root.crt -CAkey $Home/root.key -CAcreateserial; done

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do bash -c "echo 12345678 | $ChainCMD add-genesis-validator --name ${NodeNames[i]} --cert ${NodeDic[i]}/validator.crt --power 10000 --from ${Validators[i]} --home=$Home"; done


#sed -i 's/seed_mode = false/seed_mode = true/' $Home/config/config.toml
#sed -i "s/seeds = \"\"/seeds = \"$($ChainCMD tendermint show-node-id --home=${NodeDic[0]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:26656,$($ChainCMD tendermint show-node-id --home=${NodeDic[1]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:36656,$($ChainCMD tendermint show-node-id --home=${NodeDic[2]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:46656,$($ChainCMD tendermint show-node-id --home=${NodeDic[3]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:56656\"/" $Home/config/config.toml
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$($ChainCMD tendermint show-node-id --home=${NodeDic[0]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:26656,$($ChainCMD tendermint show-node-id --home=${NodeDic[1]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:36656,$($ChainCMD tendermint show-node-id --home=${NodeDic[2]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:46656,$($ChainCMD tendermint show-node-id --home=${NodeDic[3]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:56656\"/" $Home/config/config.toml
#sed -i "s/persistent_peers = \"\"/persistent_peers = \"$($ChainCMD tendermint show-node-id --home=${NodeDic[0]} | sed 's/\^M\$//')@`echo ${NodeIP[0]} | awk -F // '{print $2}'`:26656\"/" $Home/config/config.toml

echo $($ChainCMD tendermint show-node-id --home=${NodeDic[0]} | sed 's/\^M\$//')
echo $($ChainCMD tendermint show-node-id --home=${NodeDic[1]} | sed 's/\^M\$//')
echo $($ChainCMD tendermint show-node-id --home=${NodeDic[2]} | sed 's/\^M\$//')
echo $($ChainCMD tendermint show-node-id --home=${NodeDic[3]} | sed 's/\^M\$//')

for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do cp $Home/config/config.toml ${NodeDic[$i]}/config; done
for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do cp $Home/config/genesis.json ${NodeDic[$i]}/config; done
for i in `seq 0 $[ ${#Validators[*]} -1 ]`; do cp $Home/config/app.toml ${NodeDic[$i]}/config; done

sed -i 's/address = "tcp:\/\/0.0.0.0:1317"/address = "tcp:\/\/0.0.0.0:31317"/' ${NodeDic[1]}/config/app.toml
sed -i 's/address = ":8080"/address = ":38080"/' ${NodeDic[1]}/config/app.toml
sed -i 's/address = "0.0.0.0:9090"/address = "0.0.0.0:39090"/' ${NodeDic[1]}/config/app.toml
sed -i 's/address = "0.0.0.0:9091"/address = "0.0.0.0:39091"/' ${NodeDic[1]}/config/app.toml
sed -i 's/address = "0.0.0.0:8545"/address = "0.0.0.0:38545"/' ${NodeDic[1]}/config/app.toml
sed -i 's/ws-address = "0.0.0.0:8546"/ws-address = "0.0.0.0:38546"/' ${NodeDic[1]}/config/app.toml
sed -i 's/proxy_app = "tcp:\/\/127.0.0.1:26658"/proxy_app = "tcp:\/\/127.0.0.1:36658"/' ${NodeDic[1]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26657"/proxy_app = "tcp:\/\/0.0.0.0:36657"/' ${NodeDic[1]}/config/config.toml
sed -i 's/pprof_laddr = "localhost:6060"/pprof_laddr = "localhost:36060"/' ${NodeDic[1]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26656"/pprof_laddr = "tcp:\/\/0.0.0.0:36656"/' ${NodeDic[1]}/config/config.toml
sed -i 's/prometheus_listen_addr = ":26660"/prometheus_listen_addr = ":36660"/' ${NodeDic[1]}/config/config.toml

sed -i 's/address = "tcp:\/\/0.0.0.0:1317"/address = "tcp:\/\/0.0.0.0:41317"/' ${NodeDic[2]}/config/app.toml
sed -i 's/address = ":8080"/address = ":48080"/' ${NodeDic[2]}/config/app.toml
sed -i 's/address = "0.0.0.0:9090"/address = "0.0.0.0:49090"/' ${NodeDic[2]}/config/app.toml
sed -i 's/address = "0.0.0.0:9091"/address = "0.0.0.0:49091"/' ${NodeDic[2]}/config/app.toml
sed -i 's/address = "0.0.0.0:8545"/address = "0.0.0.0:48545"/' ${NodeDic[2]}/config/app.toml
sed -i 's/ws-address = "0.0.0.0:8546"/ws-address = "0.0.0.0:48546"/' ${NodeDic[2]}/config/app.toml
sed -i 's/proxy_app = "tcp:\/\/127.0.0.1:26658"/proxy_app = "tcp:\/\/127.0.0.1:46658"/' ${NodeDic[2]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26657"/proxy_app = "tcp:\/\/0.0.0.0:46657"/' ${NodeDic[2]}/config/config.toml
sed -i 's/pprof_laddr = "localhost:6060"/pprof_laddr = "localhost:46060"/' ${NodeDic[2]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26656"/pprof_laddr = "tcp:\/\/0.0.0.0:46656"/' ${NodeDic[2]}/config/config.toml
sed -i 's/prometheus_listen_addr = ":26660"/prometheus_listen_addr = ":46660"/' ${NodeDic[2]}/config/config.toml

sed -i 's/address = "tcp:\/\/0.0.0.0:1317"/address = "tcp:\/\/0.0.0.0:51317"/' ${NodeDic[3]}/config/app.toml
sed -i 's/address = ":8080"/address = ":58080"/' ${NodeDic[3]}/config/app.toml
sed -i 's/address = "0.0.0.0:9090"/address = "0.0.0.0:59090"/' ${NodeDic[3]}/config/app.toml
sed -i 's/address = "0.0.0.0:9091"/address = "0.0.0.0:59091"/' ${NodeDic[3]}/config/app.toml
sed -i 's/address = "0.0.0.0:8545"/address = "0.0.0.0:58545"/' ${NodeDic[3]}/config/app.toml
sed -i 's/ws-address = "0.0.0.0:8546"/ws-address = "0.0.0.0:58546"/' ${NodeDic[3]}/config/app.toml
sed -i 's/proxy_app = "tcp:\/\/127.0.0.1:26658"/proxy_app = "tcp:\/\/127.0.0.1:56658"/' ${NodeDic[3]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26657"/proxy_app = "tcp:\/\/0.0.0.0:56657"/' ${NodeDic[3]}/config/config.toml
sed -i 's/pprof_laddr = "localhost:6060"/pprof_laddr = "localhost:56060"/' ${NodeDic[3]}/config/config.toml
sed -i 's/laddr = "tcp:\/\/0.0.0.0:26656"/pprof_laddr = "tcp:\/\/0.0.0.0:56656"/' ${NodeDic[3]}/config/config.toml
sed -i 's/prometheus_listen_addr = ":26660"/prometheus_listen_addr = ":56660"/' ${NodeDic[3]}/config/config.toml

$ChainCMD start --pruning=nothing --home=${NodeDic[0]}
