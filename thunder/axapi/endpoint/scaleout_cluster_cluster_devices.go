

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterClusterDevices struct {
	Inst struct {

    ClusterDiscoveryTimeout ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325 `json:"cluster-discovery-timeout"`
    DeviceIdList []ScaleoutClusterClusterDevicesDeviceIdList `json:"device-id-list"`
    Enable int `json:"enable"`
    MinimumNodes ScaleoutClusterClusterDevicesMinimumNodes1326 `json:"minimum-nodes"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"cluster-devices"`
}


type ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325 struct {
    Uuid string `json:"uuid"`
}


type ScaleoutClusterClusterDevicesDeviceIdList struct {
    Ip string `json:"ip"`
    Action string `json:"action" dval:"enable"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterClusterDevicesMinimumNodes1326 struct {
    MinimumNodesNum int `json:"minimum-nodes-num"`
    Uuid string `json:"uuid"`
}

func (p *ScaleoutClusterClusterDevices) GetId() string{
    return "1"
}

func (p *ScaleoutClusterClusterDevices) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/cluster-devices"
}

func (p *ScaleoutClusterClusterDevices) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevices::Post")
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

func (p *ScaleoutClusterClusterDevices) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevices::Get")
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
func (p *ScaleoutClusterClusterDevices) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevices::Put")
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

func (p *ScaleoutClusterClusterDevices) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterClusterDevices::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
