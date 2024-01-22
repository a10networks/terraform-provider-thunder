package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSetRxtxQueue() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_set_rxtx_queue`: Set the number of Queues per port\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSetRxtxQueueCreate,
		UpdateContext: resourceSystemSetRxtxQueueUpdate,
		ReadContext:   resourceSystemSetRxtxQueueRead,
		DeleteContext: resourceSystemSetRxtxQueueDelete,

		Schema: map[string]*schema.Schema{
			"port_index": {
				Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
			},
			"rxq_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set number of new rx queues",
			},
			"txq_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set number of new tx queues",
			},
		},
	}
}
func resourceSystemSetRxtxQueueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxQueueCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxQueue(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetRxtxQueueRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSetRxtxQueueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxQueueUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxQueue(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetRxtxQueueRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSetRxtxQueueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxQueueDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxQueue(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSetRxtxQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetRxtxQueueRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetRxtxQueue(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSetRxtxQueue(d *schema.ResourceData) edpt.SystemSetRxtxQueue {
	var ret edpt.SystemSetRxtxQueue
	ret.Inst.PortIndex = d.Get("port_index").(int)
	ret.Inst.RxqSize = d.Get("rxq_size").(int)
	ret.Inst.TxqSize = d.Get("txq_size").(int)
	return ret
}
