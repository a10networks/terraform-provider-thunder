package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemLinkCapability() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_link_capability`: To enable/disable link capabilities\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemLinkCapabilityCreate,
		UpdateContext: resourceSystemLinkCapabilityUpdate,
		ReadContext:   resourceSystemLinkCapabilityRead,
		DeleteContext: resourceSystemLinkCapabilityDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable link capabilities",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemLinkCapabilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLinkCapabilityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLinkCapability(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemLinkCapabilityRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemLinkCapabilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLinkCapabilityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLinkCapability(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemLinkCapabilityRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemLinkCapabilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLinkCapabilityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLinkCapability(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemLinkCapabilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemLinkCapabilityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemLinkCapability(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemLinkCapability(d *schema.ResourceData) edpt.SystemLinkCapability {
	var ret edpt.SystemLinkCapability
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
