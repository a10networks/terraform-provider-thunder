package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeMstpInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_spanning_tree_mode_mstp_instance`: Enable spanning tree instance ID\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkSpanningTreeModeMstpInstanceCreate,
		UpdateContext: resourceNetworkSpanningTreeModeMstpInstanceUpdate,
		ReadContext:   resourceNetworkSpanningTreeModeMstpInstanceRead,
		DeleteContext: resourceNetworkSpanningTreeModeMstpInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_start": {
				Type: schema.TypeInt, Required: true, Description: "Instance ID",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Set instance priority",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_start": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
						},
						"vlan_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}
func resourceNetworkSpanningTreeModeMstpInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstpInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeMstpInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkSpanningTreeModeMstpInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstpInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeMstpInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkSpanningTreeModeMstpInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstpInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkSpanningTreeModeMstpInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstpInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkSpanningTreeModeMstpInstanceVlanList(d []interface{}) []edpt.NetworkSpanningTreeModeMstpInstanceVlanList {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeMstpInstanceVlanList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeMstpInstanceVlanList
		oi.VlanStart = in["vlan_start"].(int)
		oi.VlanEnd = in["vlan_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkSpanningTreeModeMstpInstance(d *schema.ResourceData) edpt.NetworkSpanningTreeModeMstpInstance {
	var ret edpt.NetworkSpanningTreeModeMstpInstance
	ret.Inst.InstanceStart = d.Get("instance_start").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanList = getSliceNetworkSpanningTreeModeMstpInstanceVlanList(d.Get("vlan_list").([]interface{}))
	return ret
}
