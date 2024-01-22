package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosEastWestProtection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_east_west_protection`: DDOS East-West Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosEastWestProtectionCreate,
		UpdateContext: resourceDdosEastWestProtectionUpdate,
		ReadContext:   resourceDdosEastWestProtectionRead,
		DeleteContext: resourceDdosEastWestProtectionDelete,

		Schema: map[string]*schema.Schema{
			"deployment_mode": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'L2-mode': Enable East-West Protection in Layer 2 mode.; 'L2-with-virtual-wire': Enable East-West Protection in Layer 2 mode with virtual-wire pairs.; 'L3-mode': Enable East-West Protection in Layer 3 mode.; 'disable': Disable East-West Protection.;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosEastWestProtectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEastWestProtectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEastWestProtection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEastWestProtectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosEastWestProtectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEastWestProtectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEastWestProtection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEastWestProtectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosEastWestProtectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEastWestProtectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEastWestProtection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosEastWestProtectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEastWestProtectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEastWestProtection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosEastWestProtection(d *schema.ResourceData) edpt.DdosEastWestProtection {
	var ret edpt.DdosEastWestProtection
	ret.Inst.DeploymentMode = d.Get("deployment_mode").(string)
	//omit uuid
	return ret
}
