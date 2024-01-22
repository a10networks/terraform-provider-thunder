

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatInsideSourceListAclNameList struct {
	Inst struct {

    Msl int `json:"msl"`
    Name string `json:"name"`
    Pool string `json:"pool"`
    Uuid string `json:"uuid"`

	} `json:"acl-name-list"`
}

func (p *IpNatInsideSourceListAclNameList) GetId() string{
    return p.Inst.Name
}

func (p *IpNatInsideSourceListAclNameList) getPath() string{
    return "ip/nat/inside/source/list/acl-name-list"
}

func (p *IpNatInsideSourceListAclNameList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclNameList::Post")
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

func (p *IpNatInsideSourceListAclNameList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclNameList::Get")
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
func (p *IpNatInsideSourceListAclNameList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclNameList::Put")
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

func (p *IpNatInsideSourceListAclNameList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclNameList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
