package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgH323() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_alg_h323`: Change Fixed NAT H323 ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatAlgH323Create,
		UpdateContext: resourceCgnv6FixedNatAlgH323Update,
		ReadContext:   resourceCgnv6FixedNatAlgH323Read,
		DeleteContext: resourceCgnv6FixedNatAlgH323Delete,

		Schema: map[string]*schema.Schema{
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
func resourceCgnv6FixedNatAlgH323Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgH323Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgH323(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgH323Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatAlgH323Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgH323Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgH323(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgH323Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatAlgH323Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgH323Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgH323(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatAlgH323Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgH323Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgH323(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6FixedNatAlgH323SamplingEnable(d []interface{}) []edpt.Cgnv6FixedNatAlgH323SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatAlgH323SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatAlgH323SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgH323(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgH323 {
	var ret edpt.Cgnv6FixedNatAlgH323
	ret.Inst.SamplingEnable = getSliceCgnv6FixedNatAlgH323SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
