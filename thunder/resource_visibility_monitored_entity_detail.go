package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityDetail() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitored_entity_detail`: Display Monitoring entity detail\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitoredEntityDetailCreate,
		UpdateContext: resourceVisibilityMonitoredEntityDetailUpdate,
		ReadContext:   resourceVisibilityMonitoredEntityDetailRead,
		DeleteContext: resourceVisibilityMonitoredEntityDetailDelete,

		Schema: map[string]*schema.Schema{
			"debug": {
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
func resourceVisibilityMonitoredEntityDetailCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetail(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityDetailRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitoredEntityDetailUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetail(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityDetailRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitoredEntityDetailDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetail(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitoredEntityDetailRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetail(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityMonitoredEntityDetailDebug1924(d []interface{}) edpt.VisibilityMonitoredEntityDetailDebug1924 {

	var ret edpt.VisibilityMonitoredEntityDetailDebug1924
	return ret
}

func dataToEndpointVisibilityMonitoredEntityDetail(d *schema.ResourceData) edpt.VisibilityMonitoredEntityDetail {
	var ret edpt.VisibilityMonitoredEntityDetail
	ret.Inst.Debug = getObjectVisibilityMonitoredEntityDetailDebug1924(d.Get("debug").([]interface{}))
	//omit uuid
	return ret
}
