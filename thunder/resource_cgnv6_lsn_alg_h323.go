package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgH323() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_h323`: Change LSN H323 ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgH323Create,
		UpdateContext: resourceCgnv6LsnAlgH323Update,
		ReadContext:   resourceCgnv6LsnAlgH323Read,
		DeleteContext: resourceCgnv6LsnAlgH323Delete,

		Schema: map[string]*schema.Schema{
			"h323_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable H323 ALG for LSN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'h225ras-message': H323 H225 RAS Message; 'h225cs-message': H323 H225 Call Signaling Message; 'h245ctl-message': H323 H245 Media Control Message; 'h245-tunneled': H323 H245 Tunnelled Message; 'fast-start': H323 FastStart; 'parse-error': Message Parse Error; 'message-cross-multi-packets': H323 Message Cross Multiple Packets; 'conn-ext-creation-failure': H323 Packet Rewrite Failure; 'rewrite-failure': H323 Message Rewrite Failure; 'tcp-out-of-order-drop': TCP Out-of-Order Drop; 'h245-dynamic-addr-exceed': H323 H245 Dynamic Address Exceed;",
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
func resourceCgnv6LsnAlgH323Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgH323Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgH323(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgH323Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgH323Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgH323Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgH323(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgH323Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgH323Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgH323Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgH323(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgH323Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgH323Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgH323(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgH323SamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgH323SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgH323SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgH323SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgH323(d *schema.ResourceData) edpt.Cgnv6LsnAlgH323 {
	var ret edpt.Cgnv6LsnAlgH323
	ret.Inst.H323Value = d.Get("h323_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgH323SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
