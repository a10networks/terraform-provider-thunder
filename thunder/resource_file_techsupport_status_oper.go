package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTechsupportStatusOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_techsupport_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTechsupportStatusOperCreate,
		UpdateContext: resourceFileTechsupportStatusOperUpdate,
		ReadContext:   resourceFileTechsupportStatusOperRead,
		DeleteContext: resourceFileTechsupportStatusOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}
func resourceFileTechsupportStatusOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportStatusOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportStatusOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportStatusOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTechsupportStatusOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportStatusOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportStatusOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportStatusOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTechsupportStatusOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportStatusOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportStatusOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTechsupportStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportStatusOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileTechsupportStatusOperOper(d []interface{}) edpt.FileTechsupportStatusOperOper {

	count1 := len(d)
	var ret edpt.FileTechsupportStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Message = in["message"].(string)
	}
	return ret
}

func dataToEndpointFileTechsupportStatusOper(d *schema.ResourceData) edpt.FileTechsupportStatusOper {
	var ret edpt.FileTechsupportStatusOper
	ret.Inst.Oper = getObjectFileTechsupportStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
