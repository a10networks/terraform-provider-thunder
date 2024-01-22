

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemAllVlanLimit struct {
	Inst struct {

    Bcast int `json:"bcast" dval:"5000"`
    Ipmcast int `json:"ipmcast" dval:"5000"`
    Mcast int `json:"mcast" dval:"5000"`
    UnknownUcast int `json:"unknown-ucast" dval:"5000"`
    Uuid string `json:"uuid"`

	} `json:"all-vlan-limit"`
}

func (p *SystemAllVlanLimit) GetId() string{
    return "1"
}

func (p *SystemAllVlanLimit) getPath() string{
    return "system/all-vlan-limit"
}

func (p *SystemAllVlanLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAllVlanLimit::Post")
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

func (p *SystemAllVlanLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAllVlanLimit::Get")
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
func (p *SystemAllVlanLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAllVlanLimit::Put")
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

func (p *SystemAllVlanLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAllVlanLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
