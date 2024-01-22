package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagementOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_management_oper`: Operational Status for the object management\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceManagementOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"state": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"line_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"link_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mac": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_mask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
						},
						"ipv4_default_gateway": {
							Type: schema.TypeString, Optional: true, Description: "IP gateway address",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_link_local": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_link_local_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_default_gateway": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"speed": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"duplexity": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mtu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"flow_control": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_acl": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_acl": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dhcp_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceManagementOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceManagementOperOper := setObjectInterfaceManagementOperOper(res)
		d.Set("oper", InterfaceManagementOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceManagementOperOper(ret edpt.DataInterfaceManagementOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interface":              ret.DtInterfaceManagementOper.Oper.Interface,
			"state":                  ret.DtInterfaceManagementOper.Oper.State,
			"line_protocol":          ret.DtInterfaceManagementOper.Oper.Line_protocol,
			"link_type":              ret.DtInterfaceManagementOper.Oper.Link_type,
			"mac":                    ret.DtInterfaceManagementOper.Oper.Mac,
			"ipv4_addr":              ret.DtInterfaceManagementOper.Oper.Ipv4Addr,
			"ipv4_mask":              ret.DtInterfaceManagementOper.Oper.Ipv4Mask,
			"ipv4_default_gateway":   ret.DtInterfaceManagementOper.Oper.Ipv4DefaultGateway,
			"ipv6_addr":              ret.DtInterfaceManagementOper.Oper.Ipv6Addr,
			"ipv6_prefix":            ret.DtInterfaceManagementOper.Oper.Ipv6Prefix,
			"ipv6_link_local":        ret.DtInterfaceManagementOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix": ret.DtInterfaceManagementOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_default_gateway":   ret.DtInterfaceManagementOper.Oper.Ipv6DefaultGateway,
			"speed":                  ret.DtInterfaceManagementOper.Oper.Speed,
			"duplexity":              ret.DtInterfaceManagementOper.Oper.Duplexity,
			"mtu":                    ret.DtInterfaceManagementOper.Oper.Mtu,
			"flow_control":           ret.DtInterfaceManagementOper.Oper.Flow_control,
			"ipv4_acl":               ret.DtInterfaceManagementOper.Oper.Ipv4_acl,
			"ipv6_acl":               ret.DtInterfaceManagementOper.Oper.Ipv6_acl,
			"dhcp_enabled":           ret.DtInterfaceManagementOper.Oper.Dhcp_enabled,
		},
	}
}

func getObjectInterfaceManagementOperOper(d []interface{}) edpt.InterfaceManagementOperOper {

	count1 := len(d)
	var ret edpt.InterfaceManagementOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interface = in["interface"].(string)
		ret.State = in["state"].(int)
		ret.Line_protocol = in["line_protocol"].(string)
		ret.Link_type = in["link_type"].(string)
		ret.Mac = in["mac"].(string)
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv4Mask = in["ipv4_mask"].(string)
		ret.Ipv4DefaultGateway = in["ipv4_default_gateway"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.Ipv6Prefix = in["ipv6_prefix"].(string)
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6DefaultGateway = in["ipv6_default_gateway"].(string)
		ret.Speed = in["speed"].(string)
		ret.Duplexity = in["duplexity"].(string)
		ret.Mtu = in["mtu"].(int)
		ret.Flow_control = in["flow_control"].(int)
		ret.Ipv4_acl = in["ipv4_acl"].(string)
		ret.Ipv6_acl = in["ipv6_acl"].(string)
		ret.Dhcp_enabled = in["dhcp_enabled"].(int)
	}
	return ret
}

func dataToEndpointInterfaceManagementOper(d *schema.ResourceData) edpt.InterfaceManagementOper {
	var ret edpt.InterfaceManagementOper

	ret.Oper = getObjectInterfaceManagementOperOper(d.Get("oper").([]interface{}))
	return ret
}
