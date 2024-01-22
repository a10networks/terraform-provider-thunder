package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccountingTemplateAppResources() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_accounting_template_app_resources`: Enter the application resource limits\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceAccountingTemplateAppResourcesCreate,
		UpdateContext: resourceSystemResourceAccountingTemplateAppResourcesUpdate,
		ReadContext:   resourceSystemResourceAccountingTemplateAppResourcesRead,
		DeleteContext: resourceSystemResourceAccountingTemplateAppResourcesDelete,

		Schema: map[string]*schema.Schema{
			"cache_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of cache-template allowed (cache-template count (default is max-value))",
						},
						"cache_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"client_ssl_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ssl_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value))",
						},
						"client_ssl_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"conn_reuse_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conn_reuse_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value))",
						},
						"conn_reuse_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"fast_tcp_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fast_tcp_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value))",
						},
						"fast_tcp_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"fast_udp_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fast_udp_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value))",
						},
						"fast_udp_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"fix_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fix_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of fix-template allowed (fix-template count (default is max-value))",
						},
						"fix_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_device_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_device_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-device allowed (gslb-device count (default is max-value))",
						},
						"gslb_device_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_geo_location_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_geo_location_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value))",
						},
						"gslb_geo_location_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_ip_list_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_ip_list_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value))",
						},
						"gslb_ip_list_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_policy_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_policy_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-policy allowed (gslb-policy count (default is max-value))",
						},
						"gslb_policy_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_service_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_service_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service allowed (gslb-service count (default is max-value)",
						},
						"gslb_service_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_service_ip_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_service_ip_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value))",
						},
						"gslb_service_ip_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_service_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_service_port_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value))",
						},
						"gslb_service_port_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_site_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_site_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-site allowed (gslb-site count (default is max-value))",
						},
						"gslb_site_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_svc_group_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_svc_group_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value))",
						},
						"gslb_svc_group_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-template allowed (gslb-template count (default is max-value))",
						},
						"gslb_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"gslb_zone_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_zone_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of gslb-zone allowed (gslb-zone count (default is max-value))",
						},
						"gslb_zone_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"health_monitor_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_monitor_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of health monitors allowed (health-monitor count (default is max-value))",
						},
						"health_monitor_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"http_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of http-template allowed (http-template count (default is max-value))",
						},
						"http_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"link_cost_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_cost_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of link-cost-template allowed (link-cost-template count (default is max-value))",
						},
						"link_cost_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"pbslb_entry_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pbslb_entry_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of pbslb-entry allowed (pbslb-entry count)",
						},
						"pbslb_entry_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"persist_cookie_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"persist_cookie_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value))",
						},
						"persist_cookie_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"persist_srcip_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"persist_srcip_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value))",
						},
						"persist_srcip_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"proxy_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of proxy-template allowed (server-ssl-template count (default is max-value))",
						},
						"proxy_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"real_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"real_port_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-ports allowed (real-port count (default is max-value))",
						},
						"real_port_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"real_server_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"real_server_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of real-servers allowed (real-server count (default is max-value))",
						},
						"real_server_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"server_ssl_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ssl_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value))",
						},
						"server_ssl_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"service_group_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_group_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of service groups allowed (service-group count (default is max-value))",
						},
						"service_group_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"stream_template_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stream_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of stream-template allowed (server-ssl-template count (default is max-value))",
						},
						"stream_template_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_port_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-port allowed (virtual-port count (default is max-value))",
						},
						"virtual_port_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"virtual_server_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_server_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of virtual-servers allowed (virtual-server count (default is max-value))",
						},
						"virtual_server_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSystemResourceAccountingTemplateAppResourcesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateAppResourcesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateAppResources(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateAppResourcesRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceAccountingTemplateAppResourcesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateAppResourcesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateAppResources(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateAppResourcesRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceAccountingTemplateAppResourcesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateAppResourcesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateAppResources(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceAccountingTemplateAppResourcesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateAppResourcesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateAppResources(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemResourceAccountingTemplateAppResourcesCacheTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateMinGuarantee = in["cache_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateMinGuarantee = in["client_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateMinGuarantee = in["conn_reuse_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateMinGuarantee = in["fast_tcp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateMinGuarantee = in["fast_udp_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesFixTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesFixTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesFixTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateMinGuarantee = in["fix_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbDeviceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbDeviceMax = in["gslb_device_max"].(int)
		ret.GslbDeviceMinGuarantee = in["gslb_device_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbGeoLocationMax = in["gslb_geo_location_max"].(int)
		ret.GslbGeoLocationMinGuarantee = in["gslb_geo_location_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbIpListCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbIpListCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbIpListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbIpListMax = in["gslb_ip_list_max"].(int)
		ret.GslbIpListMinGuarantee = in["gslb_ip_list_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbPolicyCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbPolicyMax = in["gslb_policy_max"].(int)
		ret.GslbPolicyMinGuarantee = in["gslb_policy_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceMax = in["gslb_service_max"].(int)
		ret.GslbServiceMinGuarantee = in["gslb_service_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServiceIpMax = in["gslb_service_ip_max"].(int)
		ret.GslbServiceIpMinGuarantee = in["gslb_service_ip_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbServicePortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbServicePortMax = in["gslb_service_port_max"].(int)
		ret.GslbServicePortMinGuarantee = in["gslb_service_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbSiteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbSiteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbSiteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSiteMax = in["gslb_site_max"].(int)
		ret.GslbSiteMinGuarantee = in["gslb_site_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSvcGroupMax = in["gslb_svc_group_max"].(int)
		ret.GslbSvcGroupMinGuarantee = in["gslb_svc_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbTemplateMax = in["gslb_template_max"].(int)
		ret.GslbTemplateMinGuarantee = in["gslb_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesGslbZoneCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesGslbZoneCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesGslbZoneCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbZoneMax = in["gslb_zone_max"].(int)
		ret.GslbZoneMinGuarantee = in["gslb_zone_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesHealthMonitorCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthMonitorMax = in["health_monitor_max"].(int)
		ret.HealthMonitorMinGuarantee = in["health_monitor_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesHttpTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateMinGuarantee = in["http_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateMinGuarantee = in["link_cost_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPbslbEntryCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PbslbEntryMax = in["pbslb_entry_max"].(int)
		ret.PbslbEntryMinGuarantee = in["pbslb_entry_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateMinGuarantee = in["persist_cookie_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateMinGuarantee = in["persist_srcip_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesProxyTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateMinGuarantee = in["proxy_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesRealPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesRealPortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesRealPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortMinGuarantee = in["real_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesRealServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesRealServerCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesRealServerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerMinGuarantee = in["real_server_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateMinGuarantee = in["server_ssl_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesServiceGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesServiceGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesServiceGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupMinGuarantee = in["service_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesStreamTemplateCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateMinGuarantee = in["stream_template_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesVirtualPortCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesVirtualPortCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesVirtualPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortMinGuarantee = in["virtual_port_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateAppResourcesVirtualServerCfg(d []interface{}) edpt.SystemResourceAccountingTemplateAppResourcesVirtualServerCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateAppResourcesVirtualServerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerMinGuarantee = in["virtual_server_min_guarantee"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceAccountingTemplateAppResources(d *schema.ResourceData) edpt.SystemResourceAccountingTemplateAppResources {
	var ret edpt.SystemResourceAccountingTemplateAppResources
	ret.Inst.CacheTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesCacheTemplateCfg(d.Get("cache_template_cfg").([]interface{}))
	ret.Inst.ClientSslTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg(d.Get("client_ssl_template_cfg").([]interface{}))
	ret.Inst.ConnReuseTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg(d.Get("conn_reuse_template_cfg").([]interface{}))
	ret.Inst.FastTcpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg(d.Get("fast_tcp_template_cfg").([]interface{}))
	ret.Inst.FastUdpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg(d.Get("fast_udp_template_cfg").([]interface{}))
	ret.Inst.FixTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesFixTemplateCfg(d.Get("fix_template_cfg").([]interface{}))
	ret.Inst.GslbDeviceCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbDeviceCfg(d.Get("gslb_device_cfg").([]interface{}))
	ret.Inst.GslbGeoLocationCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg(d.Get("gslb_geo_location_cfg").([]interface{}))
	ret.Inst.GslbIpListCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbIpListCfg(d.Get("gslb_ip_list_cfg").([]interface{}))
	ret.Inst.GslbPolicyCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbPolicyCfg(d.Get("gslb_policy_cfg").([]interface{}))
	ret.Inst.GslbServiceCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceCfg(d.Get("gslb_service_cfg").([]interface{}))
	ret.Inst.GslbServiceIpCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg(d.Get("gslb_service_ip_cfg").([]interface{}))
	ret.Inst.GslbServicePortCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbServicePortCfg(d.Get("gslb_service_port_cfg").([]interface{}))
	ret.Inst.GslbSiteCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbSiteCfg(d.Get("gslb_site_cfg").([]interface{}))
	ret.Inst.GslbSvcGroupCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg(d.Get("gslb_svc_group_cfg").([]interface{}))
	ret.Inst.GslbTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbTemplateCfg(d.Get("gslb_template_cfg").([]interface{}))
	ret.Inst.GslbZoneCfg = getObjectSystemResourceAccountingTemplateAppResourcesGslbZoneCfg(d.Get("gslb_zone_cfg").([]interface{}))
	ret.Inst.HealthMonitorCfg = getObjectSystemResourceAccountingTemplateAppResourcesHealthMonitorCfg(d.Get("health_monitor_cfg").([]interface{}))
	ret.Inst.HttpTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesHttpTemplateCfg(d.Get("http_template_cfg").([]interface{}))
	ret.Inst.LinkCostTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg(d.Get("link_cost_template_cfg").([]interface{}))
	ret.Inst.PbslbEntryCfg = getObjectSystemResourceAccountingTemplateAppResourcesPbslbEntryCfg(d.Get("pbslb_entry_cfg").([]interface{}))
	ret.Inst.PersistCookieTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg(d.Get("persist_cookie_template_cfg").([]interface{}))
	ret.Inst.PersistSrcipTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg(d.Get("persist_srcip_template_cfg").([]interface{}))
	ret.Inst.ProxyTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesProxyTemplateCfg(d.Get("proxy_template_cfg").([]interface{}))
	ret.Inst.RealPortCfg = getObjectSystemResourceAccountingTemplateAppResourcesRealPortCfg(d.Get("real_port_cfg").([]interface{}))
	ret.Inst.RealServerCfg = getObjectSystemResourceAccountingTemplateAppResourcesRealServerCfg(d.Get("real_server_cfg").([]interface{}))
	ret.Inst.ServerSslTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg(d.Get("server_ssl_template_cfg").([]interface{}))
	ret.Inst.ServiceGroupCfg = getObjectSystemResourceAccountingTemplateAppResourcesServiceGroupCfg(d.Get("service_group_cfg").([]interface{}))
	ret.Inst.StreamTemplateCfg = getObjectSystemResourceAccountingTemplateAppResourcesStreamTemplateCfg(d.Get("stream_template_cfg").([]interface{}))
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	ret.Inst.VirtualPortCfg = getObjectSystemResourceAccountingTemplateAppResourcesVirtualPortCfg(d.Get("virtual_port_cfg").([]interface{}))
	ret.Inst.VirtualServerCfg = getObjectSystemResourceAccountingTemplateAppResourcesVirtualServerCfg(d.Get("virtual_server_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
