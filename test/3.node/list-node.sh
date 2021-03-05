#!/bin/bash

#function list_ns() {
    FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

    source ../conf.env
    AUTH="Authorization: Basic $(echo -n $ApiUsername:$ApiPassword | base64)"

    echo "####################################################################"
    echo "## 3. Node: List"
    echo "####################################################################"

    INDEX=${1}

    #curl -H "${AUTH}" -sX GET http://$MyfwServer/myfw/ns/$NS_ID/clusters #| json_pp #|| return 1

    list() {
        c_CT="Content-Type: application/json"
        resp=$(curl -H "${AUTH}" -sX GET http://$MyfwServer/myfw/ns/$NS_ID/clusters/$CLUSTER_ID/nodes
        ); echo ${resp} | jq
    }

    list;
#}

#list_ns
