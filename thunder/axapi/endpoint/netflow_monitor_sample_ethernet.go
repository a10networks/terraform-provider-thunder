

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorSampleEthernet struct {
	Inst struct {

    Ifindex int `json:"ifindex"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ethernet"`
}

func (p *NetflowMonitorSampleEthernet) GetId() string{
    return strconv.Itoa(p.Inst.Ifindex)
}

func (p *NetflowMonitorSampleEthernet) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/sample/ethernet"
}

func (p *NetflowMonitorSampleEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorSampleEthernet::Post")
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

func (p *NetflowMonitorSampleEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorSampleEthernet::Get")
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
func (p *NetflowMonitorSampleEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorSampleEthernet::Put")
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

func (p *NetflowMonitorSampleEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorSampleEthernet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
