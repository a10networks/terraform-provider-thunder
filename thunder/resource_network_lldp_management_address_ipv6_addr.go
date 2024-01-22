package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkLldpManagementAddressIpv6Addr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_lldp_management_address_ipv6_addr`: Configure lldp management-address ipv6 address\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkLldpManagementAddressIpv6AddrCreate,
		UpdateContext: resourceNetworkLldpManagementAddressIpv6AddrUpdate,
		ReadContext:   resourceNetworkLldpManagementAddressIpv6AddrRead,
		DeleteContext: resourceNetworkLldpManagementAddressIpv6AddrDelete,

		Schema: map[string]*schema.Schema{
			"interface_ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_eth": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
						},
						"ipv6_ve": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ve (lldp management-address interface port number)",
						},
						"ipv6_mgmt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is ipv6 (lldp management-address ipv6 address)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkLldpManagementAddressIpv6AddrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv6AddrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv6Addr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressIpv6AddrRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkLldpManagementAddressIpv6AddrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv6AddrUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv6Addr(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressIpv6AddrRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkLldpManagementAddressIpv6AddrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv6AddrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv6Addr(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkLldpManagementAddressIpv6AddrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressIpv6AddrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressIpv6Addr(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkLldpManagementAddressIpv6AddrInterfaceIpv6(d []interface{}) edpt.NetworkLldpManagementAddressIpv6AddrInterfaceIpv6 {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressIpv6AddrInterfaceIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6Eth = in["ipv6_eth"].(int)
		ret.Ipv6Ve = in["ipv6_ve"].(int)
		ret.Ipv6Mgmt = in["ipv6_mgmt"].(int)
	}
	return ret
}

func dataToEndpointNetworkLldpManagementAddressIpv6Addr(d *schema.ResourceData) edpt.NetworkLldpManagementAddressIpv6Addr {
	var ret edpt.NetworkLldpManagementAddressIpv6Addr
	ret.Inst.InterfaceIpv6 = getObjectNetworkLldpManagementAddressIpv6AddrInterfaceIpv6(d.Get("interface_ipv6").([]interface{}))
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	return ret
}
