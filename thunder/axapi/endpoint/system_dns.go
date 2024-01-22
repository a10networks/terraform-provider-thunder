

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDns struct {
	Inst struct {

    RecursiveNameserver SystemDnsRecursiveNameserver1571 `json:"recursive-nameserver"`
    SamplingEnable []SystemDnsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type SystemDnsRecursiveNameserver1571 struct {
    FollowShared int `json:"follow-shared"`
    ServerList []SystemDnsRecursiveNameserverServerList1572 `json:"server-list"`
    Uuid string `json:"uuid"`
}


type SystemDnsRecursiveNameserverServerList1572 struct {
    Ipv4Addr string `json:"ipv4-addr"`
    V4Desc string `json:"v4-desc"`
    Ipv6Addr string `json:"ipv6-addr"`
    V6Desc string `json:"v6-desc"`
}


type SystemDnsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemDns) GetId() string{
    return "1"
}

func (p *SystemDns) getPath() string{
    return "system/dns"
}

func (p *SystemDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDns::Post")
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

func (p *SystemDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDns::Get")
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
func (p *SystemDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDns::Put")
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

func (p *SystemDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
