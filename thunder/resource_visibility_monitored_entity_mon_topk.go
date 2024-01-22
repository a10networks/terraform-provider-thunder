package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityMonTopk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitored_entity_mon_topk`: Display topk Monitoring entities\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitoredEntityMonTopkCreate,
		UpdateContext: resourceVisibilityMonitoredEntityMonTopkUpdate,
		ReadContext:   resourceVisibilityMonitoredEntityMonTopkRead,
		DeleteContext: resourceVisibilityMonitoredEntityMonTopkDelete,

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
func resourceVisibilityMonitoredEntityMonTopkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityMonTopkRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitoredEntityMonTopkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityMonTopkRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitoredEntityMonTopkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitoredEntityMonTopkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityMonitoredEntityMonTopkSources1925(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkSources1925 {

	var ret edpt.VisibilityMonitoredEntityMonTopkSources1925
	return ret
}

func dataToEndpointVisibilityMonitoredEntityMonTopk(d *schema.ResourceData) edpt.VisibilityMonitoredEntityMonTopk {
	var ret edpt.VisibilityMonitoredEntityMonTopk
	ret.Inst.Sources = getObjectVisibilityMonitoredEntityMonTopkSources1925(d.Get("sources").([]interface{}))
	//omit uuid
	return ret
}
