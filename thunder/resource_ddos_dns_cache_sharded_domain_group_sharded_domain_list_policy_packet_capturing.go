package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_sharded_domain_group_sharded_domain_list_policy_packet_capturing`: Capture packet trace of specified root zones when doing zone transfer\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingCreate,
		UpdateContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingUpdate,
		ReadContext:   resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRead,
		DeleteContext: resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingDelete,

		Schema: map[string]*schema.Schema{
			"root_zone_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"root_zone": {
							Type: schema.TypeString, Optional: true, Description: "Specify root zone to be captured",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Capture-config name",
						},
						"capture_mode": {
							Type: schema.TypeString, Optional: true, Description: "'regular': Capture packet anyway; 'capture-on-failure': Capture packet if last XFR was failed;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList(d []interface{}) []edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing(d *schema.ResourceData) edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing {
	var ret edpt.DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing
	ret.Inst.RootZoneList = getSliceDdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList(d.Get("root_zone_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
