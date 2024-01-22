package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccessListIpv4Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_access_list_ipv4_oper`: Operational Status for the object ipv4\n\n__PLACEHOLDER__",
		ReadContext: resourceAccessListIpv4OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mgmt_pkt_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"binding": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_pool_haid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_pool_group": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_pool_msl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sequence_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remark": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"icmp_type": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"icmp_code": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"svc_obj_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"geo_location_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_obj_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_host": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_host_mask": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_obj_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst_host": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst_host_mask": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dst_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"eth": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"vlan_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tcp_established": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dscp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ip_frag": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"log": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"log_transparent_sess_only": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"data_plane_hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceAccessListIpv4OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccessListIpv4OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccessListIpv4Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AccessListIpv4OperOper := setObjectAccessListIpv4OperOper(res)
		d.Set("oper", AccessListIpv4OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAccessListIpv4OperOper(ret edpt.DataAccessListIpv4Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"acl_list": setSliceAccessListIpv4OperOperAclList(ret.DtAccessListIpv4Oper.Oper.AclList),
		},
	}
}

func setSliceAccessListIpv4OperOperAclList(d []edpt.AccessListIpv4OperOperAclList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["name"] = item.Name
		in["mgmt_pkt_hit_count"] = item.MgmtPktHitCount
		in["binding"] = item.Binding
		in["nat_pool_name"] = item.NatPoolName
		in["nat_pool_haid"] = item.NatPoolHaid
		in["is_pool_group"] = item.IsPoolGroup
		in["nat_pool_msl"] = item.NatPoolMsl
		in["rule_list"] = setSliceAccessListIpv4OperOperAclListRuleList(item.RuleList)
		result = append(result, in)
	}
	return result
}

func setSliceAccessListIpv4OperOperAclListRuleList(d []edpt.AccessListIpv4OperOperAclListRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sequence_num"] = item.SequenceNum
		in["action"] = item.Action
		in["remark"] = item.Remark
		in["proto"] = item.Proto
		in["icmp_type"] = item.IcmpType
		in["icmp_code"] = item.IcmpCode
		in["svc_obj_id"] = item.SvcObjId
		in["geo_location_name"] = item.GeoLocationName
		in["src_obj_id"] = item.SrcObjId
		in["src_host"] = item.SrcHost
		in["src_host_mask"] = item.SrcHostMask
		in["src_port_start"] = item.SrcPortStart
		in["src_port_end"] = item.SrcPortEnd
		in["dst_obj_id"] = item.DstObjId
		in["dst_host"] = item.DstHost
		in["dst_host_mask"] = item.DstHostMask
		in["dst_port_start"] = item.DstPortStart
		in["dst_port_end"] = item.DstPortEnd
		in["eth"] = item.Eth
		in["trunk"] = item.Trunk
		in["vlan_id"] = item.VlanId
		in["tcp_established"] = item.TcpEstablished
		in["dscp"] = item.Dscp
		in["ip_frag"] = item.IpFrag
		in["log"] = item.Log
		in["log_transparent_sess_only"] = item.LogTransparentSessOnly
		in["data_plane_hits"] = item.DataPlaneHits
		result = append(result, in)
	}
	return result
}

func getObjectAccessListIpv4OperOper(d []interface{}) edpt.AccessListIpv4OperOper {

	count1 := len(d)
	var ret edpt.AccessListIpv4OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclList = getSliceAccessListIpv4OperOperAclList(in["acl_list"].([]interface{}))
	}
	return ret
}

func getSliceAccessListIpv4OperOperAclList(d []interface{}) []edpt.AccessListIpv4OperOperAclList {

	count1 := len(d)
	ret := make([]edpt.AccessListIpv4OperOperAclList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AccessListIpv4OperOperAclList
		oi.Id1 = in["id1"].(int)
		oi.Name = in["name"].(string)
		oi.MgmtPktHitCount = in["mgmt_pkt_hit_count"].(int)
		oi.Binding = in["binding"].(int)
		oi.NatPoolName = in["nat_pool_name"].(string)
		oi.NatPoolHaid = in["nat_pool_haid"].(int)
		oi.IsPoolGroup = in["is_pool_group"].(int)
		oi.NatPoolMsl = in["nat_pool_msl"].(int)
		oi.RuleList = getSliceAccessListIpv4OperOperAclListRuleList(in["rule_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAccessListIpv4OperOperAclListRuleList(d []interface{}) []edpt.AccessListIpv4OperOperAclListRuleList {

	count1 := len(d)
	ret := make([]edpt.AccessListIpv4OperOperAclListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AccessListIpv4OperOperAclListRuleList
		oi.SequenceNum = in["sequence_num"].(int)
		oi.Action = in["action"].(string)
		oi.Remark = in["remark"].(string)
		oi.Proto = in["proto"].(string)
		oi.IcmpType = in["icmp_type"].(int)
		oi.IcmpCode = in["icmp_code"].(int)
		oi.SvcObjId = in["svc_obj_id"].(string)
		oi.GeoLocationName = in["geo_location_name"].(string)
		oi.SrcObjId = in["src_obj_id"].(string)
		oi.SrcHost = in["src_host"].(string)
		oi.SrcHostMask = in["src_host_mask"].(string)
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		oi.DstObjId = in["dst_obj_id"].(string)
		oi.DstHost = in["dst_host"].(string)
		oi.DstHostMask = in["dst_host_mask"].(string)
		oi.DstPortStart = in["dst_port_start"].(int)
		oi.DstPortEnd = in["dst_port_end"].(int)
		oi.Eth = in["eth"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.VlanId = in["vlan_id"].(int)
		oi.TcpEstablished = in["tcp_established"].(int)
		oi.Dscp = in["dscp"].(int)
		oi.IpFrag = in["ip_frag"].(int)
		oi.Log = in["log"].(int)
		oi.LogTransparentSessOnly = in["log_transparent_sess_only"].(int)
		oi.DataPlaneHits = in["data_plane_hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAccessListIpv4Oper(d *schema.ResourceData) edpt.AccessListIpv4Oper {
	var ret edpt.AccessListIpv4Oper

	ret.Oper = getObjectAccessListIpv4OperOper(d.Get("oper").([]interface{}))
	return ret
}
