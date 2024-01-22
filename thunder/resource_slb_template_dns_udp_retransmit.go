package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsUdpRetransmit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_udp_retransmit`: DNS backend UDP retransmit policy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsUdpRetransmitCreate,
		UpdateContext: resourceSlbTemplateDnsUdpRetransmitUpdate,
		ReadContext:   resourceSlbTemplateDnsUdpRetransmitRead,
		DeleteContext: resourceSlbTemplateDnsUdpRetransmitDelete,

		Schema: map[string]*schema.Schema{
			"max_trials": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Total number of times to try DNS query to server before closing client connection, default 3",
			},
			"retry_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))",
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
func resourceSlbTemplateDnsUdpRetransmitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsUdpRetransmitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsUdpRetransmit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsUdpRetransmitRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsUdpRetransmitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsUdpRetransmitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsUdpRetransmit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsUdpRetransmitRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsUdpRetransmitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsUdpRetransmitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsUdpRetransmit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsUdpRetransmitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsUdpRetransmitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsUdpRetransmit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsUdpRetransmit(d *schema.ResourceData) edpt.SlbTemplateDnsUdpRetransmit {
	var ret edpt.SlbTemplateDnsUdpRetransmit
	ret.Inst.MaxTrials = d.Get("max_trials").(int)
	ret.Inst.RetryInterval = d.Get("retry_interval").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
