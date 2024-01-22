

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterDeviceGroups struct {
	Inst struct {

    DeviceGroupList []ScaleoutClusterDeviceGroupsDeviceGroupList `json:"device-group-list"`
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"device-groups"`
}


type ScaleoutClusterDeviceGroupsDeviceGroupList struct {
    DeviceGroup int `json:"device-group"`
    DeviceIdList []ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList `json:"device-id-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList struct {
    DeviceIdStart int `json:"device-id-start"`
    DeviceIdEnd int `json:"device-id-end"`
}

func (p *ScaleoutClusterDeviceGroups) GetId() string{
    return "1"
}

func (p *ScaleoutClusterDeviceGroups) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/device-groups"
}

func (p *ScaleoutClusterDeviceGroups) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroups::Post")
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

func (p *ScaleoutClusterDeviceGroups) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroups::Get")
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
func (p *ScaleoutClusterDeviceGroups) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroups::Put")
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

func (p *ScaleoutClusterDeviceGroups) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDeviceGroups::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
