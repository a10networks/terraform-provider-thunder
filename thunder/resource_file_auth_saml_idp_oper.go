package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAuthSamlIdpOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_auth_saml_idp_oper`: Operational Status for the object auth-saml-idp\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAuthSamlIdpOperCreate,
		UpdateContext: resourceFileAuthSamlIdpOperUpdate,
		ReadContext:   resourceFileAuthSamlIdpOperRead,
		DeleteContext: resourceFileAuthSamlIdpOperDelete,

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
func resourceFileAuthSamlIdpOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthSamlIdpOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthSamlIdpOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthSamlIdpOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAuthSamlIdpOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthSamlIdpOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthSamlIdpOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthSamlIdpOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAuthSamlIdpOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthSamlIdpOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthSamlIdpOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAuthSamlIdpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthSamlIdpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthSamlIdpOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAuthSamlIdpOperOper(d []interface{}) edpt.FileAuthSamlIdpOperOper {

	count1 := len(d)
	var ret edpt.FileAuthSamlIdpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAuthSamlIdpOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAuthSamlIdpOperOperFileList(d []interface{}) []edpt.FileAuthSamlIdpOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAuthSamlIdpOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAuthSamlIdpOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAuthSamlIdpOper(d *schema.ResourceData) edpt.FileAuthSamlIdpOper {
	var ret edpt.FileAuthSamlIdpOper
	ret.Inst.Oper = getObjectFileAuthSamlIdpOperOper(d.Get("oper").([]interface{}))
	return ret
}
