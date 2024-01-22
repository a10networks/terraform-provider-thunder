package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterClusterDevicesMinimumNodes() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_cluster_devices_minimum_nodes`: Minimum number of nodes to start scaleout\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterClusterDevicesMinimumNodesCreate,
		UpdateContext: resourceScaleoutClusterClusterDevicesMinimumNodesUpdate,
		ReadContext:   resourceScaleoutClusterClusterDevicesMinimumNodesRead,
		DeleteContext: resourceScaleoutClusterClusterDevicesMinimumNodesDelete,

		Schema: map[string]*schema.Schema{
			"minimum_nodes_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the minimum number of the node required to start service",
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
func resourceScaleoutClusterClusterDevicesMinimumNodesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesMinimumNodesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesMinimumNodes(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesMinimumNodesRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesMinimumNodesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesMinimumNodesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesMinimumNodes(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterClusterDevicesMinimumNodesRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterClusterDevicesMinimumNodesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesMinimumNodesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesMinimumNodes(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterClusterDevicesMinimumNodesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterClusterDevicesMinimumNodesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterClusterDevicesMinimumNodes(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutClusterClusterDevicesMinimumNodes(d *schema.ResourceData) edpt.ScaleoutClusterClusterDevicesMinimumNodes {
	var ret edpt.ScaleoutClusterClusterDevicesMinimumNodes
	ret.Inst.MinimumNodesNum = d.Get("minimum_nodes_num").(int)
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
