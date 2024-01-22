

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DomainGroup struct {
	Inst struct {

    DomainGroupList []DomainGroupDomainGroupList `json:"domain-group-list"`
    DomainRateList []DomainGroupDomainRateList `json:"domain-rate-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"domain-group"`
}


type DomainGroupDomainGroupList struct {
    DomainListName string `json:"domain-list-name"`
    Action string `json:"action" dval:"permit"`
}


type DomainGroupDomainRateList struct {
    DummyName string `json:"dummy-name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    DomainListList []DomainGroupDomainRateListDomainListList `json:"domain-list-list"`
}


type DomainGroupDomainRateListDomainListList struct {
    Name string `json:"name"`
    PerSuffixRate int `json:"per-suffix-rate"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DomainGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DomainGroup) getPath() string{
    return "domain-group"
}

func (p *DomainGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroup::Post")
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

func (p *DomainGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroup::Get")
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
func (p *DomainGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroup::Put")
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

func (p *DomainGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
