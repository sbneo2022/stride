#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/../config.sh

WAIT_FOR_STRING $STRIDE_LOGS "height=220 module=txindex"

# Create new connections and channels for sharing voting power between two chains
relayer_logs=${LOGS}/relayer-gaia-ics.log
relayer_exec=$(GET_VAR_VALUE RELAYER_GAIA_ICS_EXEC)
relayer_config=$STATE/relayer-gaia-ics/config
mnemonic=$(GET_VAR_VALUE     RELAYER_GAIA_ICS_MNEMONIC)
chain_name=gaia
account_name=$(GET_VAR_VALUE RELAYER_GAIA_ICS_ACCT)
coin_type=$(GET_VAR_VALUE    COSMOS_COIN_TYPE)

mkdir -p $relayer_config
chmod -R 777 $STATE/relayer-gaia-ics
cp ${DOCKERNET_HOME}/config/relayer_config_ics.yaml $relayer_config/config.yaml

printf "STRIDE <> GAIA(ICS) - Adding relayer keys..."
$relayer_exec rly keys restore stride $RELAYER_STRIDE_ICS_ACCT "$mnemonic" >> $relayer_logs 2>&1
$relayer_exec rly keys restore $chain_name $account_name "$mnemonic" --coin-type $coin_type >> $relayer_logs 2>&1
echo "Done restoring relayer keys"

printf "STRIDE <> GAIA - Creating new connections..." | tee -a $relayer_logs
$relayer_exec rly transact connection stride-gaia-ics >> $relayer_logs 2>&1
echo "Done."
sleep 10

printf "STRIDE <> GAIA - Creating new channels..." | tee -a $relayer_logs
$relayer_exec rly transact channel stride-gaia-ics --src-port consumer --dst-port provider --order ordered --version 1 >> $relayer_logs 2>&1
echo "Done."

$DOCKER_COMPOSE up -d relayer-gaia-ics
$DOCKER_COMPOSE logs -f relayer-gaia-ics | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $relayer_logs 2>&1 &