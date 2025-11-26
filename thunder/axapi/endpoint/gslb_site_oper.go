package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type GslbSiteOper struct {
	IpServerList []GslbSiteOperIpServerList `json:"ip-server-list"`

	Oper GslbSiteOperOper `json:"oper"`

	SiteName string `json:"site-name"`

	SlbDevList []GslbSiteOperSlbDevList `json:"slb-dev-list"`
}
type DataGslbSiteOper struct {
	DtGslbSiteOper GslbSiteOper `json:"site"`
}

type GslbSiteOperIpServerList struct {
	IpServerName string                       `json:"ip-server-name"`
	Oper         GslbSiteOperIpServerListOper `json:"oper"`
}

type GslbSiteOperIpServerListOper struct {
	IpServer            string                                     `json:"ip-server"`
	IpAddress           string                                     `json:"ip-address"`
	Desc                string                                     `json:"desc"`
	State               string                                     `json:"state"`
	ServiceIp           string                                     `json:"service-ip"`
	PortCount           int                                        `json:"port-count"`
	VirtualServer       int                                        `json:"virtual-server"`
	Disabled            int                                        `json:"disabled"`
	GslbProtocol        int                                        `json:"gslb-protocol"`
	LocalProtocol       int                                        `json:"local-protocol"`
	ManuallyHealthCheck int                                        `json:"manually-health-check"`
	Use_gslb_state      int                                        `json:"use_gslb_state"`
	Dynamic             int                                        `json:"dynamic"`
	Hits                int                                        `json:"hits"`
	Recent              int                                        `json:"recent"`
	DrsList             []GslbSiteOperIpServerListOperDrsList      `json:"drs-list"`
	IpServerPort        []GslbSiteOperIpServerListOperIpServerPort `json:"ip-server-port"`
}

type GslbSiteOperIpServerListOperDrsList struct {
	DrsName                string                                       `json:"drs-name"`
	DrsIpAddress           string                                       `json:"drs-ip-address"`
	DrsFqdnName            string                                       `json:"drs-fqdn-name"`
	DrsState               string                                       `json:"drs-state"`
	DrsServiceIp           string                                       `json:"drs-service-ip"`
	DrsPortCount           int                                          `json:"drs-port-count"`
	DrsVirtualServer       int                                          `json:"drs-virtual-server"`
	DrsDisabled            int                                          `json:"drs-disabled"`
	DrsGslbProtocol        int                                          `json:"drs-gslb-protocol"`
	DrsLocalProtocol       int                                          `json:"drs-local-protocol"`
	DrsManuallyHealthCheck int                                          `json:"drs-manually-health-check"`
	DrsUse_gslb_state      int                                          `json:"drs-use_gslb_state"`
	DrsDynamic             int                                          `json:"drs-dynamic"`
	DrsHits                int                                          `json:"drs-hits"`
	DrsRecent              int                                          `json:"drs-recent"`
	DrsPort                []GslbSiteOperIpServerListOperDrsListDrsPort `json:"drs-port"`
}

type GslbSiteOperIpServerListOperDrsListDrsPort struct {
	Vport         int    `json:"vport"`
	VportProtocol string `json:"vport-protocol"`
	VportState    string `json:"vport-state"`
	ServiceName   string `json:"service-name"`
}

type GslbSiteOperIpServerListOperIpServerPort struct {
	Vport         int    `json:"vport"`
	VportProtocol string `json:"vport-protocol"`
	VportState    string `json:"vport-state"`
	ServiceName   string `json:"service-name"`
}

type GslbSiteOperOper struct {
	BwCost           int                              `json:"bw-cost"`
	GslbSite         string                           `json:"gslb-site"`
	TemplateName     string                           `json:"template-name"`
	CurrCount        int                              `json:"curr-count"`
	HighestCount     int                              `json:"highest-count"`
	Unlimited        int                              `json:"unlimited"`
	Limit            int                              `json:"limit"`
	Unusable         int                              `json:"unusable"`
	Type             string                           `json:"type"`
	Len              int                              `json:"len"`
	Value            int                              `json:"value"`
	Time             int                              `json:"time"`
	State            string                           `json:"state"`
	TotalVipCurrConn int                              `json:"total-vip-curr-conn"`
	TypeLast         []GslbSiteOperOperTypeLast       `json:"type-last"`
	ClientLdnsList   []GslbSiteOperOperClientLdnsList `json:"client-ldns-list"`
}

type GslbSiteOperOperTypeLast struct {
	Type string `json:"type"`
	Last string `json:"last"`
}

