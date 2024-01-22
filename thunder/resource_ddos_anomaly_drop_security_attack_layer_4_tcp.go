package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosAnomalyDropSecurityAttackLayer4Tcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_anomaly_drop_security_attack_layer_4_tcp`: Disable (Enable) Security Attack Layer-4 TCP scrubbing\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosAnomalyDropSecurityAttackLayer4TcpCreate,
		UpdateContext: resourceDdosAnomalyDropSecurityAttackLayer4TcpUpdate,
		ReadContext:   resourceDdosAnomalyDropSecurityAttackLayer4TcpRead,
		DeleteContext: resourceDdosAnomalyDropSecurityAttackLayer4TcpDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "capture-config name (Can only configure when drop-disabled)",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the anomaly event (Can only configure when drop-disabled)",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosAnomalyDropSecurityAttackLayer4TcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer4TcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer4Tcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropSecurityAttackLayer4TcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosAnomalyDropSecurityAttackLayer4TcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer4TcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer4Tcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropSecurityAttackLayer4TcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosAnomalyDropSecurityAttackLayer4TcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer4TcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer4Tcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosAnomalyDropSecurityAttackLayer4TcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer4TcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer4Tcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosAnomalyDropSecurityAttackLayer4Tcp(d *schema.ResourceData) edpt.DdosAnomalyDropSecurityAttackLayer4Tcp {
	var ret edpt.DdosAnomalyDropSecurityAttackLayer4Tcp
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
