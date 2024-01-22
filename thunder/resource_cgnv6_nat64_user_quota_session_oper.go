package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64UserQuotaSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat64_user_quota_session_oper`: Operational Status for the object user-quota-session\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat64UserQuotaSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefix_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"icmp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lid_number": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"flags": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"session_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"prefix_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"all_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"shared_partition": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pool_shared": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_by_tcp_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_by_udp_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_by_icmp_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_by_all_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_sort_by_cli": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"display_debug": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nat64UserQuotaSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64UserQuotaSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64UserQuotaSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat64UserQuotaSessionOperOper := setObjectCgnv6Nat64UserQuotaSessionOperOper(res)
		d.Set("oper", Cgnv6Nat64UserQuotaSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat64UserQuotaSessionOperOper(ret edpt.DataCgnv6Nat64UserQuotaSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":      setSliceCgnv6Nat64UserQuotaSessionOperOperSessionList(ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.SessionList),
			"session_count":     ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.SessionCount,
			"prefix_filter":     ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.PrefixFilter,
			"inside_addr_v6":    ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.InsideAddrV6,
			"all_partitions":    ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.AllPartitions,
			"shared_partition":  ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.SharedPartition,
			"partition_name":    ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.PartitionName,
			"nat_pool_name":     ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.NatPoolName,
			"pool_shared":       ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.PoolShared,
			"top_count":         ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopCount,
			"top_by_tcp_usage":  ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopByTcpUsage,
			"top_by_udp_usage":  ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopByUdpUsage,
			"top_by_icmp_usage": ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopByIcmpUsage,
			"top_by_all_usage":  ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopByAllUsage,
			"top_sort_by_cli":   ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.TopSortByCli,
			"display_debug":     ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.DisplayDebug,
			"nat_addr":          ret.DtCgnv6Nat64UserQuotaSessionOper.Oper.NatAddr,
		},
	}
}

func setSliceCgnv6Nat64UserQuotaSessionOperOperSessionList(d []edpt.Cgnv6Nat64UserQuotaSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_address"] = item.InsideAddress
		in["prefix_len"] = item.PrefixLen
		in["nat_address"] = item.NatAddress
		in["icmp_quota"] = item.IcmpQuota
		in["udp_quota"] = item.UdpQuota
		in["tcp_quota"] = item.TcpQuota
		in["session_count"] = item.SessionCount
		in["nat_pool_name"] = item.NatPoolName
		in["lid_number"] = item.LidNumber
		in["flags"] = item.Flags
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Nat64UserQuotaSessionOperOper(d []interface{}) edpt.Cgnv6Nat64UserQuotaSessionOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Nat64UserQuotaSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceCgnv6Nat64UserQuotaSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.SessionCount = in["session_count"].(int)
		ret.PrefixFilter = in["prefix_filter"].(string)
		ret.InsideAddrV6 = in["inside_addr_v6"].(string)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.NatPoolName = in["nat_pool_name"].(string)
		ret.PoolShared = in["pool_shared"].(int)
		ret.TopCount = in["top_count"].(int)
		ret.TopByTcpUsage = in["top_by_tcp_usage"].(int)
		ret.TopByUdpUsage = in["top_by_udp_usage"].(int)
		ret.TopByIcmpUsage = in["top_by_icmp_usage"].(int)
		ret.TopByAllUsage = in["top_by_all_usage"].(int)
		ret.TopSortByCli = in["top_sort_by_cli"].(int)
		ret.DisplayDebug = in["display_debug"].(string)
		ret.NatAddr = in["nat_addr"].(string)
	}
	return ret
}

func getSliceCgnv6Nat64UserQuotaSessionOperOperSessionList(d []interface{}) []edpt.Cgnv6Nat64UserQuotaSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Nat64UserQuotaSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Nat64UserQuotaSessionOperOperSessionList
		oi.InsideAddress = in["inside_address"].(string)
		oi.PrefixLen = in["prefix_len"].(int)
		oi.NatAddress = in["nat_address"].(string)
		oi.IcmpQuota = in["icmp_quota"].(int)
		oi.UdpQuota = in["udp_quota"].(int)
		oi.TcpQuota = in["tcp_quota"].(int)
		oi.SessionCount = in["session_count"].(int)
		oi.NatPoolName = in["nat_pool_name"].(string)
		oi.LidNumber = in["lid_number"].(int)
		oi.Flags = in["flags"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Nat64UserQuotaSessionOper(d *schema.ResourceData) edpt.Cgnv6Nat64UserQuotaSessionOper {
	var ret edpt.Cgnv6Nat64UserQuotaSessionOper

	ret.Oper = getObjectCgnv6Nat64UserQuotaSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
