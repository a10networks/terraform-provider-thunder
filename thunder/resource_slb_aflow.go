package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbAflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_aflow`: Configure aFlow\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbAflowCreate,
		UpdateContext: resourceSlbAflowUpdate,
		ReadContext:   resourceSlbAflowRead,
		DeleteContext: resourceSlbAflowDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'pause_conn': Pause connection; 'pause_conn_fail': Pause connection fail; 'resume_conn': Resume connection; 'event_resume_conn': Resume conn by event; 'timer_resume_conn': Resume conn by timer; 'try_to_resume_conn': Resume conn by trying; 'retry_resume_conn': Resume conn by retry; 'error_resume_conn': Resume conn by error; 'open_new_server_conn': Open new server conn; 'reuse_server_idle_conn': Reuse idle server conn; 'inc_aflow_limit': Inc aFlow limit;",
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
func resourceSlbAflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbAflowRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbAflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbAflowRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbAflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbAflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbAflowSamplingEnable(d []interface{}) []edpt.SlbAflowSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbAflowSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAflowSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbAflow(d *schema.ResourceData) edpt.SlbAflow {
	var ret edpt.SlbAflow
	ret.Inst.SamplingEnable = getSliceSlbAflowSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
