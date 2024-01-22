package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSite() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site`: Specify a GSLB site\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteCreate,
		UpdateContext: resourceGslbSiteUpdate,
		ReadContext:   resourceGslbSiteRead,
		DeleteContext: resourceGslbSiteDelete,

		Schema: map[string]*schema.Schema{
			"active_rdt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aging_time": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Aging Time, Unit: min, default is 10",
						},
						"smooth_factor": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Factor of Smooth RDT, default is 10",
						},
						"range_factor": {
							Type: schema.TypeInt, Optional: true, Default: 25, Description: "Factor of RDT Range, default is 25 (Range Factor of Smooth RDT)",
						},
						"limit": {
							Type: schema.TypeInt, Optional: true, Default: 16383, Description: "Limit of valid RDT, default is 16383 (Limit, unit: millisecond)",
						},
						"mask": {
							Type: schema.TypeString, Optional: true, Default: "/32", Description: "Client IP subnet mask, default is 32",
						},
						"ipv6_mask": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "Client IPv6 subnet mask, default is 128",
						},
						"ignore_count": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Ignore count if RDT is out of range, default is 5",
						},
						"bind_geoloc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bind RDT to geo-location",
						},
						"overlap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable overlap for geo-location to do longest match",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"auto_map": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable DNS Auto Mapping",
			},
			"bw_cost": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify cost of band-width",
			},
			"controller": {
				Type: schema.TypeString, Optional: true, Description: "Specify the local controller for the GSLB site (Specify the hostname of the local controller)",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all servers in the GSLB site",
			},
			"easy_rdt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aging_time": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Aging Time, Unit: min, default is 10",
						},
						"smooth_factor": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Factor of Smooth RDT, default is 10",
						},
						"range_factor": {
							Type: schema.TypeInt, Optional: true, Default: 25, Description: "Factor of RDT Range, default is 25 (Range Factor of Smooth RDT)",
						},
						"limit": {
							Type: schema.TypeInt, Optional: true, Default: 16383, Description: "Limit of valid RDT, default is 16383 (Limit, unit: millisecond)",
						},
						"mask": {
							Type: schema.TypeString, Optional: true, Default: "/32", Description: "Client IP subnet mask, default is 32",
						},
						"ipv6_mask": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "Client IPv6 subnet mask, default is 128",
						},
						"ignore_count": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Ignore count if RDT is out of range, default is 5",
						},
						"bind_geoloc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bind RDT to geo-location",
						},
						"overlap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable overlap for geo-location to do longest match",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ip_server_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_server_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the real server name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the limit for bandwidth, default is unlimited",
			},
			"multiple_geo_locations": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type: schema.TypeString, Optional: true, Description: "Specify the geographic location of the GSLB site (Specify geo-location for this site)",
						},
					},
				},
			},
			"proto_aging_fast": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fast GSLB Protocol aging",
			},
			"proto_aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify GSLB Protocol aging time",
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "Specify GSLB site name",
			},
			"slb_dev_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_name": {
							Type: schema.TypeString, Required: true, Description: "Specify SLB device name",
						},
						"ip_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 address",
						},
						"domain": {
							Type: schema.TypeString, Optional: true, Description: "Device hostname",
						},
						"dev_resolve_as": {
							Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN (Default Query type); 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
						},
						"admin_preference": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify administrative preference (Specify admin-preference value,default is 100)",
						},
						"session_number": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"session_utilization": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rdt_type": {
							Type: schema.TypeString, Optional: true, Description: "'rdt': rdt; 'site-rdt': site-rdt;",
						},
						"client_ip": {
							Type: schema.TypeString, Optional: true, Description: "Specify client IP address",
						},
						"rdt_value": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Round-delay-time",
						},
						"probe_timer": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auto_detect": {
							Type: schema.TypeString, Optional: true, Default: "ip-and-port", Description: "'ip': Service IP only; 'port': Service Port only; 'ip-and-port': Both service IP and service port; 'disabled': disable auto-detect;",
						},
						"auto_map": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable DNS Auto Mapping",
						},
						"max_client": {
							Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Specify maximum number of clients, default is 32768",
						},
						"proto_aging_time": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Specify GSLB Protocol aging time, default is 60",
						},
						"proto_aging_fast": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Fast GSLB Protocol aging",
						},
						"health_check_action": {
							Type: schema.TypeString, Optional: true, Default: "health-check", Description: "'health-check': Enable health Check; 'health-check-disable': Disable health check;",
						},
						"gateway_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"proto_compatible": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Protocol in compatible mode",
						},
						"msg_format_acos_2x": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Protocol in compatible mode with a ACOS 2.x GSLB peer",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"template": {
				Type: schema.TypeString, Optional: true, Description: "Specify template to collect site information (Specify GSLB SNMP template name)",
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the threshold for limit",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify a weight for the GSLB site (Weight, default is 1)",
			},
		},
	}
}
func resourceGslbSiteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSite(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSite(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSite(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSite(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGslbSiteActiveRdt397(d []interface{}) edpt.GslbSiteActiveRdt397 {

	count1 := len(d)
	var ret edpt.GslbSiteActiveRdt397
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgingTime = in["aging_time"].(int)
		ret.SmoothFactor = in["smooth_factor"].(int)
		ret.RangeFactor = in["range_factor"].(int)
		ret.Limit = in["limit"].(int)
		ret.Mask = in["mask"].(string)
		ret.Ipv6Mask = in["ipv6_mask"].(int)
		ret.IgnoreCount = in["ignore_count"].(int)
		ret.BindGeoloc = in["bind_geoloc"].(int)
		ret.Overlap = in["overlap"].(int)
		//omit uuid
	}
	return ret
}

func getObjectGslbSiteEasyRdt398(d []interface{}) edpt.GslbSiteEasyRdt398 {

	count1 := len(d)
	var ret edpt.GslbSiteEasyRdt398
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgingTime = in["aging_time"].(int)
		ret.SmoothFactor = in["smooth_factor"].(int)
		ret.RangeFactor = in["range_factor"].(int)
		ret.Limit = in["limit"].(int)
		ret.Mask = in["mask"].(string)
		ret.Ipv6Mask = in["ipv6_mask"].(int)
		ret.IgnoreCount = in["ignore_count"].(int)
		ret.BindGeoloc = in["bind_geoloc"].(int)
		ret.Overlap = in["overlap"].(int)
		//omit uuid
	}
	return ret
}

func getSliceGslbSiteIpServerList(d []interface{}) []edpt.GslbSiteIpServerList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteIpServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteIpServerList
		oi.IpServerName = in["ip_server_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteMultipleGeoLocations(d []interface{}) []edpt.GslbSiteMultipleGeoLocations {

	count1 := len(d)
	ret := make([]edpt.GslbSiteMultipleGeoLocations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteMultipleGeoLocations
		oi.GeoLocation = in["geo_location"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevList(d []interface{}) []edpt.GslbSiteSlbDevList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevList
		oi.DeviceName = in["device_name"].(string)
		oi.IpAddress = in["ip_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.Domain = in["domain"].(string)
		oi.DevResolveAs = in["dev_resolve_as"].(string)
		oi.AdminPreference = in["admin_preference"].(int)
		oi.SessionNumber = in["session_number"].(int)
		oi.SessionUtilization = in["session_utilization"].(int)
		oi.RdtType = in["rdt_type"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.RdtValue = in["rdt_value"].(int)
		oi.ProbeTimer = in["probe_timer"].(int)
		oi.AutoDetect = in["auto_detect"].(string)
		oi.AutoMap = in["auto_map"].(int)
		oi.MaxClient = in["max_client"].(int)
		oi.ProtoAgingTime = in["proto_aging_time"].(int)
		oi.ProtoAgingFast = in["proto_aging_fast"].(int)
		oi.HealthCheckAction = in["health_check_action"].(string)
		oi.GatewayIpAddr = in["gateway_ip_addr"].(string)
		oi.ProtoCompatible = in["proto_compatible"].(int)
		oi.MsgFormatAcos2x = in["msg_format_acos_2x"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.VipServer = getObjectGslbSiteSlbDevListVipServer(in["vip_server"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteSlbDevListVipServer(d []interface{}) edpt.GslbSiteSlbDevListVipServer {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevListVipServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VipServerV4List = getSliceGslbSiteSlbDevListVipServerVipServerV4List(in["vip_server_v4_list"].([]interface{}))
		ret.VipServerV6List = getSliceGslbSiteSlbDevListVipServerVipServerV6List(in["vip_server_v6_list"].([]interface{}))
		ret.VipServerNameList = getSliceGslbSiteSlbDevListVipServerVipServerNameList(in["vip_server_name_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerV4List(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerV4List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerV4List
		oi.Ipv4 = in["ipv4"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerV6List(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerV6List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerV6List
		oi.Ipv6 = in["ipv6"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerNameList(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerNameList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerNameList
		oi.VipName = in["vip_name"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceGslbSiteSlbDevListVipServerVipServerNameListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevListVipServerVipServerNameListSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevListVipServerVipServerNameListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevListVipServerVipServerNameListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevListVipServerVipServerNameListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSite(d *schema.ResourceData) edpt.GslbSite {
	var ret edpt.GslbSite
	ret.Inst.ActiveRdt = getObjectGslbSiteActiveRdt397(d.Get("active_rdt").([]interface{}))
	ret.Inst.AutoMap = d.Get("auto_map").(int)
	ret.Inst.BwCost = d.Get("bw_cost").(int)
	ret.Inst.Controller = d.Get("controller").(string)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.EasyRdt = getObjectGslbSiteEasyRdt398(d.Get("easy_rdt").([]interface{}))
	ret.Inst.IpServerList = getSliceGslbSiteIpServerList(d.Get("ip_server_list").([]interface{}))
	ret.Inst.Limit = d.Get("limit").(int)
	ret.Inst.MultipleGeoLocations = getSliceGslbSiteMultipleGeoLocations(d.Get("multiple_geo_locations").([]interface{}))
	ret.Inst.ProtoAgingFast = d.Get("proto_aging_fast").(int)
	ret.Inst.ProtoAgingTime = d.Get("proto_aging_time").(int)
	ret.Inst.SiteName = d.Get("site_name").(string)
	ret.Inst.SlbDevList = getSliceGslbSiteSlbDevList(d.Get("slb_dev_list").([]interface{}))
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	return ret
}
