

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnEndpointIndependentMappingTcp struct {
	Inst struct {

    PortList []Cgnv6LsnEndpointIndependentMappingTcpPortList `json:"port-list"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type Cgnv6LsnEndpointIndependentMappingTcpPortList struct {
    Port int `json:"port"`
    PortEnd int `json:"port-end"`
}

func (p *Cgnv6LsnEndpointIndependentMappingTcp) GetId() string{
    return "1"
}

func (p *Cgnv6LsnEndpointIndependentMappingTcp) getPath() string{
    return "cgnv6/lsn/endpoint-independent-mapping/tcp"
}

func (p *Cgnv6LsnEndpointIndependentMappingTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentMappingTcp::Post")
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

func (p *Cgnv6LsnEndpointIndependentMappingTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentMappingTcp::Get")
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
func (p *Cgnv6LsnEndpointIndependentMappingTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentMappingTcp::Put")
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

func (p *Cgnv6LsnEndpointIndependentMappingTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnEndpointIndependentMappingTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
