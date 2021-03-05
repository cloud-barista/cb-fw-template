package model

type VM struct {
	Name         string   `json:"name"`
	Config       string   `json:"connectionName"`
	VPC          string   `json:"vNetId"`
	Subnet       string   `json:"subnetId"`
	Firewall     []string `json:"securityGroupIds"`
	SSHKey       string   `json:"sshKeyId"`
	Image        string   `json:"imageId"`
	Spec         string   `json:"specId"`
	UserAccount  string   `json:"vmUserAccount"`
	UserPassword string   `json:"vmUserPassword"`
	Description  string   `json:"description"`
	PublicIP     string   `json:"publicIP"` // output
	Credential   string   // private
	UId          string   `json:"uid"`
	Role         string   `json:"role"`
}
