package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneMappingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_mapping_oper`: Operational Status for the object mapping\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOneMappingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_mapping_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sessions": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"inside_address_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_address_ipv6": {
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
						"inside_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"shared_pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6OneToOneMappingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneMappingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneMappingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOneMappingOperOper := setObjectCgnv6OneToOneMappingOperOper(res)
		d.Set("oper", Cgnv6OneToOneMappingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOneMappingOperOper(ret edpt.DataCgnv6OneToOneMappingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_mapping_list": setSliceCgnv6OneToOneMappingOperOperSessionMappingList(ret.DtCgnv6OneToOneMappingOper.Oper.SessionMappingList),
			"total":                ret.DtCgnv6OneToOneMappingOper.Oper.Total,
			"inside_address_ipv4":  ret.DtCgnv6OneToOneMappingOper.Oper.InsideAddressIpv4,
			"inside_address_ipv6":  ret.DtCgnv6OneToOneMappingOper.Oper.InsideAddressIpv6,
			"all_partitions":       ret.DtCgnv6OneToOneMappingOper.Oper.AllPartitions,
			"shared_partition":     ret.DtCgnv6OneToOneMappingOper.Oper.SharedPartition,
			"partition_name":       ret.DtCgnv6OneToOneMappingOper.Oper.PartitionName,
			"inside_addr_start":    ret.DtCgnv6OneToOneMappingOper.Oper.InsideAddrStart,
			"inside_addr_end":      ret.DtCgnv6OneToOneMappingOper.Oper.InsideAddrEnd,
			"nat_addr_val":         ret.DtCgnv6OneToOneMappingOper.Oper.NatAddrVal,
			"nat_addr_start":       ret.DtCgnv6OneToOneMappingOper.Oper.NatAddrStart,
			"nat_addr_end":         ret.DtCgnv6OneToOneMappingOper.Oper.NatAddrEnd,
			"pool_name":            ret.DtCgnv6OneToOneMappingOper.Oper.PoolName,
			"shared_pool_name":     ret.DtCgnv6OneToOneMappingOper.Oper.SharedPoolName,
		},
	}
}

func setSliceCgnv6OneToOneMappingOperOperSessionMappingList(d []edpt.Cgnv6OneToOneMappingOperOperSessionMappingList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_ipv4_address"] = item.InsideIpv4Address
		in["inside_ipv6_address"] = item.InsideIpv6Address
		in["nat_address"] = item.NatAddress
		in["sessions"] = item.Sessions
		in["age"] = item.Age
		in["pool"] = item.Pool
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6OneToOneMappingOperOper(d []interface{}) edpt.Cgnv6OneToOneMappingOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneMappingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionMappingList = getSliceCgnv6OneToOneMappingOperOperSessionMappingList(in["session_mapping_list"].([]interface{}))
		ret.Total = in["total"].(int)
		ret.InsideAddressIpv4 = in["inside_address_ipv4"].(string)
		ret.InsideAddressIpv6 = in["inside_address_ipv6"].(string)
		ret.AllPartitions = in["all_partitions"].(int)
		ret.SharedPartition = in["shared_partition"].(int)
		ret.PartitionName = in["partition_name"].(string)
		ret.InsideAddrStart = in["inside_addr_start"].(string)
		ret.InsideAddrEnd = in["inside_addr_end"].(string)
		ret.NatAddrVal = in["nat_addr_val"].(string)
		ret.NatAddrStart = in["nat_addr_start"].(string)
		ret.NatAddrEnd = in["nat_addr_end"].(string)
		ret.PoolName = in["pool_name"].(string)
		ret.SharedPoolName = in["shared_pool_name"].(string)
	}
	return ret
}

func getSliceCgnv6OneToOneMappingOperOperSessionMappingList(d []interface{}) []edpt.Cgnv6OneToOneMappingOperOperSessionMappingList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneMappingOperOperSessionMappingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneMappingOperOperSessionMappingList
		oi.InsideIpv4Address = in["inside_ipv4_address"].(string)
		oi.InsideIpv6Address = in["inside_ipv6_address"].(string)
		oi.NatAddress = in["nat_address"].(string)
		oi.Sessions = in["sessions"].(int)
		oi.Age = in["age"].(string)
		oi.Pool = in["pool"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneMappingOper(d *schema.ResourceData) edpt.Cgnv6OneToOneMappingOper {
	var ret edpt.Cgnv6OneToOneMappingOper

	ret.Oper = getObjectCgnv6OneToOneMappingOperOper(d.Get("oper").([]interface{}))
	return ret
}
