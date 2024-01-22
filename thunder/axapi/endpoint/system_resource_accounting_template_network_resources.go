

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceAccountingTemplateNetworkResources struct {
	Inst struct {

    Ipv4AclLineCfg SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg `json:"ipv4-acl-line-cfg"`
    Ipv6AclLineCfg SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg `json:"ipv6-acl-line-cfg"`
    ObjectGroupCfg SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg `json:"object-group-cfg"`
    ObjectGroupClauseCfg SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg `json:"object-group-clause-cfg"`
    StaticArpCfg SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg `json:"static-arp-cfg"`
    StaticIpv4RouteCfg SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg `json:"static-ipv4-route-cfg"`
    StaticIpv6RouteCfg SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg `json:"static-ipv6-route-cfg"`
    StaticMacCfg SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg `json:"static-mac-cfg"`
    StaticNeighborCfg SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg `json:"static-neighbor-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"network-resources"`
}


type SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg struct {
    Ipv4AclLineMax int `json:"ipv4-acl-line-max"`
    Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg struct {
    Ipv6AclLineMax int `json:"ipv6-acl-line-max"`
    Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg struct {
    ObjectGroupMax int `json:"object-group-max"`
    ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg struct {
    ObjectGroupClauseMax int `json:"object-group-clause-max"`
    ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg struct {
    StaticArpMax int `json:"static-arp-max"`
    StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg struct {
    StaticIpv4RouteMax int `json:"static-ipv4-route-max"`
    StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg struct {
    StaticIpv6RouteMax int `json:"static-ipv6-route-max"`
    StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg struct {
    StaticMacMax int `json:"static-mac-max"`
    StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}


type SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg struct {
    StaticNeighborMax int `json:"static-neighbor-max"`
    StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}

func (p *SystemResourceAccountingTemplateNetworkResources) GetId() string{
    return "1"
}

func (p *SystemResourceAccountingTemplateNetworkResources) getPath() string{
    return "system/resource-accounting/template/" +p.Inst.Name + "/network-resources"
}

func (p *SystemResourceAccountingTemplateNetworkResources) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateNetworkResources::Post")
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

func (p *SystemResourceAccountingTemplateNetworkResources) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateNetworkResources::Get")
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
func (p *SystemResourceAccountingTemplateNetworkResources) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateNetworkResources::Put")
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

func (p *SystemResourceAccountingTemplateNetworkResources) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateNetworkResources::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
