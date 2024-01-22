package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbResourceUsage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_resource_usage`: Configure SLB Resource Usage\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbResourceUsageCreate,
		UpdateContext: resourceSlbResourceUsageUpdate,
		ReadContext:   resourceSlbResourceUsageRead,
		DeleteContext: resourceSlbResourceUsageDelete,

		Schema: map[string]*schema.Schema{
			"cache_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable HTTP Cache Templates in the System",
			},
			"client_ssl_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Client SSL Templates in the System",
			},
			"conn_reuse_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Connection reuse Templates in the System",
			},
			"fast_tcp_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Fast TCP Templates in the System",
			},
			"fast_udp_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Fast UDP Templates in the System",
			},
			"fix_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable FIX Templates in the System",
			},
			"gslb_device_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB devices in the System",
			},
			"gslb_geo_location_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB geo-location in the System",
			},
			"gslb_ip_list_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB ip-list in the System",
			},
			"gslb_policy_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB policies in the System",
			},
			"gslb_service_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB services in the System",
			},
			"gslb_service_ip_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB service-ip in the System",
			},
			"gslb_service_port_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB service-port in the System",
			},
			"gslb_site_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB sites in the System",
			},
			"gslb_svc_group_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB services in the System",
			},
			"gslb_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB templates in the System",
			},
			"gslb_zone_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total GSLB zones in the System",
			},
			"health_monitor_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Health Monitors in the System",
			},
			"http_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable HTTP Templates in the System",
			},
			"link_cost_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Link-cost Templates in the System",
			},
			"nat_pool_addr_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable NAT Pool addresses in the System (deprecated)",
			},
			"pbslb_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable pbslb entry in the System",
			},
			"pbslb_subnet_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total PBSLB Subnets in the System",
			},
			"persist_cookie_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Persistent cookie Templates in the System",
			},
			"persist_srcip_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Source IP Persistent Templates in the System",
			},
			"proxy_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Proxy Templates in the System",
			},
			"real_port_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Real Server Ports in the System",
			},
			"real_server_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Real Servers in the System",
			},
			"server_ssl_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Server SSL Templates in the System",
			},
			"service_group_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Service Groups in the System",
			},
			"slb_threshold_res_usage_percent": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 0%))",
			},
			"stream_template_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable Streaming media in the System",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_port_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Virtual Server Ports in the System",
			},
			"virtual_server_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Virtual Servers in the System",
			},
		},
	}
}
func resourceSlbResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResourceUsageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResourceUsage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbResourceUsageRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResourceUsageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResourceUsage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbResourceUsageRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResourceUsageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResourceUsage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResourceUsageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResourceUsage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbResourceUsage(d *schema.ResourceData) edpt.SlbResourceUsage {
	var ret edpt.SlbResourceUsage
	ret.Inst.CacheTemplateCount = d.Get("cache_template_count").(int)
	ret.Inst.ClientSslTemplateCount = d.Get("client_ssl_template_count").(int)
	ret.Inst.ConnReuseTemplateCount = d.Get("conn_reuse_template_count").(int)
	ret.Inst.FastTcpTemplateCount = d.Get("fast_tcp_template_count").(int)
	ret.Inst.FastUdpTemplateCount = d.Get("fast_udp_template_count").(int)
	ret.Inst.FixTemplateCount = d.Get("fix_template_count").(int)
	ret.Inst.GslbDeviceCount = d.Get("gslb_device_count").(int)
	ret.Inst.GslbGeoLocationCount = d.Get("gslb_geo_location_count").(int)
	ret.Inst.GslbIpListCount = d.Get("gslb_ip_list_count").(int)
	ret.Inst.GslbPolicyCount = d.Get("gslb_policy_count").(int)
	ret.Inst.GslbServiceCount = d.Get("gslb_service_count").(int)
	ret.Inst.GslbServiceIpCount = d.Get("gslb_service_ip_count").(int)
	ret.Inst.GslbServicePortCount = d.Get("gslb_service_port_count").(int)
	ret.Inst.GslbSiteCount = d.Get("gslb_site_count").(int)
	ret.Inst.GslbSvcGroupCount = d.Get("gslb_svc_group_count").(int)
	ret.Inst.GslbTemplateCount = d.Get("gslb_template_count").(int)
	ret.Inst.GslbZoneCount = d.Get("gslb_zone_count").(int)
	ret.Inst.HealthMonitorCount = d.Get("health_monitor_count").(int)
	ret.Inst.HttpTemplateCount = d.Get("http_template_count").(int)
	ret.Inst.LinkCostTemplateCount = d.Get("link_cost_template_count").(int)
	ret.Inst.NatPoolAddrCount = d.Get("nat_pool_addr_count").(int)
	ret.Inst.PbslbEntryCount = d.Get("pbslb_entry_count").(int)
	ret.Inst.PbslbSubnetCount = d.Get("pbslb_subnet_count").(int)
	ret.Inst.PersistCookieTemplateCount = d.Get("persist_cookie_template_count").(int)
	ret.Inst.PersistSrcipTemplateCount = d.Get("persist_srcip_template_count").(int)
	ret.Inst.ProxyTemplateCount = d.Get("proxy_template_count").(int)
	ret.Inst.RealPortCount = d.Get("real_port_count").(int)
	ret.Inst.RealServerCount = d.Get("real_server_count").(int)
	ret.Inst.ServerSslTemplateCount = d.Get("server_ssl_template_count").(int)
	ret.Inst.ServiceGroupCount = d.Get("service_group_count").(int)
	ret.Inst.SlbThresholdResUsagePercent = d.Get("slb_threshold_res_usage_percent").(int)
	ret.Inst.StreamTemplateCount = d.Get("stream_template_count").(int)
	//omit uuid
	ret.Inst.VirtualPortCount = d.Get("virtual_port_count").(int)
	ret.Inst.VirtualServerCount = d.Get("virtual_server_count").(int)
	return ret
}
