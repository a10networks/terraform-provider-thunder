package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterServiceConfigTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_service_config_template`: Scaleout template\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterServiceConfigTemplateCreate,
		UpdateContext: resourceScaleoutClusterServiceConfigTemplateUpdate,
		ReadContext:   resourceScaleoutClusterServiceConfigTemplateRead,
		DeleteContext: resourceScaleoutClusterServiceConfigTemplateDelete,

		Schema: map[string]*schema.Schema{
			"bucket_count": {
				Type: schema.TypeInt, Optional: true, Default: 256, Description: "Number of traffic buckets",
			},
			"device_group": {
				Type: schema.TypeInt, Optional: true, Description: "Device group id",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Scaleout template Name",
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
func resourceScaleoutClusterServiceConfigTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfigTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterServiceConfigTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterServiceConfigTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfigTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterServiceConfigTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterServiceConfigTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfigTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterServiceConfigTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfigTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterServiceConfigTemplate(d *schema.ResourceData) edpt.ScaleoutClusterServiceConfigTemplate {
	var ret edpt.ScaleoutClusterServiceConfigTemplate
	ret.Inst.BucketCount = d.Get("bucket_count").(int)
	ret.Inst.DeviceGroup = d.Get("device_group").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
