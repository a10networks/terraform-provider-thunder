package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryLicenseOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_license_oper`: Operational Status for the object license\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryLicenseOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"module_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"license_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"license_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"license_expiry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"remaining_period": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"is_grace": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"grace_period": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"serial_number": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryLicenseOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryLicenseOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryLicenseOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryLicenseOperOper := setObjectWebCategoryLicenseOperOper(res)
		d.Set("oper", WebCategoryLicenseOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryLicenseOperOper(ret edpt.DataWebCategoryLicenseOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"module_status":    ret.DtWebCategoryLicenseOper.Oper.ModuleStatus,
			"license_status":   ret.DtWebCategoryLicenseOper.Oper.LicenseStatus,
			"license_type":     ret.DtWebCategoryLicenseOper.Oper.LicenseType,
			"license_expiry":   ret.DtWebCategoryLicenseOper.Oper.LicenseExpiry,
			"remaining_period": ret.DtWebCategoryLicenseOper.Oper.RemainingPeriod,
			"is_grace":         ret.DtWebCategoryLicenseOper.Oper.IsGrace,
			"grace_period":     ret.DtWebCategoryLicenseOper.Oper.GracePeriod,
			"serial_number":    ret.DtWebCategoryLicenseOper.Oper.SerialNumber,
		},
	}
}

func getObjectWebCategoryLicenseOperOper(d []interface{}) edpt.WebCategoryLicenseOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryLicenseOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ModuleStatus = in["module_status"].(string)
		ret.LicenseStatus = in["license_status"].(string)
		ret.LicenseType = in["license_type"].(string)
		ret.LicenseExpiry = in["license_expiry"].(string)
		ret.RemainingPeriod = in["remaining_period"].(string)
		ret.IsGrace = in["is_grace"].(string)
		ret.GracePeriod = in["grace_period"].(string)
		ret.SerialNumber = in["serial_number"].(string)
	}
	return ret
}

func dataToEndpointWebCategoryLicenseOper(d *schema.ResourceData) edpt.WebCategoryLicenseOper {
	var ret edpt.WebCategoryLicenseOper

	ret.Oper = getObjectWebCategoryLicenseOperOper(d.Get("oper").([]interface{}))
	return ret
}
