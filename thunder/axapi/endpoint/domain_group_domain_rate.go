

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DomainGroupDomainRate struct {
	Inst struct {

    DomainListList []DomainGroupDomainRateDomainListList `json:"domain-list-list"`
    DummyName string `json:"dummy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"domain-rate"`
}


type DomainGroupDomainRateDomainListList struct {
    Name string `json:"name"`
    PerSuffixRate int `json:"per-suffix-rate"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DomainGroupDomainRate) GetId() string{
    return p.Inst.DummyName
}

func (p *DomainGroupDomainRate) getPath() string{
    return "domain-group/" +p.Inst.Name + "/domain-rate"
}

func (p *DomainGroupDomainRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroupDomainRate::Post")
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

func (p *DomainGroupDomainRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroupDomainRate::Get")
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
func (p *DomainGroupDomainRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroupDomainRate::Put")
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

func (p *DomainGroupDomainRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DomainGroupDomainRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
