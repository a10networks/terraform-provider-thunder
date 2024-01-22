package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterServiceConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_service_config`: Configure scaleout templates for SLB, CGN and VRRP\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterServiceConfigCreate,
		UpdateContext: resourceScaleoutClusterServiceConfigUpdate,
		ReadContext:   resourceScaleoutClusterServiceConfigRead,
		DeleteContext: resourceScaleoutClusterServiceConfigDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"template_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Scaleout template Name",
						},
						"bucket_count": {
							Type: schema.TypeInt, Optional: true, Default: 256, Description: "Number of traffic buckets",
						},
						"device_group": {
							Type: schema.TypeInt, Optional: true, Description: "Device group id",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
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
func resourceScaleoutClusterServiceConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterServiceConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterServiceConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterServiceConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterServiceConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterServiceConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterServiceConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterServiceConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterServiceConfigTemplateList(d []interface{}) []edpt.ScaleoutClusterServiceConfigTemplateList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterServiceConfigTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterServiceConfigTemplateList
		oi.Name = in["name"].(string)
		oi.BucketCount = in["bucket_count"].(int)
		oi.DeviceGroup = in["device_group"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterServiceConfig(d *schema.ResourceData) edpt.ScaleoutClusterServiceConfig {
	var ret edpt.ScaleoutClusterServiceConfig
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.TemplateList = getSliceScaleoutClusterServiceConfigTemplateList(d.Get("template_list").([]interface{}))
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
