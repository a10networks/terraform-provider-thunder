

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceAccountingTemplateSystemResources struct {
	Inst struct {

    BwLimitCfg SystemResourceAccountingTemplateSystemResourcesBwLimitCfg `json:"bw-limit-cfg"`
    ConcurrentSessionLimitCfg SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg `json:"concurrent-session-limit-cfg"`
    FwcpsLimitCfg SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg `json:"fwcps-limit-cfg"`
    L4SessionLimitCfg SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg `json:"l4-session-limit-cfg"`
    L4cpsLimitCfg SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg `json:"l4cps-limit-cfg"`
    L7cpsLimitCfg SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg `json:"l7cps-limit-cfg"`
    NatcpsLimitCfg SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg `json:"natcps-limit-cfg"`
    SslThroughputLimitCfg SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg `json:"ssl-throughput-limit-cfg"`
    SslcpsLimitCfg SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg `json:"sslcps-limit-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"system-resources"`
}


type SystemResourceAccountingTemplateSystemResourcesBwLimitCfg struct {
    BwLimitMax int `json:"bw-limit-max"`
    BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg struct {
    ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}


type SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg struct {
    FwcpsLimitMax int `json:"fwcps-limit-max"`
}


type SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg struct {
    L4SessionLimitMax string `json:"l4-session-limit-max"`
    L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}


type SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg struct {
    L4cpsLimitMax int `json:"l4cps-limit-max"`
}


type SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg struct {
    L7cpsLimitMax int `json:"l7cps-limit-max"`
}


type SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg struct {
    NatcpsLimitMax int `json:"natcps-limit-max"`
}


type SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg struct {
    SslThroughputLimitMax int `json:"ssl-throughput-limit-max"`
    SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg struct {
    SslcpsLimitMax int `json:"sslcps-limit-max"`
}

func (p *SystemResourceAccountingTemplateSystemResources) GetId() string{
    return "1"
}

func (p *SystemResourceAccountingTemplateSystemResources) getPath() string{
    return "system/resource-accounting/template/" +p.Inst.Name + "/system-resources"
}

func (p *SystemResourceAccountingTemplateSystemResources) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateSystemResources::Post")
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

func (p *SystemResourceAccountingTemplateSystemResources) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateSystemResources::Get")
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
func (p *SystemResourceAccountingTemplateSystemResources) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateSystemResources::Put")
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

func (p *SystemResourceAccountingTemplateSystemResources) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateSystemResources::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
