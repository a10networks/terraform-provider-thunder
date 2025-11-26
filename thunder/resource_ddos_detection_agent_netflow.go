package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentNetflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_agent_netflow`: Configure DDoS detection agent Netflow settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionAgentNetflowCreate,
		UpdateContext: resourceDdosDetectionAgentNetflowUpdate,
		ReadContext:   resourceDdosDetectionAgentNetflowRead,
		DeleteContext: resourceDdosDetectionAgentNetflowDelete,

		Schema: map[string]*schema.Schema{
			"active_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow active timeout (seconds)",
			},
			"inactive_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow inactive timeout (seconds)",
			},
			"netflow_samples_collection": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Netflow flow samples collection(default); 'disable': Disable Netflow flow samples collection;",
			},
			"netflow_sampling_rate": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure agent's netflow sampling rate",
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
func resourceDdosDetectionAgentNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentNetflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentNetflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentNetflowRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionAgentNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentNetflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentNetflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentNetflowRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionAgentNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentNetflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentNetflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionAgentNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentNetflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentNetflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionAgentNetflow(d *schema.ResourceData) edpt.DdosDetectionAgentNetflow {
	var ret edpt.DdosDetectionAgentNetflow
	ret.Inst.ActiveTimeout = d.Get("active_timeout").(int)
	ret.Inst.InactiveTimeout = d.Get("inactive_timeout").(int)
	ret.Inst.NetflowSamplesCollection = d.Get("netflow_samples_collection").(string)
	ret.Inst.NetflowSamplingRate = d.Get("netflow_sampling_rate").(int)
	//omit uuid
	ret.Inst.AgentName = d.Get("agent_name").(string)
	return ret
}
