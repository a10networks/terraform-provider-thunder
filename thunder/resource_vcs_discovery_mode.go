package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsDiscoveryMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_discovery_mode`: vcs election discovery mode\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsDiscoveryModeCreate,
		UpdateContext: resourceVcsDiscoveryModeUpdate,
		ReadContext:   resourceVcsDiscoveryModeRead,
		DeleteContext: resourceVcsDiscoveryModeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "Mcast", Description: "'Unicast': enable VCS Unicast election; 'Mcast': enable VCS Multicast election; 'Mixed': enable VCS Multicast/Unicastelection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVcsDiscoveryModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDiscoveryModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDiscoveryMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDiscoveryModeRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsDiscoveryModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDiscoveryModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDiscoveryMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDiscoveryModeRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsDiscoveryModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDiscoveryModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDiscoveryMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsDiscoveryModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDiscoveryModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDiscoveryMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsDiscoveryMode(d *schema.ResourceData) edpt.VcsDiscoveryMode {
	var ret edpt.VcsDiscoveryMode
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
