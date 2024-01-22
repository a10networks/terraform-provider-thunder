package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosRunTimeUserString() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_run_time_user_string`: Configure run time user string\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosRunTimeUserStringCreate,
		UpdateContext: resourceDdosRunTimeUserStringUpdate,
		ReadContext:   resourceDdosRunTimeUserStringRead,
		DeleteContext: resourceDdosRunTimeUserStringDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Description: "Add run time user string",
			},
		},
	}
}
func resourceDdosRunTimeUserStringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosRunTimeUserStringCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosRunTimeUserString(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosRunTimeUserStringRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosRunTimeUserStringUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosRunTimeUserStringUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosRunTimeUserString(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosRunTimeUserStringRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosRunTimeUserStringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosRunTimeUserStringDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosRunTimeUserString(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosRunTimeUserStringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosRunTimeUserStringRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosRunTimeUserString(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosRunTimeUserString(d *schema.ResourceData) edpt.DdosRunTimeUserString {
	var ret edpt.DdosRunTimeUserString
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	return ret
}
