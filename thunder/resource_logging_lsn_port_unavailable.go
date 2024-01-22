package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLsnPortUnavailable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_lsn_port_unavailable`: Set LSN system log policy when port unavailable occur\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingLsnPortUnavailableCreate,
		UpdateContext: resourceLoggingLsnPortUnavailableUpdate,
		ReadContext:   resourceLoggingLsnPortUnavailableRead,
		DeleteContext: resourceLoggingLsnPortUnavailableDelete,

		Schema: map[string]*schema.Schema{
			"ip_based": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log LSN port unavailable based on NAT IP (Default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingLsnPortUnavailableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnPortUnavailableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnPortUnavailable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnPortUnavailableRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingLsnPortUnavailableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnPortUnavailableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnPortUnavailable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnPortUnavailableRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingLsnPortUnavailableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnPortUnavailableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnPortUnavailable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingLsnPortUnavailableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnPortUnavailableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnPortUnavailable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingLsnPortUnavailable(d *schema.ResourceData) edpt.LoggingLsnPortUnavailable {
	var ret edpt.LoggingLsnPortUnavailable
	ret.Inst.IpBased = d.Get("ip_based").(int)
	//omit uuid
	return ret
}
