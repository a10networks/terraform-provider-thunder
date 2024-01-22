

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerProfile struct {
	Inst struct {

    Action string `json:"action"`
    Analytics string `json:"analytics" dval:"all"`
    AutoRestartAction string `json:"auto-restart-action" dval:"enable"`
    AvailabilityZone string `json:"availability-zone"`
    ClusterId string `json:"cluster-id"`
    ClusterName string `json:"cluster-name"`
    Host string `json:"host"`
    HostIpv6 string `json:"host-ipv6"`
    Interval int `json:"interval" dval:"3"`
    PasswordEncrypted string `json:"password-encrypted"`
    Port int `json:"port"`
    Provider1 string `json:"provider1"`
    ReSync HarmonyControllerProfileReSync400 `json:"re-sync"`
    Region string `json:"region"`
    SecretValue string `json:"secret-value"`
    ThunderMgmtIp HarmonyControllerProfileThunderMgmtIp401 `json:"thunder-mgmt-ip"`
    Tunnel HarmonyControllerProfileTunnel402 `json:"tunnel"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserName string `json:"user-name"`
    Uuid string `json:"uuid"`

	} `json:"profile"`
}


type HarmonyControllerProfileReSync400 struct {
    SchemaRegistry int `json:"schema-registry"`
    AnalyticsBus int `json:"analytics-bus"`
}


type HarmonyControllerProfileThunderMgmtIp401 struct {
    IpAddress string `json:"ip-address"`
    Ipv6Addr string `json:"ipv6-addr"`
    Uuid string `json:"uuid"`
}


type HarmonyControllerProfileTunnel402 struct {
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}

func (p *HarmonyControllerProfile) GetId() string{
    return "1"
}

func (p *HarmonyControllerProfile) getPath() string{
    return "harmony-controller/profile"
}

func (p *HarmonyControllerProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfile::Post")
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

func (p *HarmonyControllerProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfile::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *HarmonyControllerProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfile::Put")
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

func (p *HarmonyControllerProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
