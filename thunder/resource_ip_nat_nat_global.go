package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatNatGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_nat_global`: Debug statistics for IP NAT\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatNatGlobalCreate,
		UpdateContext: resourceIpNatNatGlobalUpdate,
		ReadContext:   resourceIpNatNatGlobalRead,
		DeleteContext: resourceIpNatNatGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'cross_cpu_helper_created': Cross CPU Helper Created; 'cross_cpu_helper_free': Cross CPU Helper Free; 'cross_cpu_sent': Cross CPU Helper Packets Sent; 'cross_cpu_rcv': Cross CPU Helper Packets Received; 'cross_cpu_helper_nat_pool_standby': Cross CPU Helper Standby; 'cross_cpu_helper_cpu_mismatch': Cross CPU Helper CPU Mismatch; 'cross_cpu_bad_l3': Cross CPU Unsupported L3; 'cross_cpu_bad_l4': Cross CPU Unsupported L4; 'cross_cpu_no_session': Cross CPU No Session Found; 'cross_cpu_helper_deleted': Cross CPU Helper Deleted; 'cross_cpu_helper_free_retry_lookup': Cross CPU Helper Free Retry Lookup; 'cross_cpu_helper_free_not_found': Cross CPU Helper Free Not Found;",
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
func resourceIpNatNatGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatNatGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatNatGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatNatGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatNatGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatNatGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatNatGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatNatGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatNatGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatNatGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatNatGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatNatGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatNatGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatNatGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpNatNatGlobalSamplingEnable(d []interface{}) []edpt.IpNatNatGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpNatNatGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatNatGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpNatNatGlobal(d *schema.ResourceData) edpt.IpNatNatGlobal {
	var ret edpt.IpNatNatGlobal
	ret.Inst.SamplingEnable = getSliceIpNatNatGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
