package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_src_based_policy_policy_class_list_oper`: Operational Status for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperRead,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper := setObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_connections":        ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.CurrentConnections,
			"is_connections_exceed":      ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.IsConnectionsExceed,
			"connection_limit":           ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.ConnectionLimit,
			"current_connection_rate":    ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.CurrentConnectionRate,
			"is_connection_rate_exceed":  ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.IsConnectionRateExceed,
			"connection_rate_limit":      ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.ConnectionRateLimit,
			"current_packet_rate":        ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.CurrentPacketRate,
			"is_packet_rate_exceed":      ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.IsPacketRateExceed,
			"packet_rate_limit":          ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.PacketRateLimit,
			"current_kbit_rate":          ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.CurrentKbitRate,
			"is_kbit_rate_exceed":        ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.IsKbitRateExceed,
			"kbit_rate_limit":            ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.KbitRateLimit,
			"current_frag_packet_rate":   ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.CurrentFragPacketRate,
			"is_frag_packet_rate_exceed": ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.IsFragPacketRateExceed,
			"frag_packet_rate_limit":     ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.FragPacketRateLimit,
			"debug_str":                  ret.DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper.Oper.DebugStr,
		},
	}
}

func getObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper
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

func dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOper

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	return ret
}
