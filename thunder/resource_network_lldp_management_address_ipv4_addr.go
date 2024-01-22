package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkLldpManagementAddressIpv4Addr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_lldp_management_address_ipv4_addr`: Configure lldp management-address ipv4 address\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkLldpManagementAddressIpv4AddrCreate,
		UpdateContext: resourceNetworkLldpManagementAddressIpv4AddrUpdate,
		ReadContext:   resourceNetworkLldpManagementAddressIpv4AddrRead,
		DeleteContext: resourceNetworkLldpManagementAddressIpv4AddrDelete,

		Schema: map[string]*schema.Schema{
			"interface_ipv4": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_eth": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
						},
						"ipv4_ve": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ve (lldp management-address interface port number)",
						},
						"ipv4_mgmt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
						},
					},
				},
			},
			"ipv4": {
				Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is ipv4 (lldp management-address ipv4 address)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkLldpManagementAddressIpv4AddrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv4AddrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv4Addr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressIpv4AddrRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkLldpManagementAddressIpv4AddrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv4AddrUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv4Addr(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressIpv4AddrRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkLldpManagementAddressIpv4AddrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv4AddrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv4Addr(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkLldpManagementAddressIpv4AddrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv4AddrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv4Addr(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkLldpManagementAddressIpv4AddrInterfaceIpv4(d []interface{}) edpt.NetworkLldpManagementAddressIpv4AddrInterfaceIpv4 {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressIpv4AddrInterfaceIpv4
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Eth = in["ipv4_eth"].(int)
		ret.Ipv4Ve = in["ipv4_ve"].(int)
		ret.Ipv4Mgmt = in["ipv4_mgmt"].(int)
	}
	return ret
}

func dataToEndpointNetworkLldpManagementAddressIpv4Addr(d *schema.ResourceData) edpt.NetworkLldpManagementAddressIpv4Addr {
	var ret edpt.NetworkLldpManagementAddressIpv4Addr
	ret.Inst.InterfaceIpv4 = getObjectNetworkLldpManagementAddressIpv4AddrInterfaceIpv4(d.Get("interface_ipv4").([]interface{}))
	ret.Inst.Ipv4 = d.Get("ipv4").(string)
	//omit uuid
	return ret
}
