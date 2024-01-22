

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowCollectorTemplateOper struct {
    
    Oper NetflowCollectorTemplateOperOper `json:"oper"`

}
type DataNetflowCollectorTemplateOper struct {
    DtNetflowCollectorTemplateOper NetflowCollectorTemplateOper `json:"template"`
}


type NetflowCollectorTemplateOperOper struct {
    NfTemplateList []NetflowCollectorTemplateOperOperNfTemplateList `json:"nf-template-list"`
}


type NetflowCollectorTemplateOperOperNfTemplateList struct {
    ExporterAddress string `json:"exporter-address"`
    ObservationDomainId int `json:"observation-domain-id"`
    TemplateId int `json:"template-id"`
    TemplateType string `json:"template-type"`
    NflowVersion int `json:"nflow-version"`
    FieldCount int `json:"field-count"`
    SecondsToExpire int `json:"seconds-to-expire"`
    PartitionId int `json:"partition-id"`
}

func (p *NetflowCollectorTemplateOper) GetId() string{
    return "1"
}

func (p *NetflowCollectorTemplateOper) getPath() string{
    return "netflow/collector/template/oper"
}

func (p *NetflowCollectorTemplateOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetflowCollectorTemplateOper,error) {
logger.Println("NetflowCollectorTemplateOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetflowCollectorTemplateOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
