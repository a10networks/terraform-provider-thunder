

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VcsDevice struct {
	Inst struct {

    AffinityVrrpAVrid int `json:"affinity-vrrp-a-vrid"`
    Device int `json:"device"`
    Enable int `json:"enable"`
    EthernetCfg []VcsDeviceEthernetCfg `json:"ethernet-cfg"`
    Management int `json:"management"`
    Priority int `json:"priority"`
    TrunkCfg []VcsDeviceTrunkCfg `json:"trunk-cfg"`
    Ttl int `json:"ttl" dval:"64"`
    UnicastPort int `json:"unicast-port" dval:"41216"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VeCfg []VcsDeviceVeCfg `json:"ve-cfg"`

	} `json:"device"`
}


type VcsDeviceEthernetCfg struct {
    Ethernet int `json:"ethernet"`
}


type VcsDeviceTrunkCfg struct {
    Trunk int `json:"trunk"`
}


type VcsDeviceVeCfg struct {
    Ve int `json:"ve"`
}

func (p *VcsDevice) GetId() string{
    return strconv.Itoa(p.Inst.Device)
}

func (p *VcsDevice) getPath() string{
    return "vcs/device"
}

func (p *VcsDevice) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsDevice::Post")
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

func (p *VcsDevice) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsDevice::Get")
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
func (p *VcsDevice) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsDevice::Put")
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

func (p *VcsDevice) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsDevice::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
