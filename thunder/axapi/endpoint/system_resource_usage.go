

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceUsage struct {
	Inst struct {

    AflexTableEntryCount int `json:"aflex-table-entry-count"`
    AuthPortalHtmlFileSize int `json:"auth-portal-html-file-size" dval:"20"`
    AuthPortalImageFileSize int `json:"auth-portal-image-file-size" dval:"6"`
    AuthSessionCount int `json:"auth-session-count"`
    AuthzPolicyNumber int `json:"authz-policy-number"`
    ClassListAcEntryCount int `json:"class-list-ac-entry-count"`
    ClassListEntryCount int `json:"class-list-entry-count"`
    ClassListIpv6AddrCount int `json:"class-list-ipv6-addr-count"`
    IpsecSaNumber int `json:"ipsec-sa-number"`
    L4SessionCount int `json:"l4-session-count"`
    MaxAflexAuthzCollectionNumber int `json:"max-aflex-authz-collection-number" dval:"512"`
    MaxAflexFileSize int `json:"max-aflex-file-size" dval:"32"`
    NatPoolAddrCount int `json:"nat-pool-addr-count"`
    RadiusTableSize int `json:"radius-table-size"`
    RamCacheMemoryLimit int `json:"ram-cache-memory-limit"`
    SslContextMemory int `json:"ssl-context-memory" dval:"2048"`
    SslDmaMemory int `json:"ssl-dma-memory" dval:"256"`
    Uuid string `json:"uuid"`
    Visibility SystemResourceUsageVisibility1645 `json:"visibility"`

	} `json:"resource-usage"`
}


type SystemResourceUsageVisibility1645 struct {
    MonitoredEntityCount int `json:"monitored-entity-count"`
    Uuid string `json:"uuid"`
}

func (p *SystemResourceUsage) GetId() string{
    return "1"
}

func (p *SystemResourceUsage) getPath() string{
    return "system/resource-usage"
}

func (p *SystemResourceUsage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceUsage::Post")
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

func (p *SystemResourceUsage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceUsage::Get")
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
func (p *SystemResourceUsage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceUsage::Put")
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

func (p *SystemResourceUsage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceUsage::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
