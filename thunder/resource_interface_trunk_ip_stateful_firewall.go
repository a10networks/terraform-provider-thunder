package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpStatefulFirewall() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ip_stateful_firewall`: Configure Stateful Firewall direction\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpStatefulFirewallCreate,
		UpdateContext: resourceInterfaceTrunkIpStatefulFirewallUpdate,
		ReadContext:   resourceInterfaceTrunkIpStatefulFirewallRead,
		DeleteContext: resourceInterfaceTrunkIpStatefulFirewallDelete,

		Schema: map[string]*schema.Schema{
			"access_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
			},
			"acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "ACL id",
			},
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
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
func resourceInterfaceTrunkIpStatefulFirewallCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpStatefulFirewallCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpStatefulFirewall(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpStatefulFirewallRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpStatefulFirewallUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpStatefulFirewallUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpStatefulFirewall(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpStatefulFirewallRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpStatefulFirewallDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpStatefulFirewallDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpStatefulFirewall(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpStatefulFirewallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpStatefulFirewallRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpStatefulFirewall(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkIpStatefulFirewall(d *schema.ResourceData) edpt.InterfaceTrunkIpStatefulFirewall {
	var ret edpt.InterfaceTrunkIpStatefulFirewall
	ret.Inst.AccessList = d.Get("access_list").(int)
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.Outside = d.Get("outside").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
