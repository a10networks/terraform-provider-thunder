package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetTrunkGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_trunk_group`: Trunk Group Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetTrunkGroupCreate,
		UpdateContext: resourceInterfaceEthernetTrunkGroupUpdate,
		ReadContext:   resourceInterfaceEthernetTrunkGroupRead,
		DeleteContext: resourceInterfaceEthernetTrunkGroupDelete,

		Schema: map[string]*schema.Schema{
			"admin_key": {
				Type: schema.TypeInt, Optional: true, Description: "LACP admin key (Admin key value)",
			},
			"mode": {
				Type: schema.TypeString, Optional: true, Default: "active", Description: "'active': enable initiation of LACP negotiation on a port(default); 'passive': disable initiation of LACP negotiation on a port;",
			},
			"port_priority": {
				Type: schema.TypeInt, Optional: true, Description: "Set LACP priority for a port (LACP port priority)",
			},
			"timeout": {
				Type: schema.TypeString, Optional: true, Default: "long", Description: "'long': Set LACP long timeout (default); 'short': Set LACP short timeout;",
			},
			"trunk_number": {
				Type: schema.TypeInt, Required: true, Description: "Trunk Number",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Default: "static", Description: "'static': Static (default); 'lacp': lacp; 'lacp-udld': lacp-udld;",
			},
			"udld_timeout_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fast": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "fast timeout in unit of milli-seconds(default 1000)",
						},
						"slow": {
							Type: schema.TypeInt, Optional: true, Description: "slow timeout in unit of seconds",
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
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetTrunkGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetTrunkGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetTrunkGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetTrunkGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetTrunkGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetTrunkGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetTrunkGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetTrunkGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetTrunkGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetTrunkGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetTrunkGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetTrunkGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetTrunkGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetTrunkGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetTrunkGroupUdldTimeoutCfg(d []interface{}) edpt.InterfaceEthernetTrunkGroupUdldTimeoutCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetTrunkGroupUdldTimeoutCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Fast = in["fast"].(int)
		ret.Slow = in["slow"].(int)
	}
	return ret
}

func dataToEndpointInterfaceEthernetTrunkGroup(d *schema.ResourceData) edpt.InterfaceEthernetTrunkGroup {
	var ret edpt.InterfaceEthernetTrunkGroup
	ret.Inst.AdminKey = d.Get("admin_key").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.PortPriority = d.Get("port_priority").(int)
	ret.Inst.Timeout = d.Get("timeout").(string)
	ret.Inst.TrunkNumber = d.Get("trunk_number").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UdldTimeoutCfg = getObjectInterfaceEthernetTrunkGroupUdldTimeoutCfg(d.Get("udld_timeout_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
