

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions struct {
	Inst struct {

    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"reachability-options"`
}

func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) GetId() string{
    return "1"
}

func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device/session-sync/reachability-options"
}

func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions::Post")
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

func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions::Get")
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
func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions::Put")
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

func (p *ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
