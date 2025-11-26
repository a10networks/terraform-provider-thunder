package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelDebugEncapTableOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_overlay_tunnel_debug_encap_table_oper`: Operational Status for the object encap-table\n\n__PLACEHOLDER__",
		ReadContext: resourceOverlayTunnelDebugEncapTableOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vtep": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hindex": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_vtep_ip": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 address",
									},
									"dst_vtep_ip": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 address",
									},
									"src_vtep_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dst_vtep_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_ipv6": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_vtep_mac": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"encap_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vtep_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vtep_vnp_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lifname": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partname": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_static": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tot_ip_encap": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tot_ipv6_encap": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceOverlayTunnelDebugEncapTableOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelDebugEncapTableOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelDebugEncapTableOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		OverlayTunnelDebugEncapTableOperOper := setObjectOverlayTunnelDebugEncapTableOperOper(res)
		d.Set("oper", OverlayTunnelDebugEncapTableOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectOverlayTunnelDebugEncapTableOperOper(ret edpt.DataOverlayTunnelDebugEncapTableOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vtep":           setSliceOverlayTunnelDebugEncapTableOperOperVtep(ret.DtOverlayTunnelDebugEncapTableOper.Oper.Vtep),
			"tot_ip_encap":   ret.DtOverlayTunnelDebugEncapTableOper.Oper.Tot_ip_encap,
			"tot_ipv6_encap": ret.DtOverlayTunnelDebugEncapTableOper.Oper.Tot_ipv6_encap,
		},
	}
}

func setSliceOverlayTunnelDebugEncapTableOperOperVtep(d []edpt.OverlayTunnelDebugEncapTableOperOperVtep) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["index"] = item.Index
		in["hindex"] = item.Hindex
		in["src_vtep_ip"] = item.Src_vtep_ip
		in["dst_vtep_ip"] = item.Dst_vtep_ip
		in["src_vtep_ipv6"] = item.Src_vtep_ipv6
		in["dst_vtep_ipv6"] = item.Dst_vtep_ipv6
		in["is_ipv6"] = item.Is_ipv6
		in["dst_vtep_mac"] = item.Dst_vtep_mac
		in["encap_type"] = item.Encap_type
		in["vtep_id"] = item.Vtep_id
		in["vtep_vnp_id"] = item.Vtep_vnp_id
		in["lifname"] = item.Lifname
		in["partname"] = item.Partname
		in["vlan"] = item.Vlan
		in["is_static"] = item.Is_static
		in["age"] = item.Age
		result = append(result, in)
	}
	return result
}

func getObjectOverlayTunnelDebugEncapTableOperOper(d []interface{}) edpt.OverlayTunnelDebugEncapTableOperOper {

	count1 := len(d)
	var ret edpt.OverlayTunnelDebugEncapTableOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vtep = getSliceOverlayTunnelDebugEncapTableOperOperVtep(in["vtep"].([]interface{}))
		ret.Tot_ip_encap = in["tot_ip_encap"].(int)
		ret.Tot_ipv6_encap = in["tot_ipv6_encap"].(int)
	}
	return ret
}

func getSliceOverlayTunnelDebugEncapTableOperOperVtep(d []interface{}) []edpt.OverlayTunnelDebugEncapTableOperOperVtep {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelDebugEncapTableOperOperVtep, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelDebugEncapTableOperOperVtep
		oi.Index = in["index"].(int)
		oi.Hindex = in["hindex"].(int)
		oi.Src_vtep_ip = in["src_vtep_ip"].(string)
		oi.Dst_vtep_ip = in["dst_vtep_ip"].(string)
		oi.Src_vtep_ipv6 = in["src_vtep_ipv6"].(string)
		oi.Dst_vtep_ipv6 = in["dst_vtep_ipv6"].(string)
		oi.Is_ipv6 = in["is_ipv6"].(int)
		oi.Dst_vtep_mac = in["dst_vtep_mac"].(string)
		oi.Encap_type = in["encap_type"].(string)
		oi.Vtep_id = in["vtep_id"].(int)
		oi.Vtep_vnp_id = in["vtep_vnp_id"].(int)
		oi.Lifname = in["lifname"].(string)
		oi.Partname = in["partname"].(string)
		oi.Vlan = in["vlan"].(int)
		oi.Is_static = in["is_static"].(int)
		oi.Age = in["age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointOverlayTunnelDebugEncapTableOper(d *schema.ResourceData) edpt.OverlayTunnelDebugEncapTableOper {
	var ret edpt.OverlayTunnelDebugEncapTableOper

	ret.Oper = getObjectOverlayTunnelDebugEncapTableOperOper(d.Get("oper").([]interface{}))
	return ret
}
