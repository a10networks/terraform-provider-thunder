package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkArpStatic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_arp_static`: static ARP Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkArpStaticCreate,
		UpdateContext: resourceNetworkArpStaticUpdate,
		ReadContext:   resourceNetworkArpStaticRead,
		DeleteContext: resourceNetworkArpStaticDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Port Value)",
			},
			"ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IP address",
			},
			"mac_addr": {
				Type: schema.TypeString, Optional: true, Description: "MAC address",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk group",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN ID",
			},
		},
	}
}
func resourceNetworkArpStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpStaticCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpStatic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkArpStaticRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkArpStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpStaticUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpStatic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkArpStaticRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkArpStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpStaticDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpStatic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkArpStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpStaticRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpStatic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkArpStatic(d *schema.ResourceData) edpt.NetworkArpStatic {
	var ret edpt.NetworkArpStatic
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.MacAddr = d.Get("mac_addr").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
