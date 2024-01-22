package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatFullConeSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_full_cone_session_oper`: Operational Status for the object full-cone-session\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatFullConeSessionOperRead,

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
									"inside_v6_address": {
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
									"eim": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"eif": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
						"nat44_total_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat64_total_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dslite_total_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"session_type": {
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
						"inside_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_end": {
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

func resourceCgnv6FixedNatFullConeSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatFullConeSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatFullConeSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatFullConeSessionOperOper := setObjectCgnv6FixedNatFullConeSessionOperOper(res)
		d.Set("oper", Cgnv6FixedNatFullConeSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatFullConeSessionOperOper(ret edpt.DataCgnv6FixedNatFullConeSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":               setSliceCgnv6FixedNatFullConeSessionOperOperSessionList(ret.DtCgnv6FixedNatFullConeSessionOper.Oper.SessionList),
			"session_count":              ret.DtCgnv6FixedNatFullConeSessionOper.Oper.SessionCount,
			"nat44_total_session_count":  ret.DtCgnv6FixedNatFullConeSessionOper.Oper.Nat44TotalSessionCount,
			"nat64_total_session_count":  ret.DtCgnv6FixedNatFullConeSessionOper.Oper.Nat64TotalSessionCount,
			"dslite_total_session_count": ret.DtCgnv6FixedNatFullConeSessionOper.Oper.DsliteTotalSessionCount,
			"session_type":               ret.DtCgnv6FixedNatFullConeSessionOper.Oper.SessionType,
			"all_partitions":             ret.DtCgnv6FixedNatFullConeSessionOper.Oper.AllPartitions,
			"shared_partition":           ret.DtCgnv6FixedNatFullConeSessionOper.Oper.SharedPartition,
			"partition_name":             ret.DtCgnv6FixedNatFullConeSessionOper.Oper.PartitionName,
			"inside_addr":                ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddr,
			"inside_addr_start":          ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddrStart,
			"inside_addr_end":            ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddrEnd,
			"inside_addr_v6":             ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddrV6,
			"inside_addr_v6_start":       ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddrV6Start,
			"inside_addr_v6_end":         ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsideAddrV6End,
			"nat_addr":                   ret.DtCgnv6FixedNatFullConeSessionOper.Oper.NatAddr,
			"nat_addr_start":             ret.DtCgnv6FixedNatFullConeSessionOper.Oper.NatAddrStart,
			"nat_addr_end":               ret.DtCgnv6FixedNatFullConeSessionOper.Oper.NatAddrEnd,
			"inside_port":                ret.DtCgnv6FixedNatFullConeSessionOper.Oper.InsidePort,
			"nat_port":                   ret.DtCgnv6FixedNatFullConeSessionOper.Oper.NatPort,
			"pcp":                        ret.DtCgnv6FixedNatFullConeSessionOper.Oper.Pcp,
			"graceful":                   ret.DtCgnv6FixedNatFullConeSessionOper.Oper.Graceful,
			"debug_session":              ret.DtCgnv6FixedNatFullConeSessionOper.Oper.DebugSession,
		},
	}
}

func setSliceCgnv6FixedNatFullConeSessionOperOperSessionList(d []edpt.Cgnv6FixedNatFullConeSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["inside_v6_address"] = item.InsideV6Address
		in["inside_address"] = item.InsideAddress
		in["inside_port"] = item.InsidePort
		in["nat_address"] = item.NatAddress
		in["nat_port"] = item.NatPort
		in["eim"] = item.Eim
		in["eif"] = item.Eif
		in["cpu"] = item.Cpu
		in["age"] = item.Age
		in["flags"] = item.Flags
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6FixedNatFullConeSessionOperOper(d []interface{}) edpt.Cgnv6FixedNatFullConeSessionOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatFullConeSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceCgnv6FixedNatFullConeSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.SessionCount = in["session_count"].(int)
		ret.Nat44TotalSessionCount = in["nat44_total_session_count"].(int)
		ret.Nat64TotalSessionCount = in["nat64_total_session_count"].(int)
		ret.DsliteTotalSessionCount = in["dslite_total_session_count"].(int)
		ret.SessionType = in["session_type"].(string)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddr = in["inside_addr"].(string)
		ret.InsideAddrStart = in["inside_addr_start"].(string)
		ret.InsideAddrEnd = in["inside_addr_end"].(string)
		ret.InsideAddrV6 = in["inside_addr_v6"].(string)
		ret.InsideAddrV6Start = in["inside_addr_v6_start"].(string)
		ret.InsideAddrV6End = in["inside_addr_v6_end"].(string)
		ret.NatAddr = in["nat_addr"].(string)
		ret.NatAddrStart = in["nat_addr_start"].(string)
		ret.NatAddrEnd = in["nat_addr_end"].(string)
		ret.InsidePort = in["inside_port"].(int)
		ret.NatPort = in["nat_port"].(int)
		ret.Pcp = in["pcp"].(int)
		ret.Graceful = in["graceful"].(int)
		ret.DebugSession = in["debug_session"].(int)
	}
	return ret
}

func getSliceCgnv6FixedNatFullConeSessionOperOperSessionList(d []interface{}) []edpt.Cgnv6FixedNatFullConeSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatFullConeSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatFullConeSessionOperOperSessionList
		oi.Protocol = in["protocol"].(string)
		oi.InsideV6Address = in["inside_v6_address"].(string)
		oi.InsideAddress = in["inside_address"].(string)
		oi.InsidePort = in["inside_port"].(int)
		oi.NatAddress = in["nat_address"].(string)
		oi.NatPort = in["nat_port"].(int)
		oi.Eim = in["eim"].(int)
		oi.Eif = in["eif"].(int)
		oi.Cpu = in["cpu"].(int)
		oi.Age = in["age"].(string)
		oi.Flags = in["flags"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatFullConeSessionOper(d *schema.ResourceData) edpt.Cgnv6FixedNatFullConeSessionOper {
	var ret edpt.Cgnv6FixedNatFullConeSessionOper

	ret.Oper = getObjectCgnv6FixedNatFullConeSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
