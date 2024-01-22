package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterDeviceGroupsDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_device_groups_device_group`: Configure scaleout device groups\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterDeviceGroupsDeviceGroupCreate,
		UpdateContext: resourceScaleoutClusterDeviceGroupsDeviceGroupUpdate,
		ReadContext:   resourceScaleoutClusterDeviceGroupsDeviceGroupRead,
		DeleteContext: resourceScaleoutClusterDeviceGroupsDeviceGroupDelete,

		Schema: map[string]*schema.Schema{
			"device_group": {
				Type: schema.TypeInt, Required: true, Description: "scaleout device group",
			},
			"device_id_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"device_id_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
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
func resourceScaleoutClusterDeviceGroupsDeviceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsDeviceGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroupsDeviceGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDeviceGroupsDeviceGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterDeviceGroupsDeviceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsDeviceGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroupsDeviceGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDeviceGroupsDeviceGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterDeviceGroupsDeviceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsDeviceGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroupsDeviceGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterDeviceGroupsDeviceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsDeviceGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroupsDeviceGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList
		oi.DeviceIdStart = in["device_id_start"].(int)
		oi.DeviceIdEnd = in["device_id_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterDeviceGroupsDeviceGroup(d *schema.ResourceData) edpt.ScaleoutClusterDeviceGroupsDeviceGroup {
	var ret edpt.ScaleoutClusterDeviceGroupsDeviceGroup
	ret.Inst.DeviceGroup = d.Get("device_group").(int)
	ret.Inst.DeviceIdList = getSliceScaleoutClusterDeviceGroupsDeviceGroupDeviceIdList(d.Get("device_id_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
