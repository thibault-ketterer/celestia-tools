
# rm -rf ~./<path-to-node-store>/data

sudo sysctl -w net.core.rmem_max=2500000
# celestia light start --p2p.network blockspacerace --core.ip <address> --gateway --gateway.addr <ip-address> --gateway.port 26659
network=blockspacerace
echo init 
celestia light init --p2p.network $network
