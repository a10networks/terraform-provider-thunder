

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnEndpointIndependentFilteringTcp struct {
	Inst struct {

    PortList []Cgnv6LsnEndpointIndependentFilteringTcpPortList `json:"port-list"`
    SessionLimit int `json:"session-limit" dval:"65535"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type Cgnv6LsnEndpointIndependentFilteringTcpPortList struct {
    Port int `json:"port"`
    PortEnd int `json:"port-end"`
}

func (p *Cgnv6LsnEndpointIndependentFilteringTcp) GetId() string{
    return "1"
}

func (p *Cgnv6LsnEndpointIndependentFilteringTcp) getPath() string{
    return "cgnv6/lsn/endpoint-independent-filtering/tcp"
}

func (p *Cgnv6LsnEndpointIndependentFilteringTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentFilteringTcp::Post")
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

func (p *Cgnv6LsnEndpointIndependentFilteringTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentFilteringTcp::Get")
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
func (p *Cgnv6LsnEndpointIndependentFilteringTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentFilteringTcp::Put")
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

func (p *Cgnv6LsnEndpointIndependentFilteringTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentFilteringTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
