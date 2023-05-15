#!/bin/bash
set -euo pipefail

celestia_node_dir="/home/ubuntu/celestia-node"
node_type="light"

# setup SERVICE if needed

cat | sudo tee /etc/systemd/system/celestia-${node_type}d.service >/dev/null <<EOF

[Unit]
Description=celestia-lightd Light Node
After=network-online.target
 
[Service]
User=ubuntu
ExecStart=$celestia_node_dir/start.sh
Restart=on-failure
RestartSec=3
LimitNOFILE=4096
 
[Install]
WantedBy=multi-user.target
EOF

echo "updating system config"
sudo systemctl daemon-reload
sudo systemctl enable celestia-${node_type}d.service
echo
echo "stopping daemon"
sudo systemctl stop celestia-${node_type}d.service

cd $celestia_node_dir

git fetch --tags
last_tag=$(git tag | sort -V | tail -1)
if [ -z "$last_tag" ];then
	echo "last tag not found, fix the script"
	exit 1
fi
echo
echo "last tag is [$last_tag]"
echo
echo "checkout"
git checkout "${last_tag}"

echo
echo "building from source if needed"
make build
echo
echo "installing"
sudo make install

echo
echo "update config"
if [ "$node_type" == "light" ];then
	celestia light config-update --p2p.network blockspacerace
fi

echo
echo "start service"
sudo systemctl start celestia-${node_type}d.service

echo
sudo systemctl status celestia-${node_type}d.service --no-pager

echo
celestia version

echo
echo "upgrade finished"
echo "have a good day"
echo
