package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemNdiscRa() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ndisc_ra`: Neighbor discovery and RA counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemNdiscRaCreate,
		UpdateContext: resourceSystemNdiscRaUpdate,
		ReadContext:   resourceSystemNdiscRaRead,
		DeleteContext: resourceSystemNdiscRaDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'good_recv': Good Router Solicitations (R.S.) Received; 'periodic_sent': Periodic Router Advertisements (R.A.) Sent; 'rate_limit': R.S. Rate Limited; 'bad_hop_limit': R.S. Bad Hop Limit; 'truncated': R.S. Truncated; 'bad_icmpv6_csum': R.S. Bad ICMPv6 Checksum; 'bad_icmpv6_code': R.S. Unknown ICMPv6 Code; 'bad_icmpv6_option': R.S. Bad ICMPv6 Option; 'l2_addr_and_unspec': R.S. Src Link-Layer Option and Unspecified Address; 'no_free_buffers': No Free Buffers to send R.A.;",
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
func resourceSystemNdiscRaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNdiscRaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNdiscRa(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemNdiscRaRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemNdiscRaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNdiscRaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNdiscRa(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemNdiscRaRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemNdiscRaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNdiscRaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNdiscRa(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemNdiscRaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNdiscRaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNdiscRa(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemNdiscRaSamplingEnable(d []interface{}) []edpt.SystemNdiscRaSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemNdiscRaSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemNdiscRaSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemNdiscRa(d *schema.ResourceData) edpt.SystemNdiscRa {
	var ret edpt.SystemNdiscRa
	ret.Inst.SamplingEnable = getSliceSystemNdiscRaSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
