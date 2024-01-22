package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_icmp`: Display ICMP statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIcmpCreate,
		UpdateContext: resourceSystemIcmpUpdate,
		ReadContext:   resourceSystemIcmpRead,
		DeleteContext: resourceSystemIcmpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Total number; 'inmsgs': In Messages; 'inerrors': In Errors; 'indestunreachs': In Destination Unreachable; 'intimeexcds': In TTL Exceeds; 'inparmprobs': In Parameter Problem; 'insrcquenchs': In Source Quench Error; 'inredirects': In Redirects; 'inechos': In Echo requests; 'inechoreps': In Echo replies; 'intimestamps': In Timestamp; 'intimestampreps': In Timestamp Rep; 'inaddrmasks': In Address Masks; 'inaddrmaskreps': In Address Mask Rep; 'outmsgs': Out Message; 'outerrors': Out Errors; 'outdestunreachs': Out Destination Unreachable; 'outtimeexcds': Out TTL Exceeds; 'outparmprobs': Out Parameter Problem; 'outsrcquenchs': Out Source Quench Error; 'outredirects': Out Redirects; 'outechos': Out Echo Requests; 'outechoreps': Out Echo Replies; 'outtimestamps': Out Time Stamp; 'outtimestampreps': Out Time Stamp Rep; 'outaddrmasks': Out Address Mask; 'outaddrmaskreps': Out Address Mask Rep;",
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
func resourceSystemIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIcmpSamplingEnable(d []interface{}) []edpt.SystemIcmpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIcmpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIcmp(d *schema.ResourceData) edpt.SystemIcmp {
	var ret edpt.SystemIcmp
	ret.Inst.SamplingEnable = getSliceSystemIcmpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
