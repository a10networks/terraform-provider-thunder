package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterClusterDevicesDeviceId() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_cluster_devices_device_id`: Configure Scaleout devices\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterClusterDevicesDeviceIdCreate,
		UpdateContext: resourceScaleoutClusterClusterDevicesDeviceIdUpdate,
		ReadContext:   resourceScaleoutClusterClusterDevicesDeviceIdRead,
		DeleteContext: resourceScaleoutClusterClusterDevicesDeviceIdDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"device_id": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "",
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
func resourceScaleoutClusterClusterDevicesDeviceIdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesDeviceIdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesDeviceId(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesDeviceIdRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesDeviceIdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesDeviceIdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesDeviceId(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesDeviceIdRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterClusterDevicesDeviceIdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesDeviceIdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesDeviceId(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesDeviceIdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesDeviceIdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesDeviceId(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterClusterDevicesDeviceId(d *schema.ResourceData) edpt.ScaleoutClusterClusterDevicesDeviceId {
	var ret edpt.ScaleoutClusterClusterDevicesDeviceId
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DeviceId = d.Get("device_id").(int)
	ret.Inst.Ip = d.Get("ip").(string)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
