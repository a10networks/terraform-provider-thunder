package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetSpanningTree() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_spanning_tree`: Spanning tree port configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetSpanningTreeCreate,
		UpdateContext: resourceInterfaceEthernetSpanningTreeUpdate,
		ReadContext:   resourceInterfaceEthernetSpanningTreeRead,
		DeleteContext: resourceInterfaceEthernetSpanningTreeDelete,

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
func resourceInterfaceEthernetSpanningTreeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetSpanningTreeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetSpanningTree(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetSpanningTreeRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetSpanningTreeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetSpanningTreeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetSpanningTree(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetSpanningTreeRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetSpanningTreeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetSpanningTreeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetSpanningTree(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetSpanningTreeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetSpanningTreeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetSpanningTree(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceEthernetSpanningTreeInstanceList(d []interface{}) []edpt.InterfaceEthernetSpanningTreeInstanceList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetSpanningTreeInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetSpanningTreeInstanceList
		oi.InstanceStart = in["instance_start"].(int)
		oi.MstpPathCost = in["mstp_path_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceEthernetSpanningTree(d *schema.ResourceData) edpt.InterfaceEthernetSpanningTree {
	var ret edpt.InterfaceEthernetSpanningTree
	ret.Inst.AdminEdge = d.Get("admin_edge").(int)
	ret.Inst.AutoEdge = d.Get("auto_edge").(int)
	ret.Inst.InstanceList = getSliceInterfaceEthernetSpanningTreeInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.PathCost = d.Get("path_cost").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
