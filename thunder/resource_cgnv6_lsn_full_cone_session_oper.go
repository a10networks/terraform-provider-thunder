package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnFullConeSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_full_cone_session_oper`: Operational Status for the object full-cone-session\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnFullConeSessionOperRead,

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
						"inside_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_end": {
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

func resourceCgnv6LsnFullConeSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnFullConeSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnFullConeSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnFullConeSessionOperOper := setObjectCgnv6LsnFullConeSessionOperOper(res)
		d.Set("oper", Cgnv6LsnFullConeSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnFullConeSessionOperOper(ret edpt.DataCgnv6LsnFullConeSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":        setSliceCgnv6LsnFullConeSessionOperOperSessionList(ret.DtCgnv6LsnFullConeSessionOper.Oper.SessionList),
			"session_count":       ret.DtCgnv6LsnFullConeSessionOper.Oper.SessionCount,
			"total_session_count": ret.DtCgnv6LsnFullConeSessionOper.Oper.TotalSessionCount,
			"all_partitions":      ret.DtCgnv6LsnFullConeSessionOper.Oper.AllPartitions,
			"shared_partition":    ret.DtCgnv6LsnFullConeSessionOper.Oper.SharedPartition,
			"partition_name":      ret.DtCgnv6LsnFullConeSessionOper.Oper.PartitionName,
			"inside_addr":         ret.DtCgnv6LsnFullConeSessionOper.Oper.InsideAddr,
			"inside_addr_start":   ret.DtCgnv6LsnFullConeSessionOper.Oper.InsideAddrStart,
			"inside_addr_end":     ret.DtCgnv6LsnFullConeSessionOper.Oper.InsideAddrEnd,
			"nat_addr":            ret.DtCgnv6LsnFullConeSessionOper.Oper.NatAddr,
			"nat_addr_start":      ret.DtCgnv6LsnFullConeSessionOper.Oper.NatAddrStart,
			"nat_addr_end":        ret.DtCgnv6LsnFullConeSessionOper.Oper.NatAddrEnd,
			"inside_port":         ret.DtCgnv6LsnFullConeSessionOper.Oper.InsidePort,
			"nat_port":            ret.DtCgnv6LsnFullConeSessionOper.Oper.NatPort,
			"nat_pool_name":       ret.DtCgnv6LsnFullConeSessionOper.Oper.NatPoolName,
			"pool_shared":         ret.DtCgnv6LsnFullConeSessionOper.Oper.PoolShared,
			"pcp":                 ret.DtCgnv6LsnFullConeSessionOper.Oper.Pcp,
			"graceful":            ret.DtCgnv6LsnFullConeSessionOper.Oper.Graceful,
			"debug_session":       ret.DtCgnv6LsnFullConeSessionOper.Oper.DebugSession,
		},
	}
}

func setSliceCgnv6LsnFullConeSessionOperOperSessionList(d []edpt.Cgnv6LsnFullConeSessionOperOperSessionList) []map[string]interface{} {
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

func getObjectCgnv6LsnFullConeSessionOperOper(d []interface{}) edpt.Cgnv6LsnFullConeSessionOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnFullConeSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceCgnv6LsnFullConeSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.SessionCount = in["session_count"].(int)
		ret.TotalSessionCount = in["total_session_count"].(int)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddr = in["inside_addr"].(string)
		ret.InsideAddrStart = in["inside_addr_start"].(string)
		ret.InsideAddrEnd = in["inside_addr_end"].(string)
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

func getSliceCgnv6LsnFullConeSessionOperOperSessionList(d []interface{}) []edpt.Cgnv6LsnFullConeSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnFullConeSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnFullConeSessionOperOperSessionList
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

func dataToEndpointCgnv6LsnFullConeSessionOper(d *schema.ResourceData) edpt.Cgnv6LsnFullConeSessionOper {
	var ret edpt.Cgnv6LsnFullConeSessionOper

	ret.Oper = getObjectCgnv6LsnFullConeSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
