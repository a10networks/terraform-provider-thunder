package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslHwMemory() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ssl_hw_memory`: Configure the size of SSL hardward memory blocks\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSslHwMemoryCreate,
		UpdateContext: resourceSystemSslHwMemoryUpdate,
		ReadContext:   resourceSystemSslHwMemoryRead,
		DeleteContext: resourceSystemSslHwMemoryDelete,

		Schema: map[string]*schema.Schema{
			"mem_block_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mem_block": {
							Type: schema.TypeString, Optional: true, Description: "'ssl_mem': Shared SSL memory; 'ssl_context': Cipher session context memory for SSL engine 1; 'ssl_context_2': Cipher session context memory for SSL engine 2; 'ssl_context_3': Cipher session context memory for SSL engine 3; 'ssl_context_4': Cipher session context memory for SSL engine 4; 'ssl_context_5': Cipher session context memory for SSL engine 5; 'ssl_context_6': Cipher session context memory for SSL engine 6; 'ssl_context_7': Cipher session context memory for SSL engine 7; 'ssl_context_8': Cipher session context memory for SSL engine 8; 'ssl_context_9': Cipher session context memory for SSL engine 9; 'ssl_context_10': Cipher session context memory for SSL engine 10; 'ssl_context_11': Cipher session context memory for SSL engine 11; 'ssl_context_12': Cipher session context memory for SSL engine 12; 'ssl_context_13': Cipher session context memory for SSL engine 13; 'ssl_context_14': Cipher session context memory for SSL engine 14; 'ssl_context_15': Cipher session context memory for SSL engine 15; 'ssl_context_16': Cipher session context memory for SSL engine 16;",
						},
						"size": {
							Type: schema.TypeInt, Optional: true, Description: "Size of the block",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSslHwMemoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslHwMemoryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslHwMemory(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslHwMemoryRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSslHwMemoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslHwMemoryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslHwMemory(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslHwMemoryRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSslHwMemoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslHwMemoryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslHwMemory(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSslHwMemoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslHwMemoryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslHwMemory(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemSslHwMemoryMemBlockCfg(d []interface{}) []edpt.SystemSslHwMemoryMemBlockCfg {

	count1 := len(d)
	ret := make([]edpt.SystemSslHwMemoryMemBlockCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSslHwMemoryMemBlockCfg
		oi.MemBlock = in["mem_block"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemSslHwMemory(d *schema.ResourceData) edpt.SystemSslHwMemory {
	var ret edpt.SystemSslHwMemory
	ret.Inst.MemBlockCfg = getSliceSystemSslHwMemoryMemBlockCfg(d.Get("mem_block_cfg").([]interface{}))
	//omit uuid
	return ret
}
