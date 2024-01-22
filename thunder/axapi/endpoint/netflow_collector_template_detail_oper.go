

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowCollectorTemplateDetailOper struct {
    
    Oper NetflowCollectorTemplateDetailOperOper `json:"oper"`

}
type DataNetflowCollectorTemplateDetailOper struct {
    DtNetflowCollectorTemplateDetailOper NetflowCollectorTemplateDetailOper `json:"template-detail"`
}


type NetflowCollectorTemplateDetailOperOper struct {
    NfTemplateList []NetflowCollectorTemplateDetailOperOperNfTemplateList `json:"nf-template-list"`
}


type NetflowCollectorTemplateDetailOperOperNfTemplateList struct {
    ExporterAddress string `json:"exporter-address"`
    ObservationDomainId int `json:"observation-domain-id"`
    TemplateId int `json:"template-id"`
    TemplateType string `json:"template-type"`
    NflowVersion int `json:"nflow-version"`
    FieldCount int `json:"field-count"`
    SecondsToExpire int `json:"seconds-to-expire"`
    PartitionId int `json:"partition-id"`
    TemplateFieldList []NetflowCollectorTemplateDetailOperOperNfTemplateListTemplateFieldList `json:"template-field-list"`
}


type NetflowCollectorTemplateDetailOperOperNfTemplateListTemplateFieldList struct {
    Id1 int `json:"id1"`
    Length int `json:"length"`
    EnterpriseField string `json:"enterprise-field"`
    VariableLength string `json:"variable-length"`
}

func (p *NetflowCollectorTemplateDetailOper) GetId() string{
    return "1"
}

func (p *NetflowCollectorTemplateDetailOper) getPath() string{
    return "netflow/collector/template-detail/oper"
}

func (p *NetflowCollectorTemplateDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetflowCollectorTemplateDetailOper,error) {
logger.Println("NetflowCollectorTemplateDetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetflowCollectorTemplateDetailOper
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
