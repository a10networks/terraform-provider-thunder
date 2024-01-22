package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_pool_addr_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_server_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_server_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_server_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_port_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_port_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"real_port_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_group_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_group_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_group_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_port_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_port_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_port_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_server_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_server_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_server_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fix_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fix_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fix_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"proxy_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"proxy_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"proxy_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"conn_reuse_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"conn_reuse_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"conn_reuse_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_tcp_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_tcp_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_tcp_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_udp_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_udp_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fast_udp_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"client_ssl_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"client_ssl_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"client_ssl_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_ssl_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_ssl_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_ssl_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"link_cost_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"link_cost_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"link_cost_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_entry_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_entry_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_entry_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stream_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stream_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stream_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_cookie_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_cookie_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_cookie_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_srcip_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_srcip_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_srcip_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"health_monitor_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"health_monitor_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"health_monitor_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_subnet_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_subnet_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pbslb_subnet_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_site_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_site_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_site_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_device_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_device_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_device_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_ip_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_ip_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_ip_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_port_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_port_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_port_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_zone_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_zone_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_zone_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_service_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_policy_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_policy_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_policy_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_geo_location_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_geo_location_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_geo_location_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_ip_list_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_ip_list_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_ip_list_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_template_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_template_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_template_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_svcgroup_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_svcgroup_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gslb_svcgroup_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_threshold_res_usage_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_threshold_res_usage_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_threshold_res_usage_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbResourceUsageOperOper := setObjectSlbResourceUsageOperOper(res)
		d.Set("oper", SlbResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbResourceUsageOperOper(ret edpt.DataSlbResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nat_pool_addr_min":               ret.DtSlbResourceUsageOper.Oper.NatPoolAddrMin,
			"nat_pool_addr_max":               ret.DtSlbResourceUsageOper.Oper.NatPoolAddrMax,
			"nat_pool_addr_default":           ret.DtSlbResourceUsageOper.Oper.NatPoolAddrDefault,
			"real_server_min":                 ret.DtSlbResourceUsageOper.Oper.RealServerMin,
			"real_server_max":                 ret.DtSlbResourceUsageOper.Oper.RealServerMax,
			"real_server_default":             ret.DtSlbResourceUsageOper.Oper.RealServerDefault,
			"real_port_min":                   ret.DtSlbResourceUsageOper.Oper.RealPortMin,
			"real_port_max":                   ret.DtSlbResourceUsageOper.Oper.RealPortMax,
			"real_port_default":               ret.DtSlbResourceUsageOper.Oper.RealPortDefault,
			"service_group_min":               ret.DtSlbResourceUsageOper.Oper.ServiceGroupMin,
			"service_group_max":               ret.DtSlbResourceUsageOper.Oper.ServiceGroupMax,
			"service_group_default":           ret.DtSlbResourceUsageOper.Oper.ServiceGroupDefault,
			"virtual_port_min":                ret.DtSlbResourceUsageOper.Oper.VirtualPortMin,
			"virtual_port_max":                ret.DtSlbResourceUsageOper.Oper.VirtualPortMax,
			"virtual_port_default":            ret.DtSlbResourceUsageOper.Oper.VirtualPortDefault,
			"virtual_server_min":              ret.DtSlbResourceUsageOper.Oper.VirtualServerMin,
			"virtual_server_max":              ret.DtSlbResourceUsageOper.Oper.VirtualServerMax,
			"virtual_server_default":          ret.DtSlbResourceUsageOper.Oper.VirtualServerDefault,
			"http_template_min":               ret.DtSlbResourceUsageOper.Oper.HttpTemplateMin,
			"http_template_max":               ret.DtSlbResourceUsageOper.Oper.HttpTemplateMax,
			"http_template_default":           ret.DtSlbResourceUsageOper.Oper.HttpTemplateDefault,
			"fix_template_min":                ret.DtSlbResourceUsageOper.Oper.FixTemplateMin,
			"fix_template_max":                ret.DtSlbResourceUsageOper.Oper.FixTemplateMax,
			"fix_template_default":            ret.DtSlbResourceUsageOper.Oper.FixTemplateDefault,
			"proxy_template_min":              ret.DtSlbResourceUsageOper.Oper.ProxyTemplateMin,
			"proxy_template_max":              ret.DtSlbResourceUsageOper.Oper.ProxyTemplateMax,
			"proxy_template_default":          ret.DtSlbResourceUsageOper.Oper.ProxyTemplateDefault,
			"conn_reuse_template_min":         ret.DtSlbResourceUsageOper.Oper.ConnReuseTemplateMin,
			"conn_reuse_template_max":         ret.DtSlbResourceUsageOper.Oper.ConnReuseTemplateMax,
			"conn_reuse_template_default":     ret.DtSlbResourceUsageOper.Oper.ConnReuseTemplateDefault,
			"fast_tcp_template_min":           ret.DtSlbResourceUsageOper.Oper.FastTcpTemplateMin,
			"fast_tcp_template_max":           ret.DtSlbResourceUsageOper.Oper.FastTcpTemplateMax,
			"fast_tcp_template_default":       ret.DtSlbResourceUsageOper.Oper.FastTcpTemplateDefault,
			"fast_udp_template_min":           ret.DtSlbResourceUsageOper.Oper.FastUdpTemplateMin,
			"fast_udp_template_max":           ret.DtSlbResourceUsageOper.Oper.FastUdpTemplateMax,
			"fast_udp_template_default":       ret.DtSlbResourceUsageOper.Oper.FastUdpTemplateDefault,
			"client_ssl_template_min":         ret.DtSlbResourceUsageOper.Oper.ClientSslTemplateMin,
			"client_ssl_template_max":         ret.DtSlbResourceUsageOper.Oper.ClientSslTemplateMax,
			"client_ssl_template_default":     ret.DtSlbResourceUsageOper.Oper.ClientSslTemplateDefault,
			"server_ssl_template_min":         ret.DtSlbResourceUsageOper.Oper.ServerSslTemplateMin,
			"server_ssl_template_max":         ret.DtSlbResourceUsageOper.Oper.ServerSslTemplateMax,
			"server_ssl_template_default":     ret.DtSlbResourceUsageOper.Oper.ServerSslTemplateDefault,
			"link_cost_template_min":          ret.DtSlbResourceUsageOper.Oper.LinkCostTemplateMin,
			"link_cost_template_max":          ret.DtSlbResourceUsageOper.Oper.LinkCostTemplateMax,
			"link_cost_template_default":      ret.DtSlbResourceUsageOper.Oper.LinkCostTemplateDefault,
			"pbslb_entry_min":                 ret.DtSlbResourceUsageOper.Oper.PbslbEntryMin,
			"pbslb_entry_max":                 ret.DtSlbResourceUsageOper.Oper.PbslbEntryMax,
			"pbslb_entry_default":             ret.DtSlbResourceUsageOper.Oper.PbslbEntryDefault,
			"stream_template_min":             ret.DtSlbResourceUsageOper.Oper.StreamTemplateMin,
			"stream_template_max":             ret.DtSlbResourceUsageOper.Oper.StreamTemplateMax,
			"stream_template_default":         ret.DtSlbResourceUsageOper.Oper.StreamTemplateDefault,
			"persist_cookie_template_min":     ret.DtSlbResourceUsageOper.Oper.PersistCookieTemplateMin,
			"persist_cookie_template_max":     ret.DtSlbResourceUsageOper.Oper.PersistCookieTemplateMax,
			"persist_cookie_template_default": ret.DtSlbResourceUsageOper.Oper.PersistCookieTemplateDefault,
			"persist_srcip_template_min":      ret.DtSlbResourceUsageOper.Oper.PersistSrcipTemplateMin,
			"persist_srcip_template_max":      ret.DtSlbResourceUsageOper.Oper.PersistSrcipTemplateMax,
			"persist_srcip_template_default":  ret.DtSlbResourceUsageOper.Oper.PersistSrcipTemplateDefault,
			"health_monitor_count_min":        ret.DtSlbResourceUsageOper.Oper.HealthMonitorCountMin,
			"health_monitor_count_max":        ret.DtSlbResourceUsageOper.Oper.HealthMonitorCountMax,
			"health_monitor_count_default":    ret.DtSlbResourceUsageOper.Oper.HealthMonitorCountDefault,
			"pbslb_subnet_count_min":          ret.DtSlbResourceUsageOper.Oper.PbslbSubnetCountMin,
			"pbslb_subnet_count_max":          ret.DtSlbResourceUsageOper.Oper.PbslbSubnetCountMax,
			"pbslb_subnet_count_default":      ret.DtSlbResourceUsageOper.Oper.PbslbSubnetCountDefault,
			"gslb_site_count_min":             ret.DtSlbResourceUsageOper.Oper.GslbSiteCountMin,
			"gslb_site_count_max":             ret.DtSlbResourceUsageOper.Oper.GslbSiteCountMax,
			"gslb_site_count_default":         ret.DtSlbResourceUsageOper.Oper.GslbSiteCountDefault,
			"gslb_device_count_min":           ret.DtSlbResourceUsageOper.Oper.GslbDeviceCountMin,
			"gslb_device_count_max":           ret.DtSlbResourceUsageOper.Oper.GslbDeviceCountMax,
			"gslb_device_count_default":       ret.DtSlbResourceUsageOper.Oper.GslbDeviceCountDefault,
			"gslb_service_ip_count_min":       ret.DtSlbResourceUsageOper.Oper.GslbServiceIpCountMin,
			"gslb_service_ip_count_max":       ret.DtSlbResourceUsageOper.Oper.GslbServiceIpCountMax,
			"gslb_service_ip_count_default":   ret.DtSlbResourceUsageOper.Oper.GslbServiceIpCountDefault,
			"gslb_service_port_count_min":     ret.DtSlbResourceUsageOper.Oper.GslbServicePortCountMin,
			"gslb_service_port_count_max":     ret.DtSlbResourceUsageOper.Oper.GslbServicePortCountMax,
			"gslb_service_port_count_default": ret.DtSlbResourceUsageOper.Oper.GslbServicePortCountDefault,
			"gslb_zone_count_min":             ret.DtSlbResourceUsageOper.Oper.GslbZoneCountMin,
			"gslb_zone_count_max":             ret.DtSlbResourceUsageOper.Oper.GslbZoneCountMax,
			"gslb_zone_count_default":         ret.DtSlbResourceUsageOper.Oper.GslbZoneCountDefault,
			"gslb_service_count_min":          ret.DtSlbResourceUsageOper.Oper.GslbServiceCountMin,
			"gslb_service_count_max":          ret.DtSlbResourceUsageOper.Oper.GslbServiceCountMax,
			"gslb_service_count_default":      ret.DtSlbResourceUsageOper.Oper.GslbServiceCountDefault,
			"gslb_policy_count_min":           ret.DtSlbResourceUsageOper.Oper.GslbPolicyCountMin,
			"gslb_policy_count_max":           ret.DtSlbResourceUsageOper.Oper.GslbPolicyCountMax,
			"gslb_policy_count_default":       ret.DtSlbResourceUsageOper.Oper.GslbPolicyCountDefault,
			"gslb_geo_location_count_min":     ret.DtSlbResourceUsageOper.Oper.GslbGeoLocationCountMin,
			"gslb_geo_location_count_max":     ret.DtSlbResourceUsageOper.Oper.GslbGeoLocationCountMax,
			"gslb_geo_location_count_default": ret.DtSlbResourceUsageOper.Oper.GslbGeoLocationCountDefault,
			"gslb_ip_list_count_min":          ret.DtSlbResourceUsageOper.Oper.GslbIpListCountMin,
			"gslb_ip_list_count_max":          ret.DtSlbResourceUsageOper.Oper.GslbIpListCountMax,
			"gslb_ip_list_count_default":      ret.DtSlbResourceUsageOper.Oper.GslbIpListCountDefault,
			"gslb_template_count_min":         ret.DtSlbResourceUsageOper.Oper.GslbTemplateCountMin,
			"gslb_template_count_max":         ret.DtSlbResourceUsageOper.Oper.GslbTemplateCountMax,
			"gslb_template_count_default":     ret.DtSlbResourceUsageOper.Oper.GslbTemplateCountDefault,
			"gslb_svcgroup_count_min":         ret.DtSlbResourceUsageOper.Oper.GslbSvcgroupCountMin,
			"gslb_svcgroup_count_max":         ret.DtSlbResourceUsageOper.Oper.GslbSvcgroupCountMax,
			"gslb_svcgroup_count_default":     ret.DtSlbResourceUsageOper.Oper.GslbSvcgroupCountDefault,
			"cache_template_min":              ret.DtSlbResourceUsageOper.Oper.CacheTemplateMin,
			"cache_template_max":              ret.DtSlbResourceUsageOper.Oper.CacheTemplateMax,
			"cache_template_default":          ret.DtSlbResourceUsageOper.Oper.CacheTemplateDefault,
			"slb_threshold_res_usage_default": ret.DtSlbResourceUsageOper.Oper.SlbThresholdResUsageDefault,
			"slb_threshold_res_usage_min":     ret.DtSlbResourceUsageOper.Oper.SlbThresholdResUsageMin,
			"slb_threshold_res_usage_max":     ret.DtSlbResourceUsageOper.Oper.SlbThresholdResUsageMax,
		},
	}
}

