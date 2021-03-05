#!/bin/bash

#function create_ns() {
    FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

	source ../conf.env
	AUTH="Authorization: Basic $(echo -n $ApiUsername:$ApiPassword | base64)"

	echo "####################################################################"
	echo "## 3. Node: Add"
	echo "####################################################################"

	INDEX=${1}

    create() {
        c_CT="Content-Type: application/json"
        resp=$(curl -H "${AUTH}" -sX POST http://$MyfwServer/myfw/ns/$NS_ID/clusters/$CLUSTER_ID/nodes -H "${c_CT}" -d @- <<EOF
        {
            "config"                : "myfw-test-config",
            "workerNodeCount"       : 2,
            "workerNodeSpec"        : "myfw-test-spec" 
        }
EOF
        ); echo ${resp} | jq
    }

    create;
#}

#create_ns