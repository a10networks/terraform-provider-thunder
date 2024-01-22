package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWire() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire`: Configure Virtual Wire\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireCreate,
		UpdateContext: resourceNetworkVirtualWireUpdate,
		ReadContext:   resourceNetworkVirtualWireRead,
		DeleteContext: resourceNetworkVirtualWireDelete,

		Schema: map[string]*schema.Schema{
			"eth1": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"eth2": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"group1": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"group2": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"trunk1": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"trunk2": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_wire_id": {
				Type: schema.TypeInt, Required: true, Description: "virtual wire id",
			},
		},
	}
}
func resourceNetworkVirtualWireCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWire(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWire(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWire(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWire(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkVirtualWire(d *schema.ResourceData) edpt.NetworkVirtualWire {
	var ret edpt.NetworkVirtualWire
	ret.Inst.Eth1 = d.Get("eth1").(int)
	ret.Inst.Eth2 = d.Get("eth2").(int)
	ret.Inst.Group1 = d.Get("group1").(int)
	ret.Inst.Group2 = d.Get("group2").(int)
	ret.Inst.Trunk1 = d.Get("trunk1").(int)
	ret.Inst.Trunk2 = d.Get("trunk2").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VirtualWireId = d.Get("virtual_wire_id").(int)
	return ret
}