func getObjectSlbResourceUsageOperOper(d []interface{}) edpt.SlbResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.SlbResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatPoolAddrMin = in["nat_pool_addr_min"].(int)
		ret.NatPoolAddrMax = in["nat_pool_addr_max"].(int)
		ret.NatPoolAddrDefault = in["nat_pool_addr_default"].(int)
		ret.RealServerMin = in["real_server_min"].(int)
		ret.RealServerMax = in["real_server_max"].(int)
		ret.RealServerDefault = in["real_server_default"].(int)
		ret.RealPortMin = in["real_port_min"].(int)
		ret.RealPortMax = in["real_port_max"].(int)
		ret.RealPortDefault = in["real_port_default"].(int)
		ret.ServiceGroupMin = in["service_group_min"].(int)
		ret.ServiceGroupMax = in["service_group_max"].(int)
		ret.ServiceGroupDefault = in["service_group_default"].(int)
		ret.VirtualPortMin = in["virtual_port_min"].(int)
		ret.VirtualPortMax = in["virtual_port_max"].(int)
		ret.VirtualPortDefault = in["virtual_port_default"].(int)
		ret.VirtualServerMin = in["virtual_server_min"].(int)
		ret.VirtualServerMax = in["virtual_server_max"].(int)
		ret.VirtualServerDefault = in["virtual_server_default"].(int)
		ret.HttpTemplateMin = in["http_template_min"].(int)
		ret.HttpTemplateMax = in["http_template_max"].(int)
		ret.HttpTemplateDefault = in["http_template_default"].(int)
		ret.FixTemplateMin = in["fix_template_min"].(int)
		ret.FixTemplateMax = in["fix_template_max"].(int)
		ret.FixTemplateDefault = in["fix_template_default"].(int)
		ret.ProxyTemplateMin = in["proxy_template_min"].(int)
		ret.ProxyTemplateMax = in["proxy_template_max"].(int)
		ret.ProxyTemplateDefault = in["proxy_template_default"].(int)
		ret.ConnReuseTemplateMin = in["conn_reuse_template_min"].(int)
		ret.ConnReuseTemplateMax = in["conn_reuse_template_max"].(int)
		ret.ConnReuseTemplateDefault = in["conn_reuse_template_default"].(int)
		ret.FastTcpTemplateMin = in["fast_tcp_template_min"].(int)
		ret.FastTcpTemplateMax = in["fast_tcp_template_max"].(int)
		ret.FastTcpTemplateDefault = in["fast_tcp_template_default"].(int)
		ret.FastUdpTemplateMin = in["fast_udp_template_min"].(int)
		ret.FastUdpTemplateMax = in["fast_udp_template_max"].(int)
		ret.FastUdpTemplateDefault = in["fast_udp_template_default"].(int)
		ret.ClientSslTemplateMin = in["client_ssl_template_min"].(int)
		ret.ClientSslTemplateMax = in["client_ssl_template_max"].(int)
		ret.ClientSslTemplateDefault = in["client_ssl_template_default"].(int)
		ret.ServerSslTemplateMin = in["server_ssl_template_min"].(int)
		ret.ServerSslTemplateMax = in["server_ssl_template_max"].(int)
		ret.ServerSslTemplateDefault = in["server_ssl_template_default"].(int)
		ret.LinkCostTemplateMin = in["link_cost_template_min"].(int)
		ret.LinkCostTemplateMax = in["link_cost_template_max"].(int)
		ret.LinkCostTemplateDefault = in["link_cost_template_default"].(int)
		ret.PbslbEntryMin = in["pbslb_entry_min"].(int)
		ret.PbslbEntryMax = in["pbslb_entry_max"].(int)
		ret.PbslbEntryDefault = in["pbslb_entry_default"].(int)
		ret.StreamTemplateMin = in["stream_template_min"].(int)
		ret.StreamTemplateMax = in["stream_template_max"].(int)
		ret.StreamTemplateDefault = in["stream_template_default"].(int)
		ret.PersistCookieTemplateMin = in["persist_cookie_template_min"].(int)
		ret.PersistCookieTemplateMax = in["persist_cookie_template_max"].(int)
		ret.PersistCookieTemplateDefault = in["persist_cookie_template_default"].(int)
		ret.PersistSrcipTemplateMin = in["persist_srcip_template_min"].(int)
		ret.PersistSrcipTemplateMax = in["persist_srcip_template_max"].(int)
		ret.PersistSrcipTemplateDefault = in["persist_srcip_template_default"].(int)
		ret.HealthMonitorCountMin = in["health_monitor_count_min"].(int)
		ret.HealthMonitorCountMax = in["health_monitor_count_max"].(int)
		ret.HealthMonitorCountDefault = in["health_monitor_count_default"].(int)
		ret.PbslbSubnetCountMin = in["pbslb_subnet_count_min"].(int)
		ret.PbslbSubnetCountMax = in["pbslb_subnet_count_max"].(int)
		ret.PbslbSubnetCountDefault = in["pbslb_subnet_count_default"].(int)
		ret.GslbSiteCountMin = in["gslb_site_count_min"].(int)
		ret.GslbSiteCountMax = in["gslb_site_count_max"].(int)
		ret.GslbSiteCountDefault = in["gslb_site_count_default"].(int)
		ret.GslbDeviceCountMin = in["gslb_device_count_min"].(int)
		ret.GslbDeviceCountMax = in["gslb_device_count_max"].(int)
		ret.GslbDeviceCountDefault = in["gslb_device_count_default"].(int)
		ret.GslbServiceIpCountMin = in["gslb_service_ip_count_min"].(int)
		ret.GslbServiceIpCountMax = in["gslb_service_ip_count_max"].(int)
		ret.GslbServiceIpCountDefault = in["gslb_service_ip_count_default"].(int)
		ret.GslbServicePortCountMin = in["gslb_service_port_count_min"].(int)
		ret.GslbServicePortCountMax = in["gslb_service_port_count_max"].(int)
		ret.GslbServicePortCountDefault = in["gslb_service_port_count_default"].(int)
		ret.GslbZoneCountMin = in["gslb_zone_count_min"].(int)
		ret.GslbZoneCountMax = in["gslb_zone_count_max"].(int)
		ret.GslbZoneCountDefault = in["gslb_zone_count_default"].(int)
		ret.GslbServiceCountMin = in["gslb_service_count_min"].(int)
		ret.GslbServiceCountMax = in["gslb_service_count_max"].(int)
		ret.GslbServiceCountDefault = in["gslb_service_count_default"].(int)
		ret.GslbPolicyCountMin = in["gslb_policy_count_min"].(int)
		ret.GslbPolicyCountMax = in["gslb_policy_count_max"].(int)
		ret.GslbPolicyCountDefault = in["gslb_policy_count_default"].(int)
		ret.GslbGeoLocationCountMin = in["gslb_geo_location_count_min"].(int)
		ret.GslbGeoLocationCountMax = in["gslb_geo_location_count_max"].(int)
		ret.GslbGeoLocationCountDefault = in["gslb_geo_location_count_default"].(int)
		ret.GslbIpListCountMin = in["gslb_ip_list_count_min"].(int)
		ret.GslbIpListCountMax = in["gslb_ip_list_count_max"].(int)
		ret.GslbIpListCountDefault = in["gslb_ip_list_count_default"].(int)
		ret.GslbTemplateCountMin = in["gslb_template_count_min"].(int)
		ret.GslbTemplateCountMax = in["gslb_template_count_max"].(int)
		ret.GslbTemplateCountDefault = in["gslb_template_count_default"].(int)
		ret.GslbSvcgroupCountMin = in["gslb_svcgroup_count_min"].(int)
		ret.GslbSvcgroupCountMax = in["gslb_svcgroup_count_max"].(int)
		ret.GslbSvcgroupCountDefault = in["gslb_svcgroup_count_default"].(int)
		ret.CacheTemplateMin = in["cache_template_min"].(int)
		ret.CacheTemplateMax = in["cache_template_max"].(int)
		ret.CacheTemplateDefault = in["cache_template_default"].(int)
		ret.SlbThresholdResUsageDefault = in["slb_threshold_res_usage_default"].(int)
		ret.SlbThresholdResUsageMin = in["slb_threshold_res_usage_min"].(int)
		ret.SlbThresholdResUsageMax = in["slb_threshold_res_usage_max"].(int)
	}
	return ret
}

func dataToEndpointSlbResourceUsageOper(d *schema.ResourceData) edpt.SlbResourceUsageOper {
	var ret edpt.SlbResourceUsageOper

	ret.Oper = getObjectSlbResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
