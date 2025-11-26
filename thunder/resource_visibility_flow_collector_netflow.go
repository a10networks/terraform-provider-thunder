package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_flow_collector_netflow`: NetFlow/IPFIX collector\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityFlowCollectorNetflowCreate,
		UpdateContext: resourceVisibilityFlowCollectorNetflowUpdate,
		ReadContext:   resourceVisibilityFlowCollectorNetflowRead,
		DeleteContext: resourceVisibilityFlowCollectorNetflowDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'pkts-rcvd': Total nflow packets received; 'v9-templates-created': Total v9 templates created; 'v9-templates-deleted': Total v9 templates deleted; 'v10-templates-created': Total v10(IPFIX) templates created; 'v10-templates-deleted': Total v10(IPFIX) templates deleted; 'template-drop-exceeded': Total templates dropped because of maximum limit; 'template-drop-out-of-memory': Total templates dropped becuase of out of memory; 'frag-dropped': Total nflow fragment packets dropped; 'agent-not-found': nflow pkts from not configured agents; 'version-not-supported': nflow version not supported; 'unknown-dir': nflow sample direction is unknown;",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'templates-added-to-delq': Netflow templates added to the delete queue; 'templates-removed-from-delq': Netflow templates removed from the delete queue;",
									},
								},
							},
						},
						"detail": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
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
func resourceVisibilityFlowCollectorNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorNetflowRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityFlowCollectorNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorNetflowRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityFlowCollectorNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityFlowCollectorNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityFlowCollectorNetflowSamplingEnable(d []interface{}) []edpt.VisibilityFlowCollectorNetflowSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowTemplate2036(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplate2036 {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplate2036
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceVisibilityFlowCollectorNetflowTemplateSamplingEnable2037(in["sampling_enable"].([]interface{}))
		ret.Detail = getObjectVisibilityFlowCollectorNetflowTemplateDetail2038(in["detail"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateSamplingEnable2037(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable2037 {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable2037, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable2037
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowTemplateDetail2038(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateDetail2038 {

	var ret edpt.VisibilityFlowCollectorNetflowTemplateDetail2038
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflow(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflow {
	var ret edpt.VisibilityFlowCollectorNetflow
	ret.Inst.SamplingEnable = getSliceVisibilityFlowCollectorNetflowSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Template = getObjectVisibilityFlowCollectorNetflowTemplate2036(d.Get("template").([]interface{}))
	//omit uuid
	return ret
}
