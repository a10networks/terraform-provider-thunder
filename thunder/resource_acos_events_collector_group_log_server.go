package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsCollectorGroupLogServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_collector_group_log_server`: Configure log server DNS Name/ip-address and optional port\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsCollectorGroupLogServerCreate,
		UpdateContext: resourceAcosEventsCollectorGroupLogServerUpdate,
		ReadContext:   resourceAcosEventsCollectorGroupLogServerRead,
		DeleteContext: resourceAcosEventsCollectorGroupLogServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Member name",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsCollectorGroupLogServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupLogServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroupLogServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsCollectorGroupLogServerRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsCollectorGroupLogServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupLogServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroupLogServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsCollectorGroupLogServerRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsCollectorGroupLogServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupLogServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroupLogServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsCollectorGroupLogServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupLogServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroupLogServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsCollectorGroupLogServer(d *schema.ResourceData) edpt.AcosEventsCollectorGroupLogServer {
	var ret edpt.AcosEventsCollectorGroupLogServer
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Port = d.Get("port").(int)
	//omit uuid
	return ret
}
