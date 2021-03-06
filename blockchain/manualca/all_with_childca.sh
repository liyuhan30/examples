#!/usr/bin/env bash

# 结合了 ca链的脚本
# FIXME : 如何使用了中间证书的话,tls 需要更改,tls 应该是tlsintermediatecerts下的文件作为ca.crt,所以 关于 ${DOMAIN}下的tlsca下的文件
# 不能直接使用 caOrganization下的xxx-ca 的ca-cert.pem或者ca-chain.pem代替

./cleanAll.sh

export GOPATH=/Users/joker/go
export DIRECTORY_NAME=/examples/blockchain/manualca && \
export DOMAIN=demo.com && \
export CERTIFICATE_DOMAIN=demo-com && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/{artifacts,crypto-config,network} && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/{caOrganizations,ordererOrganizations,peerOrganizations} && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/{tls-ca,order-ca,org-ca} && \
./start_all_ts.sh && \

sleep 5 && \

fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:4052  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name orderer0.${DOMAIN} --id.secret order0pw --id.type order --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name orderer1.${DOMAIN} --id.secret order1pw --id.type order --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name orderer2.${DOMAIN} --id.secret order2pw --id.type order --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name orderer3.${DOMAIN} --id.secret order3pw --id.type order --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name orderer4.${DOMAIN} --id.secret order4pw --id.type order --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name peer1.${DOMAIN} --id.secret peer1pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true "' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \
fabric-ca-client register -d --id.name peer2.${DOMAIN} --id.secret peer2pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true "' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4052 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/admin/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem && \


mkdir -p  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/{ca,orderers,tlsca,users}/ && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/order-ca/ca-chain.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \

fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/cacerts/*.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/signcerts/*.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/server.key  && \
echo  \
'NodeOUs:
    Enable: true
    ClientOUIdentifier:
      Certificate: cacerts/0-0-0-0-4053.pem
      OrganizationalUnitIdentifier: client
    PeerOUIdentifier:
      Certificate: cacerts/0-0-0-0-4053.pem
      OrganizationalUnitIdentifier: peer
    AdminOUIdentifier:
      Certificate: cacerts/0-0-0-0-4053.pem
      OrganizationalUnitIdentifier: admin
    OrdererOUIdentifier:
      Certificate: cacerts/0-0-0-0-4053.pem
      OrganizationalUnitIdentifier: orderer ' > ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml && \

fabric-ca-client register -d --id.name orderer0.${DOMAIN} --id.secret orderer0pw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client register -d --id.name orderer1.${DOMAIN} --id.secret orderer1pw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client register -d --id.name orderer2.${DOMAIN} --id.secret orderer2pw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client register -d --id.name orderer3.${DOMAIN} --id.secret orderer3pw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client register -d --id.name orderer4.${DOMAIN} --id.secret orderer4pw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' --id.attrs '"hf.IntermediateCA=true"' -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

fabric-ca-client register -d --id.name Admin@${DOMAIN} --id.secret adminpw --id.type admin --id.attrs "hf.Registrar.Roles=admin,hf.Registrar.Attributes=*,hf.Revoker=true,hf.GenCRL=true,admin=true:ecert,abac.init=true:ecert,hf.IntermediateCA=true" -u https://0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

mkdir -p  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/{orderer0.${DOMAIN},orderer1.${DOMAIN},orderer2.${DOMAIN},orderer3.${DOMAIN},orderer4.${DOMAIN}} && \

fabric-ca-client enroll -d -u https://orderer0.${DOMAIN}:orderer0pw@0.0.0.0:4053  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer1.${DOMAIN}:orderer1pw@0.0.0.0:4053  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer2.${DOMAIN}:orderer2pw@0.0.0.0:4053  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer3.${DOMAIN}:orderer3pw@0.0.0.0:4053  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer4.${DOMAIN}:orderer4pw@0.0.0.0:4053  -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

fabric-ca-client enroll -d -u https://orderer0.${DOMAIN}:order0pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts orderer0.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer1.${DOMAIN}:order1pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts orderer1.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer2.${DOMAIN}:order2pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts orderer2.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer3.${DOMAIN}:order3pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts orderer3.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \
fabric-ca-client enroll -d -u https://orderer4.${DOMAIN}:order4pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts orderer4.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/tlsca/ca.crt && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/tlscacerts/*.pem  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/server.key && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/tlscacerts/*.pem  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/server.key && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/tlscacerts/*.pem  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/server.key && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/tlscacerts/*.pem  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/server.key && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/tlscacerts/*.pem  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/server.key && \

mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/msp/tlscacerts && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/msp/tlscacerts && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/msp/tlscacerts && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/msp/tlscacerts && \
mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/msp/tlscacerts && \


cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/msp/tlscacerts/tlsca.pem && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/msp/tlscacerts/tlsca.pem && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/msp/tlscacerts/tlsca.pem && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/msp/tlscacerts/tlsca.pem && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/msp/tlscacerts/tlsca.pem && \


mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/tlscacerts  && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/msp/tlscacerts/tlsca.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/tlscacerts/ && \


cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer0.${DOMAIN}/msp/ && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer1.${DOMAIN}/msp/ && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer2.${DOMAIN}/msp/ && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer3.${DOMAIN}/msp/ && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/orderers/orderer4.${DOMAIN}/msp/ && \

fabric-ca-client enroll -d -u https://Admin@${DOMAIN}:adminpw@0.0.0.0:4053 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/ordererOrganizations/${DOMAIN}/ca/ca.crt && \

mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/{ca,tlsca,peers,users} && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/org-ca/ca-chain.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ && \
cp  ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/caOrganizations/tls-ca/ca-chain.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/tlsca/ && \

fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/cacerts/*.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/signcerts/*.pem ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/server.key && \

fabric-ca-client register -d --id.name peer1.${DOMAIN} --id.secret peer1pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer,hf.IntermediateCA=true,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' -u https://0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
fabric-ca-client register -d --id.name peer2.${DOMAIN} --id.secret peer2pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer,hf.IntermediateCA=true,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' -u https://0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
fabric-ca-client register -d --id.name Admin@${DOMAIN} --id.secret adminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin,hf.IntermediateCA=true,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' -u https://0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
fabric-ca-client register -d --id.name User0@${DOMAIN} --id.secret user0pw --id.type client -u https://0.0.0.0:4054 --id.attrs '"hf.Registrar.Roles=client,hf.IntermediateCA=true,hf.GenCRL=true,hf.Revoker=true,hf.AffiliationMgr=true"' -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \

mkdir -p ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/{peer1.${DOMAIN},peer2.${DOMAIN}} && \

fabric-ca-client enroll -d -u https://peer1.${DOMAIN}:peer1pw@0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
fabric-ca-client enroll -d -u https://peer2.${DOMAIN}:peer2pw@0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \
fabric-ca-client enroll -d -u https://Admin@${DOMAIN}:adminpw@0.0.0.0:4054 -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/ca/ca-chain.pem && \

echo  \
 'NodeOUs:
     Enable: true
     ClientOUIdentifier:
       Certificate: cacerts/0-0-0-0-4054.pem
       OrganizationalUnitIdentifier: client
     PeerOUIdentifier:
       Certificate: cacerts/0-0-0-0-4054.pem
       OrganizationalUnitIdentifier: peer
     AdminOUIdentifier:
       Certificate: cacerts/0-0-0-0-4054.pem
       OrganizationalUnitIdentifier: admin
     OrdererOUIdentifier:
       Certificate: cacerts/0-0-0-0-4054.pem
       OrganizationalUnitIdentifier: orderer ' > ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/config.yaml && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/msp && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/msp && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/msp/config.yaml ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp && \

mv ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/users/Admin@${DOMAIN}/msp/signcerts/Admin@${DOMAIN}-cert.pem && \

fabric-ca-client enroll -d -u https://peer1.${DOMAIN}:peer1pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts peer1.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/tlsca/ca-chain.pem && \
fabric-ca-client enroll -d -u https://peer2.${DOMAIN}:peer2pw@0.0.0.0:4052 --enrollment.profile tls --csr.hosts peer2.${DOMAIN} -M ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls --tls.certfiles ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/tlsca/ca-chain.pem && \

cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer2.${DOMAIN}/tls/server.key && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/tlscacerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/ca.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/signcerts/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/server.crt && \
cp ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/keystore/* ${GOPATH}/src/${DIRECTORY_NAME}/crypto-config/peerOrganizations/${DOMAIN}/peers/peer1.${DOMAIN}/tls/server.key
