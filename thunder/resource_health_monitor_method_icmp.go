package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_icmp`: ICMP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodIcmpCreate,
		UpdateContext: resourceHealthMonitorMethodIcmpUpdate,
		ReadContext:   resourceHealthMonitorMethodIcmpRead,
		DeleteContext: resourceHealthMonitorMethodIcmpDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "ICMP type",
			},
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify IPv4 address of destination behind monitored node",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address of destination behind monitored node",
			},
			"transparent": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply transparent mode",
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
func resourceHealthMonitorMethodIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodIcmp(d *schema.ResourceData) edpt.HealthMonitorMethodIcmp {
	var ret edpt.HealthMonitorMethodIcmp
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.Transparent = d.Get("transparent").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
