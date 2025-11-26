package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSslCrlOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ssl_crl_oper`: Operational Status for the object ssl-crl\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSslCrlOperCreate,
		UpdateContext: resourceFileSslCrlOperUpdate,
		ReadContext:   resourceFileSslCrlOperRead,
		DeleteContext: resourceFileSslCrlOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"issuer": {
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
func resourceFileSslCrlOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCrlOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCrlOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCrlOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslCrlOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCrlOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCrlOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCrlOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSslCrlOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCrlOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCrlOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSslCrlOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCrlOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCrlOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileSslCrlOperOper(d []interface{}) edpt.FileSslCrlOperOper {

	count1 := len(d)
	var ret edpt.FileSslCrlOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.FileList = getSliceFileSslCrlOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileSslCrlOperOperFileList(d []interface{}) []edpt.FileSslCrlOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileSslCrlOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileSslCrlOperOperFileList
		oi.File = in["file"].(string)
		oi.Issuer = in["issuer"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileSslCrlOper(d *schema.ResourceData) edpt.FileSslCrlOper {
	var ret edpt.FileSslCrlOper
	ret.Inst.Oper = getObjectFileSslCrlOperOper(d.Get("oper").([]interface{}))
	return ret
}
