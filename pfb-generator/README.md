celestia PayForBlob generator

# tldr
    go run submit.go <randomint> <YOURNODEIP> | sh

# PfB generator

call it like this

- change the number 12133 to some random of your choice
- change the IP to your light node IP


    `go run submit.go seed  IP`

    `go run submit.go 12133 celestia.lankou.org | sh`

it will generate the curl command to call like this

    curl -X POST -d '{"namespace_id": "26bafbbbe9ec6532", "data": "935347d9e8a1a011bc489b16843d5f3a", "gas_limit": 80000, "fee": 2000}' http://celestia.lankou.org:26659/submit_pfb

so you just have to pipe it to a shell to execute

    go run  submit.go 12133 celestia.lankou.org | sh

it will generate the output

    {"height":232396,"txhash":"FF587EDA72D79B2C49EBBC6D88A73B2E336B4E57200F6463097661F501DD0B54","data":"122A0A282F63656C65737469612E626C6F622E76312E4D7367506179466F72426C6F6273526573706F6E7365","raw_log":"[{\"msg_index\":0,\"events\":[{\"type\":\"celestia.blob.v1.EventPayForBlobs\",\"attributes\":[{\"key\":\"blob_sizes\",\"va.........

