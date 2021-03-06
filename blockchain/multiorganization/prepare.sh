#!/usr/bin/env bash
# 第一个参数为有几个channel
# 第二个参数为有几个org
artifactsDir=${PWD}/artifacts
if [[ -d ${artifactsDir} ]];then
    rm -rf ./artifacts/*
else
    mkdir  -p ${artifactsDir}
fi



limit=2
orgLimit=5
if [[ -n ${1} ]];then
limit=${1}
fi
if [[ -n ${2} ]];then
orgLimit=${2}
fi

echo "初始化创世块"
configtxgen  -configPath ${PWD} --profile OrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block
echo "solo 创世块"
configtxgen -configPath ${PWD} --profile OrdererSoloGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.solo.genesis.block
echo "生成channel的配置信息"

for (( i=1; i<=${limit}; i++ ))
do
    channelName=DemoChannel${i}
    txFileName=demochannel${i}.tx
    asChannelName=demochannel${i}
    configtxgen  -configPath ${PWD} --profile ${channelName}  -outputCreateChannelTx ./artifacts/${txFileName} -channelID ${asChannelName}
    for (( j=1; j<=${orgLimit}; j++ ))
    do
        orgMsp=Org${j}MSP
        anchorTxName=org${j}mspanchors.tx
        configtxgen  -configPath ${PWD} --profile ${channelName} -outputAnchorPeersUpdate ./artifacts/${anchorTxName} -channelID ${asChannelName} -asOrg ${orgMsp}
    done
done

