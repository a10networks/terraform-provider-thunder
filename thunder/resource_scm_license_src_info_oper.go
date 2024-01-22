package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScmLicenseSrcInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scm_license_src_info_oper`: Operational Status for the object license-src-info\n\n__PLACEHOLDER__",
		ReadContext: resourceScmLicenseSrcInfoOperRead,

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
						"source1": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source1_module_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source1_module": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source1_expiry": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source1_notes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"source2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source2_module_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source2_module": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source2_expiry": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source2_notes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"source3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source3_module_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source3_module": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source3_expiry": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source3_notes": {
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

func resourceScmLicenseSrcInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScmLicenseSrcInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScmLicenseSrcInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScmLicenseSrcInfoOperOper := setObjectScmLicenseSrcInfoOperOper(res)
		d.Set("oper", ScmLicenseSrcInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScmLicenseSrcInfoOperOper(ret edpt.DataScmLicenseSrcInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			//omit uuid
			"usb_uuid":            ret.DtScmLicenseSrcInfoOper.Oper.UsbUuid,
			"billing_serial":      ret.DtScmLicenseSrcInfoOper.Oper.BillingSerial,
			"token":               ret.DtScmLicenseSrcInfoOper.Oper.Token,
			"product":             ret.DtScmLicenseSrcInfoOper.Oper.Product,
			"platform":            ret.DtScmLicenseSrcInfoOper.Oper.Platform,
			"burst":               ret.DtScmLicenseSrcInfoOper.Oper.Burst,
			"version":             ret.DtScmLicenseSrcInfoOper.Oper.Version,
			"glm_ping_interval":   ret.DtScmLicenseSrcInfoOper.Oper.GlmPingInterval,
			"source1":             ret.DtScmLicenseSrcInfoOper.Oper.Source1,
			"source1_module_list": setSliceScmLicenseSrcInfoOperOperSource1ModuleList(ret.DtScmLicenseSrcInfoOper.Oper.Source1ModuleList),
			"source2":             ret.DtScmLicenseSrcInfoOper.Oper.Source2,
			"source2_module_list": setSliceScmLicenseSrcInfoOperOperSource2ModuleList(ret.DtScmLicenseSrcInfoOper.Oper.Source2ModuleList),
			"source3":             ret.DtScmLicenseSrcInfoOper.Oper.Source3,
			"source3_module_list": setSliceScmLicenseSrcInfoOperOperSource3ModuleList(ret.DtScmLicenseSrcInfoOper.Oper.Source3ModuleList),
			"hw_serialno":         ret.DtScmLicenseSrcInfoOper.Oper.HwSerialno,
			"product_desc":        ret.DtScmLicenseSrcInfoOper.Oper.ProductDesc,
		},
	}
}

func setSliceScmLicenseSrcInfoOperOperSource1ModuleList(d []edpt.ScmLicenseSrcInfoOperOperSource1ModuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["source1_module"] = item.Source1Module
		in["source1_expiry"] = item.Source1Expiry
		in["source1_notes"] = item.Source1Notes
		result = append(result, in)
	}
	return result
}

func setSliceScmLicenseSrcInfoOperOperSource2ModuleList(d []edpt.ScmLicenseSrcInfoOperOperSource2ModuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["source2_module"] = item.Source2Module
		in["source2_expiry"] = item.Source2Expiry
		in["source2_notes"] = item.Source2Notes
		result = append(result, in)
	}
	return result
}

func setSliceScmLicenseSrcInfoOperOperSource3ModuleList(d []edpt.ScmLicenseSrcInfoOperOperSource3ModuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["source3_module"] = item.Source3Module
		in["source3_expiry"] = item.Source3Expiry
		in["source3_notes"] = item.Source3Notes
		result = append(result, in)
	}
	return result
}

func getObjectScmLicenseSrcInfoOperOper(d []interface{}) edpt.ScmLicenseSrcInfoOperOper {

	count1 := len(d)
	var ret edpt.ScmLicenseSrcInfoOperOper
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
		ret.Source1 = in["source1"].(string)
		ret.Source1ModuleList = getSliceScmLicenseSrcInfoOperOperSource1ModuleList(in["source1_module_list"].([]interface{}))
		ret.Source2 = in["source2"].(string)
		ret.Source2ModuleList = getSliceScmLicenseSrcInfoOperOperSource2ModuleList(in["source2_module_list"].([]interface{}))
		ret.Source3 = in["source3"].(string)
		ret.Source3ModuleList = getSliceScmLicenseSrcInfoOperOperSource3ModuleList(in["source3_module_list"].([]interface{}))
		ret.HwSerialno = in["hw_serialno"].(string)
		ret.ProductDesc = in["product_desc"].(string)
	}
	return ret
}

func getSliceScmLicenseSrcInfoOperOperSource1ModuleList(d []interface{}) []edpt.ScmLicenseSrcInfoOperOperSource1ModuleList {

	count1 := len(d)
	ret := make([]edpt.ScmLicenseSrcInfoOperOperSource1ModuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScmLicenseSrcInfoOperOperSource1ModuleList
		oi.Source1Module = in["source1_module"].(string)
		oi.Source1Expiry = in["source1_expiry"].(string)
		oi.Source1Notes = in["source1_notes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScmLicenseSrcInfoOperOperSource2ModuleList(d []interface{}) []edpt.ScmLicenseSrcInfoOperOperSource2ModuleList {

	count1 := len(d)
	ret := make([]edpt.ScmLicenseSrcInfoOperOperSource2ModuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScmLicenseSrcInfoOperOperSource2ModuleList
		oi.Source2Module = in["source2_module"].(string)
		oi.Source2Expiry = in["source2_expiry"].(string)
		oi.Source2Notes = in["source2_notes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScmLicenseSrcInfoOperOperSource3ModuleList(d []interface{}) []edpt.ScmLicenseSrcInfoOperOperSource3ModuleList {

	count1 := len(d)
	ret := make([]edpt.ScmLicenseSrcInfoOperOperSource3ModuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScmLicenseSrcInfoOperOperSource3ModuleList
		oi.Source3Module = in["source3_module"].(string)
		oi.Source3Expiry = in["source3_expiry"].(string)
		oi.Source3Notes = in["source3_notes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScmLicenseSrcInfoOper(d *schema.ResourceData) edpt.ScmLicenseSrcInfoOper {
	var ret edpt.ScmLicenseSrcInfoOper

	ret.Oper = getObjectScmLicenseSrcInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
