package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScmLicenseinfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scm_licenseinfo_oper`: Operational Status for the object licenseinfo\n\n__PLACEHOLDER__",
		ReadContext: resourceScmLicenseinfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "",
						},
						"usb_uuid": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"billing_serial": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"token": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"product": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"platform": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"burst": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"glm_ping_interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"module_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"module": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expiry": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"notes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"hw_serialno": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"product_desc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScmLicenseinfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScmLicenseinfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScmLicenseinfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScmLicenseinfoOperOper := setObjectScmLicenseinfoOperOper(res)
		d.Set("oper", ScmLicenseinfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScmLicenseinfoOperOper(ret edpt.DataScmLicenseinfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			//omit uuid
			"usb_uuid":          ret.DtScmLicenseinfoOper.Oper.UsbUuid,
			"billing_serial":    ret.DtScmLicenseinfoOper.Oper.BillingSerial,
			"token":             ret.DtScmLicenseinfoOper.Oper.Token,
			"product":           ret.DtScmLicenseinfoOper.Oper.Product,
			"platform":          ret.DtScmLicenseinfoOper.Oper.Platform,
			"burst":             ret.DtScmLicenseinfoOper.Oper.Burst,
			"version":           ret.DtScmLicenseinfoOper.Oper.Version,
			"glm_ping_interval": ret.DtScmLicenseinfoOper.Oper.GlmPingInterval,
			"module_list":       setSliceScmLicenseinfoOperOperModuleList(ret.DtScmLicenseinfoOper.Oper.ModuleList),
			"hw_serialno":       ret.DtScmLicenseinfoOper.Oper.HwSerialno,
			"product_desc":      ret.DtScmLicenseinfoOper.Oper.ProductDesc,
		},
	}
}

func setSliceScmLicenseinfoOperOperModuleList(d []edpt.ScmLicenseinfoOperOperModuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["module"] = item.Module
		in["expiry"] = item.Expiry
		in["notes"] = item.Notes
		result = append(result, in)
	}
	return result
}

func getObjectScmLicenseinfoOperOper(d []interface{}) edpt.ScmLicenseinfoOperOper {

	count1 := len(d)
	var ret edpt.ScmLicenseinfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.UsbUuid = in["usb_uuid"].(string)
		ret.BillingSerial = in["billing_serial"].(string)
		ret.Token = in["token"].(string)
		ret.Product = in["product"].(string)
		ret.Platform = in["platform"].(string)
		ret.Burst = in["burst"].(string)
		ret.Version = in["version"].(string)
		ret.GlmPingInterval = in["glm_ping_interval"].(int)
		ret.ModuleList = getSliceScmLicenseinfoOperOperModuleList(in["module_list"].([]interface{}))
		ret.HwSerialno = in["hw_serialno"].(string)
		ret.ProductDesc = in["product_desc"].(string)
	}
	return ret
}

func getSliceScmLicenseinfoOperOperModuleList(d []interface{}) []edpt.ScmLicenseinfoOperOperModuleList {

	count1 := len(d)
	ret := make([]edpt.ScmLicenseinfoOperOperModuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScmLicenseinfoOperOperModuleList
		oi.Module = in["module"].(string)
		oi.Expiry = in["expiry"].(string)
		oi.Notes = in["notes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScmLicenseinfoOper(d *schema.ResourceData) edpt.ScmLicenseinfoOper {
	var ret edpt.ScmLicenseinfoOper

	ret.Oper = getObjectScmLicenseinfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
