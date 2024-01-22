package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSetRxtxDescSize() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_set_rxtx_desc_size`: Set the number of Descriptor size per queue\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSetRxtxDescSizeCreate,
		UpdateContext: resourceSystemSetRxtxDescSizeUpdate,
		ReadContext:   resourceSystemSetRxtxDescSizeRead,
		DeleteContext: resourceSystemSetRxtxDescSizeDelete,

		Schema: map[string]*schema.Schema{
			"port_index": {
				Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
			},
			"rxd_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set new rx-descriptor size",
			},
			"txd_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set new tx-descriptor size",
			},
		},
	}
}
func resourceSystemSetRxtxDescSizeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxDescSizeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxDescSize(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetRxtxDescSizeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSetRxtxDescSizeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxDescSizeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxDescSize(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetRxtxDescSizeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSetRxtxDescSizeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxDescSizeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxDescSize(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSetRxtxDescSizeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxDescSizeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxDescSize(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSetRxtxDescSize(d *schema.ResourceData) edpt.SystemSetRxtxDescSize {
	var ret edpt.SystemSetRxtxDescSize
	ret.Inst.PortIndex = d.Get("port_index").(int)
	ret.Inst.RxdSize = d.Get("rxd_size").(int)
	ret.Inst.TxdSize = d.Get("txd_size").(int)
	return ret
}
