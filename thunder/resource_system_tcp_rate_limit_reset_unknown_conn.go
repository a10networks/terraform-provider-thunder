package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpRateLimitResetUnknownConn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_tcp_rate_limit_reset_unknown_conn`: Configure rate limit for reset-unknown-conn\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTcpRateLimitResetUnknownConnCreate,
		UpdateContext: resourceSystemTcpRateLimitResetUnknownConnUpdate,
		ReadContext:   resourceSystemTcpRateLimitResetUnknownConnRead,
		DeleteContext: resourceSystemTcpRateLimitResetUnknownConnDelete,

		Schema: map[string]*schema.Schema{
			"log_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
			},
			"pkt_rate_for_reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTcpRateLimitResetUnknownConnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRateLimitResetUnknownConnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpRateLimitResetUnknownConn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTcpRateLimitResetUnknownConnRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTcpRateLimitResetUnknownConnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRateLimitResetUnknownConnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpRateLimitResetUnknownConn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTcpRateLimitResetUnknownConnRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTcpRateLimitResetUnknownConnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRateLimitResetUnknownConnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpRateLimitResetUnknownConn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTcpRateLimitResetUnknownConnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRateLimitResetUnknownConnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpRateLimitResetUnknownConn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTcpRateLimitResetUnknownConn(d *schema.ResourceData) edpt.SystemTcpRateLimitResetUnknownConn {
	var ret edpt.SystemTcpRateLimitResetUnknownConn
	ret.Inst.LogForResetUnknownConn = d.Get("log_for_reset_unknown_conn").(int)
	ret.Inst.PktRateForResetUnknownConn = d.Get("pkt_rate_for_reset_unknown_conn").(int)
	//omit uuid
	return ret
}
