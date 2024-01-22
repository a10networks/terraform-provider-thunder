package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowCollectorTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_collector_template`: Configure NetFlow/IPFIX collector templates\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowCollectorTemplateCreate,
		UpdateContext: resourceNetflowCollectorTemplateUpdate,
		ReadContext:   resourceNetflowCollectorTemplateRead,
		DeleteContext: resourceNetflowCollectorTemplateDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'v9-templates-created': Netflow v9 Templates Created; 'v9-templates-deleted': Netflow v9 Templates Deleted; 'v10-templates-created': Netflow v10(IPFIX) Templates Created; 'v10-templates-deleted': Netflow v10(IPFIX) Templates Deleted; 'template-drop-exceeded': Netflow Templates Dropped Maximum Limit Reached; 'template-drop-out-of-memory': Netflow Templates Dropped OOM; 'templates-added-to-delq': Netflow Templates Added To Delete Queue; 'templates-removed-from-delq': Netflow Templates Removed From Delete Queue;",
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
func resourceNetflowCollectorTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCollectorTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowCollectorTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCollectorTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowCollectorTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowCollectorTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetflowCollectorTemplateSamplingEnable(d []interface{}) []edpt.NetflowCollectorTemplateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetflowCollectorTemplateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowCollectorTemplateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowCollectorTemplate(d *schema.ResourceData) edpt.NetflowCollectorTemplate {
	var ret edpt.NetflowCollectorTemplate
	ret.Inst.SamplingEnable = getSliceNetflowCollectorTemplateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
