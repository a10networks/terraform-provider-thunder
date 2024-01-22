package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrackingTemplateTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_tracking_template_template`: Configure tracking template to be used by scaleout\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateDelete,

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
func resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateThresholdCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateThresholdCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateThresholdCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateThresholdCfg
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateTemplate(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplate {
	var ret edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplate
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.ThresholdCfg = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateThresholdCfg(d.Get("threshold_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
