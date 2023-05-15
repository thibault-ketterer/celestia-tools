
NODE_TYPE=light
AUTH_TOKEN=$(celestia $NODE_TYPE auth admin --p2p.network blockspacerace)
echo $AUTH_TOKEN
echo $AUTH_TOKEN > authtoken

curl -X POST \
	-H "Authorization: Bearer $AUTH_TOKEN" \
	-H 'Content-Type: application/json' \
	-d '{"jsonrpc":"2.0","id":0,"method":"p2p.Info","params":[]}' \
	http://localhost:26658 |jq | tee auth.json
