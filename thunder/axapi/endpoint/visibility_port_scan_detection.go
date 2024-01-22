

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPortScanDetection struct {
	Inst struct {

    Events int `json:"events" dval:"10"`
    Interval int `json:"interval" dval:"60"`
    Uuid string `json:"uuid"`
    V4List string `json:"v4-list"`
    V6List string `json:"v6-list"`

	} `json:"port-scan-detection"`
}

func (p *VisibilityPortScanDetection) GetId() string{
    return "1"
}

func (p *VisibilityPortScanDetection) getPath() string{
    return "visibility/port-scan-detection"
}

func (p *VisibilityPortScanDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPortScanDetection::Post")
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

func (p *VisibilityPortScanDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPortScanDetection::Get")
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
func (p *VisibilityPortScanDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPortScanDetection::Put")
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

func (p *VisibilityPortScanDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPortScanDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
