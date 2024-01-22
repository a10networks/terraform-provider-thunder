package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_interface_ethernet`: VRRP-A interface ethernet\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAInterfaceEthernetCreate,
		UpdateContext: resourceVrrpAInterfaceEthernetUpdate,
		ReadContext:   resourceVrrpAInterfaceEthernetRead,
		DeleteContext: resourceVrrpAInterfaceEthernetDelete,

		Schema: map[string]*schema.Schema{
			"both": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "both a router and server interface",
			},
			"ethernet_val": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet Interface",
			},
			"no_heartbeat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "do not send out heartbeat packet from this interface",
			},
			"router_interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "interface to upstream router",
			},
			"server_interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "interface to real server",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
						},
					},
				},
			},
		},
	}
}
func resourceVrrpAInterfaceEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceEthernetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceEthernet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAInterfaceEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAInterfaceEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceEthernetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceEthernet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAInterfaceEthernetRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAInterfaceEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceEthernetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceEthernet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAInterfaceEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceEthernetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceEthernet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVrrpAInterfaceEthernetVlanCfg(d []interface{}) []edpt.VrrpAInterfaceEthernetVlanCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAInterfaceEthernetVlanCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAInterfaceEthernetVlanCfg
		oi.Vlan = in["vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAInterfaceEthernet(d *schema.ResourceData) edpt.VrrpAInterfaceEthernet {
	var ret edpt.VrrpAInterfaceEthernet
	ret.Inst.Both = d.Get("both").(int)
	ret.Inst.EthernetVal = d.Get("ethernet_val").(int)
	ret.Inst.NoHeartbeat = d.Get("no_heartbeat").(int)
	ret.Inst.RouterInterface = d.Get("router_interface").(int)
	ret.Inst.ServerInterface = d.Get("server_interface").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanCfg = getSliceVrrpAInterfaceEthernetVlanCfg(d.Get("vlan_cfg").([]interface{}))
	return ret
}
