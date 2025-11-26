package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceUsageVisibility() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_usage_visibility`: Configure System Resource Usage for visibility\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceUsageVisibilityCreate,
		UpdateContext: resourceSystemResourceUsageVisibilityUpdate,
		ReadContext:   resourceSystemResourceUsageVisibilityRead,
		DeleteContext: resourceSystemResourceUsageVisibilityDelete,

		Schema: map[string]*schema.Schema{
			"monitored_entity_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total number of monitored entities for visibility",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemResourceUsageVisibilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageVisibilityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsageVisibility(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceUsageVisibilityRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceUsageVisibilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageVisibilityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsageVisibility(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceUsageVisibilityRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceUsageVisibilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageVisibilityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsageVisibility(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceUsageVisibilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageVisibilityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsageVisibility(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemResourceUsageVisibility(d *schema.ResourceData) edpt.SystemResourceUsageVisibility {
	var ret edpt.SystemResourceUsageVisibility
	ret.Inst.MonitoredEntityCount = d.Get("monitored_entity_count").(int)
	//omit uuid
	return ret
}
