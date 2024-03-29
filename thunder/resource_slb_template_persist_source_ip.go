package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePersistSourceIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_persist_source_ip`: Source IP persistence\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePersistSourceIpCreate,
		UpdateContext: resourceSlbTemplatePersistSourceIpUpdate,
		ReadContext:   resourceSlbTemplatePersistSourceIpRead,
		DeleteContext: resourceSlbTemplatePersistSourceIpDelete,

		Schema: map[string]*schema.Schema{
			"dont_honor_conn_rules": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not observe connection rate rules",
			},
			"enforce_higher_priority": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce to use high priority node if available",
			},
			"hash_persist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use hash value of source IP address",
			},
			"incl_dst_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include destination IP on the persist",
			},
			"incl_sport": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include source port on the persist",
			},
			"match_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persistence type",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Source IP persistence template name",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask",
			},
			"netmask6": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask",
			},
			"primary_port": {
				Type: schema.TypeInt, Optional: true, Description: "Primary port to create the persist session",
			},
			"scan_all_members": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist with SCAN of all members",
			},
			"server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist to the same server, default is port",
			},
			"service_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist within the same service group",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Persistence timeout (in minutes)",
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
func resourceSlbTemplatePersistSourceIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSourceIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSourceIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistSourceIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePersistSourceIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSourceIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSourceIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistSourceIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePersistSourceIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSourceIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSourceIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePersistSourceIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSourceIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSourceIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePersistSourceIp(d *schema.ResourceData) edpt.SlbTemplatePersistSourceIp {
	var ret edpt.SlbTemplatePersistSourceIp
	ret.Inst.DontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	ret.Inst.EnforceHigherPriority = d.Get("enforce_higher_priority").(int)
	ret.Inst.HashPersist = d.Get("hash_persist").(int)
	ret.Inst.InclDstIp = d.Get("incl_dst_ip").(int)
	ret.Inst.InclSport = d.Get("incl_sport").(int)
	ret.Inst.MatchType = d.Get("match_type").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.Netmask6 = d.Get("netmask6").(int)
	ret.Inst.PrimaryPort = d.Get("primary_port").(int)
	ret.Inst.ScanAllMembers = d.Get("scan_all_members").(int)
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.ServiceGroup = d.Get("service_group").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
