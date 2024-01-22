

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateHttpMalformedHttp struct {
	Inst struct {

    MalformedHttp string `json:"malformed-http"`
    MalformedHttpAction string `json:"malformed-http-action"`
    MalformedHttpActionListName string `json:"malformed-http-action-list-name"`
    MalformedHttpBadChunkMonEnabled int `json:"malformed-http-bad-chunk-mon-enabled"`
    MalformedHttpMaxContentLength int `json:"malformed-http-max-content-length" dval:"4294967295"`
    MalformedHttpMaxHeaderNameSize int `json:"malformed-http-max-header-name-size" dval:"64"`
    MalformedHttpMaxLineSize int `json:"malformed-http-max-line-size" dval:"32512"`
    MalformedHttpMaxNumHeaders int `json:"malformed-http-max-num-headers" dval:"90"`
    MalformedHttpMaxReqLineSize int `json:"malformed-http-max-req-line-size" dval:"32512"`
    Uuid string `json:"uuid"`
    HttpTmplName string 

	} `json:"malformed-http"`
}

func (p *DdosZoneTemplateHttpMalformedHttp) GetId() string{
    return "1"
}

func (p *DdosZoneTemplateHttpMalformedHttp) getPath() string{
    return "ddos/zone-template/http/" +p.Inst.HttpTmplName + "/malformed-http"
}

func (p *DdosZoneTemplateHttpMalformedHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpMalformedHttp::Post")
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

func (p *DdosZoneTemplateHttpMalformedHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpMalformedHttp::Get")
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
func (p *DdosZoneTemplateHttpMalformedHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpMalformedHttp::Put")
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

func (p *DdosZoneTemplateHttpMalformedHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpMalformedHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
