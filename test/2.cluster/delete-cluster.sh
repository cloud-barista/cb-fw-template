#!/bin/bash

#function delete_ns() {
    FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

    source ../conf.env
    AUTH="Authorization: Basic $(echo -n $ApiUsername:$ApiPassword | base64)"

    echo "####################################################################"
    echo "## 2. Cluster: Delete"
    echo "####################################################################"

    INDEX=${1}

    #curl -H "${AUTH}" -sX DELETE http://$TumblebugServer/tumblebug/ns/$NS_ID | json_pp #|| return 1

    delete() {
        c_CT="Content-Type: application/json"
        resp=$(curl -H "${AUTH}" -sX DELETE http://$MyfwServer/myfw/ns/$NS_ID/clusters/$CLUSTER_ID
        ); echo ${resp} | jq
    }

    delete;
#}

#delete_ns