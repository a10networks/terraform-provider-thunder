package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkSpanningTree() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_spanning_tree`: help Spanning tree port configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkSpanningTreeCreate,
		UpdateContext: resourceInterfaceTrunkSpanningTreeUpdate,
		ReadContext:   resourceInterfaceTrunkSpanningTreeRead,
		DeleteContext: resourceInterfaceTrunkSpanningTreeDelete,

		Schema: map[string]*schema.Schema{
			"admin_edge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable admin-edge",
			},
			"auto_edge": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable auto-edge",
			},
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_start": {
							Type: schema.TypeInt, Optional: true, Description: "Instance ID",
						},
						"mstp_path_cost": {
							Type: schema.TypeInt, Optional: true, Description: "Path cost (Limit)",
						},
					},
				},
			},
			"path_cost": {
				Type: schema.TypeInt, Optional: true, Description: "Path cost (Limit)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkSpanningTreeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkSpanningTreeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkSpanningTree(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkSpanningTreeRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkSpanningTreeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkSpanningTreeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkSpanningTree(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkSpanningTreeRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkSpanningTreeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkSpanningTreeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkSpanningTree(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkSpanningTreeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkSpanningTreeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkSpanningTree(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceTrunkSpanningTreeInstanceList(d []interface{}) []edpt.InterfaceTrunkSpanningTreeInstanceList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkSpanningTreeInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkSpanningTreeInstanceList
		oi.InstanceStart = in["instance_start"].(int)
		oi.MstpPathCost = in["mstp_path_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTrunkSpanningTree(d *schema.ResourceData) edpt.InterfaceTrunkSpanningTree {
	var ret edpt.InterfaceTrunkSpanningTree
	ret.Inst.AdminEdge = d.Get("admin_edge").(int)
	ret.Inst.AutoEdge = d.Get("auto_edge").(int)
	ret.Inst.InstanceList = getSliceInterfaceTrunkSpanningTreeInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.PathCost = d.Get("path_cost").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
