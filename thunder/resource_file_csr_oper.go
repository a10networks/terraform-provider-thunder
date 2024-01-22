package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileCsrOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_csr_oper`: Operational Status for the object csr\n\n__PLACEHOLDER__",
		CreateContext: resourceFileCsrOperCreate,
		UpdateContext: resourceFileCsrOperUpdate,
		ReadContext:   resourceFileCsrOperRead,
		DeleteContext: resourceFileCsrOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_csr": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"common_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"organization": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"sortby_name": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}
func resourceFileCsrOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCsrOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileCsrOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCsrOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileCsrOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileCsrOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileCsrOperOper(d []interface{}) edpt.FileCsrOperOper {

	count1 := len(d)
	var ret edpt.FileCsrOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslCsr = getSliceFileCsrOperOperSslCsr(in["ssl_csr"].([]interface{}))
		ret.SortbyName = in["sortby_name"].(int)
	}
	return ret
}

func getSliceFileCsrOperOperSslCsr(d []interface{}) []edpt.FileCsrOperOperSslCsr {

	count1 := len(d)
	ret := make([]edpt.FileCsrOperOperSslCsr, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileCsrOperOperSslCsr
		oi.Name = in["name"].(string)
		oi.Type = in["type"].(string)
		oi.CommonName = in["common_name"].(string)
		oi.Organization = in["organization"].(string)
		oi.Subject = in["subject"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileCsrOper(d *schema.ResourceData) edpt.FileCsrOper {
	var ret edpt.FileCsrOper
	ret.Inst.Oper = getObjectFileCsrOperOper(d.Get("oper").([]interface{}))
	return ret
}
