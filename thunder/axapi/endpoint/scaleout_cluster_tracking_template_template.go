

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterTrackingTemplateTemplate struct {
	Inst struct {

    Template string `json:"template"`
    ThresholdCfg []ScaleoutClusterTrackingTemplateTemplateThresholdCfg `json:"threshold-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"template"`
}


type ScaleoutClusterTrackingTemplateTemplateThresholdCfg struct {
    Threshold int `json:"threshold"`
    Action string `json:"action"`
}

func (p *ScaleoutClusterTrackingTemplateTemplate) GetId() string{
    return p.Inst.Template
}

func (p *ScaleoutClusterTrackingTemplateTemplate) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/tracking-template/template"
}

func (p *ScaleoutClusterTrackingTemplateTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterTrackingTemplateTemplate::Post")
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

func (p *ScaleoutClusterTrackingTemplateTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterTrackingTemplateTemplate::Get")
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
func (p *ScaleoutClusterTrackingTemplateTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterTrackingTemplateTemplate::Put")
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

func (p *ScaleoutClusterTrackingTemplateTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterTrackingTemplateTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
