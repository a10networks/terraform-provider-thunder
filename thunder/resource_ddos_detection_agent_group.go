package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_agent_group`: Configure DDoS detection agent group\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionAgentGroupCreate,
		UpdateContext: resourceDdosDetectionAgentGroupUpdate,
		ReadContext:   resourceDdosDetectionAgentGroupRead,
		DeleteContext: resourceDdosDetectionAgentGroupDelete,

		Schema: map[string]*schema.Schema{
			"agent": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_name": {
							Type: schema.TypeString, Optional: true, Description: "detection agent name",
						},
					},
				},
			},
			"agent_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name for the agent-group",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionAgentGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionAgentGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionAgentGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionAgentGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDetectionAgentGroupAgent(d []interface{}) []edpt.DdosDetectionAgentGroupAgent {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentGroupAgent, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentGroupAgent
		oi.AgentName = in["agent_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionAgentGroup(d *schema.ResourceData) edpt.DdosDetectionAgentGroup {
	var ret edpt.DdosDetectionAgentGroup
	ret.Inst.Agent = getSliceDdosDetectionAgentGroupAgent(d.Get("agent").([]interface{}))
	ret.Inst.AgentGroupName = d.Get("agent_group_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
