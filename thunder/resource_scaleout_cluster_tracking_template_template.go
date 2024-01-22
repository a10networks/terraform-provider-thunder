package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterTrackingTemplateTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_tracking_template_template`: Configure tracking template to be used by scaleout [OBSOLETED!]\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterTrackingTemplateTemplateCreate,
		UpdateContext: resourceScaleoutClusterTrackingTemplateTemplateUpdate,
		ReadContext:   resourceScaleoutClusterTrackingTemplateTemplateRead,
		DeleteContext: resourceScaleoutClusterTrackingTemplateTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template": {
				Type: schema.TypeString, Required: true, Description: "bind tracking template name",
			},
			"threshold_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
							Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;",
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
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterTrackingTemplateTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterTrackingTemplateTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterTrackingTemplateTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterTrackingTemplateTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterTrackingTemplateTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterTrackingTemplateTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterTrackingTemplateTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterTrackingTemplateTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterTrackingTemplateTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterTrackingTemplateTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterTrackingTemplateTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterTrackingTemplateTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterTrackingTemplateTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterTrackingTemplateTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterTrackingTemplateTemplateThresholdCfg(d []interface{}) []edpt.ScaleoutClusterTrackingTemplateTemplateThresholdCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterTrackingTemplateTemplateThresholdCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterTrackingTemplateTemplateThresholdCfg
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterTrackingTemplateTemplate(d *schema.ResourceData) edpt.ScaleoutClusterTrackingTemplateTemplate {
	var ret edpt.ScaleoutClusterTrackingTemplateTemplate
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.ThresholdCfg = getSliceScaleoutClusterTrackingTemplateTemplateThresholdCfg(d.Get("threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
