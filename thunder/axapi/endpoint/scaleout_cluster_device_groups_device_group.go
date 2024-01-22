

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterDeviceGroupsDeviceGroup struct {
	Inst struct {

    DeviceGroup int `json:"device-group"`
    DeviceIdList []ScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList `json:"device-id-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"device-group"`
}


type ScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList struct {
    DeviceIdStart int `json:"device-id-start"`
    DeviceIdEnd int `json:"device-id-end"`
}

func (p *ScaleoutClusterDeviceGroupsDeviceGroup) GetId() string{
    return strconv.Itoa(p.Inst.DeviceGroup)
}

func (p *ScaleoutClusterDeviceGroupsDeviceGroup) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/device-groups/device-group"
}

func (p *ScaleoutClusterDeviceGroupsDeviceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroupsDeviceGroup::Post")
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

func (p *ScaleoutClusterDeviceGroupsDeviceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroupsDeviceGroup::Get")
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
func (p *ScaleoutClusterDeviceGroupsDeviceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroupsDeviceGroup::Put")
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

func (p *ScaleoutClusterDeviceGroupsDeviceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroupsDeviceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
