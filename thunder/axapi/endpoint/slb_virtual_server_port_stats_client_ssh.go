

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats48 struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStats48Stats `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbVirtualServerPortStats48Stats struct {
    ClientSsh SlbVirtualServerPortStats48StatsClientSsh `json:"client-ssh"`
}


type SlbVirtualServerPortStats48StatsClientSsh struct {
    Successful_handshakes int `json:"successful_handshakes"`
    Failed_handshakes int `json:"failed_handshakes"`
    Forwarding_errors int `json:"forwarding_errors"`
}

func (p *SlbVirtualServerPortStats48) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats48) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?client-ssh=true"
}

func (p *SlbVirtualServerPortStats48) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats48::Post")
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

func (p *SlbVirtualServerPortStats48) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats48::Get")
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
func (p *SlbVirtualServerPortStats48) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats48::Put")
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

func (p *SlbVirtualServerPortStats48) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats48::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
