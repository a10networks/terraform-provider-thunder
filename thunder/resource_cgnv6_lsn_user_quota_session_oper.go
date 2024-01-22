package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnUserQuotaSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_user_quota_session_oper`: Operational Status for the object user-quota-session\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnUserQuotaSessionOperRead,

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
						"all_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"shared_partition": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr": {
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
						"nat_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnUserQuotaSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnUserQuotaSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnUserQuotaSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnUserQuotaSessionOperOper := setObjectCgnv6LsnUserQuotaSessionOperOper(res)
		d.Set("oper", Cgnv6LsnUserQuotaSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnUserQuotaSessionOperOper(ret edpt.DataCgnv6LsnUserQuotaSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":      setSliceCgnv6LsnUserQuotaSessionOperOperSessionList(ret.DtCgnv6LsnUserQuotaSessionOper.Oper.SessionList),
			"session_count":     ret.DtCgnv6LsnUserQuotaSessionOper.Oper.SessionCount,
			"all_partitions":    ret.DtCgnv6LsnUserQuotaSessionOper.Oper.AllPartitions,
			"shared_partition":  ret.DtCgnv6LsnUserQuotaSessionOper.Oper.SharedPartition,
			"partition_name":    ret.DtCgnv6LsnUserQuotaSessionOper.Oper.PartitionName,
			"inside_addr":       ret.DtCgnv6LsnUserQuotaSessionOper.Oper.InsideAddr,
			"nat_pool_name":     ret.DtCgnv6LsnUserQuotaSessionOper.Oper.NatPoolName,
			"pool_shared":       ret.DtCgnv6LsnUserQuotaSessionOper.Oper.PoolShared,
			"top_count":         ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopCount,
			"top_by_tcp_usage":  ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopByTcpUsage,
			"top_by_udp_usage":  ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopByUdpUsage,
			"top_by_icmp_usage": ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopByIcmpUsage,
			"top_by_all_usage":  ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopByAllUsage,
			"top_sort_by_cli":   ret.DtCgnv6LsnUserQuotaSessionOper.Oper.TopSortByCli,
			"nat_addr":          ret.DtCgnv6LsnUserQuotaSessionOper.Oper.NatAddr,
		},
	}
}

func setSliceCgnv6LsnUserQuotaSessionOperOperSessionList(d []edpt.Cgnv6LsnUserQuotaSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_address"] = item.InsideAddress
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

func getObjectCgnv6LsnUserQuotaSessionOperOper(d []interface{}) edpt.Cgnv6LsnUserQuotaSessionOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnUserQuotaSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceCgnv6LsnUserQuotaSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.SessionCount = in["session_count"].(int)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddr = in["inside_addr"].(string)
		ret.NatPoolName = in["nat_pool_name"].(string)
		ret.PoolShared = in["pool_shared"].(int)
		ret.TopCount = in["top_count"].(int)
		ret.TopByTcpUsage = in["top_by_tcp_usage"].(int)
		ret.TopByUdpUsage = in["top_by_udp_usage"].(int)
		ret.TopByIcmpUsage = in["top_by_icmp_usage"].(int)
		ret.TopByAllUsage = in["top_by_all_usage"].(int)
		ret.TopSortByCli = in["top_sort_by_cli"].(int)
		ret.NatAddr = in["nat_addr"].(string)
	}
	return ret
}

func getSliceCgnv6LsnUserQuotaSessionOperOperSessionList(d []interface{}) []edpt.Cgnv6LsnUserQuotaSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnUserQuotaSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnUserQuotaSessionOperOperSessionList
		oi.InsideAddress = in["inside_address"].(string)
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

func dataToEndpointCgnv6LsnUserQuotaSessionOper(d *schema.ResourceData) edpt.Cgnv6LsnUserQuotaSessionOper {
	var ret edpt.Cgnv6LsnUserQuotaSessionOper

	ret.Oper = getObjectCgnv6LsnUserQuotaSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
