package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAuthPortalOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_auth_portal_oper`: Operational Status for the object auth-portal\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAuthPortalOperCreate,
		UpdateContext: resourceFileAuthPortalOperUpdate,
		ReadContext:   resourceFileAuthPortalOperRead,
		DeleteContext: resourceFileAuthPortalOperDelete,

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
func resourceFileAuthPortalOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthPortalOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAuthPortalOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthPortalOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAuthPortalOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAuthPortalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAuthPortalOperOper(d []interface{}) edpt.FileAuthPortalOperOper {

	count1 := len(d)
	var ret edpt.FileAuthPortalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAuthPortalOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAuthPortalOperOperFileList(d []interface{}) []edpt.FileAuthPortalOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAuthPortalOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAuthPortalOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAuthPortalOper(d *schema.ResourceData) edpt.FileAuthPortalOper {
	var ret edpt.FileAuthPortalOper
	ret.Inst.Oper = getObjectFileAuthPortalOperOper(d.Get("oper").([]interface{}))
	return ret
}
