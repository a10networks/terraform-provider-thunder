package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbActiveRdt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_active_rdt`: Active RDT Options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbActiveRdtCreate,
		UpdateContext: resourceGslbActiveRdtUpdate,
		ReadContext:   resourceGslbActiveRdtRead,
		DeleteContext: resourceGslbActiveRdtDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Specify Query Domain (Specify Domain Name)",
			},
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Using ICMP",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify Query Interval, unit: second, default is 1",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify local port to send probe packet, default is 0 (no port)",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify Retry Count, default is 3",
			},
			"sleep": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify Sleep Time when query fail, unit: second, default is 3",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 3000, Description: "Specify Query Timeout, unit: msec, default is 3000",
			},
			"track": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Specify Tracking Time, unit: second, default is 60",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbActiveRdtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbActiveRdtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbActiveRdt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbActiveRdtRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbActiveRdtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbActiveRdtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbActiveRdt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbActiveRdtRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbActiveRdtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbActiveRdtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbActiveRdt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbActiveRdtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbActiveRdtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbActiveRdt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbActiveRdt(d *schema.ResourceData) edpt.GslbActiveRdt {
	var ret edpt.GslbActiveRdt
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.Sleep = d.Get("sleep").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.Track = d.Get("track").(int)
	//omit uuid
	return ret
}
