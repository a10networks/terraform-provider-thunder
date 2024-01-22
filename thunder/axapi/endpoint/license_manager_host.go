

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LicenseManagerHost struct {
	Inst struct {

    HostIpv4 string `json:"host-ipv4"`
    HostIpv6 string `json:"host-ipv6"`
    Port int `json:"port" dval:"443"`
    Uuid string `json:"uuid"`

	} `json:"host"`
}

func (p *LicenseManagerHost) GetId() string{
    return p.Inst.HostIpv4+"+"+p.Inst.HostIpv6
}

func (p *LicenseManagerHost) getPath() string{
    return "license-manager/host"
}

func (p *LicenseManagerHost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerHost::Post")
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

func (p *LicenseManagerHost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerHost::Get")
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
func (p *LicenseManagerHost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerHost::Put")
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

func (p *LicenseManagerHost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerHost::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
