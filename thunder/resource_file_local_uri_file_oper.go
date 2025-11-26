package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileLocalUriFileOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_local_uri_file_oper`: Operational Status for the object local-uri-file\n\n__PLACEHOLDER__",
		CreateContext: resourceFileLocalUriFileOperCreate,
		UpdateContext: resourceFileLocalUriFileOperUpdate,
		ReadContext:   resourceFileLocalUriFileOperRead,
		DeleteContext: resourceFileLocalUriFileOperDelete,

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
									"url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"template": {
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
func resourceFileLocalUriFileOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLocalUriFileOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLocalUriFileOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLocalUriFileOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileLocalUriFileOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLocalUriFileOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLocalUriFileOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLocalUriFileOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileLocalUriFileOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLocalUriFileOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLocalUriFileOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileLocalUriFileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLocalUriFileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLocalUriFileOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileLocalUriFileOperOper(d []interface{}) edpt.FileLocalUriFileOperOper {

	count1 := len(d)
	var ret edpt.FileLocalUriFileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileLocalUriFileOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileLocalUriFileOperOperFileList(d []interface{}) []edpt.FileLocalUriFileOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileLocalUriFileOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileLocalUriFileOperOperFileList
		oi.File = in["file"].(string)
		oi.Url = in["url"].(string)
		oi.Template = in["template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileLocalUriFileOper(d *schema.ResourceData) edpt.FileLocalUriFileOper {
	var ret edpt.FileLocalUriFileOper
	ret.Inst.Oper = getObjectFileLocalUriFileOperOper(d.Get("oper").([]interface{}))
	return ret
}
