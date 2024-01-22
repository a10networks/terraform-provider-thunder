package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ips_global`: IPS profile\n\n__PLACEHOLDER__",
		CreateContext: resourceIpsGlobalCreate,
		UpdateContext: resourceIpsGlobalUpdate,
		ReadContext:   resourceIpsGlobalRead,
		DeleteContext: resourceIpsGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ips_matched_total': IPS Matched Total; 'ips_matched_http': IPS Matched HTTP; 'ips_matched_dns': IPS Matched DNS; 'ips_matched_other': IPS Matched Other; 'ips_matched_action_pass': IPS Matched Action Pass; 'ips_matched_action_drop': IPS Matched Action Drop; 'ips_matched_action_blacklist': IPS Matched Action Blacklist; 'ips_matched_severity_high': IPS Matched Severity High; 'ips_matched_severity_medium': IPS Matched Severity Medium; 'ips_matched_severity_low': IPS Matched Severity Low;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpsGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpsGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceIpsGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpsGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceIpsGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpsGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpsGlobalSamplingEnable(d []interface{}) []edpt.IpsGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpsGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpsGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpsGlobal(d *schema.ResourceData) edpt.IpsGlobal {
	var ret edpt.IpsGlobal
	ret.Inst.SamplingEnable = getSliceIpsGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
