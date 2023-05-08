celestia PayForBlob generator

# tldr
    git clone https://github.com/thibault-ketterer/celestia-tools/

    cd celestia-tools/pfb-generator

    go run submit.go <randomint> <YOURNODEIP> | sh

# PfB generator

call it like this

- change the number 12133 to some random of your choice
- change the IP to your light node IP


    `go run submit.go seed  IP`

    `go run ui-submit.go 12133123 celestia.lankou.org`

it will generate the output

    {"height":232396,"txhash":"FF587EDA72D79B2C49EBBC6D88A73B2E336B4E57200F6463097661F501DD0B54","data":"122A0A282F63656C65737469612E626C6F622E76312E4D7367506179466F72426C6F6273526573706F6E7365","raw_log":"[{\"msg_index\":0,\"events\":[{\"type\":\"celestia.blob.v1.EventPayForBlobs\",\"attributes\":[{\"key\":\"blob_sizes\",\"va.........

