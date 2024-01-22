

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDeviceExcludeInterfaces struct {
	Inst struct {

    EthCfg []ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg `json:"eth-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg `json:"loopback-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg `json:"trunk-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg `json:"ve-cfg"`
    ClusterId string 

	} `json:"exclude-interfaces"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg struct {
    Ve int `json:"ve"`
}

func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) GetId() string{
    return "1"
}

func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device/exclude-interfaces"
}

func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceExcludeInterfaces::Post")
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

func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceExcludeInterfaces::Get")
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
func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceExcludeInterfaces::Put")
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

func (p *ScaleoutClusterLocalDeviceExcludeInterfaces) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceExcludeInterfaces::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
