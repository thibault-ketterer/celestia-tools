#!/bin/bash

# rm -rf ~./<path-to-node-store>/data

# sudo sysctl -w net.core.rmem_max=2500000
network=blockspacerace

# echo init 
# celestia light init --p2p.network $network

echo start 

myip="148.113.6.168"
# myip="localhost"

for rpc in https://rpc-blockspacerace.pops.one \
https://rpc-1.celestia.nodes.guru \
https://rpc-2.celestia.nodes.guru \
https://celestia-testnet.rpc.kjnodes.com \
https://celestia.rpc.waynewayner.de \
https://rpc-blockspacerace.mzonder.com \
https://rpc-t.celestia.nodestake.top \
https://rpc-blockspacerace.ryabina.io \
https://celest-archive.rpc.theamsolutions.info \
https://blockspacerace-rpc.chainode.tech \
https://rpc-blockspacerace.suntzu.pro \
https://public.celestia.w3hitchhiker.com \
https://rpc.celestia.stakewith.us \
https://celestia-rpc.validatrium.club \
https://celrace-rpc.easy2stake.com \
;do
	figlet ${rpc:7:15}
	if timeout 4 curl -s $rpc > /dev/null; then
		echo "rpc START [$rpc] is up"
		celestia light start --core.ip $rpc/ --gateway \
			--gateway.addr "$myip" --gateway.port 26659 \
			--metrics.tls=false --metrics --metrics.endpoint otel.celestia.tools:4318 \
			--p2p.network $network
	else
		echo "rpc [$rpc] is down"
		continue
	fi
#	exit 1
done