type GslbSiteOperOperClientLdnsList struct {
	ClientIp   string `json:"client-ip"`
	Age        int    `json:"age"`
	Type       string `json:"type"`
	RdtSample1 int    `json:"rdt-sample1"`
	RdtSample2 int    `json:"rdt-sample2"`
	RdtSample3 int    `json:"rdt-sample3"`
	RdtSample4 int    `json:"rdt-sample4"`
	RdtSample5 int    `json:"rdt-sample5"`
	RdtSample6 int    `json:"rdt-sample6"`
	RdtSample7 int    `json:"rdt-sample7"`
	RdtSample8 int    `json:"rdt-sample8"`
}

type GslbSiteOperSlbDevList struct {
	DeviceName string                          `json:"device-name"`
	Oper       GslbSiteOperSlbDevListOper      `json:"oper"`
	VipServer  GslbSiteOperSlbDevListVipServer `json:"vip-server"`
}

type GslbSiteOperSlbDevListOper struct {
	FqdnBased            int                                          `json:"fqdn-based"`
	Dev_name             string                                       `json:"dev_name"`
	DynamicDevList       []GslbSiteOperSlbDevListOperDynamicDevList   `json:"dynamic-dev-list"`
	DynVipserverList     []GslbSiteOperSlbDevListOperDynVipserverList `json:"dyn-vipserver-list"`
	Dev_ip               string                                       `json:"dev_ip"`
	Dev_attr             string                                       `json:"dev_attr"`
	Dev_admin_preference int                                          `json:"dev_admin_preference"`
	Dev_session_num      int                                          `json:"dev_session_num"`
	Dev_session_util     int                                          `json:"dev_session_util"`
	Dev_gw_state         string                                       `json:"dev_gw_state"`
	Dev_ip_cnt           int                                          `json:"dev_ip_cnt"`
	DevState             string                                       `json:"dev-state"`
	DevCreationType      int                                          `json:"dev-creation-type"`
	ClientLdnsList       []GslbSiteOperSlbDevListOperClientLdnsList   `json:"client-ldns-list"`
}

type GslbSiteOperSlbDevListOperDynamicDevList struct {
	DynDevName             string `json:"dyn-dev-name"`
	DynDevIp               string `json:"dyn-dev-ip"`
	DynDevInheritVipserver int    `json:"dyn-dev-inherit-vipserver"`
}

type GslbSiteOperSlbDevListOperDynVipserverList struct {
	DynSvrIp    string                                               `json:"dyn-svr-ip"`
	DynSvrState string                                               `json:"dyn-svr-state"`
	DynSvrHits  int                                                  `json:"dyn-svr-hits"`
	PortList    []GslbSiteOperSlbDevListOperDynVipserverListPortList `json:"port-list"`
}

type GslbSiteOperSlbDevListOperDynVipserverListPortList struct {
	PortNum      int    `json:"port-num"`
	PortProtocol string `json:"port-protocol"`
	PortState    string `json:"port-state"`
}

type GslbSiteOperSlbDevListOperClientLdnsList struct {
	ClientIp   string `json:"client-ip"`
	Age        int    `json:"age"`
	Type       string `json:"type"`
	RdtSample1 int    `json:"rdt-sample1"`
	RdtSample2 int    `json:"rdt-sample2"`
	RdtSample3 int    `json:"rdt-sample3"`
	RdtSample4 int    `json:"rdt-sample4"`
	RdtSample5 int    `json:"rdt-sample5"`
	RdtSample6 int    `json:"rdt-sample6"`
	RdtSample7 int    `json:"rdt-sample7"`
	RdtSample8 int    `json:"rdt-sample8"`
}

type GslbSiteOperSlbDevListVipServer struct {
	Oper              GslbSiteOperSlbDevListVipServerOper                `json:"oper"`
	VipServerV4List   []GslbSiteOperSlbDevListVipServerVipServerV4List   `json:"vip-server-v4-list"`
	VipServerV6List   []GslbSiteOperSlbDevListVipServerVipServerV6List   `json:"vip-server-v6-list"`
	VipServerNameList []GslbSiteOperSlbDevListVipServerVipServerNameList `json:"vip-server-name-list"`
}

type GslbSiteOperSlbDevListVipServerOper struct {
}

type GslbSiteOperSlbDevListVipServerVipServerV4List struct {
	Ipv4 string                                             `json:"ipv4"`
	Oper GslbSiteOperSlbDevListVipServerVipServerV4ListOper `json:"oper"`
}

