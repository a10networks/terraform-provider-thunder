

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterClusterDevicesDeviceId struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    DeviceId int `json:"device-id"`
    Ip string `json:"ip"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"device-id"`
}

func (p *ScaleoutClusterClusterDevicesDeviceId) GetId() string{
    return strconv.Itoa(p.Inst.DeviceId)
}

func (p *ScaleoutClusterClusterDevicesDeviceId) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/cluster-devices/device-id"
}

func (p *ScaleoutClusterClusterDevicesDeviceId) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevicesDeviceId::Post")
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

func (p *ScaleoutClusterClusterDevicesDeviceId) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevicesDeviceId::Get")
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
func (p *ScaleoutClusterClusterDevicesDeviceId) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevicesDeviceId::Put")
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

func (p *ScaleoutClusterClusterDevicesDeviceId) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevicesDeviceId::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
