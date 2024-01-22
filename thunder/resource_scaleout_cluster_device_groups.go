package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterDeviceGroups() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_device_groups`: Configure scaleout device groups\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterDeviceGroupsCreate,
		UpdateContext: resourceScaleoutClusterDeviceGroupsUpdate,
		ReadContext:   resourceScaleoutClusterDeviceGroupsRead,
		DeleteContext: resourceScaleoutClusterDeviceGroupsDelete,

		Schema: map[string]*schema.Schema{
			"device_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
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
func resourceScaleoutClusterDeviceGroupsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroups(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDeviceGroupsRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterDeviceGroupsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroups(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterDeviceGroupsRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterDeviceGroupsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroups(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterDeviceGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDeviceGroupsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterDeviceGroups(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupList(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupList
		oi.DeviceGroup = in["device_group"].(int)
		oi.DeviceIdList = getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList(in["device_id_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList
		oi.DeviceIdStart = in["device_id_start"].(int)
		oi.DeviceIdEnd = in["device_id_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterDeviceGroups(d *schema.ResourceData) edpt.ScaleoutClusterDeviceGroups {
	var ret edpt.ScaleoutClusterDeviceGroups
	ret.Inst.DeviceGroupList = getSliceScaleoutClusterDeviceGroupsDeviceGroupList(d.Get("device_group_list").([]interface{}))
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
