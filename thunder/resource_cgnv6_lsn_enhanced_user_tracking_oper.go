package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnEnhancedUserTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_enhanced_user_tracking_oper`: Operational Status for the object enhanced-user-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnEnhancedUserTrackingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tcp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icmp_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_peak": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_session_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"upl_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"upl_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwl_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dwl_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"user_count": {
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
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnEnhancedUserTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEnhancedUserTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEnhancedUserTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnEnhancedUserTrackingOperOper := setObjectCgnv6LsnEnhancedUserTrackingOperOper(res)
		d.Set("oper", Cgnv6LsnEnhancedUserTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnEnhancedUserTrackingOperOper(ret edpt.DataCgnv6LsnEnhancedUserTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"user_list":        setSliceCgnv6LsnEnhancedUserTrackingOperOperUserList(ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.UserList),
			"user_count":       ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.UserCount,
			"all_partitions":   ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.AllPartitions,
			"shared_partition": ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.SharedPartition,
			"partition_name":   ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.PartitionName,
			"inside_addr":      ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.InsideAddr,
			"nat_pool_name":    ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.NatPoolName,
			"pool_shared":      ret.DtCgnv6LsnEnhancedUserTrackingOper.Oper.PoolShared,
		},
	}
}

func setSliceCgnv6LsnEnhancedUserTrackingOperOperUserList(d []edpt.Cgnv6LsnEnhancedUserTrackingOperOperUserList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_address"] = item.InsideAddress
		in["nat_address"] = item.NatAddress
		in["tcp_quota"] = item.TcpQuota
		in["udp_quota"] = item.UdpQuota
		in["icmp_quota"] = item.IcmpQuota
		in["session_count"] = item.SessionCount
		in["tcp_peak"] = item.TcpPeak
		in["udp_peak"] = item.UdpPeak
		in["icmp_peak"] = item.IcmpPeak
		in["session_peak"] = item.SessionPeak
		in["total_session_count"] = item.TotalSessionCount
		in["upl_packets"] = item.UplPackets
		in["upl_bytes"] = item.UplBytes
		in["dwl_packets"] = item.DwlPackets
		in["dwl_bytes"] = item.DwlBytes
		in["nat_pool_name"] = item.NatPoolName
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6LsnEnhancedUserTrackingOperOper(d []interface{}) edpt.Cgnv6LsnEnhancedUserTrackingOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnEnhancedUserTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserList = getSliceCgnv6LsnEnhancedUserTrackingOperOperUserList(in["user_list"].([]interface{}))
		ret.UserCount = in["user_count"].(int)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddr = in["inside_addr"].(string)
		ret.NatPoolName = in["nat_pool_name"].(string)
		ret.PoolShared = in["pool_shared"].(int)
	}
	return ret
}

func getSliceCgnv6LsnEnhancedUserTrackingOperOperUserList(d []interface{}) []edpt.Cgnv6LsnEnhancedUserTrackingOperOperUserList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnEnhancedUserTrackingOperOperUserList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnEnhancedUserTrackingOperOperUserList
		oi.InsideAddress = in["inside_address"].(string)
		oi.NatAddress = in["nat_address"].(string)
		oi.TcpQuota = in["tcp_quota"].(int)
		oi.UdpQuota = in["udp_quota"].(int)
		oi.IcmpQuota = in["icmp_quota"].(int)
		oi.SessionCount = in["session_count"].(int)
		oi.TcpPeak = in["tcp_peak"].(int)
		oi.UdpPeak = in["udp_peak"].(int)
		oi.IcmpPeak = in["icmp_peak"].(int)
		oi.SessionPeak = in["session_peak"].(int)
		oi.TotalSessionCount = in["total_session_count"].(int)
		oi.UplPackets = in["upl_packets"].(int)
		oi.UplBytes = in["upl_bytes"].(int)
		oi.DwlPackets = in["dwl_packets"].(int)
		oi.DwlBytes = in["dwl_bytes"].(int)
		oi.NatPoolName = in["nat_pool_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnEnhancedUserTrackingOper(d *schema.ResourceData) edpt.Cgnv6LsnEnhancedUserTrackingOper {
	var ret edpt.Cgnv6LsnEnhancedUserTrackingOper

	ret.Oper = getObjectCgnv6LsnEnhancedUserTrackingOperOper(d.Get("oper").([]interface{}))
	return ret
}
