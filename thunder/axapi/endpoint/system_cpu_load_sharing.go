

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuLoadSharing struct {
	Inst struct {

    CpuUsage SystemCpuLoadSharingCpuUsage `json:"cpu-usage"`
    Disable int `json:"disable"`
    Others int `json:"others"`
    PacketsPerSecond SystemCpuLoadSharingPacketsPerSecond `json:"packets-per-second"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Uuid string `json:"uuid"`

	} `json:"cpu-load-sharing"`
}


type SystemCpuLoadSharingCpuUsage struct {
    Low int `json:"low" dval:"60"`
    High int `json:"high" dval:"75"`
}


type SystemCpuLoadSharingPacketsPerSecond struct {
    Min int `json:"min" dval:"100000"`
}

func (p *SystemCpuLoadSharing) GetId() string{
    return "1"
}

func (p *SystemCpuLoadSharing) getPath() string{
    return "system/cpu-load-sharing"
}

func (p *SystemCpuLoadSharing) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuLoadSharing::Post")
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

func (p *SystemCpuLoadSharing) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuLoadSharing::Get")
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
func (p *SystemCpuLoadSharing) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuLoadSharing::Put")
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

func (p *SystemCpuLoadSharing) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuLoadSharing::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
