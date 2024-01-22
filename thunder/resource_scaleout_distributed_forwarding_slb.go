package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDistributedForwardingSlb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_distributed_forwarding_slb`: Enable Scaleout distributed-forwarding for slb\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutDistributedForwardingSlbCreate,
		UpdateContext: resourceScaleoutDistributedForwardingSlbUpdate,
		ReadContext:   resourceScaleoutDistributedForwardingSlbRead,
		DeleteContext: resourceScaleoutDistributedForwardingSlbDelete,

		Schema: map[string]*schema.Schema{
			"slb_value": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable SLB; 'disable': Disable SLB;",
			},
			"tcp_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "configure packet threshold value to offload TCP sessions(default 5)",
			},
			"udp_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "configure packet threshold value to offload UDP sessions(default 5)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutDistributedForwardingSlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingSlbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingSlb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingSlbRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutDistributedForwardingSlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingSlbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingSlb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingSlbRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutDistributedForwardingSlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingSlbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingSlb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutDistributedForwardingSlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingSlbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingSlb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutDistributedForwardingSlb(d *schema.ResourceData) edpt.ScaleoutDistributedForwardingSlb {
	var ret edpt.ScaleoutDistributedForwardingSlb
	ret.Inst.SlbValue = d.Get("slb_value").(string)
	ret.Inst.TcpThreshold = d.Get("tcp_threshold").(int)
	ret.Inst.UdpThreshold = d.Get("udp_threshold").(int)
	//omit uuid
	return ret
}
