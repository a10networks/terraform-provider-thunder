package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_agent_snmp`: Configure DDoS detection agent snmp\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionAgentSnmpCreate,
		UpdateContext: resourceDdosDetectionAgentSnmpUpdate,
		ReadContext:   resourceDdosDetectionAgentSnmpRead,
		DeleteContext: resourceDdosDetectionAgentSnmpDelete,

		Schema: map[string]*schema.Schema{
			"community_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's community-string for SNMP",
			},
			"ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv4 address for SNMP",
			},
			"refresh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "refresh SNMP information",
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
func resourceDdosDetectionAgentSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionAgentSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentSnmpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionAgentSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionAgentSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionAgentSnmp(d *schema.ResourceData) edpt.DdosDetectionAgentSnmp {
	var ret edpt.DdosDetectionAgentSnmp
	ret.Inst.CommunityString = d.Get("community_string").(string)
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Refresh = d.Get("refresh").(int)
	//omit uuid
	ret.Inst.AgentName = d.Get("agent_name").(string)
	return ret
}
