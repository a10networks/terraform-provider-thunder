package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemQueuingBuffer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_queuing_buffer`: Enable/Disable micro-burst traffic support\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemQueuingBufferCreate,
		UpdateContext: resourceSystemQueuingBufferUpdate,
		ReadContext:   resourceSystemQueuingBufferRead,
		DeleteContext: resourceSystemQueuingBufferDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable micro-burst traffic support",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemQueuingBufferCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQueuingBufferCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQueuingBuffer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemQueuingBufferRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemQueuingBufferUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQueuingBufferUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQueuingBuffer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemQueuingBufferRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemQueuingBufferDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQueuingBufferDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQueuingBuffer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemQueuingBufferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQueuingBufferRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQueuingBuffer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemQueuingBuffer(d *schema.ResourceData) edpt.SystemQueuingBuffer {
	var ret edpt.SystemQueuingBuffer
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
