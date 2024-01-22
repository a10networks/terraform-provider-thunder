package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkMacAddressStatic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_mac_address_static`: Static MAC commands\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkMacAddressStaticCreate,
		UpdateContext: resourceNetworkMacAddressStaticUpdate,
		ReadContext:   resourceNetworkMacAddressStaticRead,
		DeleteContext: resourceNetworkMacAddressStaticDelete,

		Schema: map[string]*schema.Schema{
			"dest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Trap MAC with this DA to CPU",
			},
			"mac": {
				Type: schema.TypeString, Required: true, Description: "Configure a Static MAC address",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet Port on which the Address is applicable (Port Value)",
			},
			"redirect_dummy_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy for redirect use",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN Id",
			},
		},
	}
}
func resourceNetworkMacAddressStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAddressStaticCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAddressStatic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkMacAddressStaticRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkMacAddressStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAddressStaticUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAddressStatic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkMacAddressStaticRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkMacAddressStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAddressStaticDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAddressStatic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkMacAddressStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAddressStaticRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAddressStatic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkMacAddressStatic(d *schema.ResourceData) edpt.NetworkMacAddressStatic {
	var ret edpt.NetworkMacAddressStatic
	ret.Inst.Dest = d.Get("dest").(int)
	ret.Inst.Mac = d.Get("mac").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.RedirectDummyMac = d.Get("redirect_dummy_mac").(int)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
