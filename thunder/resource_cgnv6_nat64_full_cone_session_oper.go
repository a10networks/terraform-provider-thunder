package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64FullConeSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat64_full_cone_session_oper`: Operational Status for the object full-cone-session\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat64FullConeSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"outbound": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inbound": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeString, Optional: true, Description: "",
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
						"total_session_count": {
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
						"inside_addr_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_v6_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_v6_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pool_shared": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pcp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"graceful": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"debug_session": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nat64FullConeSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FullConeSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FullConeSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat64FullConeSessionOperOper := setObjectCgnv6Nat64FullConeSessionOperOper(res)
		d.Set("oper", Cgnv6Nat64FullConeSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat64FullConeSessionOperOper(ret edpt.DataCgnv6Nat64FullConeSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":         setSliceCgnv6Nat64FullConeSessionOperOperSessionList(ret.DtCgnv6Nat64FullConeSessionOper.Oper.SessionList),
			"session_count":        ret.DtCgnv6Nat64FullConeSessionOper.Oper.SessionCount,
			"total_session_count":  ret.DtCgnv6Nat64FullConeSessionOper.Oper.TotalSessionCount,
			"all_partitions":       ret.DtCgnv6Nat64FullConeSessionOper.Oper.AllPartitions,
			"shared_partition":     ret.DtCgnv6Nat64FullConeSessionOper.Oper.SharedPartition,
			"partition_name":       ret.DtCgnv6Nat64FullConeSessionOper.Oper.PartitionName,
			"inside_addr_v6":       ret.DtCgnv6Nat64FullConeSessionOper.Oper.InsideAddrV6,
			"inside_addr_v6_start": ret.DtCgnv6Nat64FullConeSessionOper.Oper.InsideAddrV6Start,
			"inside_addr_v6_end":   ret.DtCgnv6Nat64FullConeSessionOper.Oper.InsideAddrV6End,
			"nat_addr":             ret.DtCgnv6Nat64FullConeSessionOper.Oper.NatAddr,
			"nat_addr_start":       ret.DtCgnv6Nat64FullConeSessionOper.Oper.NatAddrStart,
			"nat_addr_end":         ret.DtCgnv6Nat64FullConeSessionOper.Oper.NatAddrEnd,
			"inside_port":          ret.DtCgnv6Nat64FullConeSessionOper.Oper.InsidePort,
			"nat_port":             ret.DtCgnv6Nat64FullConeSessionOper.Oper.NatPort,
			"nat_pool_name":        ret.DtCgnv6Nat64FullConeSessionOper.Oper.NatPoolName,
			"pool_shared":          ret.DtCgnv6Nat64FullConeSessionOper.Oper.PoolShared,
			"pcp":                  ret.DtCgnv6Nat64FullConeSessionOper.Oper.Pcp,
			"graceful":             ret.DtCgnv6Nat64FullConeSessionOper.Oper.Graceful,
			"debug_session":        ret.DtCgnv6Nat64FullConeSessionOper.Oper.DebugSession,
		},
	}
}

func setSliceCgnv6Nat64FullConeSessionOperOperSessionList(d []edpt.Cgnv6Nat64FullConeSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["inside_address"] = item.InsideAddress
		in["inside_port"] = item.InsidePort
		in["nat_address"] = item.NatAddress
		in["nat_port"] = item.NatPort
		in["outbound"] = item.Outbound
		in["inbound"] = item.Inbound
		in["nat_pool_name"] = item.NatPoolName
		in["cpu"] = item.Cpu
		in["age"] = item.Age
		in["flags"] = item.Flags
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6Nat64FullConeSessionOperOper(d []interface{}) edpt.Cgnv6Nat64FullConeSessionOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6Nat64FullConeSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceCgnv6Nat64FullConeSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.SessionCount = in["session_count"].(int)
		ret.TotalSessionCount = in["total_session_count"].(int)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddrV6 = in["inside_addr_v6"].(string)
		ret.InsideAddrV6Start = in["inside_addr_v6_start"].(string)
		ret.InsideAddrV6End = in["inside_addr_v6_end"].(string)
		ret.NatAddr = in["nat_addr"].(string)
		ret.NatAddrStart = in["nat_addr_start"].(string)
		ret.NatAddrEnd = in["nat_addr_end"].(string)
		ret.InsidePort = in["inside_port"].(int)
		ret.NatPort = in["nat_port"].(int)
		ret.NatPoolName = in["nat_pool_name"].(string)
		ret.PoolShared = in["pool_shared"].(int)
		ret.Pcp = in["pcp"].(int)
		ret.Graceful = in["graceful"].(int)
		ret.DebugSession = in["debug_session"].(int)
	}
	return ret
}

func getSliceCgnv6Nat64FullConeSessionOperOperSessionList(d []interface{}) []edpt.Cgnv6Nat64FullConeSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Nat64FullConeSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Nat64FullConeSessionOperOperSessionList
		oi.Protocol = in["protocol"].(string)
		oi.InsideAddress = in["inside_address"].(string)
		oi.InsidePort = in["inside_port"].(int)
		oi.NatAddress = in["nat_address"].(string)
		oi.NatPort = in["nat_port"].(int)
		oi.Outbound = in["outbound"].(int)
		oi.Inbound = in["inbound"].(int)
		oi.NatPoolName = in["nat_pool_name"].(string)
		oi.Cpu = in["cpu"].(int)
		oi.Age = in["age"].(string)
		oi.Flags = in["flags"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Nat64FullConeSessionOper(d *schema.ResourceData) edpt.Cgnv6Nat64FullConeSessionOper {
	var ret edpt.Cgnv6Nat64FullConeSessionOper

	ret.Oper = getObjectCgnv6Nat64FullConeSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
