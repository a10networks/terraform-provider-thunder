package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntitySecondaryMonTopk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitored_entity_secondary_mon_topk`: Display topk secondary Monitoring entities\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitoredEntitySecondaryMonTopkCreate,
		UpdateContext: resourceVisibilityMonitoredEntitySecondaryMonTopkUpdate,
		ReadContext:   resourceVisibilityMonitoredEntitySecondaryMonTopkRead,
		DeleteContext: resourceVisibilityMonitoredEntitySecondaryMonTopkDelete,

		Schema: map[string]*schema.Schema{
			"sources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitoredEntitySecondaryMonTopkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntitySecondaryMonTopkRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitoredEntitySecondaryMonTopkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntitySecondaryMonTopkRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitoredEntitySecondaryMonTopkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitoredEntitySecondaryMonTopkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkSources1926(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkSources1926 {

	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkSources1926
	return ret
}

func dataToEndpointVisibilityMonitoredEntitySecondaryMonTopk(d *schema.ResourceData) edpt.VisibilityMonitoredEntitySecondaryMonTopk {
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopk
	ret.Inst.Sources = getObjectVisibilityMonitoredEntitySecondaryMonTopkSources1926(d.Get("sources").([]interface{}))
	//omit uuid
	return ret
}
