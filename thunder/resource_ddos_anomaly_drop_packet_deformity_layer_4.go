package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosAnomalyDropPacketDeformityLayer4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_anomaly_drop_packet_deformity_layer_4`: Disable (Enable) Layer-4 Packet Deformities\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosAnomalyDropPacketDeformityLayer4Create,
		UpdateContext: resourceDdosAnomalyDropPacketDeformityLayer4Update,
		ReadContext:   resourceDdosAnomalyDropPacketDeformityLayer4Read,
		DeleteContext: resourceDdosAnomalyDropPacketDeformityLayer4Delete,

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
func resourceDdosAnomalyDropPacketDeformityLayer4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropPacketDeformityLayer4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosAnomalyDropPacketDeformityLayer4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropPacketDeformityLayer4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosAnomalyDropPacketDeformityLayer4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosAnomalyDropPacketDeformityLayer4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosAnomalyDropPacketDeformityLayer4(d *schema.ResourceData) edpt.DdosAnomalyDropPacketDeformityLayer4 {
	var ret edpt.DdosAnomalyDropPacketDeformityLayer4
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
