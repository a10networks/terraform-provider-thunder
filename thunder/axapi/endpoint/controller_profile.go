package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ControllerProfile struct {
	Inst struct {
		Action string `json:"action"`

		Analytics string `json:"analytics" dval:"disable"`

		AnalyticsData int `json:"analytics-data"`

		ApiKey string `json:"api-key"`

		AutoRestartAction string `json:"auto-restart-action" dval:"enable"`

		AvailabilityZone string `json:"availability-zone"`

		ClusterId string `json:"cluster-id"`

		ClusterName string `json:"cluster-name"`

		Force ControllerProfileForce140 `json:"force"`

		Host string `json:"host"`

		HostIpv6 string `json:"host-ipv6"`

		Interval int `json:"interval" dval:"3"`

		Organization string `json:"organization"`

		PasswordEncrypted string `json:"password-encrypted"`

		Port int `json:"port"`

		ReSync ControllerProfileReSync141 `json:"re-sync"`

		Region string `json:"region"`

		SecretValue string `json:"secret-value"`

		ThunderMgmtIp ControllerProfileThunderMgmtIp142 `json:"thunder-mgmt-ip"`

		Tunnel ControllerProfileTunnel143 `json:"tunnel"`

		UseMgmtPort int `json:"use-mgmt-port"`

		UserName string `json:"user-name"`

		Uuid string `json:"uuid"`
	} `json:"profile"`
}

type ControllerProfileForce140 struct {
	Deregister int `json:"deregister"`
}

type ControllerProfileReSync141 struct {
	SchemaRegistry int `json:"schema-registry"`
	AnalyticsBus   int `json:"analytics-bus"`
}

type ControllerProfileThunderMgmtIp142 struct {
	IpAddress string `json:"ip-address"`
	Ipv6Addr  string `json:"ipv6-addr"`
	Uuid      string `json:"uuid"`
}

type ControllerProfileTunnel143 struct {
	Action string `json:"action" dval:"disable"`
	Uuid   string `json:"uuid"`
}

func (p *ControllerProfile) GetId() string {
	return "1"
}

func (p *ControllerProfile) getPath() string {
	return "controller/profile"
}

func (p *ControllerProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfile::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *ControllerProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfile::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *ControllerProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfile::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *ControllerProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ControllerProfile::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
