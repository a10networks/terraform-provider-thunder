

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDeviceTrafficRedirection struct {
	Inst struct {

    FollowShared int `json:"follow-shared"`
    Interfaces ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333 `json:"interfaces"`
    ReachabilityOptions ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338 `json:"reachability-options"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"traffic-redirection"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333 struct {
    EthCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338 struct {
    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
}

func (p *ScaleoutClusterLocalDeviceTrafficRedirection) GetId() string{
    return "1"
}

func (p *ScaleoutClusterLocalDeviceTrafficRedirection) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device/traffic-redirection"
}

func (p *ScaleoutClusterLocalDeviceTrafficRedirection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrafficRedirection::Post")
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

func (p *ScaleoutClusterLocalDeviceTrafficRedirection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrafficRedirection::Get")
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
func (p *ScaleoutClusterLocalDeviceTrafficRedirection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrafficRedirection::Put")
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

func (p *ScaleoutClusterLocalDeviceTrafficRedirection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterLocalDeviceTrafficRedirection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
