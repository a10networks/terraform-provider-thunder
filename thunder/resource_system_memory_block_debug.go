package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMemoryBlockDebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_memory_block_debug`: Set memory block debug\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMemoryBlockDebugCreate,
		UpdateContext: resourceSystemMemoryBlockDebugUpdate,
		ReadContext:   resourceSystemMemoryBlockDebugRead,
		DeleteContext: resourceSystemMemoryBlockDebugDelete,

		Schema: map[string]*schema.Schema{
			"assert_block": {
				Type: schema.TypeInt, Optional: true, Default: 65536, Description: "Over size block allocation (Assert memory block over size (default: 65536))",
			},
			"first_blk": {
				Type: schema.TypeInt, Optional: true, Default: 8192, Description: "First memory block ascending order (default: 8192) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
			},
			"fourth_blk": {
				Type: schema.TypeInt, Optional: true, Default: 65536, Description: "Fourth memory block (default: 65536) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
			},
			"pktdump_block": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable pktdump Oversize block request packet",
			},
			"second_blk": {
				Type: schema.TypeInt, Optional: true, Default: 16384, Description: "Second memory block (default: 16384) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
			},
			"third_blk": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Third memory block (default: 32768) (Memory blocks 32,64,128,256,512,1K,2K,4K,8K,16K,32K,64K)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemMemoryBlockDebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryBlockDebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryBlockDebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMemoryBlockDebugRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMemoryBlockDebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryBlockDebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryBlockDebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMemoryBlockDebugRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMemoryBlockDebugDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryBlockDebugDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryBlockDebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMemoryBlockDebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryBlockDebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryBlockDebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMemoryBlockDebug(d *schema.ResourceData) edpt.SystemMemoryBlockDebug {
	var ret edpt.SystemMemoryBlockDebug
	ret.Inst.AssertBlock = d.Get("assert_block").(int)
	ret.Inst.FirstBlk = d.Get("first_blk").(int)
	ret.Inst.FourthBlk = d.Get("fourth_blk").(int)
	ret.Inst.PktdumpBlock = d.Get("pktdump_block").(int)
	ret.Inst.SecondBlk = d.Get("second_blk").(int)
	ret.Inst.ThirdBlk = d.Get("third_blk").(int)
	//omit uuid
	return ret
}
