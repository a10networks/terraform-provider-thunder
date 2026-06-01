package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnSystemStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_system_status_oper`: Operational Status for the object system-status\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnSystemStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lsn_cps": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_sessions_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_sessions_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp_sessions_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp_sessions_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tcp_nat_ports_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tcp_nat_ports_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"udp_nat_ports_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"udp_nat_ports_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_entries_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_entries_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnSystemStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnSystemStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnSystemStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnSystemStatusOperOper := setObjectCgnv6LsnSystemStatusOperOper(res)
		d.Set("oper", Cgnv6LsnSystemStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnSystemStatusOperOper(ret edpt.DataCgnv6LsnSystemStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"lsn_cps":             ret.DtCgnv6LsnSystemStatusOper.Oper.LsnCps,
			"data_sessions_used":  ret.DtCgnv6LsnSystemStatusOper.Oper.DataSessionsUsed,
			"data_sessions_free":  ret.DtCgnv6LsnSystemStatusOper.Oper.DataSessionsFree,
			"smp_sessions_used":   ret.DtCgnv6LsnSystemStatusOper.Oper.SmpSessionsUsed,
			"smp_sessions_free":   ret.DtCgnv6LsnSystemStatusOper.Oper.SmpSessionsFree,
			"tcp_nat_ports_used":  ret.DtCgnv6LsnSystemStatusOper.Oper.TcpNatPortsUsed,
			"tcp_nat_ports_free":  ret.DtCgnv6LsnSystemStatusOper.Oper.TcpNatPortsFree,
			"udp_nat_ports_used":  ret.DtCgnv6LsnSystemStatusOper.Oper.UdpNatPortsUsed,
			"udp_nat_ports_free":  ret.DtCgnv6LsnSystemStatusOper.Oper.UdpNatPortsFree,
			"radius_entries_used": ret.DtCgnv6LsnSystemStatusOper.Oper.RadiusEntriesUsed,
			"radius_entries_free": ret.DtCgnv6LsnSystemStatusOper.Oper.RadiusEntriesFree,
		},
	}
}

func getObjectCgnv6LsnSystemStatusOperOper(d []interface{}) edpt.Cgnv6LsnSystemStatusOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnSystemStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LsnCps = in["lsn_cps"].(int)
		ret.DataSessionsUsed = in["data_sessions_used"].(int)
		ret.DataSessionsFree = in["data_sessions_free"].(int)
		ret.SmpSessionsUsed = in["smp_sessions_used"].(int)
		ret.SmpSessionsFree = in["smp_sessions_free"].(int)
		ret.TcpNatPortsUsed = in["tcp_nat_ports_used"].(int)
		ret.TcpNatPortsFree = in["tcp_nat_ports_free"].(int)
		ret.UdpNatPortsUsed = in["udp_nat_ports_used"].(int)
		ret.UdpNatPortsFree = in["udp_nat_ports_free"].(int)
		ret.RadiusEntriesUsed = in["radius_entries_used"].(int)
		ret.RadiusEntriesFree = in["radius_entries_free"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnSystemStatusOper(d *schema.ResourceData) edpt.Cgnv6LsnSystemStatusOper {
	var ret edpt.Cgnv6LsnSystemStatusOper

	ret.Oper = getObjectCgnv6LsnSystemStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
