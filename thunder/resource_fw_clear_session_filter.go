package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwClearSessionFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_clear_session_filter`: Enable clear L4 session filter for fw (Default: Disable)\n\n__PLACEHOLDER__",
		CreateContext: resourceFwClearSessionFilterCreate,
		UpdateContext: resourceFwClearSessionFilterUpdate,
		ReadContext:   resourceFwClearSessionFilterRead,
		DeleteContext: resourceFwClearSessionFilterDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': Disable clear L4 session filter for fw (Default: disabled); 'enable': Enable clear L4 session filter for fw;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwClearSessionFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClearSessionFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClearSessionFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwClearSessionFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceFwClearSessionFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClearSessionFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClearSessionFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwClearSessionFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceFwClearSessionFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClearSessionFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClearSessionFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwClearSessionFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClearSessionFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClearSessionFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwClearSessionFilter(d *schema.ResourceData) edpt.FwClearSessionFilter {
	var ret edpt.FwClearSessionFilter
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
