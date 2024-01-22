

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureCaptureConfig struct {
	Inst struct {

    ConcurrentCaptures int `json:"concurrent-captures"`
    ConcurrentCapturesAge int `json:"concurrent-captures-age" dval:"1"`
    ConcurrentConnPerCapture int `json:"concurrent-conn-per-capture" dval:"100"`
    ConcurrentConnTag int `json:"concurrent-conn-tag"`
    CreatePcapFilesNow int `json:"create-pcap-files-now"`
    Disable int `json:"disable"`
    DisableAutoMerge int `json:"disable-auto-merge"`
    EnableContinuousGlobalCapture int `json:"enable-continuous-global-capture"`
    FileCount int `json:"file-count" dval:"10"`
    FileSize int `json:"file-size" dval:"1"`
    KeepPcapFilesAfterMerge int `json:"keep-pcap-files-after-merge"`
    Name string `json:"name"`
    NumberOfPacketsPerCapture int `json:"number-of-packets-per-capture"`
    NumberOfPacketsPerConn int `json:"number-of-packets-per-conn"`
    NumberOfPacketsTotal int `json:"number-of-packets-total"`
    PacketLength int `json:"packet-length" dval:"128"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"capture-config"`
}

func (p *VisibilityPacketCaptureCaptureConfig) GetId() string{
    return p.Inst.Name
}

func (p *VisibilityPacketCaptureCaptureConfig) getPath() string{
    return "visibility/packet-capture/capture-config"
}

func (p *VisibilityPacketCaptureCaptureConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureCaptureConfig::Post")
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

func (p *VisibilityPacketCaptureCaptureConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureCaptureConfig::Get")
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
func (p *VisibilityPacketCaptureCaptureConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureCaptureConfig::Put")
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

func (p *VisibilityPacketCaptureCaptureConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureCaptureConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
