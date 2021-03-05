package service

import (
	"fmt"

	"github.com/cloud-barista/cb-fw-template/pkg/core/common"
	"github.com/cloud-barista/cb-fw-template/pkg/core/model"
	"github.com/cloud-barista/cb-fw-template/pkg/core/model/tumblebug"
	"github.com/cloud-barista/cb-fw-template/pkg/utils/lang"

	logger "github.com/sirupsen/logrus"
)

func ListCluster(namespace string) (*model.ClusterList, error) {
	clusters := model.NewClusterList(namespace)

	err := clusters.SelectList()
	if err != nil {
		return nil, err
	}

	return clusters, nil
}

func GetCluster(namespace string, clusterName string) (*model.Cluster, error) {
	cluster := model.NewCluster(namespace, clusterName)
	err := cluster.Select()
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func CreateCluster(namespace string, req *model.ClusterReq) (*model.Cluster, error) {

	clusterName := req.Name

	check, _ := CheckCluster(namespace, clusterName)

	if check {
		err := fmt.Errorf("The cluster " + clusterName + " already exists.")
		return nil, err
	}

	cluster := model.NewCluster(namespace, clusterName)
	cluster.UId = lang.GetUid()
	mcisName := clusterName

	// vm
	var VMs []model.VM
	req.WorkerNodeCount = 1
	for i := 0; i < req.WorkerNodeCount; i++ {
		vm := tumblebug.NewTVm(namespace, mcisName)
		vm.VM = model.VM{
			Name:         lang.GetNodeName(clusterName, "master"), //lang.GetNodeName(clusterName, spec.Role),
			Config:       "cb-myfw-test-config",                   //req.Config,
			VPC:          "cb-myfw-test-vpc",                      //vpc.Name,
			Subnet:       "cb-myfw-test-subnet",                   //vpc.Subnets[0].Name,
			Firewall:     []string{"cb-myfw-test-firewall"},       //[]string{fw.Name},
			SSHKey:       "cb-myfw-test-sshkey",                   //sshKey.Name,
			Image:        "cb-myfw-test-image",                    //image.Name,
			Spec:         "cb-myfw-test-spec",                     //spec.Name,
			UserAccount:  "cb-myfw-test-useraccount",              //account,
			UserPassword: "",
			Description:  "",
			Credential:   "cb-myfw-test-credential", //sshKey.PrivateKey,
			Role:         "cb-myfw-test-role",       //spec.Role,
		}

		/*
			// vm 생성
			logger.Infof("start create VM (mcisname=%s, nodename=%s)", mcisName, vm.VM.Name)
			err := vm.POST()
			if err != nil {
				logger.Warnf("create VM error (mcisname=%s, nodename=%s)", mcisName, vm.VM.Name)
				return nil, err
			}
		*/
		VMs = append(VMs, vm.VM)
		logger.Infof("create VM OK.. (mcisname=%s, nodename=%s)", mcisName, vm.VM.Name)
	}

	// 결과값 저장
	var nodes []model.Node
	cluster.MCIS = mcisName
	for _, vm := range VMs { // range mcis.VMs {
		node := model.NewNodeVM(namespace, cluster.Name, vm)
		node.UId = lang.GetUid()

		// insert node in store
		nodes = append(nodes, *node)
		err := node.Insert()
		if err != nil {
			return nil, err
		}
	}
	err := cluster.Insert()
	if err != nil {
		return nil, err
	}

	cluster.Complete()
	cluster.Nodes = nodes

	return cluster, nil
}

func DeleteCluster(namespace string, clusterName string) (*model.Status, error) {
	// Here comes the required actions of DeleteCluster

	//mcisName := clusterName //cluster 이름과 동일하게 (임시)

	status := model.NewStatus()
	status.Code = model.STATUS_UNKNOWN

	check, _ := CheckCluster(namespace, clusterName)

	if !check {
		err := fmt.Errorf("The cluster " + clusterName + " does not exist.")
		return nil, err
	}

	/*
		// 1. delete mcis
		logger.Infof("start delete MCIS (name=%s)", mcisName)
		mcis := tumblebug.NewMCIS(namespace, mcisName)
		exist, err := mcis.GET()
		if err != nil {
			return status, err
		}
		if exist {
			if err = mcis.DELETE(); err != nil {
				return status, err
			}
			// sleep 이후 확인하는 로직 추가 필요
			logger.Infof("delete MCIS OK.. (name=%s)", mcisName)
			status.Code = model.STATUS_SUCCESS
			status.Message = "success"
	*/
	cluster := model.NewCluster(namespace, clusterName)
	if err := cluster.Delete(); err != nil {
		status.Message = "delete success but cannot delete from the store"
		return status, nil
	}
	/*
		} else {
			status.Code = model.STATUS_NOT_EXIST
			logger.Infof("delete MCIS skip (cannot find).. (name=%s)", mcisName)
		}
	*/

	return status, nil
}

func CheckCluster(namespace string, clusterName string) (bool, error) {
	logger.Infof("CheckCluster(namespace=%s, clusterName=%s)", namespace, clusterName)

	// Check parameters' emptiness
	if namespace == "" {
		err := fmt.Errorf("CheckCluster failed; namespace given is null.")
		return false, err
	} else if clusterName == "" {
		err := fmt.Errorf("CheckCluster failed; clusterName given is null.")
		return false, err
	}

	key := lang.GetStoreClusterKey(namespace, clusterName)
	//fmt.Println(key)

	keyValue, _ := common.CBStore.Get(key)
	if keyValue != nil {
		return true, nil
	}
	return false, nil

}
