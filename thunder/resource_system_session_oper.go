package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_session_oper`: Operational Status for the object session\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"source_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"source_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v6_source_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"v6_dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"uie": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_source_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_dest_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_source_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_dest_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_source_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_ipv6_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_v6_source_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_v6_dest_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"persist_v6_source_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"persist_v6_dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"helper_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_ipv4": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sip_source_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sip_dest_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sip_source_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sip_dest_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sip_source_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sip_dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemSessionOperOper := setObjectSystemSessionOperOper(res)
		d.Set("oper", SystemSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemSessionOperOper(ret edpt.DataSystemSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all":                    ret.DtSystemSessionOper.Oper.All,
			"filter_name":            ret.DtSystemSessionOper.Oper.FilterName,
			"ipv4":                   ret.DtSystemSessionOper.Oper.Ipv4,
			"source_v4_addr":         ret.DtSystemSessionOper.Oper.SourceV4Addr,
			"dest_v4_addr":           ret.DtSystemSessionOper.Oper.DestV4Addr,
			"source_port":            ret.DtSystemSessionOper.Oper.SourcePort,
			"dest_port":              ret.DtSystemSessionOper.Oper.DestPort,
			"ipv6":                   ret.DtSystemSessionOper.Oper.Ipv6,
			"source_v6_addr":         ret.DtSystemSessionOper.Oper.SourceV6Addr,
			"dest_v6_addr":           ret.DtSystemSessionOper.Oper.DestV6Addr,
			"v6_source_port":         ret.DtSystemSessionOper.Oper.V6SourcePort,
			"v6_dest_port":           ret.DtSystemSessionOper.Oper.V6DestPort,
			"persist":                ret.DtSystemSessionOper.Oper.Persist,
			"uie":                    ret.DtSystemSessionOper.Oper.Uie,
			"persist_type":           ret.DtSystemSessionOper.Oper.PersistType,
			"persist_source_v4_addr": ret.DtSystemSessionOper.Oper.PersistSourceV4Addr,
			"persist_dest_v4_addr":   ret.DtSystemSessionOper.Oper.PersistDestV4Addr,
			"persist_source_v6_addr": ret.DtSystemSessionOper.Oper.PersistSourceV6Addr,
			"persist_dest_v6_addr":   ret.DtSystemSessionOper.Oper.PersistDestV6Addr,
			"persist_source_port":    ret.DtSystemSessionOper.Oper.PersistSourcePort,
			"persist_dest_port":      ret.DtSystemSessionOper.Oper.PersistDestPort,
			"persist_ipv6":           ret.DtSystemSessionOper.Oper.PersistIpv6,
			"persist_ipv6_type":      ret.DtSystemSessionOper.Oper.PersistIpv6Type,
			"persist_v6_source_addr": ret.DtSystemSessionOper.Oper.PersistV6SourceAddr,
			"persist_v6_dest_addr":   ret.DtSystemSessionOper.Oper.PersistV6DestAddr,
			"persist_v6_source_port": ret.DtSystemSessionOper.Oper.PersistV6SourcePort,
			"persist_v6_dest_port":   ret.DtSystemSessionOper.Oper.PersistV6DestPort,
			"fw":                     ret.DtSystemSessionOper.Oper.Fw,
			"helper_sessions":        ret.DtSystemSessionOper.Oper.HelperSessions,
			"fw_ipv4":                ret.DtSystemSessionOper.Oper.FwIpv4,
			"fw_ipv6":                ret.DtSystemSessionOper.Oper.FwIpv6,
			"sip":                    ret.DtSystemSessionOper.Oper.Sip,
			"sip_source_v4_addr":     ret.DtSystemSessionOper.Oper.SipSourceV4Addr,
			"sip_dest_v4_addr":       ret.DtSystemSessionOper.Oper.SipDestV4Addr,
			"sip_source_v6_addr":     ret.DtSystemSessionOper.Oper.SipSourceV6Addr,
			"sip_dest_v6_addr":       ret.DtSystemSessionOper.Oper.SipDestV6Addr,
			"sip_source_port":        ret.DtSystemSessionOper.Oper.SipSourcePort,
			"sip_dest_port":          ret.DtSystemSessionOper.Oper.SipDestPort,
		},
	}
}

func getObjectSystemSessionOperOper(d []interface{}) edpt.SystemSessionOperOper {

	count1 := len(d)
	var ret edpt.SystemSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.FilterName = in["filter_name"].(string)
		ret.Ipv4 = in["ipv4"].(int)
		ret.SourceV4Addr = in["source_v4_addr"].(string)
		ret.DestV4Addr = in["dest_v4_addr"].(string)
		ret.SourcePort = in["source_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.Ipv6 = in["ipv6"].(int)
		ret.SourceV6Addr = in["source_v6_addr"].(string)
		ret.DestV6Addr = in["dest_v6_addr"].(string)
		ret.V6SourcePort = in["v6_source_port"].(int)
		ret.V6DestPort = in["v6_dest_port"].(int)
		ret.Persist = in["persist"].(int)
		ret.Uie = in["uie"].(int)
		ret.PersistType = in["persist_type"].(string)
		ret.PersistSourceV4Addr = in["persist_source_v4_addr"].(string)
		ret.PersistDestV4Addr = in["persist_dest_v4_addr"].(string)
		ret.PersistSourceV6Addr = in["persist_source_v6_addr"].(string)
		ret.PersistDestV6Addr = in["persist_dest_v6_addr"].(string)
		ret.PersistSourcePort = in["persist_source_port"].(int)
		ret.PersistDestPort = in["persist_dest_port"].(int)
		ret.PersistIpv6 = in["persist_ipv6"].(int)
		ret.PersistIpv6Type = in["persist_ipv6_type"].(string)
		ret.PersistV6SourceAddr = in["persist_v6_source_addr"].(string)
		ret.PersistV6DestAddr = in["persist_v6_dest_addr"].(string)
		ret.PersistV6SourcePort = in["persist_v6_source_port"].(int)
		ret.PersistV6DestPort = in["persist_v6_dest_port"].(int)
		ret.Fw = in["fw"].(int)
		ret.HelperSessions = in["helper_sessions"].(int)
		ret.FwIpv4 = in["fw_ipv4"].(int)
		ret.FwIpv6 = in["fw_ipv6"].(int)
		ret.Sip = in["sip"].(int)
		ret.SipSourceV4Addr = in["sip_source_v4_addr"].(string)
		ret.SipDestV4Addr = in["sip_dest_v4_addr"].(string)
		ret.SipSourceV6Addr = in["sip_source_v6_addr"].(string)
		ret.SipDestV6Addr = in["sip_dest_v6_addr"].(string)
		ret.SipSourcePort = in["sip_source_port"].(int)
		ret.SipDestPort = in["sip_dest_port"].(int)
	}
	return ret
}

func dataToEndpointSystemSessionOper(d *schema.ResourceData) edpt.SystemSessionOper {
	var ret edpt.SystemSessionOper

	ret.Oper = getObjectSystemSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
