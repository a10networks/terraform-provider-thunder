

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate struct {
	Inst struct {

    Action string `json:"action"`
    MultiTemplate string `json:"multi-template"`
    Template []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate `json:"template"`
    Threshold int `json:"threshold"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"multi-template"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate struct {
    TemplateName string `json:"template-name"`
    PartitionName string `json:"partition-name"`
}

func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) GetId() string{
    return p.Inst.MultiTemplate
}

func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device/tracking-template/multi-template"
}

func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate::Post")
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

func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate::Get")
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
func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate::Put")
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

func (p *ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
