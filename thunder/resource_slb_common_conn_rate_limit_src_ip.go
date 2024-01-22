package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonConnRateLimitSrcIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_conn_rate_limit_src_ip`: Set connection limit based on source IP address\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonConnRateLimitSrcIpCreate,
		UpdateContext: resourceSlbCommonConnRateLimitSrcIpUpdate,
		ReadContext:   resourceSlbCommonConnRateLimitSrcIpRead,
		DeleteContext: resourceSlbCommonConnRateLimitSrcIpDelete,

		Schema: map[string]*schema.Schema{
			"disable_ipv6_support": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"exceed_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action if threshold exceeded",
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Description: "Set max connections per period",
			},
			"limit_period": {
				Type: schema.TypeString, Optional: true, Description: "'100': 100 ms; '1000': 1000 ms;",
			},
			"lock_out": {
				Type: schema.TypeInt, Optional: true, Description: "Set lockout period in seconds if threshold exceeded",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log if threshold exceeded",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': Set TCP connection rate limit; 'udp': Set UDP packet rate limit;",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold shared amongst all virtual ports",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonConnRateLimitSrcIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonConnRateLimitSrcIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonConnRateLimitSrcIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonConnRateLimitSrcIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonConnRateLimitSrcIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonConnRateLimitSrcIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonConnRateLimitSrcIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonConnRateLimitSrcIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonConnRateLimitSrcIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonConnRateLimitSrcIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonConnRateLimitSrcIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonConnRateLimitSrcIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonConnRateLimitSrcIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonConnRateLimitSrcIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonConnRateLimitSrcIp(d *schema.ResourceData) edpt.SlbCommonConnRateLimitSrcIp {
	var ret edpt.SlbCommonConnRateLimitSrcIp
	ret.Inst.DisableIpv6Support = d.Get("disable_ipv6_support").(int)
	ret.Inst.ExceedAction = d.Get("exceed_action").(int)
	ret.Inst.Limit = d.Get("limit").(int)
	ret.Inst.LimitPeriod = d.Get("limit_period").(string)
	ret.Inst.LockOut = d.Get("lock_out").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	//omit uuid
	return ret
}
