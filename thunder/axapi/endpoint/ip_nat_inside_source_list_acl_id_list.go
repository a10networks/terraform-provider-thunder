

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpNatInsideSourceListAclIdList struct {
	Inst struct {

    AclId int `json:"acl-id"`
    Msl int `json:"msl"`
    Pool string `json:"pool"`
    Uuid string `json:"uuid"`

	} `json:"acl-id-list"`
}

func (p *IpNatInsideSourceListAclIdList) GetId() string{
    return strconv.Itoa(p.Inst.AclId)
}

func (p *IpNatInsideSourceListAclIdList) getPath() string{
    return "ip/nat/inside/source/list/acl-id-list"
}

func (p *IpNatInsideSourceListAclIdList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclIdList::Post")
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

func (p *IpNatInsideSourceListAclIdList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclIdList::Get")
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
func (p *IpNatInsideSourceListAclIdList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclIdList::Put")
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

func (p *IpNatInsideSourceListAclIdList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceListAclIdList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
