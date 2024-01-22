

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodKerberosKdc struct {
	Inst struct {

    KerberosCfg HealthMonitorMethodKerberosKdcKerberosCfg `json:"kerberos-cfg"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"kerberos-kdc"`
}


type HealthMonitorMethodKerberosKdcKerberosCfg struct {
    Kinit int `json:"kinit"`
    KinitPricipalName string `json:"kinit-pricipal-name"`
    KinitPassword string `json:"kinit-password"`
    KinitEncrypted string `json:"kinit-encrypted"`
    KinitKdc string `json:"kinit-kdc"`
    TcpOnly int `json:"tcp-only"`
    Kadmin int `json:"kadmin"`
    KadminRealm string `json:"kadmin-realm"`
    KadminPricipalName string `json:"kadmin-pricipal-name"`
    KadminPassword string `json:"kadmin-password"`
    KadminEncrypted string `json:"kadmin-encrypted"`
    KadminServer string `json:"kadmin-server"`
    KadminKdc string `json:"kadmin-kdc"`
    Kpasswd int `json:"kpasswd"`
    KpasswdPricipalName string `json:"kpasswd-pricipal-name"`
    KpasswdPassword string `json:"kpasswd-password"`
    KpasswdEncrypted string `json:"kpasswd-encrypted"`
    KpasswdServer string `json:"kpasswd-server"`
    KpasswdKdc string `json:"kpasswd-kdc"`
}

func (p *HealthMonitorMethodKerberosKdc) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodKerberosKdc) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/kerberos-kdc"
}

func (p *HealthMonitorMethodKerberosKdc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodKerberosKdc::Post")
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

func (p *HealthMonitorMethodKerberosKdc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodKerberosKdc::Get")
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
func (p *HealthMonitorMethodKerberosKdc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodKerberosKdc::Put")
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

func (p *HealthMonitorMethodKerberosKdc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodKerberosKdc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
