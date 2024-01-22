

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureAutomatedCapturesOper struct {
    
    Oper VisibilityPacketCaptureAutomatedCapturesOperOper `json:"oper"`

}
type DataVisibilityPacketCaptureAutomatedCapturesOper struct {
    DtVisibilityPacketCaptureAutomatedCapturesOper VisibilityPacketCaptureAutomatedCapturesOper `json:"automated-captures"`
}


type VisibilityPacketCaptureAutomatedCapturesOperOper struct {
    CaptureName string `json:"capture-name"`
    AutomatedCaptureConfig []VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig `json:"automated-capture-config"`
}


type VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig struct {
    AutomatedCaptureConfigLine string `json:"automated-capture-config-line"`
}

func (p *VisibilityPacketCaptureAutomatedCapturesOper) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureAutomatedCapturesOper) getPath() string{
    return "visibility/packet-capture/automated-captures/oper"
}

func (p *VisibilityPacketCaptureAutomatedCapturesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPacketCaptureAutomatedCapturesOper,error) {
logger.Println("VisibilityPacketCaptureAutomatedCapturesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPacketCaptureAutomatedCapturesOper
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
