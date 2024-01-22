package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPolling() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling`: Configure sFlow counter polling on specified source\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingCreate,
		UpdateContext: resourceSflowPollingUpdate,
		ReadContext:   resourceSflowPollingRead,
		DeleteContext: resourceSflowPollingDelete,

		Schema: map[string]*schema.Schema{
			"a10_proprietary": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_deprecated_counters": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export deprecated counters",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_usage": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Polling CPU usage",
			},
			"ddos": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable sflow polling for DDOS statistics; 'disable': Disable sflow polling for DDOS statistics;",
						},
						"compatibility3_0": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDOS sflow polling 3.0/3.1 compatibility mode",
						},
						"address_byte_order_host": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export sflow address field in host byte order",
						},
						"compatibility2_9": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDOS sflow polling 2.9 compatibility mode",
						},
						"dns_cache_zone_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable polling for dns cache per instance and per zone statistics",
						},
						"enable_anomaly_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Polling for system wide anomaly statistics",
						},
						"dyn_entry_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable polling for dynamic entry statistics",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface to sample",
						},
						"eth_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface to sample",
						},
					},
				},
			},
			"ethernet_ext_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": {
							Type: schema.TypeInt, Required: true, Description: "Ethernet interface to poll",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ethernet_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": {
							Type: schema.TypeInt, Required: true, Description: "Ethernet interface to poll",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"http": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling HTTP counters; 'disable': Disable polling HTTP counters;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"http_counter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Polling HTTP counters",
			},
			"system_health": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_health_usage": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling system health information; 'disable': Disable polling system health information;",
						},
						"per_control_cpu_usage": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling control cpu; 'disable': Disable polling control cpu usage;",
						},
						"per_data_cpu_usage": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling data cpu; 'disable': Disable polling data cpu usage;",
						},
						"license_statistics": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling license statistics; 'disable': Disable polling license statistics;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve_start": {
							Type: schema.TypeInt, Optional: true, Description: "VE interface to sample",
						},
						"ve_end": {
							Type: schema.TypeInt, Optional: true, Description: "VE interface to sample",
						},
					},
				},
			},
		},
	}
}
func resourceSflowPollingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPolling(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPolling(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPolling(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPolling(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSflowPollingA10Proprietary1402(d []interface{}) edpt.SflowPollingA10Proprietary1402 {

	count1 := len(d)
	var ret edpt.SflowPollingA10Proprietary1402
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExportDeprecatedCounters = in["export_deprecated_counters"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSflowPollingDdos1403(d []interface{}) edpt.SflowPollingDdos1403 {

	count1 := len(d)
	var ret edpt.SflowPollingDdos1403
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Toggle = in["toggle"].(string)
		ret.Compatibility3_0 = in["compatibility3_0"].(int)
		ret.AddressByteOrderHost = in["address_byte_order_host"].(int)
		ret.Compatibility2_9 = in["compatibility2_9"].(int)
		ret.DnsCacheZoneStats = in["dns_cache_zone_stats"].(int)
		ret.EnableAnomalyStats = in["enable_anomaly_stats"].(int)
		ret.DynEntryStats = in["dyn_entry_stats"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSflowPollingEthList(d []interface{}) []edpt.SflowPollingEthList {

	count1 := len(d)
	ret := make([]edpt.SflowPollingEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowPollingEthList
		oi.EthStart = in["eth_start"].(int)
		oi.EthEnd = in["eth_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSflowPollingEthernetExtList(d []interface{}) []edpt.SflowPollingEthernetExtList {

	count1 := len(d)
	ret := make([]edpt.SflowPollingEthernetExtList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowPollingEthernetExtList
		oi.Start = in["start"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSflowPollingEthernetList(d []interface{}) []edpt.SflowPollingEthernetList {

	count1 := len(d)
	ret := make([]edpt.SflowPollingEthernetList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowPollingEthernetList
		oi.Start = in["start"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSflowPollingHttp1404(d []interface{}) edpt.SflowPollingHttp1404 {

	count1 := len(d)
	var ret edpt.SflowPollingHttp1404
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Toggle = in["toggle"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSflowPollingSystemHealth1405(d []interface{}) edpt.SflowPollingSystemHealth1405 {

	count1 := len(d)
	var ret edpt.SflowPollingSystemHealth1405
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemHealthUsage = in["system_health_usage"].(string)
		ret.PerControlCpuUsage = in["per_control_cpu_usage"].(string)
		ret.PerDataCpuUsage = in["per_data_cpu_usage"].(string)
		ret.LicenseStatistics = in["license_statistics"].(string)
		//omit uuid
	}
	return ret
}

func getSliceSflowPollingVeList(d []interface{}) []edpt.SflowPollingVeList {

	count1 := len(d)
	ret := make([]edpt.SflowPollingVeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowPollingVeList
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSflowPolling(d *schema.ResourceData) edpt.SflowPolling {
	var ret edpt.SflowPolling
	ret.Inst.A10Proprietary = getObjectSflowPollingA10Proprietary1402(d.Get("a10_proprietary").([]interface{}))
	ret.Inst.CpuUsage = d.Get("cpu_usage").(int)
	ret.Inst.Ddos = getObjectSflowPollingDdos1403(d.Get("ddos").([]interface{}))
	ret.Inst.EthList = getSliceSflowPollingEthList(d.Get("eth_list").([]interface{}))
	ret.Inst.EthernetExtList = getSliceSflowPollingEthernetExtList(d.Get("ethernet_ext_list").([]interface{}))
	ret.Inst.EthernetList = getSliceSflowPollingEthernetList(d.Get("ethernet_list").([]interface{}))
	ret.Inst.Http = getObjectSflowPollingHttp1404(d.Get("http").([]interface{}))
	ret.Inst.HttpCounter = d.Get("http_counter").(int)
	ret.Inst.SystemHealth = getObjectSflowPollingSystemHealth1405(d.Get("system_health").([]interface{}))
	//omit uuid
	ret.Inst.VeList = getSliceSflowPollingVeList(d.Get("ve_list").([]interface{}))
	return ret
}