type GslbSiteOperSlbDevListVipServerVipServerV4ListOper struct {
	Dev_vip_addr        string                                                             `json:"dev_vip_addr"`
	Dev_vip_state       string                                                             `json:"dev_vip_state"`
	NodeName            string                                                             `json:"node-name"`
	ServiceIp           string                                                             `json:"service-ip"`
	PortCount           int                                                                `json:"port-count"`
	VirtualServer       int                                                                `json:"virtual-server"`
	Disabled            int                                                                `json:"disabled"`
	GslbProtocol        int                                                                `json:"gslb-protocol"`
	LocalProtocol       int                                                                `json:"local-protocol"`
	ManuallyHealthCheck int                                                                `json:"manually-health-check"`
	Use_gslb_state      int                                                                `json:"use_gslb_state"`
	Dynamic             int                                                                `json:"dynamic"`
	Shared              int                                                                `json:"shared"`
	Hits                int                                                                `json:"hits"`
	Recent              int                                                                `json:"recent"`
	DevVipPortList      []GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList `json:"dev-vip-port-list"`
}

type GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList struct {
	DevVipPortNum         int    `json:"dev-vip-port-num"`
	DevVipPortState       string `json:"dev-vip-port-state"`
	DevVipPortDevCurrConn int    `json:"dev-vip-port-dev-curr-conn"`
	DevVipPortProtocol    string `json:"dev-vip-port-protocol"`
	DevVipPortServiceName string `json:"dev-vip-port-service-name"`
}

type GslbSiteOperSlbDevListVipServerVipServerV6List struct {
	Ipv6 string                                             `json:"ipv6"`
	Oper GslbSiteOperSlbDevListVipServerVipServerV6ListOper `json:"oper"`
}

type GslbSiteOperSlbDevListVipServerVipServerV6ListOper struct {
	Dev_vip_addr        string                                                             `json:"dev_vip_addr"`
	Dev_vip_state       string                                                             `json:"dev_vip_state"`
	NodeName            string                                                             `json:"node-name"`
	ServiceIp           string                                                             `json:"service-ip"`
	PortCount           int                                                                `json:"port-count"`
	VirtualServer       int                                                                `json:"virtual-server"`
	Disabled            int                                                                `json:"disabled"`
	GslbProtocol        int                                                                `json:"gslb-protocol"`
	LocalProtocol       int                                                                `json:"local-protocol"`
	ManuallyHealthCheck int                                                                `json:"manually-health-check"`
	Use_gslb_state      int                                                                `json:"use_gslb_state"`
	Dynamic             int                                                                `json:"dynamic"`
	Shared              int                                                                `json:"shared"`
	Hits                int                                                                `json:"hits"`
	Recent              int                                                                `json:"recent"`
	DevVipPortList      []GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList `json:"dev-vip-port-list"`
}

type GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList struct {
	DevVipPortNum         int    `json:"dev-vip-port-num"`
	DevVipPortState       string `json:"dev-vip-port-state"`
	DevVipPortDevCurrConn int    `json:"dev-vip-port-dev-curr-conn"`
	DevVipPortProtocol    string `json:"dev-vip-port-protocol"`
	DevVipPortServiceName string `json:"dev-vip-port-service-name"`
}

type GslbSiteOperSlbDevListVipServerVipServerNameList struct {
	VipName string                                               `json:"vip-name"`
	Oper    GslbSiteOperSlbDevListVipServerVipServerNameListOper `json:"oper"`
}

type GslbSiteOperSlbDevListVipServerVipServerNameListOper struct {
	Dev_vip_addr        string                                                               `json:"dev_vip_addr"`
	Dev_vip_state       string                                                               `json:"dev_vip_state"`
	NodeName            string                                                               `json:"node-name"`
	ServiceIp           string                                                               `json:"service-ip"`
	PortCount           int                                                                  `json:"port-count"`
	VirtualServer       int                                                                  `json:"virtual-server"`
	Disabled            int                                                                  `json:"disabled"`
	GslbProtocol        int                                                                  `json:"gslb-protocol"`
	LocalProtocol       int                                                                  `json:"local-protocol"`
	ManuallyHealthCheck int                                                                  `json:"manually-health-check"`
	Use_gslb_state      int                                                                  `json:"use_gslb_state"`
	Dynamic             int                                                                  `json:"dynamic"`
	Shared              int                                                                  `json:"shared"`
	Hits                int                                                                  `json:"hits"`
	Recent              int                                                                  `json:"recent"`
	DevVipPortList      []GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList `json:"dev-vip-port-list"`
}

type GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList struct {
	DevVipPortNum         int    `json:"dev-vip-port-num"`
	DevVipPortState       string `json:"dev-vip-port-state"`
	DevVipPortDevCurrConn int    `json:"dev-vip-port-dev-curr-conn"`
	DevVipPortProtocol    string `json:"dev-vip-port-protocol"`
	DevVipPortServiceName string `json:"dev-vip-port-service-name"`
}

func (p *GslbSiteOper) GetId() string {
	return "1"
}

func (p *GslbSiteOper) getPath() string {
	return "gslb/site/" + p.SiteName + "/oper"
}

func (p *GslbSiteOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteOper, error) {
	logger.Println("GslbSiteOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbSiteOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
