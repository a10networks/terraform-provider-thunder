package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileLicenseOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_license_oper`: Operational Status for the object license\n\n__PLACEHOLDER__",
		CreateContext: resourceFileLicenseOperCreate,
		UpdateContext: resourceFileLicenseOperUpdate,
		ReadContext:   resourceFileLicenseOperRead,
		DeleteContext: resourceFileLicenseOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"feature_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"feature_installed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expire_date": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"notice": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"temporary": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sn": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bandwidth": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file_name": {
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
func resourceFileLicenseOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLicenseOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLicenseOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLicenseOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileLicenseOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLicenseOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLicenseOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLicenseOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileLicenseOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLicenseOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLicenseOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileLicenseOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLicenseOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLicenseOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileLicenseOperOper(d []interface{}) edpt.FileLicenseOperOper {

	count1 := len(d)
	var ret edpt.FileLicenseOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostId = in["host_id"].(string)
		ret.FeatureList = getSliceFileLicenseOperOperFeatureList(in["feature_list"].([]interface{}))
		ret.FileList = getSliceFileLicenseOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileLicenseOperOperFeatureList(d []interface{}) []edpt.FileLicenseOperOperFeatureList {

	count1 := len(d)
	ret := make([]edpt.FileLicenseOperOperFeatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileLicenseOperOperFeatureList
		oi.FeatureInstalled = in["feature_installed"].(string)
		oi.Version = in["version"].(string)
		oi.ExpireDate = in["expire_date"].(string)
		oi.Notice = in["notice"].(string)
		oi.Temporary = in["temporary"].(string)
		oi.Sn = in["sn"].(string)
		oi.Bandwidth = in["bandwidth"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFileLicenseOperOperFileList(d []interface{}) []edpt.FileLicenseOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileLicenseOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileLicenseOperOperFileList
		oi.File_name = in["file_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileLicenseOper(d *schema.ResourceData) edpt.FileLicenseOper {
	var ret edpt.FileLicenseOper
	ret.Inst.Oper = getObjectFileLicenseOperOper(d.Get("oper").([]interface{}))
	return ret
}
