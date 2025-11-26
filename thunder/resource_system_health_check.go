package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_health_check`: Layer-2 path health monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemHealthCheckCreate,
		UpdateContext: resourceSystemHealthCheckUpdate,
		ReadContext:   resourceSystemHealthCheckRead,
		DeleteContext: resourceSystemHealthCheckDelete,

		Schema: map[string]*schema.Schema{
			"l2bfd_multiplier": {
				Type: schema.TypeInt, Optional: true, Description: "Multiplier value used to compute holddown (value used to multiply the interval (default: 4))",
			},
			"l2bfd_rx_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum receive interval capability (Milliseconds (default: 800))",
			},
			"l2bfd_tx_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Transmit interval between BFD packets",
			},
			"l2hm_hc_name": {
				Type: schema.TypeString, Required: true, Description: "Monitor Name",
			},
			"method_l2bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Method is l2bfd",
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
func resourceSystemHealthCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHealthCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHealthCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHealthCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemHealthCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHealthCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHealthCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHealthCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemHealthCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHealthCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHealthCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHealthCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHealthCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemHealthCheck(d *schema.ResourceData) edpt.SystemHealthCheck {
	var ret edpt.SystemHealthCheck
	ret.Inst.L2bfdMultiplier = d.Get("l2bfd_multiplier").(int)
	ret.Inst.L2bfdRxInterval = d.Get("l2bfd_rx_interval").(int)
	ret.Inst.L2bfdTxInterval = d.Get("l2bfd_tx_interval").(int)
	ret.Inst.L2hmHcName = d.Get("l2hm_hc_name").(string)
	ret.Inst.MethodL2bfd = d.Get("method_l2bfd").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
