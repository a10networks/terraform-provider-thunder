package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAuthJwksOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_auth_jwks_oper`: Operational Status for the object auth-jwks\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAuthJwksOperCreate,
		UpdateContext: resourceFileAuthJwksOperUpdate,
		ReadContext:   resourceFileAuthJwksOperRead,
		DeleteContext: resourceFileAuthJwksOperDelete,

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
func resourceFileAuthJwksOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthJwksOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthJwksOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthJwksOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAuthJwksOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthJwksOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthJwksOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthJwksOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAuthJwksOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthJwksOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthJwksOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAuthJwksOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthJwksOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthJwksOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAuthJwksOperOper(d []interface{}) edpt.FileAuthJwksOperOper {

	count1 := len(d)
	var ret edpt.FileAuthJwksOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAuthJwksOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAuthJwksOperOperFileList(d []interface{}) []edpt.FileAuthJwksOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAuthJwksOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAuthJwksOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAuthJwksOper(d *schema.ResourceData) edpt.FileAuthJwksOper {
	var ret edpt.FileAuthJwksOper
	ret.Inst.Oper = getObjectFileAuthJwksOperOper(d.Get("oper").([]interface{}))
	return ret
}
