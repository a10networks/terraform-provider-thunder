package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSystemDefaultLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_system_default_limit`: Specify DDOS system-default limits\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSystemDefaultLimitCreate,
		UpdateContext: resourceDdosSystemDefaultLimitUpdate,
		ReadContext:   resourceDdosSystemDefaultLimitRead,
		DeleteContext: resourceDdosSystemDefaultLimitDelete,

		Schema: map[string]*schema.Schema{
			"default_bit_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Default Kibit (kibibit / 1024-bit) rate limit",
			},
			"default_conn_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Default Connection limit",
			},
			"default_conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Default Connection rate limit",
			},
			"default_frag_pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Default Fragmented packet rate limit",
			},
			"default_over_limit_action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Silently Drop the new connection / new packet when it exceeds limit",
						},
					},
				},
			},
			"default_pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Default Packet rate limit",
			},
			"limit_type": {
				Type: schema.TypeString, Required: true, Description: "'dst-entry': dst-entry; 'dst-icmp': dst-icmp; 'dst-other': dst-other; 'dst-tcp': dst-tcp; 'dst-udp': dst-udp; 'src-entry': src-entry; 'src-icmp': src-icmp; 'src-other': src-other; 'src-tcp': src-tcp; 'src-udp': src-udp;",
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
func resourceDdosSystemDefaultLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefaultLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSystemDefaultLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSystemDefaultLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefaultLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSystemDefaultLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSystemDefaultLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefaultLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSystemDefaultLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefaultLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosSystemDefaultLimitDefaultOverLimitAction(d []interface{}) edpt.DdosSystemDefaultLimitDefaultOverLimitAction {

	count1 := len(d)
	var ret edpt.DdosSystemDefaultLimitDefaultOverLimitAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Drop = in["drop"].(int)
	}
	return ret
}

func dataToEndpointDdosSystemDefaultLimit(d *schema.ResourceData) edpt.DdosSystemDefaultLimit {
	var ret edpt.DdosSystemDefaultLimit
	ret.Inst.DefaultBitRateLimit = d.Get("default_bit_rate_limit").(int)
	ret.Inst.DefaultConnLimit = d.Get("default_conn_limit").(int)
	ret.Inst.DefaultConnRateLimit = d.Get("default_conn_rate_limit").(int)
	ret.Inst.DefaultFragPktRateLimit = d.Get("default_frag_pkt_rate_limit").(int)
	ret.Inst.DefaultOverLimitAction = getObjectDdosSystemDefaultLimitDefaultOverLimitAction(d.Get("default_over_limit_action").([]interface{}))
	ret.Inst.DefaultPktRateLimit = d.Get("default_pkt_rate_limit").(int)
	ret.Inst.LimitType = d.Get("limit_type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
