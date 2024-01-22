

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDeviceSessionSyncInterfaces struct {
	Inst struct {

    EthCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg `json:"eth-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg `json:"loopback-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg `json:"trunk-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg `json:"ve-cfg"`
    ClusterId string 

	} `json:"interfaces"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg struct {
    Ve int `json:"ve"`
}

func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) GetId() string{
    return "1"
}

func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device/session-sync/interfaces"
}

func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncInterfaces::Post")
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

func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncInterfaces::Get")
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
func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncInterfaces::Put")
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

func (p *ScaleoutClusterLocalDeviceSessionSyncInterfaces) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceSessionSyncInterfaces::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
