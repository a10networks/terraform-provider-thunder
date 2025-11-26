package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsOffload() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_offload`: Virtual Chassis System offload settings\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsOffloadCreate,
		UpdateContext: resourceVcsOffloadUpdate,
		ReadContext:   resourceVcsOffloadRead,
		DeleteContext: resourceVcsOffloadDelete,

		Schema: map[string]*schema.Schema{
			"application": {
				Type: schema.TypeString, Required: true, Description: "'upgrade': Image transfer during upgrade;",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'scp': Use scp;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVcsOffloadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsOffloadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsOffload(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsOffloadRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsOffloadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsOffloadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsOffload(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsOffloadRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsOffloadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsOffloadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsOffload(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsOffloadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsOffloadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsOffload(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsOffload(d *schema.ResourceData) edpt.VcsOffload {
	var ret edpt.VcsOffload
	ret.Inst.Application = d.Get("application").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	return ret
}
