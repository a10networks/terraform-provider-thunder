

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsCollectorGroupLogServer struct {
	Inst struct {

    Name string `json:"name"`
    Port int `json:"port"`
    Uuid string `json:"uuid"`

	} `json:"log-server"`
}

func (p *AcosEventsCollectorGroupLogServer) GetId() string{
    return p.Inst.Name+"+"+strconv.Itoa(p.Inst.Port)
}

func (p *AcosEventsCollectorGroupLogServer) getPath() string{
    return "acos-events/collector-group/"+p.Inst.Name+"/log-server"
}

func (p *AcosEventsCollectorGroupLogServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroupLogServer::Post")
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

func (p *AcosEventsCollectorGroupLogServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroupLogServer::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *AcosEventsCollectorGroupLogServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroupLogServer::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *AcosEventsCollectorGroupLogServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroupLogServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
