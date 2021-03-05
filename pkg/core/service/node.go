package service

import (
	"fmt"

	"github.com/cloud-barista/cb-fw-template/pkg/core/common"
	"github.com/cloud-barista/cb-fw-template/pkg/core/model"
	"github.com/cloud-barista/cb-fw-template/pkg/core/model/tumblebug"
	"github.com/cloud-barista/cb-fw-template/pkg/utils/lang"
	logger "github.com/sirupsen/logrus"
)

func ListNode(namespace string, clusterName string) (*model.NodeList, error) {
	nodes := model.NewNodeList(namespace, clusterName)
	err := nodes.SelectList()
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func GetNode(namespace string, clusterName string, nodeName string) (*model.Node, error) {
	node := model.NewNode(namespace, clusterName, nodeName)
	err := node.Select()
	if err != nil {
		return nil, err
	}

	return node, nil
}

func AddNode(namespace string, clusterName string, req *model.NodeReq) (*model.NodeList, error) {
	mcisName := clusterName

	// Here comes the required actions of AddNode

	// vm
	var VMs []model.VM
	req.WorkerNodeCount = 1
	for i := 0; i < req.WorkerNodeCount; i++ {
		vm := tumblebug.NewTVm(namespace, mcisName)
		vm.VM = model.VM{
			Name:         lang.GetNodeName(clusterName, "worker"), //lang.GetNodeName(clusterName, spec.Role),
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

	// insert to store
	nodes := model.NewNodeList(namespace, clusterName)
	for _, vm := range VMs {
		node := model.NewNodeVM(namespace, clusterName, vm)
		err := node.Insert()
		if err != nil {
			return nil, err
		}
		nodes.Items = append(nodes.Items, *node)
	}

	return nodes, nil
}

func RemoveNode(namespace string, clusterName string, nodeName string) (*model.Status, error) {
	logger.Infof("RemoveNode(namespace=%s, clusterName=%s, nodeName=%s)", namespace, clusterName, nodeName)
	status := model.NewStatus()
	status.Code = model.STATUS_UNKNOWN

	check, _ := CheckNode(namespace, clusterName, nodeName)

	if !check {
		err := fmt.Errorf("The node " + nodeName + " does not exist.")
		return nil, err
	}

	// Here comes the required actions of RemoveNode

	/*
		// delete vm
		vm := tumblebug.NewTVm(namespace, clusterName)
		vm.VM.Name = nodeName
		err := vm.DELETE()
		if err != nil {
			status.Message = "delete vm failed"
			return status, err
		}
	*/

	// delete node from store
	node := model.NewNode(namespace, clusterName, nodeName)
	if err := node.Delete(); err != nil {
		status.Message = err.Error()
		return status, nil
	}

	status.Code = model.STATUS_SUCCESS
	status.Message = "success"

	return status, nil
}

func CheckNode(namespace string, clusterName string, nodeName string) (bool, error) {
	logger.Infof("CheckNode(namespace=%s, clusterName=%s, nodeName=%s)", namespace, clusterName, nodeName)

	// Check parameters' emptiness
	if namespace == "" {
		err := fmt.Errorf("CheckNode failed; namespace given is null.")
		return false, err
	} else if clusterName == "" {
		err := fmt.Errorf("CheckNode failed; clusterName given is null.")
		return false, err
	} else if nodeName == "" {
		err := fmt.Errorf("CheckNode failed; nodeName given is null.")
		return false, err
	}

	key := lang.GetStoreNodeKey(namespace, clusterName, nodeName)
	//fmt.Println(key)

	keyValue, _ := common.CBStore.Get(key)
	if keyValue != nil {
		return true, nil
	}
	return false, nil

}
