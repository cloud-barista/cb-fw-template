#!/bin/bash
FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

    source ../conf.env
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "get-node.sh [GCP/AWS] <clsuter name> <node name>"
	echo "    ./get-node.sh GCP cb-cluster cb-gcp-cluster-test-0-w-asdf"
	exit 0; 
fi


# ------------------------------------------------------------------------------
# const
c_URL_MYFW="http://localhost:4321/myfw"
c_CT="Content-Type: application/json"


# -----------------------------------------------------------------
# parameter

# 1. CSP
if [ "$#" -gt 0 ]; then v_CSP="$1"; else	v_CSP="${CSP}"; fi
if [ "${v_CSP}" == "" ]; then 
	read -e -p "Cloud ? [AWS(default) or GCP] : "  v_CSP
fi
if [ "${v_CSP}" == "" ]; then v_CSP="AWS"; fi
if [ "${v_CSP}" != "GCP" ] && [ "${v_CSP}" != "AWS" ]; then echo "[ERROR] missing <cloud>"; exit -1;fi

# PREFIX
if [ "${v_CSP}" == "GCP" ]; then 
	v_PREFIX="cb-gcp"
else
	v_PREFIX="cb-aws"
fi

# 2. Cluster Name
if [ "$#" -gt 1 ]; then v_CLUSTER_NAME="$2"; else	v_METHOD="${CLUSTER_NAME}"; fi
if [ "${v_CLUSTER_NAME}" == "" ]; then 
	read -e -p "Cluster name  ? : "  v_CLUSTER_NAME
fi
if [ "${v_CLUSTER_NAME}" == "" ]; then echo "[ERROR] missing <cluster name>"; exit -1; fi

# 3. Node Name
if [ "$#" -gt 2 ]; then v_NODE_NAME="$3"; else	v_METHOD="${NODE_NAME}"; fi
if [ "${v_NODE_NAME}" == "" ]; then 
	read -e -p "Node name  ? : "  v_NODE_NAME
fi
if [ "${v_NODE_NAME}" == "" ]; then echo "[ERROR] missing <cluster name>"; exit -1; fi

NM_NAMESPACE="${v_PREFIX}-namespace"
c_URL_MYFW_NS="${c_URL_MYFW}/ns/${NM_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Prefix                     is '${v_PREFIX}'"
echo "- Cluster name               is '${v_CLUSTER_NAME}'"
echo "- Node name               	 is '${v_NODE_NAME}'"
echo "- Namespace                  is '${NM_NAMESPACE}'"


# ------------------------------------------------------------------------------
# Delete a cluster
get() {

	curl -sX GET http://$MyfwServer/myfw/ns/$NS_ID/clusters/$CLUSTER_ID/nodes/${v_NODE_NAME}    -H "${c_CT}" | jq;

}

# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	get;
fi
