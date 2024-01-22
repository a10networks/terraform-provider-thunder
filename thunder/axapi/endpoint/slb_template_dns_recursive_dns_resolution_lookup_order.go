

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsRecursiveDnsResolutionLookupOrder struct {
	Inst struct {

    QueryType []SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType `json:"query-type"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"lookup-order"`
}


type SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType struct {
    StrQueryType string `json:"str-query-type"`
    NumQueryType int `json:"num-query-type"`
    Order string `json:"order"`
}

func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) GetId() string{
    return "1"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/recursive-dns-resolution/lookup-order"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionLookupOrder::Post")
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

func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionLookupOrder::Get")
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
func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionLookupOrder::Put")
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

func (p *SlbTemplateDnsRecursiveDnsResolutionLookupOrder) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionLookupOrder::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
