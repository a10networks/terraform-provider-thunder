package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectionIpv6SrcHashMaskBitsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_protection_ipv6_src_hash_mask_bits_oper`: Operational Status for the object ipv6-src-hash-mask-bits\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosProtectionIpv6SrcHashMaskBitsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"offsets": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mask_bit_offset_1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mask_bit_offset_2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mask_bit_offset_3": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mask_bit_offset_4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mask_bit_offset_5": {
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

func resourceDdosProtectionIpv6SrcHashMaskBitsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionIpv6SrcHashMaskBitsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionIpv6SrcHashMaskBitsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosProtectionIpv6SrcHashMaskBitsOperOper := setObjectDdosProtectionIpv6SrcHashMaskBitsOperOper(res)
		d.Set("oper", DdosProtectionIpv6SrcHashMaskBitsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosProtectionIpv6SrcHashMaskBitsOperOper(ret edpt.DataDdosProtectionIpv6SrcHashMaskBitsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"offsets": setSliceDdosProtectionIpv6SrcHashMaskBitsOperOperOffsets(ret.DtDdosProtectionIpv6SrcHashMaskBitsOper.Oper.Offsets),
		},
	}
}

func setSliceDdosProtectionIpv6SrcHashMaskBitsOperOperOffsets(d []edpt.DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["mask_bit_offset_1"] = item.MaskBitOffset1
		in["mask_bit_offset_2"] = item.MaskBitOffset2
		in["mask_bit_offset_3"] = item.MaskBitOffset3
		in["mask_bit_offset_4"] = item.MaskBitOffset4
		in["mask_bit_offset_5"] = item.MaskBitOffset5
		result = append(result, in)
	}
	return result
}

func getObjectDdosProtectionIpv6SrcHashMaskBitsOperOper(d []interface{}) edpt.DdosProtectionIpv6SrcHashMaskBitsOperOper {

	count1 := len(d)
	var ret edpt.DdosProtectionIpv6SrcHashMaskBitsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Offsets = getSliceDdosProtectionIpv6SrcHashMaskBitsOperOperOffsets(in["offsets"].([]interface{}))
	}
	return ret
}

func getSliceDdosProtectionIpv6SrcHashMaskBitsOperOperOffsets(d []interface{}) []edpt.DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets {

	count1 := len(d)
	ret := make([]edpt.DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets
		oi.MaskBitOffset1 = in["mask_bit_offset_1"].(int)
		oi.MaskBitOffset2 = in["mask_bit_offset_2"].(int)
		oi.MaskBitOffset3 = in["mask_bit_offset_3"].(int)
		oi.MaskBitOffset4 = in["mask_bit_offset_4"].(int)
		oi.MaskBitOffset5 = in["mask_bit_offset_5"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosProtectionIpv6SrcHashMaskBitsOper(d *schema.ResourceData) edpt.DdosProtectionIpv6SrcHashMaskBitsOper {
	var ret edpt.DdosProtectionIpv6SrcHashMaskBitsOper

	ret.Oper = getObjectDdosProtectionIpv6SrcHashMaskBitsOperOper(d.Get("oper").([]interface{}))
	return ret
}
