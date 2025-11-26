package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentSflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_agent_sflow`: Configure DDoS detection agent sFlow settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionAgentSflowCreate,
		UpdateContext: resourceDdosDetectionAgentSflowUpdate,
		ReadContext:   resourceDdosDetectionAgentSflowRead,
		DeleteContext: resourceDdosDetectionAgentSflowDelete,

		Schema: map[string]*schema.Schema{
			"sflow_pkt_samples_collection": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable sflow packet samples collection(default); 'disable': Disable sflow packet samples collection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"agent_name": {
				Type: schema.TypeString, Required: true, Description: "AgentName",
			},
		},
	}
}
func resourceDdosDetectionAgentSflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentSflowRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionAgentSflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentSflowRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionAgentSflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionAgentSflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionAgentSflow(d *schema.ResourceData) edpt.DdosDetectionAgentSflow {
	var ret edpt.DdosDetectionAgentSflow
	ret.Inst.SflowPktSamplesCollection = d.Get("sflow_pkt_samples_collection").(string)
	//omit uuid
	ret.Inst.AgentName = d.Get("agent_name").(string)
	return ret
}
