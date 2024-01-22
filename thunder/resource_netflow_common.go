package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowCommon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_common`: NetFlow/IPFIX common settings\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowCommonCreate,
		UpdateContext: resourceNetflowCommonUpdate,
		ReadContext:   resourceNetflowCommonRead,
		DeleteContext: resourceNetflowCommonDelete,

		Schema: map[string]*schema.Schema{
			"max_packet_queue_time": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "Configure netflow packet queue time (Max packet queue time(*20ms). Default:50( *20ms = 1s)))",
			},
			"nat44_tpl_1001": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"reset_time_on_flow_record": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset session start time to current time on each flow timeout export for long-lasting session (default: disabled).",
			},
			"selector_algorithm_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm_name": {
							Type: schema.TypeString, Required: true, Description: "'random': random;",
						},
						"sampling_population": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure sampling population for random algorithm",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceNetflowCommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCommonRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowCommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCommonRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowCommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetflowCommonSelectorAlgorithmList(d []interface{}) []edpt.NetflowCommonSelectorAlgorithmList {

	count1 := len(d)
	ret := make([]edpt.NetflowCommonSelectorAlgorithmList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowCommonSelectorAlgorithmList
		oi.AlgorithmName = in["algorithm_name"].(string)
		oi.SamplingPopulation = in["sampling_population"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowCommon(d *schema.ResourceData) edpt.NetflowCommon {
	var ret edpt.NetflowCommon
	ret.Inst.MaxPacketQueueTime = d.Get("max_packet_queue_time").(int)
	ret.Inst.Nat44_tpl_1001 = d.Get("nat44_tpl_1001").(int)
	ret.Inst.ResetTimeOnFlowRecord = d.Get("reset_time_on_flow_record").(int)
	ret.Inst.SelectorAlgorithmList = getSliceNetflowCommonSelectorAlgorithmList(d.Get("selector_algorithm_list").([]interface{}))
	//omit uuid
	return ret
}
