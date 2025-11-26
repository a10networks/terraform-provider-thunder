package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileDnssecDsOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_dnssec_ds_oper`: Operational Status for the object dnssec-ds\n\n__PLACEHOLDER__",
		CreateContext: resourceFileDnssecDsOperCreate,
		UpdateContext: resourceFileDnssecDsOperUpdate,
		ReadContext:   resourceFileDnssecDsOperRead,
		DeleteContext: resourceFileDnssecDsOperDelete,

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
									"parent": {
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
func resourceFileDnssecDsOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDsOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDsOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDnssecDsOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileDnssecDsOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDsOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDsOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDnssecDsOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileDnssecDsOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDsOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDsOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileDnssecDsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDsOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileDnssecDsOperOper(d []interface{}) edpt.FileDnssecDsOperOper {

	count1 := len(d)
	var ret edpt.FileDnssecDsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileDnssecDsOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileDnssecDsOperOperFileList(d []interface{}) []edpt.FileDnssecDsOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileDnssecDsOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileDnssecDsOperOperFileList
		oi.File = in["file"].(string)
		oi.Parent = in["parent"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileDnssecDsOper(d *schema.ResourceData) edpt.FileDnssecDsOper {
	var ret edpt.FileDnssecDsOper
	ret.Inst.Oper = getObjectFileDnssecDsOperOper(d.Get("oper").([]interface{}))
	return ret
}
