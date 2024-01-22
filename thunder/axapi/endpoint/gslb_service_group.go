

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceGroup struct {
	Inst struct {

    DependencySite int `json:"dependency-site"`
    Disable int `json:"disable"`
    DisableSiteList []GslbServiceGroupDisableSiteList `json:"disable-site-list"`
    Member []GslbServiceGroupMember `json:"member"`
    PersistentAgingTime int `json:"persistent-aging-time" dval:"5"`
    PersistentIpv6Mask int `json:"persistent-ipv6-mask" dval:"128"`
    PersistentMask string `json:"persistent-mask" dval:"/32"`
    PersistentSite int `json:"persistent-site"`
    ServiceGroupName string `json:"service-group-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-group"`
}


type GslbServiceGroupDisableSiteList struct {
    DisableSite string `json:"disable-site"`
}


type GslbServiceGroupMember struct {
    MemberName string `json:"member-name"`
}

func (p *GslbServiceGroup) GetId() string{
    return p.Inst.ServiceGroupName
}

func (p *GslbServiceGroup) getPath() string{
    return "gslb/service-group"
}

func (p *GslbServiceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceGroup::Post")
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

func (p *GslbServiceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceGroup::Get")
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
func (p *GslbServiceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceGroup::Put")
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

func (p *GslbServiceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
