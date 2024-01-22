package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmp6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_icmp6`: Display ICMPv6 statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIcmp6Create,
		UpdateContext: resourceSystemIcmp6Update,
		ReadContext:   resourceSystemIcmp6Read,
		DeleteContext: resourceSystemIcmp6Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'in_msgs': In messages; 'in_errors': In Errors; 'in_dest_un_reach': In Destunation Unreachable; 'in_pkt_too_big': In Packet too big; 'in_time_exceeds': In TTL Exceeds; 'in_param_prob': In Parameter Problem; 'in_echoes': In Echo requests; 'in_exho_reply': In Echo replies; 'in_grp_mem_query': In Group member query; 'in_grp_mem_resp': In Group member reply; 'in_grp_mem_reduction': In Group member reduction; 'in_router_sol': In Router solicitation; 'in_ra': In Router advertisement; 'in_ns': In neighbor solicitation; 'in_na': In neighbor advertisement; 'in_redirect': In Redirects; 'out_msg': Out Messages; 'out_dst_un_reach': Out Destination Unreachable; 'out_pkt_too_big': Out Packet too big; 'out_time_exceeds': Out TTL Exceeds; 'out_param_prob': Out Parameter Problem; 'out_echo_req': Out Echo requests; 'out_echo_replies': Out Echo replies; 'out_rs': Out Router solicitation; 'out_ra': Out Router advertisement; 'out_ns': Out neighbor solicitation; 'out_na': Out neighbor advertisement; 'out_redirects': Out Redirects; 'out_mem_resp': Out Group member reply; 'out_mem_reductions': Out Group member reduction; 'err_rs': Error Router solicitation; 'err_ra': Error Router advertisement; 'err_ns': Error Neighbor solicitation; 'err_na': Error Neighbor advertisement; 'err_redirects': Error Redirects; 'err_echoes': Error Echo requests; 'err_echo_replies': Error Echo replies;",
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
func resourceSystemIcmp6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmp6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmp6Read(ctx, d, meta)
	}
	return diags
}

func resourceSystemIcmp6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmp6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmp6Read(ctx, d, meta)
	}
	return diags
}
func resourceSystemIcmp6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmp6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIcmp6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmp6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIcmp6SamplingEnable(d []interface{}) []edpt.SystemIcmp6SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIcmp6SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmp6SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIcmp6(d *schema.ResourceData) edpt.SystemIcmp6 {
	var ret edpt.SystemIcmp6
	ret.Inst.SamplingEnable = getSliceSystemIcmp6SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
