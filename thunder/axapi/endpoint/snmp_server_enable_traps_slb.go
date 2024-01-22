

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSlb struct {
	Inst struct {

    All int `json:"all"`
    ApplicationBufferLimit int `json:"application-buffer-limit"`
    BwRateLimitExceed int `json:"bw-rate-limit-exceed"`
    BwRateLimitResume int `json:"bw-rate-limit-resume"`
    GatewayDown int `json:"gateway-down"`
    GatewayUp int `json:"gateway-up"`
    ServerConnLimit int `json:"server-conn-limit"`
    ServerConnResume int `json:"server-conn-resume"`
    ServerDisabled int `json:"server-disabled"`
    ServerDown int `json:"server-down"`
    ServerSelectionFailure int `json:"server-selection-failure"`
    ServerUp int `json:"server-up"`
    ServiceConnLimit int `json:"service-conn-limit"`
    ServiceConnResume int `json:"service-conn-resume"`
    ServiceDown int `json:"service-down"`
    ServiceGroupDown int `json:"service-group-down"`
    ServiceGroupMemberDown int `json:"service-group-member-down"`
    ServiceGroupMemberUp int `json:"service-group-member-up"`
    ServiceGroupUp int `json:"service-group-up"`
    ServiceUp int `json:"service-up"`
    Uuid string `json:"uuid"`
    VipConnlimit int `json:"vip-connlimit"`
    VipConnratelimit int `json:"vip-connratelimit"`
    VipDown int `json:"vip-down"`
    VipPortConnlimit int `json:"vip-port-connlimit"`
    VipPortConnratelimit int `json:"vip-port-connratelimit"`
    VipPortDown int `json:"vip-port-down"`
    VipPortUp int `json:"vip-port-up"`
    VipUp int `json:"vip-up"`

	} `json:"slb"`
}

func (p *SnmpServerEnableTrapsSlb) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSlb) getPath() string{
    return "snmp-server/enable/traps/slb"
}

func (p *SnmpServerEnableTrapsSlb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlb::Post")
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

func (p *SnmpServerEnableTrapsSlb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlb::Get")
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
func (p *SnmpServerEnableTrapsSlb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlb::Put")
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

func (p *SnmpServerEnableTrapsSlb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
