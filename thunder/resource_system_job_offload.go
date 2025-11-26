package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemJobOffload() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_job_offload`: job offload counter\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemJobOffloadCreate,
		UpdateContext: resourceSystemJobOffloadUpdate,
		ReadContext:   resourceSystemJobOffloadRead,
		DeleteContext: resourceSystemJobOffloadDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'jobs': Current Jobs; 'submit': Jobs Submitted; 'receive': Jobs Received; 'execute': Jobs Executed; 'snt_home': Jobs Sent Back Home; 'rcv_home': Jobs Received Home; 'complete': Jobs Completed; 'fail_submit': Jobs Failed to Submit; 'q_no_space': No More Space in Q; 'fail_execute': Failed to Execute Job; 'fail_complete': Failed to Complete Job;",
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
func resourceSystemJobOffloadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffload(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemJobOffloadRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemJobOffloadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffload(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemJobOffloadRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemJobOffloadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffload(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemJobOffloadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffload(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemJobOffloadSamplingEnable(d []interface{}) []edpt.SystemJobOffloadSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemJobOffloadSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemJobOffloadSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemJobOffload(d *schema.ResourceData) edpt.SystemJobOffload {
	var ret edpt.SystemJobOffload
	ret.Inst.SamplingEnable = getSliceSystemJobOffloadSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
