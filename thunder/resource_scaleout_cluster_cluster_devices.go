package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterClusterDevices() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_cluster_devices`: Configure devices in the cluster\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterClusterDevicesCreate,
		UpdateContext: resourceScaleoutClusterClusterDevicesUpdate,
		ReadContext:   resourceScaleoutClusterClusterDevicesRead,
		DeleteContext: resourceScaleoutClusterClusterDevicesDelete,

		Schema: map[string]*schema.Schema{
			"cluster_discovery_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"device_id_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"minimum_nodes": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"minimum_nodes_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the minimum number of the node required to start service",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceScaleoutClusterClusterDevicesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevices(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevices(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterClusterDevicesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevices(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevices(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325(d []interface{}) edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325 {

	var ret edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325
	return ret
}

func getSliceScaleoutClusterClusterDevicesDeviceIdList(d []interface{}) []edpt.ScaleoutClusterClusterDevicesDeviceIdList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterClusterDevicesDeviceIdList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterClusterDevicesDeviceIdList
		oi.Ip = in["ip"].(string)
		oi.Action = in["action"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterClusterDevicesMinimumNodes1326(d []interface{}) edpt.ScaleoutClusterClusterDevicesMinimumNodes1326 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterClusterDevicesMinimumNodes1326
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinimumNodesNum = in["minimum_nodes_num"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointScaleoutClusterClusterDevices(d *schema.ResourceData) edpt.ScaleoutClusterClusterDevices {
	var ret edpt.ScaleoutClusterClusterDevices
	ret.Inst.ClusterDiscoveryTimeout = getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1325(d.Get("cluster_discovery_timeout").([]interface{}))
	ret.Inst.DeviceIdList = getSliceScaleoutClusterClusterDevicesDeviceIdList(d.Get("device_id_list").([]interface{}))
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.MinimumNodes = getObjectScaleoutClusterClusterDevicesMinimumNodes1326(d.Get("minimum_nodes").([]interface{}))
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
