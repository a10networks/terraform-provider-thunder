

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodLdap struct {
	Inst struct {

    Acceptnotfound int `json:"AcceptNotFound"`
    Acceptresref int `json:"AcceptResRef"`
    Basedn string `json:"BaseDN"`
    Ldap int `json:"ldap"`
    LdapBinddn string `json:"ldap-binddn"`
    LdapEncrypted string `json:"ldap-encrypted"`
    LdapPassword int `json:"ldap-password"`
    LdapPasswordString string `json:"ldap-password-string"`
    LdapPort int `json:"ldap-port" dval:"389"`
    LdapQuery string `json:"ldap-query"`
    LdapRunSearch int `json:"ldap-run-search"`
    LdapSecurity string `json:"ldap-security"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ldap"`
}

func (p *HealthMonitorMethodLdap) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodLdap) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/ldap"
}

func (p *HealthMonitorMethodLdap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodLdap::Post")
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

func (p *HealthMonitorMethodLdap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodLdap::Get")
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
func (p *HealthMonitorMethodLdap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodLdap::Put")
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

func (p *HealthMonitorMethodLdap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodLdap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
