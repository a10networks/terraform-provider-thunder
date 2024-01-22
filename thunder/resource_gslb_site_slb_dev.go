package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDev() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_slb_dev`: Specify a SLB device for the GSLB site\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteSlbDevCreate,
		UpdateContext: resourceGslbSiteSlbDevUpdate,
		ReadContext:   resourceGslbSiteSlbDevRead,
		DeleteContext: resourceGslbSiteSlbDevDelete,

		Schema: map[string]*schema.Schema{
			"admin_preference": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify administrative preference (Specify admin-preference value,default is 100)",
			},
			"auto_detect": {
				Type: schema.TypeString, Optional: true, Default: "ip-and-port", Description: "'ip': Service IP only; 'port': Service Port only; 'ip-and-port': Both service IP and service port; 'disabled': disable auto-detect;",
			},
			"auto_map": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable DNS Auto Mapping",
			},
			"client_ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify client IP address",
			},
			"dev_resolve_as": {
				Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN (Default Query type); 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
			},
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "Specify SLB device name",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Device hostname",
			},
			"gateway_ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"health_check_action": {
				Type: schema.TypeString, Optional: true, Default: "health-check", Description: "'health-check': Enable health Check; 'health-check-disable': Disable health check;",
			},
			"ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address",
			},
			"max_client": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Specify maximum number of clients, default is 32768",
			},
			"msg_format_acos_2x": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Protocol in compatible mode with a ACOS 2.x GSLB peer",
			},
			"probe_timer": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"proto_aging_fast": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Fast GSLB Protocol aging",
			},
			"proto_aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Specify GSLB Protocol aging time, default is 60",
			},
			"proto_compatible": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Protocol in compatible mode",
			},
			"rdt_type": {
				Type: schema.TypeString, Optional: true, Description: "'rdt': rdt; 'site-rdt': site-rdt;",
			},
			"rdt_value": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Round-delay-time",
			},
			"session_number": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"session_utilization": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip_server_v4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type: schema.TypeString, Required: true, Description: "Specify IP address",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'dev_vip_hits': Number of times the service-ip was selected; 'dev_vip_recent': Recent hits;",
												},
											},
										},
									},
								},
							},
						},
						"vip_server_v6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6": {
										Type: schema.TypeString, Required: true, Description: "Specify IP address (IPv6 address)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'dev_vip_hits': Number of times the service-ip was selected; 'dev_vip_recent': Recent hits;",
												},
											},
										},
									},
								},
							},
						},
						"vip_server_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_name": {
										Type: schema.TypeString, Required: true, Description: "Specify a VIP name for the SLB device",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'dev_vip_hits': Number of times the service-ip was selected; 'dev_vip_recent': Recent hits;",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}
func resourceGslbSiteSlbDevCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDev(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteSlbDevUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDev(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteSlbDevDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDev(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteSlbDevRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDev(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGslbSiteSlbDevVipServer396(d []interface{}) edpt.GslbSiteSlbDevVipServer396 {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevVipServer396
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VipServerV4List = getSliceGslbSiteSlbDevVipServerVipServerV4List(in["vip_server_v4_list"].([]interface{}))
		ret.VipServerV6List = getSliceGslbSiteSlbDevVipServerVipServerV6List(in["vip_server_v6_list"].([]interface{}))
		ret.VipServerNameList = getSliceGslbSiteSlbDevVipServerVipServerNameList(in["vip_server_name_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerV4List(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV4List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV4List
		oi.Ipv4 = in["ipv4"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerV4ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerV4ListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV4ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV4ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV4ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerV6List(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV6List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV6List
		oi.Ipv6 = in["ipv6"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerV6ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerV6ListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV6ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV6ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV6ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerNameList(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerNameList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerNameList
		oi.VipName = in["vip_name"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerNameListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevVipServerVipServerNameListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerNameListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerNameListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerNameListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDev(d *schema.ResourceData) edpt.GslbSiteSlbDev {
	var ret edpt.GslbSiteSlbDev
	ret.Inst.AdminPreference = d.Get("admin_preference").(int)
	ret.Inst.AutoDetect = d.Get("auto_detect").(string)
	ret.Inst.AutoMap = d.Get("auto_map").(int)
	ret.Inst.ClientIp = d.Get("client_ip").(string)
	ret.Inst.DevResolveAs = d.Get("dev_resolve_as").(string)
	ret.Inst.DeviceName = d.Get("device_name").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.GatewayIpAddr = d.Get("gateway_ip_addr").(string)
	ret.Inst.HealthCheckAction = d.Get("health_check_action").(string)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.MaxClient = d.Get("max_client").(int)
	ret.Inst.MsgFormatAcos2x = d.Get("msg_format_acos_2x").(int)
	ret.Inst.ProbeTimer = d.Get("probe_timer").(int)
	ret.Inst.ProtoAgingFast = d.Get("proto_aging_fast").(int)
	ret.Inst.ProtoAgingTime = d.Get("proto_aging_time").(int)
	ret.Inst.ProtoCompatible = d.Get("proto_compatible").(int)
	ret.Inst.RdtType = d.Get("rdt_type").(string)
	ret.Inst.RdtValue = d.Get("rdt_value").(int)
	ret.Inst.SessionNumber = d.Get("session_number").(int)
	ret.Inst.SessionUtilization = d.Get("session_utilization").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VipServer = getObjectGslbSiteSlbDevVipServer396(d.Get("vip_server").([]interface{}))
	ret.Inst.SiteName = d.Get("site_name").(string)
	return ret
}
