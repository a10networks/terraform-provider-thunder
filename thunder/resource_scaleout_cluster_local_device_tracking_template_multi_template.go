package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_tracking_template_multi_template`: Configure multi tracking template to be used by scaleout\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;",
			},
			"multi_template": {
				Type: schema.TypeString, Required: true, Description: "bind multi tracking template name",
			},
			"template": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_name": {
							Type: schema.TypeString, Optional: true, Description: "bind tracking template name",
						},
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "Partition name",
						},
					},
				},
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
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
func resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate
		oi.TemplateName = in["template_name"].(string)
		oi.PartitionName = in["partition_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate {
	var ret edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplate
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.MultiTemplate = d.Get("multi_template").(string)
	ret.Inst.Template = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateTemplate(d.Get("template").([]interface{}))
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
