package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflowTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_flow_collector_netflow_template`: NetFlow/IPFIX collector templates\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityFlowCollectorNetflowTemplateCreate,
		UpdateContext: resourceVisibilityFlowCollectorNetflowTemplateUpdate,
		ReadContext:   resourceVisibilityFlowCollectorNetflowTemplateRead,
		DeleteContext: resourceVisibilityFlowCollectorNetflowTemplateDelete,

		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityFlowCollectorNetflowTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorNetflowTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityFlowCollectorNetflowTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityFlowCollectorNetflowTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityFlowCollectorNetflowTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityFlowCollectorNetflowTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityFlowCollectorNetflowTemplateDetail1910(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateDetail1910 {

	var ret edpt.VisibilityFlowCollectorNetflowTemplateDetail1910
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateSamplingEnable(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflowTemplate(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflowTemplate {
	var ret edpt.VisibilityFlowCollectorNetflowTemplate
	ret.Inst.Detail = getObjectVisibilityFlowCollectorNetflowTemplateDetail1910(d.Get("detail").([]interface{}))
	ret.Inst.SamplingEnable = getSliceVisibilityFlowCollectorNetflowTemplateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
