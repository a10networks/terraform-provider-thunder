

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflowTemplateDetailOper struct {
    
    Oper VisibilityFlowCollectorNetflowTemplateDetailOperOper `json:"oper"`

}
type DataVisibilityFlowCollectorNetflowTemplateDetailOper struct {
    DtVisibilityFlowCollectorNetflowTemplateDetailOper VisibilityFlowCollectorNetflowTemplateDetailOper `json:"detail"`
}


type VisibilityFlowCollectorNetflowTemplateDetailOperOper struct {
    NfTemplateList []VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList `json:"nf-template-list"`
}


type VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList struct {
    ExporterAddress string `json:"exporter-address"`
    ObservationDomainId int `json:"observation-domain-id"`
    TemplateId int `json:"template-id"`
    NflowVersion int `json:"nflow-version"`
    FieldCount int `json:"field-count"`
    SecondsToExpire int `json:"seconds-to-expire"`
    PartitionId int `json:"partition-id"`
    TemplateFieldList []VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList `json:"template-field-list"`
}


type VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList struct {
    Id1 int `json:"id1"`
    Length int `json:"length"`
    EnterpriseField string `json:"enterprise-field"`
    VariableLength string `json:"variable-length"`
}

func (p *VisibilityFlowCollectorNetflowTemplateDetailOper) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflowTemplateDetailOper) getPath() string{
    return "visibility/flow-collector/netflow/template/detail/oper"
}

func (p *VisibilityFlowCollectorNetflowTemplateDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFlowCollectorNetflowTemplateDetailOper,error) {
logger.Println("VisibilityFlowCollectorNetflowTemplateDetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFlowCollectorNetflowTemplateDetailOper
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
