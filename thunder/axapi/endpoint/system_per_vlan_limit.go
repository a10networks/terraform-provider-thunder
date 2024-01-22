

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPerVlanLimit struct {
	Inst struct {

    Bcast int `json:"bcast" dval:"1000"`
    Ipmcast int `json:"ipmcast" dval:"1000"`
    Mcast int `json:"mcast" dval:"1000"`
    UnknownUcast int `json:"unknown-ucast" dval:"1000"`
    Uuid string `json:"uuid"`

	} `json:"per-vlan-limit"`
}

func (p *SystemPerVlanLimit) GetId() string{
    return "1"
}

func (p *SystemPerVlanLimit) getPath() string{
    return "system/per-vlan-limit"
}

func (p *SystemPerVlanLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPerVlanLimit::Post")
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

func (p *SystemPerVlanLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPerVlanLimit::Get")
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
func (p *SystemPerVlanLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPerVlanLimit::Put")
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

func (p *SystemPerVlanLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPerVlanLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
