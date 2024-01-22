package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_alg_pptp`: PPTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatAlgPptpCreate,
		UpdateContext: resourceIpNatAlgPptpUpdate,
		ReadContext:   resourceIpNatAlgPptpRead,
		DeleteContext: resourceIpNatAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"pptp": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': Disable PPTP NAT ALG; 'enable': Enable PPTP NAT ALG;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'current-smp-sessions': current-smp-sessions; 'current-gre-sessions': current-gre-sessions; 'smp-session-creation-failure': smp-session-creation-failure; 'truncated-pns-message': truncated-pns-message; 'truncated-pac-message': truncated-pac-message; 'mismatched-pns-call-id': mismatched-pns-call-id; 'mismatched-pac-call-id': mismatched-pac-call-id; 'retransmitted-pns-message': retransmitted-pns-message; 'retransmitted-pac-message': retransmitted-pac-message; 'truncated-gre-packet': truncated-gre-packet; 'unknown-gre-version': unknown-gre-version; 'no-matching-gre-session': no-matching-gre-session;",
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
func resourceIpNatAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpNatAlgPptpSamplingEnable(d []interface{}) []edpt.IpNatAlgPptpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpNatAlgPptpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatAlgPptpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpNatAlgPptp(d *schema.ResourceData) edpt.IpNatAlgPptp {
	var ret edpt.IpNatAlgPptp
	ret.Inst.Pptp = d.Get("pptp").(string)
	ret.Inst.SamplingEnable = getSliceIpNatAlgPptpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
