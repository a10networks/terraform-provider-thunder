package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosAnomalyDropSecurityAttackLayer3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_anomaly_drop_security_attack_layer_3`: Disable (Enable) Security Attack Layer-3 scrubbing\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosAnomalyDropSecurityAttackLayer3Create,
		UpdateContext: resourceDdosAnomalyDropSecurityAttackLayer3Update,
		ReadContext:   resourceDdosAnomalyDropSecurityAttackLayer3Read,
		DeleteContext: resourceDdosAnomalyDropSecurityAttackLayer3Delete,

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
func resourceDdosAnomalyDropSecurityAttackLayer3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropSecurityAttackLayer3Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosAnomalyDropSecurityAttackLayer3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropSecurityAttackLayer3Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosAnomalyDropSecurityAttackLayer3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosAnomalyDropSecurityAttackLayer3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropSecurityAttackLayer3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropSecurityAttackLayer3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosAnomalyDropSecurityAttackLayer3(d *schema.ResourceData) edpt.DdosAnomalyDropSecurityAttackLayer3 {
	var ret edpt.DdosAnomalyDropSecurityAttackLayer3
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
