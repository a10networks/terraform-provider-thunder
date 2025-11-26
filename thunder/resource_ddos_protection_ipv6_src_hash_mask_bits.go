package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectionIpv6SrcHashMaskBits() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_protection_ipv6_src_hash_mask_bits`: IPv6 src hashing configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosProtectionIpv6SrcHashMaskBitsCreate,
		UpdateContext: resourceDdosProtectionIpv6SrcHashMaskBitsUpdate,
		ReadContext:   resourceDdosProtectionIpv6SrcHashMaskBitsRead,
		DeleteContext: resourceDdosProtectionIpv6SrcHashMaskBitsDelete,

		Schema: map[string]*schema.Schema{
			"mask_bit_offset_1": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
			},
			"mask_bit_offset_2": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
			},
			"mask_bit_offset_3": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
			},
			"mask_bit_offset_4": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
			},
			"mask_bit_offset_5": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosProtectionIpv6SrcHashMaskBitsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionIpv6SrcHashMaskBitsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionIpv6SrcHashMaskBits(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionIpv6SrcHashMaskBitsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosProtectionIpv6SrcHashMaskBitsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionIpv6SrcHashMaskBitsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionIpv6SrcHashMaskBits(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionIpv6SrcHashMaskBitsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosProtectionIpv6SrcHashMaskBitsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionIpv6SrcHashMaskBitsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionIpv6SrcHashMaskBits(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosProtectionIpv6SrcHashMaskBitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionIpv6SrcHashMaskBitsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionIpv6SrcHashMaskBits(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosProtectionIpv6SrcHashMaskBits(d *schema.ResourceData) edpt.DdosProtectionIpv6SrcHashMaskBits {
	var ret edpt.DdosProtectionIpv6SrcHashMaskBits
	ret.Inst.MaskBitOffset1 = d.Get("mask_bit_offset_1").(int)
	ret.Inst.MaskBitOffset2 = d.Get("mask_bit_offset_2").(int)
	ret.Inst.MaskBitOffset3 = d.Get("mask_bit_offset_3").(int)
	ret.Inst.MaskBitOffset4 = d.Get("mask_bit_offset_4").(int)
	ret.Inst.MaskBitOffset5 = d.Get("mask_bit_offset_5").(int)
	//omit uuid
	return ret
}
