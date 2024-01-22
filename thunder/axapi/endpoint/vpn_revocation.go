

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnRevocation struct {
	Inst struct {

    Ca string `json:"ca"`
    Crl VpnRevocationCrl `json:"crl"`
    Name string `json:"name"`
    Ocsp VpnRevocationOcsp `json:"ocsp"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"revocation"`
}


type VpnRevocationCrl struct {
    CrlPri string `json:"crl-pri"`
    CrlSec string `json:"crl-sec"`
}


type VpnRevocationOcsp struct {
    OcspPri string `json:"ocsp-pri"`
    OcspSec string `json:"ocsp-sec"`
}

func (p *VpnRevocation) GetId() string{
    return p.Inst.Name
}

func (p *VpnRevocation) getPath() string{
    return "vpn/revocation"
}

func (p *VpnRevocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnRevocation::Post")
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

func (p *VpnRevocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnRevocation::Get")
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
func (p *VpnRevocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnRevocation::Put")
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

func (p *VpnRevocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnRevocation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
