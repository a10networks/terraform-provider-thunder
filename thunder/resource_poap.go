package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_poap`: Set POAP; when next boot, if startup-config is empty, system will enter POAP Mode\n\n__PLACEHOLDER__",
		CreateContext: resourcePoapCreate,
		UpdateContext: resourcePoapUpdate,
		ReadContext:   resourcePoapRead,
		DeleteContext: resourcePoapDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable POAP mode; 'disable': Disable POAP mode;",
			},
		},
	}
}
func resourcePoapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePoapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPoap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePoapRead(ctx, d, meta)
	}
	return diags
}

func resourcePoapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePoapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPoap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePoapRead(ctx, d, meta)
	}
	return diags
}
func resourcePoapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePoapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPoap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePoapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePoapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPoap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPoap(d *schema.ResourceData) edpt.Poap {
	var ret edpt.Poap
	ret.Inst.Action = d.Get("action").(string)
	return ret
}
