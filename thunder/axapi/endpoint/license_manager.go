

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LicenseManager struct {
	Inst struct {

    BandwidthBase int `json:"bandwidth-base"`
    BandwidthUnrestricted int `json:"bandwidth-unrestricted"`
    Connect LicenseManagerConnect1046 `json:"connect"`
    HostList []LicenseManagerHostList `json:"host-list"`
    InstanceName string `json:"instance-name"`
    Interval int `json:"interval"`
    NgWafModule LicenseManagerNgWafModule1047 `json:"ng-waf-module"`
    Overage LicenseManagerOverage1048 `json:"overage"`
    ReminderList []LicenseManagerReminderList `json:"reminder-list"`
    Sn string `json:"sn"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"license-manager"`
}


type LicenseManagerConnect1046 struct {
    Connect int `json:"connect"`
    Uuid string `json:"uuid"`
}


type LicenseManagerHostList struct {
    HostIpv4 string `json:"host-ipv4"`
    HostIpv6 string `json:"host-ipv6"`
    Port int `json:"port" dval:"443"`
    Uuid string `json:"uuid"`
}


type LicenseManagerNgWafModule1047 struct {
    AccessKeyId string `json:"access-key-id"`
    SecretAccessKey string `json:"secret-access-key"`
}


type LicenseManagerOverage1048 struct {
    Days int `json:"days"`
    Hours int `json:"hours"`
    Minutes int `json:"minutes"`
    Seconds int `json:"seconds"`
    Gb int `json:"gb"`
    Mb int `json:"mb"`
    Kb int `json:"kb"`
    Bytes int `json:"bytes"`
    Uuid string `json:"uuid"`
}


type LicenseManagerReminderList struct {
    ReminderValue int `json:"reminder-value"`
    Uuid string `json:"uuid"`
}

func (p *LicenseManager) GetId() string{
    return "1"
}

func (p *LicenseManager) getPath() string{
    return "license-manager"
}

func (p *LicenseManager) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManager::Post")
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

func (p *LicenseManager) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManager::Get")
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
func (p *LicenseManager) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManager::Put")
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

func (p *LicenseManager) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManager::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
