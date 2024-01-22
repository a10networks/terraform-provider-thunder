package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkLldpManagementAddressDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_lldp_management_address_dns`: Configure lldp management-address dns address\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkLldpManagementAddressDnsCreate,
		UpdateContext: resourceNetworkLldpManagementAddressDnsUpdate,
		ReadContext:   resourceNetworkLldpManagementAddressDnsRead,
		DeleteContext: resourceNetworkLldpManagementAddressDnsDelete,

		Schema: map[string]*schema.Schema{
			"dns": {
				Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is dns (lldp management-address dns address)",
			},
			"interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface management (lldp management-address interface port number)",
						},
						"management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkLldpManagementAddressDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkLldpManagementAddressDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpManagementAddressDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkLldpManagementAddressDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkLldpManagementAddressDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpManagementAddressDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldpManagementAddressDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkLldpManagementAddressDnsInterface(d []interface{}) edpt.NetworkLldpManagementAddressDnsInterface {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressDnsInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Ve = in["ve"].(int)
		ret.Management = in["management"].(int)
	}
	return ret
}

func dataToEndpointNetworkLldpManagementAddressDns(d *schema.ResourceData) edpt.NetworkLldpManagementAddressDns {
	var ret edpt.NetworkLldpManagementAddressDns
	ret.Inst.Dns = d.Get("dns").(string)
	ret.Inst.Interface = getObjectNetworkLldpManagementAddressDnsInterface(d.Get("interface").([]interface{}))
	//omit uuid
	return ret
}
