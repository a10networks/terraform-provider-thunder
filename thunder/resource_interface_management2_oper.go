package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagement2Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_management2_oper`: Operational Status for the object management2\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceManagement2OperRead,

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

func resourceInterfaceManagement2OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagement2OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement2Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceManagement2OperOper := setObjectInterfaceManagement2OperOper(res)
		d.Set("oper", InterfaceManagement2OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceManagement2OperOper(ret edpt.DataInterfaceManagement2Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interface":              ret.DtInterfaceManagement2Oper.Oper.Interface,
			"state":                  ret.DtInterfaceManagement2Oper.Oper.State,
			"line_protocol":          ret.DtInterfaceManagement2Oper.Oper.Line_protocol,
			"link_type":              ret.DtInterfaceManagement2Oper.Oper.Link_type,
			"mac":                    ret.DtInterfaceManagement2Oper.Oper.Mac,
			"ipv4_addr":              ret.DtInterfaceManagement2Oper.Oper.Ipv4Addr,
			"ipv4_mask":              ret.DtInterfaceManagement2Oper.Oper.Ipv4Mask,
			"ipv4_default_gateway":   ret.DtInterfaceManagement2Oper.Oper.Ipv4DefaultGateway,
			"ipv6_addr":              ret.DtInterfaceManagement2Oper.Oper.Ipv6Addr,
			"ipv6_prefix":            ret.DtInterfaceManagement2Oper.Oper.Ipv6Prefix,
			"ipv6_link_local":        ret.DtInterfaceManagement2Oper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix": ret.DtInterfaceManagement2Oper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_default_gateway":   ret.DtInterfaceManagement2Oper.Oper.Ipv6DefaultGateway,
			"speed":                  ret.DtInterfaceManagement2Oper.Oper.Speed,
			"duplexity":              ret.DtInterfaceManagement2Oper.Oper.Duplexity,
			"mtu":                    ret.DtInterfaceManagement2Oper.Oper.Mtu,
			"flow_control":           ret.DtInterfaceManagement2Oper.Oper.Flow_control,
			"ipv4_acl":               ret.DtInterfaceManagement2Oper.Oper.Ipv4_acl,
			"ipv6_acl":               ret.DtInterfaceManagement2Oper.Oper.Ipv6_acl,
			"dhcp_enabled":           ret.DtInterfaceManagement2Oper.Oper.Dhcp_enabled,
		},
	}
}

func getObjectInterfaceManagement2OperOper(d []interface{}) edpt.InterfaceManagement2OperOper {

	count1 := len(d)
	var ret edpt.InterfaceManagement2OperOper
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

func dataToEndpointInterfaceManagement2Oper(d *schema.ResourceData) edpt.InterfaceManagement2Oper {
	var ret edpt.InterfaceManagement2Oper

	ret.Oper = getObjectInterfaceManagement2OperOper(d.Get("oper").([]interface{}))
	return ret
}
