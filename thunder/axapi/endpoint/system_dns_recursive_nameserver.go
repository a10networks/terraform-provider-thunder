

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDnsRecursiveNameserver struct {
	Inst struct {

    FollowShared int `json:"follow-shared"`
    ServerList []SystemDnsRecursiveNameserverServerList `json:"server-list"`
    Uuid string `json:"uuid"`

	} `json:"recursive-nameserver"`
}


type SystemDnsRecursiveNameserverServerList struct {
    Ipv4Addr string `json:"ipv4-addr"`
    V4Desc string `json:"v4-desc"`
    Ipv6Addr string `json:"ipv6-addr"`
    V6Desc string `json:"v6-desc"`
}

func (p *SystemDnsRecursiveNameserver) GetId() string{
    return "1"
}

func (p *SystemDnsRecursiveNameserver) getPath() string{
    return "system/dns/recursive-nameserver"
}

func (p *SystemDnsRecursiveNameserver) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDnsRecursiveNameserver::Post")
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

func (p *SystemDnsRecursiveNameserver) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDnsRecursiveNameserver::Get")
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
func (p *SystemDnsRecursiveNameserver) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDnsRecursiveNameserver::Put")
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

func (p *SystemDnsRecursiveNameserver) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDnsRecursiveNameserver::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
