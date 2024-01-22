

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats53 struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStats53Stats `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbVirtualServerPortStats53Stats struct {
    ServerSsh SlbVirtualServerPortStats53StatsServerSsh `json:"server-ssh"`
}


type SlbVirtualServerPortStats53StatsServerSsh struct {
    Successful_handshakes int `json:"successful_handshakes"`
    Failed_handshakes int `json:"failed_handshakes"`
    Forwarding_errors int `json:"forwarding_errors"`
}

func (p *SlbVirtualServerPortStats53) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats53) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?server-ssh=true"
}

func (p *SlbVirtualServerPortStats53) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats53::Post")
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

func (p *SlbVirtualServerPortStats53) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats53::Get")
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
func (p *SlbVirtualServerPortStats53) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats53::Put")
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

func (p *SlbVirtualServerPortStats53) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats53::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
