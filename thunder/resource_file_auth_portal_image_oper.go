package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAuthPortalImageOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_auth_portal_image_oper`: Operational Status for the object auth-portal-image\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAuthPortalImageOperCreate,
		UpdateContext: resourceFileAuthPortalImageOperUpdate,
		ReadContext:   resourceFileAuthPortalImageOperRead,
		DeleteContext: resourceFileAuthPortalImageOperDelete,

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
func resourceFileAuthPortalImageOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalImageOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalImageOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthPortalImageOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAuthPortalImageOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalImageOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalImageOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAuthPortalImageOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAuthPortalImageOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalImageOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalImageOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAuthPortalImageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAuthPortalImageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAuthPortalImageOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAuthPortalImageOperOper(d []interface{}) edpt.FileAuthPortalImageOperOper {

	count1 := len(d)
	var ret edpt.FileAuthPortalImageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAuthPortalImageOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAuthPortalImageOperOperFileList(d []interface{}) []edpt.FileAuthPortalImageOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAuthPortalImageOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAuthPortalImageOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAuthPortalImageOper(d *schema.ResourceData) edpt.FileAuthPortalImageOper {
	var ret edpt.FileAuthPortalImageOper
	ret.Inst.Oper = getObjectFileAuthPortalImageOperOper(d.Get("oper").([]interface{}))
	return ret
}
