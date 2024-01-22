package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_udp`: L4 UDP switch config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateUdpCreate,
		UpdateContext: resourceSlbTemplateUdpUpdate,
		ReadContext:   resourceSlbTemplateUdpRead,
		DeleteContext: resourceSlbTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Description: "short age (in sec), default is 31",
			},
			"avp": {
				Type: schema.TypeString, Optional: true, Description: "'4': NAS-IP-address; '8': Framed-IP-Address;",
			},
			"disable_clear_session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable immediate clearing of session",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)",
			},
			"immediate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Immediate Removal after Transaction",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Fast UDP Template Name",
			},
			"qos": {
				Type: schema.TypeInt, Optional: true, Description: "QOS level (number)",
			},
			"radius_lb_method_hash_type": {
				Type: schema.TypeString, Optional: true, Description: "'ip': IP-Hash; 'ipv6': IPv6-Hash;",
			},
			"re_select_if_server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-select another server if service port is down",
			},
			"short": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Short lived session",
			},
			"stateless_conn_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v6avp": {
				Type: schema.TypeString, Optional: true, Description: "'168': Framed-IPv6-Address; '97': Framed-IPv6-PrefixFramed-IPv6-Prefix;",
			},
		},
	}
}
func resourceSlbTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateUdp(d *schema.ResourceData) edpt.SlbTemplateUdp {
	var ret edpt.SlbTemplateUdp
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.Avp = d.Get("avp").(string)
	ret.Inst.DisableClearSession = d.Get("disable_clear_session").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.Immediate = d.Get("immediate").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Qos = d.Get("qos").(int)
	ret.Inst.RadiusLbMethodHashType = d.Get("radius_lb_method_hash_type").(string)
	ret.Inst.ReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	ret.Inst.Short = d.Get("short").(int)
	ret.Inst.StatelessConnTimeout = d.Get("stateless_conn_timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.V6avp = d.Get("v6avp").(string)
	return ret
}
