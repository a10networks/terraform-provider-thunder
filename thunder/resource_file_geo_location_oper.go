package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileGeoLocationOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_geo_location_oper`: Operational Status for the object geo-location\n\n__PLACEHOLDER__",
		CreateContext: resourceFileGeoLocationOperCreate,
		UpdateContext: resourceFileGeoLocationOperUpdate,
		ReadContext:   resourceFileGeoLocationOperRead,
		DeleteContext: resourceFileGeoLocationOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
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
func resourceFileGeoLocationOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileGeoLocationOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGeoLocationOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileGeoLocationOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileGeoLocationOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGeoLocationOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGeoLocationOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileGeoLocationOperOper(d []interface{}) edpt.FileGeoLocationOperOper {

	count1 := len(d)
	var ret edpt.FileGeoLocationOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileGeoLocationOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileGeoLocationOperOperFileList(d []interface{}) []edpt.FileGeoLocationOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileGeoLocationOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileGeoLocationOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileGeoLocationOper(d *schema.ResourceData) edpt.FileGeoLocationOper {
	var ret edpt.FileGeoLocationOper
	ret.Inst.Oper = getObjectFileGeoLocationOperOper(d.Get("oper").([]interface{}))
	return ret
}
