package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorSflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_flow_collector_sflow`: sFlow collector\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityFlowCollectorSflowCreate,
		UpdateContext: resourceVisibilityFlowCollectorSflowUpdate,
		ReadContext:   resourceVisibilityFlowCollectorSflowRead,
		DeleteContext: resourceVisibilityFlowCollectorSflowDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'pkts-received': Total sflow pkts received; 'frag-dropped': Total sflow fragment packets droppped; 'agent-not-found': sflow pkts from not configured agents; 'version-not-supported': sflow version not supported; 'unknown-dir': sflow sample direction is unknown;",
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
func resourceVisibilityFlowCollectorSflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorSflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorSflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorSflowRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityFlowCollectorSflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorSflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorSflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorSflowRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityFlowCollectorSflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorSflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorSflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityFlowCollectorSflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorSflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorSflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityFlowCollectorSflowSamplingEnable(d []interface{}) []edpt.VisibilityFlowCollectorSflowSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorSflowSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorSflowSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorSflow(d *schema.ResourceData) edpt.VisibilityFlowCollectorSflow {
	var ret edpt.VisibilityFlowCollectorSflow
	ret.Inst.SamplingEnable = getSliceVisibilityFlowCollectorSflowSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
