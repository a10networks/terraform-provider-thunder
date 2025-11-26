package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireHealthCheckOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_virtual_wire_health_check_oper`: Operational Status for the object virtual-wire-health-check\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVirtualWireHealthCheckOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vlan_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN ID, specify 1 for untagged traffic",
			},
		},
	}
}

func resourceNetworkVirtualWireHealthCheckOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheckOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVirtualWireHealthCheckOperOper := setObjectNetworkVirtualWireHealthCheckOperOper(res)
		d.Set("oper", NetworkVirtualWireHealthCheckOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVirtualWireHealthCheckOperOper(ret edpt.DataNetworkVirtualWireHealthCheckOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_state": ret.DtNetworkVirtualWireHealthCheckOper.Oper.EntryState,
			"vlan_state":  ret.DtNetworkVirtualWireHealthCheckOper.Oper.VlanState,
		},
	}
}

func getObjectNetworkVirtualWireHealthCheckOperOper(d []interface{}) edpt.NetworkVirtualWireHealthCheckOperOper {

	count1 := len(d)
	var ret edpt.NetworkVirtualWireHealthCheckOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryState = in["entry_state"].(string)
		ret.VlanState = in["vlan_state"].(string)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireHealthCheckOper(d *schema.ResourceData) edpt.NetworkVirtualWireHealthCheckOper {
	var ret edpt.NetworkVirtualWireHealthCheckOper

	ret.Oper = getObjectNetworkVirtualWireHealthCheckOperOper(d.Get("oper").([]interface{}))

	ret.Vlan = d.Get("vlan").(int)
	return ret
}
