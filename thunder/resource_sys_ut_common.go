package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtCommon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_common`: Common\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtCommonCreate,
		UpdateContext: resourceSysUtCommonUpdate,
		ReadContext:   resourceSysUtCommonRead,
		DeleteContext: resourceSysUtCommonDelete,

		Schema: map[string]*schema.Schema{
			"delay": {
				Type: schema.TypeInt, Optional: true, Description: "wait time in seconds after each run",
			},
			"proceed_on_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run test even in case of event failure",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSysUtCommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtCommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtCommon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtCommonRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtCommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtCommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtCommon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtCommonRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtCommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtCommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtCommon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtCommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtCommon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtCommon(d *schema.ResourceData) edpt.SysUtCommon {
	var ret edpt.SysUtCommon
	ret.Inst.Delay = d.Get("delay").(int)
	ret.Inst.ProceedOnError = d.Get("proceed_on_error").(int)
	//omit uuid
	return ret
}
