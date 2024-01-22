

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflowTemplateOper struct {
    
    Detail VisibilityFlowCollectorNetflowTemplateOperDetail `json:"detail"`
    Oper VisibilityFlowCollectorNetflowTemplateOperOper `json:"oper"`

}
type DataVisibilityFlowCollectorNetflowTemplateOper struct {
    DtVisibilityFlowCollectorNetflowTemplateOper VisibilityFlowCollectorNetflowTemplateOper `json:"template"`
}


type VisibilityFlowCollectorNetflowTemplateOperDetail struct {
    Oper VisibilityFlowCollectorNetflowTemplateOperDetailOper `json:"oper"`
}


type VisibilityFlowCollectorNetflowTemplateOperDetailOper struct {
    NfTemplateList []VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList `json:"nf-template-list"`
}


type VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList struct {
    ExporterAddress string `json:"exporter-address"`
    ObservationDomainId int `json:"observation-domain-id"`
    TemplateId int `json:"template-id"`
    NflowVersion int `json:"nflow-version"`
    FieldCount int `json:"field-count"`
    SecondsToExpire int `json:"seconds-to-expire"`
    PartitionId int `json:"partition-id"`
    TemplateFieldList []VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList `json:"template-field-list"`
}


type VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList struct {
    Id1 int `json:"id1"`
    Length int `json:"length"`
    EnterpriseField string `json:"enterprise-field"`
    VariableLength string `json:"variable-length"`
}


type VisibilityFlowCollectorNetflowTemplateOperOper struct {
    NfTemplateList []VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList `json:"nf-template-list"`
}


type VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList struct {
    ExporterAddress string `json:"exporter-address"`
    ObservationDomainId int `json:"observation-domain-id"`
    TemplateId int `json:"template-id"`
    NflowVersion int `json:"nflow-version"`
    FieldCount int `json:"field-count"`
    SecondsToExpire int `json:"seconds-to-expire"`
    PartitionId int `json:"partition-id"`
}

func (p *VisibilityFlowCollectorNetflowTemplateOper) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflowTemplateOper) getPath() string{
    return "visibility/flow-collector/netflow/template/oper"
}

func (p *VisibilityFlowCollectorNetflowTemplateOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFlowCollectorNetflowTemplateOper,error) {
logger.Println("VisibilityFlowCollectorNetflowTemplateOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFlowCollectorNetflowTemplateOper
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
