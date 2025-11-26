package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPsuInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_psu_info_oper`: Operational Status for the object psu-info\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPsuInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"psu_info": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"upper_left_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"upper_right_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lower_left_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lower_right_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"right_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"left_serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSystemPsuInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPsuInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPsuInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPsuInfoOperOper := setObjectSystemPsuInfoOperOper(res)
		d.Set("oper", SystemPsuInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPsuInfoOperOper(ret edpt.DataSystemPsuInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"psu_info": setObjectSystemPsuInfoOperOperPsuInfo(ret.DtSystemPsuInfoOper.Oper.PsuInfo),
		},
	}
}

func setObjectSystemPsuInfoOperOperPsuInfo(d edpt.SystemPsuInfoOperOperPsuInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["error"] = d.Error

	in["upper_left_serial_number"] = d.UpperLeftSerialNumber

	in["upper_right_serial_number"] = d.UpperRightSerialNumber

	in["lower_left_serial_number"] = d.LowerLeftSerialNumber

	in["lower_right_serial_number"] = d.LowerRightSerialNumber

	in["right_serial_number"] = d.RightSerialNumber

	in["left_serial_number"] = d.LeftSerialNumber
	result = append(result, in)
	return result
}

func getObjectSystemPsuInfoOperOper(d []interface{}) edpt.SystemPsuInfoOperOper {

	count1 := len(d)
	var ret edpt.SystemPsuInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PsuInfo = getObjectSystemPsuInfoOperOperPsuInfo(in["psu_info"].([]interface{}))
	}
	return ret
}

func getObjectSystemPsuInfoOperOperPsuInfo(d []interface{}) edpt.SystemPsuInfoOperOperPsuInfo {

	count1 := len(d)
	var ret edpt.SystemPsuInfoOperOperPsuInfo
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(string)
		ret.UpperLeftSerialNumber = in["upper_left_serial_number"].(string)
		ret.UpperRightSerialNumber = in["upper_right_serial_number"].(string)
		ret.LowerLeftSerialNumber = in["lower_left_serial_number"].(string)
		ret.LowerRightSerialNumber = in["lower_right_serial_number"].(string)
		ret.RightSerialNumber = in["right_serial_number"].(string)
		ret.LeftSerialNumber = in["left_serial_number"].(string)
	}
	return ret
}

func dataToEndpointSystemPsuInfoOper(d *schema.ResourceData) edpt.SystemPsuInfoOper {
	var ret edpt.SystemPsuInfoOper

	ret.Oper = getObjectSystemPsuInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
