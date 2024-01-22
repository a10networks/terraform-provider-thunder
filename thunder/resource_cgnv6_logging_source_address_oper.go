package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingSourceAddressOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_logging_source_address_oper`: Operational Status for the object source-address\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LoggingSourceAddressOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_address_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ref_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceCgnv6LoggingSourceAddressOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingSourceAddressOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingSourceAddressOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LoggingSourceAddressOperOper := setObjectCgnv6LoggingSourceAddressOperOper(res)
		d.Set("oper", Cgnv6LoggingSourceAddressOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LoggingSourceAddressOperOper(ret edpt.DataCgnv6LoggingSourceAddressOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"source_address_list": setObjectCgnv6LoggingSourceAddressOperOperSourceAddressList(ret.DtCgnv6LoggingSourceAddressOper.Oper.SourceAddressList),
		},
	}
}

func setObjectCgnv6LoggingSourceAddressOperOperSourceAddressList(d edpt.Cgnv6LoggingSourceAddressOperOperSourceAddressList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["source_address"] = d.SourceAddress

	in["ref_count"] = d.RefCount

	in["tcp_total"] = d.TcpTotal

	in["tcp_allocated"] = d.TcpAllocated

	in["tcp_freed"] = d.TcpFreed

	in["tcp_failed"] = d.TcpFailed

	in["udp_total"] = d.UdpTotal

	in["udp_allocated"] = d.UdpAllocated

	in["udp_freed"] = d.UdpFreed

	in["udp_failed"] = d.UdpFailed
	result = append(result, in)
	return result
}

func getObjectCgnv6LoggingSourceAddressOperOper(d []interface{}) edpt.Cgnv6LoggingSourceAddressOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingSourceAddressOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourceAddressList = getObjectCgnv6LoggingSourceAddressOperOperSourceAddressList(in["source_address_list"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6LoggingSourceAddressOperOperSourceAddressList(d []interface{}) edpt.Cgnv6LoggingSourceAddressOperOperSourceAddressList {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingSourceAddressOperOperSourceAddressList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourceAddress = in["source_address"].(string)
		ret.RefCount = in["ref_count"].(int)
		ret.TcpTotal = in["tcp_total"].(int)
		ret.TcpAllocated = in["tcp_allocated"].(int)
		ret.TcpFreed = in["tcp_freed"].(int)
		ret.TcpFailed = in["tcp_failed"].(int)
		ret.UdpTotal = in["udp_total"].(int)
		ret.UdpAllocated = in["udp_allocated"].(int)
		ret.UdpFreed = in["udp_freed"].(int)
		ret.UdpFailed = in["udp_failed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LoggingSourceAddressOper(d *schema.ResourceData) edpt.Cgnv6LoggingSourceAddressOper {
	var ret edpt.Cgnv6LoggingSourceAddressOper

	ret.Oper = getObjectCgnv6LoggingSourceAddressOperOper(d.Get("oper").([]interface{}))
	return ret
}
