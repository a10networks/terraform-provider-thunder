package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScmLicensestatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scm_licensestatus_oper`: Operational Status for the object licensestatus\n\n__PLACEHOLDER__",
		ReadContext: resourceScmLicensestatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_org_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"primary_org_email": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"license_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"license_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"license_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"org_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
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

func resourceScmLicensestatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScmLicensestatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScmLicensestatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScmLicensestatusOperOper := setObjectScmLicensestatusOperOper(res)
		d.Set("oper", ScmLicensestatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScmLicensestatusOperOper(ret edpt.DataScmLicensestatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"primary_org_id":    ret.DtScmLicensestatusOper.Oper.PrimaryOrgId,
			"primary_org_email": ret.DtScmLicensestatusOper.Oper.PrimaryOrgEmail,
			"license_list":      setSliceScmLicensestatusOperOperLicenseList(ret.DtScmLicensestatusOper.Oper.LicenseList),
		},
	}
}

func setSliceScmLicensestatusOperOperLicenseList(d []edpt.ScmLicensestatusOperOperLicenseList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["license_type"] = item.LicenseType
		in["license_name"] = item.LicenseName
		in["org_name"] = item.OrgName
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func getObjectScmLicensestatusOperOper(d []interface{}) edpt.ScmLicensestatusOperOper {

	count1 := len(d)
	var ret edpt.ScmLicensestatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryOrgId = in["primary_org_id"].(string)
		ret.PrimaryOrgEmail = in["primary_org_email"].(string)
		ret.LicenseList = getSliceScmLicensestatusOperOperLicenseList(in["license_list"].([]interface{}))
	}
	return ret
}

func getSliceScmLicensestatusOperOperLicenseList(d []interface{}) []edpt.ScmLicensestatusOperOperLicenseList {

	count1 := len(d)
	ret := make([]edpt.ScmLicensestatusOperOperLicenseList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScmLicensestatusOperOperLicenseList
		oi.LicenseType = in["license_type"].(string)
		oi.LicenseName = in["license_name"].(string)
		oi.OrgName = in["org_name"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScmLicensestatusOper(d *schema.ResourceData) edpt.ScmLicensestatusOper {
	var ret edpt.ScmLicensestatusOper

	ret.Oper = getObjectScmLicensestatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
