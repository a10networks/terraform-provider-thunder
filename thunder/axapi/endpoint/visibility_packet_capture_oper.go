

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureOper struct {
    
    AutomatedCaptures VisibilityPacketCaptureOperAutomatedCaptures `json:"automated-captures"`
    Oper VisibilityPacketCaptureOperOper `json:"oper"`

}
type DataVisibilityPacketCaptureOper struct {
    DtVisibilityPacketCaptureOper VisibilityPacketCaptureOper `json:"packet-capture"`
}


type VisibilityPacketCaptureOperAutomatedCaptures struct {
    Oper VisibilityPacketCaptureOperAutomatedCapturesOper `json:"oper"`
}


type VisibilityPacketCaptureOperAutomatedCapturesOper struct {
    CaptureName string `json:"capture-name"`
    AutomatedCaptureConfig []VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig `json:"automated-capture-config"`
}


type VisibilityPacketCaptureOperAutomatedCapturesOperAutomatedCaptureConfig struct {
    AutomatedCaptureConfigLine string `json:"automated-capture-config-line"`
}


type VisibilityPacketCaptureOperOper struct {
    PacketCaptureFiles int `json:"packet-capture-files"`
    MemoryUsage int `json:"memory-usage"`
    TotalNumberOfFiles int `json:"total-number-of-files"`
    PacketCaptureFileNameList []VisibilityPacketCaptureOperOperPacketCaptureFileNameList `json:"packet-capture-file-name-list"`
    TotalMemory string `json:"total-memory"`
    UsedMemory string `json:"used-memory"`
}


type VisibilityPacketCaptureOperOperPacketCaptureFileNameList struct {
    PacketCaptureFileName string `json:"packet-capture-file-name"`
    FileSize int `json:"file-size"`
    LastModified string `json:"last-modified"`
}

func (p *VisibilityPacketCaptureOper) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureOper) getPath() string{
    return "visibility/packet-capture/oper"
}

func (p *VisibilityPacketCaptureOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPacketCaptureOper,error) {
logger.Println("VisibilityPacketCaptureOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPacketCaptureOper
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
