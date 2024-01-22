package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgEsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_esp`: NAT64 ESP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgEspCreate,
		UpdateContext: resourceCgnv6Nat64AlgEspUpdate,
		ReadContext:   resourceCgnv6Nat64AlgEspRead,
		DeleteContext: resourceCgnv6Nat64AlgEspDelete,

		Schema: map[string]*schema.Schema{
			"esp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable NAT64 ESP ALG; 'enable-with-ctrl': Enable ESP NAT64 ALG with control session;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': ESP Sessions Created; 'nat-ip-conflict': NAT IP Conflict;",
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
func resourceCgnv6Nat64AlgEspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgEspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgEsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgEspRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgEspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgEspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgEsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgEspRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgEspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgEspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgEsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgEspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgEspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgEsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Nat64AlgEspSamplingEnable(d []interface{}) []edpt.Cgnv6Nat64AlgEspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Nat64AlgEspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Nat64AlgEspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Nat64AlgEsp(d *schema.ResourceData) edpt.Cgnv6Nat64AlgEsp {
	var ret edpt.Cgnv6Nat64AlgEsp
	ret.Inst.EspEnable = d.Get("esp_enable").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6Nat64AlgEspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
