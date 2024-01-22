package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccountingTemplateSystemResources() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_accounting_template_system_resources`: Enter the system resource limits\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceAccountingTemplateSystemResourcesCreate,
		UpdateContext: resourceSystemResourceAccountingTemplateSystemResourcesUpdate,
		ReadContext:   resourceSystemResourceAccountingTemplateSystemResourcesRead,
		DeleteContext: resourceSystemResourceAccountingTemplateSystemResourcesDelete,

		Schema: map[string]*schema.Schema{
			"bw_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bw_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default))",
						},
						"bw_limit_watermark_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
						},
					},
				},
			},
			"concurrent_session_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"concurrent_session_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default))",
						},
					},
				},
			},
			"fwcps_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fwcps_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the Firewall cps limit (Firewall cps limit (no limits applied by default))",
						},
					},
				},
			},
			"l4_session_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_session_limit_max": {
							Type: schema.TypeString, Optional: true, Description: "Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision))",
						},
						"l4_session_limit_min_guarantee": {
							Type: schema.TypeString, Optional: true, Default: "0", Description: "minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99)",
						},
					},
				},
			},
			"l4cps_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4cps_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the L4 cps limit (L4 cps limit (no limits applied by default))",
						},
					},
				},
			},
			"l7cps_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l7cps_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "L7cps-limit (L7 cps limit (no limits applied by default))",
						},
					},
				},
			},
			"natcps_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"natcps_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the Nat cps limit (NAT cps limit (no limits applied by default))",
						},
					},
				},
			},
			"ssl_throughput_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_throughput_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default))",
						},
						"ssl_throughput_limit_watermark_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable watermark (90% drop, keep existing sessions, drop  new sessions)",
						},
					},
				},
			},
			"sslcps_limit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sslcps_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the SSL cps limit (SSL cps limit (no limits applied by default))",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSystemResourceAccountingTemplateSystemResourcesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateSystemResourcesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateSystemResources(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateSystemResourcesRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceAccountingTemplateSystemResourcesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateSystemResourcesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateSystemResources(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateSystemResourcesRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceAccountingTemplateSystemResourcesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateSystemResourcesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateSystemResources(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceAccountingTemplateSystemResourcesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateSystemResourcesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateSystemResources(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemResourceAccountingTemplateSystemResourcesBwLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesBwLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesBwLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BwLimitMax = in["bw_limit_max"].(int)
		ret.BwLimitWatermarkDisable = in["bw_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentSessionLimitMax = in["concurrent_session_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwcpsLimitMax = in["fwcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionLimitMax = in["l4_session_limit_max"].(string)
		ret.L4SessionLimitMinGuarantee = in["l4_session_limit_min_guarantee"].(string)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4cpsLimitMax = in["l4cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L7cpsLimitMax = in["l7cps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatcpsLimitMax = in["natcps_limit_max"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslThroughputLimitMax = in["ssl_throughput_limit_max"].(int)
		ret.SslThroughputLimitWatermarkDisable = in["ssl_throughput_limit_watermark_disable"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg(d []interface{}) edpt.SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslcpsLimitMax = in["sslcps_limit_max"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceAccountingTemplateSystemResources(d *schema.ResourceData) edpt.SystemResourceAccountingTemplateSystemResources {
	var ret edpt.SystemResourceAccountingTemplateSystemResources
	ret.Inst.BwLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesBwLimitCfg(d.Get("bw_limit_cfg").([]interface{}))
	ret.Inst.ConcurrentSessionLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg(d.Get("concurrent_session_limit_cfg").([]interface{}))
	ret.Inst.FwcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg(d.Get("fwcps_limit_cfg").([]interface{}))
	ret.Inst.L4SessionLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg(d.Get("l4_session_limit_cfg").([]interface{}))
	ret.Inst.L4cpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg(d.Get("l4cps_limit_cfg").([]interface{}))
	ret.Inst.L7cpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg(d.Get("l7cps_limit_cfg").([]interface{}))
	ret.Inst.NatcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg(d.Get("natcps_limit_cfg").([]interface{}))
	ret.Inst.SslThroughputLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg(d.Get("ssl_throughput_limit_cfg").([]interface{}))
	ret.Inst.SslcpsLimitCfg = getObjectSystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg(d.Get("sslcps_limit_cfg").([]interface{}))
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
