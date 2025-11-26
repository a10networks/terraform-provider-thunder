package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_oper`: Operational Status for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperRead,

		Schema: map[string]*schema.Schema{
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current_connections": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_connections_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"connection_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_connection_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_connection_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"connection_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_packet_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_packet_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"packet_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_kbit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"kbit_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_frag_packet_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_frag_packet_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"frag_packet_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"debug_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper := setObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper(ret edpt.DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_connections":        ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.CurrentConnections,
			"is_connections_exceed":      ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.IsConnectionsExceed,
			"connection_limit":           ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.ConnectionLimit,
			"current_connection_rate":    ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.CurrentConnectionRate,
			"is_connection_rate_exceed":  ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.IsConnectionRateExceed,
			"connection_rate_limit":      ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.ConnectionRateLimit,
			"current_packet_rate":        ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.CurrentPacketRate,
			"is_packet_rate_exceed":      ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.IsPacketRateExceed,
			"packet_rate_limit":          ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.PacketRateLimit,
			"current_kbit_rate":          ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.CurrentKbitRate,
			"is_kbit_rate_exceed":        ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.IsKbitRateExceed,
			"kbit_rate_limit":            ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.KbitRateLimit,
			"current_frag_packet_rate":   ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.CurrentFragPacketRate,
			"is_frag_packet_rate_exceed": ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.IsFragPacketRateExceed,
			"frag_packet_rate_limit":     ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.FragPacketRateLimit,
			"debug_str":                  ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper.Oper.DebugStr,
		},
	}
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrentConnections = in["current_connections"].(int)
		ret.IsConnectionsExceed = in["is_connections_exceed"].(int)
		ret.ConnectionLimit = in["connection_limit"].(int)
		ret.CurrentConnectionRate = in["current_connection_rate"].(int)
		ret.IsConnectionRateExceed = in["is_connection_rate_exceed"].(int)
		ret.ConnectionRateLimit = in["connection_rate_limit"].(int)
		ret.CurrentPacketRate = in["current_packet_rate"].(int)
		ret.IsPacketRateExceed = in["is_packet_rate_exceed"].(int)
		ret.PacketRateLimit = in["packet_rate_limit"].(int)
		ret.CurrentKbitRate = in["current_kbit_rate"].(int)
		ret.IsKbitRateExceed = in["is_kbit_rate_exceed"].(int)
		ret.KbitRateLimit = in["kbit_rate_limit"].(int)
		ret.CurrentFragPacketRate = in["current_frag_packet_rate"].(int)
		ret.IsFragPacketRateExceed = in["is_frag_packet_rate_exceed"].(int)
		ret.FragPacketRateLimit = in["frag_packet_rate_limit"].(int)
		ret.DebugStr = in["debug_str"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper {
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOper

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Oper = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)
	return ret
}
