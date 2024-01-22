

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkBpduFwdGroup struct {
	Inst struct {

    BpduFwdGroupNumber int `json:"bpdu-fwd-group-number"`
    EthernetList []NetworkBpduFwdGroupEthernetList `json:"ethernet-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"bpdu-fwd-group"`
}


type NetworkBpduFwdGroupEthernetList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *NetworkBpduFwdGroup) GetId() string{
    return strconv.Itoa(p.Inst.BpduFwdGroupNumber)
}

func (p *NetworkBpduFwdGroup) getPath() string{
    return "network/bpdu-fwd-group"
}

func (p *NetworkBpduFwdGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBpduFwdGroup::Post")
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

func (p *NetworkBpduFwdGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBpduFwdGroup::Get")
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
func (p *NetworkBpduFwdGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBpduFwdGroup::Put")
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

func (p *NetworkBpduFwdGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBpduFwdGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
