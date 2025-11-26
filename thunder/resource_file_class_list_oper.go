package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileClassListOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_class_list_oper`: Operational Status for the object class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceFileClassListOperCreate,
		UpdateContext: resourceFileClassListOperUpdate,
		ReadContext:   resourceFileClassListOperRead,
		DeleteContext: resourceFileClassListOperDelete,

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
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"location": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"size": {
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
func resourceFileClassListOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileClassListOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileClassListOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileClassListOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileClassListOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileClassListOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileClassListOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileClassListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileClassListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileClassListOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileClassListOperOper(d []interface{}) edpt.FileClassListOperOper {

	count1 := len(d)
	var ret edpt.FileClassListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileClassListOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileClassListOperOperFileList(d []interface{}) []edpt.FileClassListOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileClassListOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileClassListOperOperFileList
		oi.File = in["file"].(string)
		oi.Type = in["type"].(string)
		oi.Location = in["location"].(string)
		oi.UserTag = in["user_tag"].(string)
		oi.Size = in["size"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileClassListOper(d *schema.ResourceData) edpt.FileClassListOper {
	var ret edpt.FileClassListOper
	ret.Inst.Oper = getObjectFileClassListOperOper(d.Get("oper").([]interface{}))
	return ret
}
