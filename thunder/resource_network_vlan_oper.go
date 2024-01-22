package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_vlan_oper`: Operational Status for the object vlan\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkVlanOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ve_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_shared_vlan": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"un_tagg_eth_ports": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ports": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tagg_eth_ports": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ports": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"un_tagg_logical_ports": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ports": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tagg_logical_ports": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ports": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"span_tree": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"vlan_num": {
				Type: schema.TypeInt, Required: true, Description: "VLAN number",
			},
		},
	}
}

func resourceNetworkVlanOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkVlanOperOper := setObjectNetworkVlanOperOper(res)
		d.Set("oper", NetworkVlanOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkVlanOperOper(ret edpt.DataNetworkVlanOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vlan_name":             ret.DtNetworkVlanOper.Oper.Vlan_name,
			"ve_num":                ret.DtNetworkVlanOper.Oper.Ve_num,
			"is_shared_vlan":        ret.DtNetworkVlanOper.Oper.Is_shared_vlan,
			"un_tagg_eth_ports":     setObjectNetworkVlanOperOperUn_tagg_eth_ports(ret.DtNetworkVlanOper.Oper.Un_tagg_eth_ports),
			"tagg_eth_ports":        setObjectNetworkVlanOperOperTagg_eth_ports(ret.DtNetworkVlanOper.Oper.Tagg_eth_ports),
			"un_tagg_logical_ports": setObjectNetworkVlanOperOperUn_tagg_logical_ports(ret.DtNetworkVlanOper.Oper.Un_tagg_logical_ports),
			"tagg_logical_ports":    setObjectNetworkVlanOperOperTagg_logical_ports(ret.DtNetworkVlanOper.Oper.Tagg_logical_ports),
			"span_tree":             ret.DtNetworkVlanOper.Oper.SpanTree,
		},
	}
}

func setObjectNetworkVlanOperOperUn_tagg_eth_ports(d edpt.NetworkVlanOperOperUn_tagg_eth_ports) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ports"] = d.Ports
	result = append(result, in)
	return result
}

func setObjectNetworkVlanOperOperTagg_eth_ports(d edpt.NetworkVlanOperOperTagg_eth_ports) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ports"] = d.Ports
	result = append(result, in)
	return result
}

func setObjectNetworkVlanOperOperUn_tagg_logical_ports(d edpt.NetworkVlanOperOperUn_tagg_logical_ports) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ports"] = d.Ports
	result = append(result, in)
	return result
}

func setObjectNetworkVlanOperOperTagg_logical_ports(d edpt.NetworkVlanOperOperTagg_logical_ports) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ports"] = d.Ports
	result = append(result, in)
	return result
}

func getObjectNetworkVlanOperOper(d []interface{}) edpt.NetworkVlanOperOper {

	count1 := len(d)
	var ret edpt.NetworkVlanOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vlan_name = in["vlan_name"].(string)
		ret.Ve_num = in["ve_num"].(int)
		ret.Is_shared_vlan = in["is_shared_vlan"].(string)
		ret.Un_tagg_eth_ports = getObjectNetworkVlanOperOperUn_tagg_eth_ports(in["un_tagg_eth_ports"].([]interface{}))
		ret.Tagg_eth_ports = getObjectNetworkVlanOperOperTagg_eth_ports(in["tagg_eth_ports"].([]interface{}))
		ret.Un_tagg_logical_ports = getObjectNetworkVlanOperOperUn_tagg_logical_ports(in["un_tagg_logical_ports"].([]interface{}))
		ret.Tagg_logical_ports = getObjectNetworkVlanOperOperTagg_logical_ports(in["tagg_logical_ports"].([]interface{}))
		ret.SpanTree = in["span_tree"].(string)
	}
	return ret
}

func getObjectNetworkVlanOperOperUn_tagg_eth_ports(d []interface{}) edpt.NetworkVlanOperOperUn_tagg_eth_ports {

	count1 := len(d)
	var ret edpt.NetworkVlanOperOperUn_tagg_eth_ports
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ports = in["ports"].(int)
	}
	return ret
}

func getObjectNetworkVlanOperOperTagg_eth_ports(d []interface{}) edpt.NetworkVlanOperOperTagg_eth_ports {

	count1 := len(d)
	var ret edpt.NetworkVlanOperOperTagg_eth_ports
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ports = in["ports"].(int)
	}
	return ret
}

func getObjectNetworkVlanOperOperUn_tagg_logical_ports(d []interface{}) edpt.NetworkVlanOperOperUn_tagg_logical_ports {

	count1 := len(d)
	var ret edpt.NetworkVlanOperOperUn_tagg_logical_ports
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ports = in["ports"].(int)
	}
	return ret
}

func getObjectNetworkVlanOperOperTagg_logical_ports(d []interface{}) edpt.NetworkVlanOperOperTagg_logical_ports {

	count1 := len(d)
	var ret edpt.NetworkVlanOperOperTagg_logical_ports
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ports = in["ports"].(int)
	}
	return ret
}

func dataToEndpointNetworkVlanOper(d *schema.ResourceData) edpt.NetworkVlanOper {
	var ret edpt.NetworkVlanOper

	ret.Oper = getObjectNetworkVlanOperOper(d.Get("oper").([]interface{}))

	ret.VlanNum = d.Get("vlan_num").(int)
	return ret
}
