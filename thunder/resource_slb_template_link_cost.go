package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkCost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_link_cost`: Server Link-Cost Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLinkCostCreate,
		UpdateContext: resourceSlbTemplateLinkCostUpdate,
		ReadContext:   resourceSlbTemplateLinkCostRead,
		DeleteContext: resourceSlbTemplateLinkCostDelete,

		Schema: map[string]*schema.Schema{
			"bandwidth_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Interval duration in seconds",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Link-Cost Template Name",
			},
			"overage_bandwidth_cost": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure overage bandwidth cost on real server (Overage bandwidth cost per Mbpi)",
			},
			"prepaid_bandwidth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure prepaid bandwidth on real server (Prepaid Bandwidth in Mbpi)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'link_total_fwd_bytes': Total bytes fwd for link; 'interval_fwd_bytes': Link Cost bytes processed in forward direction per interval; 'link_total_conn': Link Cost total connection; 'interval_avg_throughput': Link Cost average throughput per interval; 'interval_max_throughput': Link Cost max throughput per interval;",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateLinkCostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkCostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkCost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkCostRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLinkCostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkCostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkCost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkCostRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLinkCostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkCostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkCost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLinkCostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkCostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkCost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateLinkCostSamplingEnable(d []interface{}) []edpt.SlbTemplateLinkCostSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateLinkCostSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateLinkCostSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateLinkCost(d *schema.ResourceData) edpt.SlbTemplateLinkCost {
	var ret edpt.SlbTemplateLinkCost
	ret.Inst.BandwidthInterval = d.Get("bandwidth_interval").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OverageBandwidthCost = d.Get("overage_bandwidth_cost").(int)
	ret.Inst.PrepaidBandwidth = d.Get("prepaid_bandwidth").(int)
	ret.Inst.SamplingEnable = getSliceSlbTemplateLinkCostSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
