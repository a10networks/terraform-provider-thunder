package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpsec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipsec`: Configure Crypto Cores for IPsec processing\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpsecCreate,
		UpdateContext: resourceSystemIpsecUpdate,
		ReadContext:   resourceSystemIpsecRead,
		DeleteContext: resourceSystemIpsecDelete,

		Schema: map[string]*schema.Schema{
			"crypto_core": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto cores assigned for IPsec processing",
			},
			"crypto_mem": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Crypto memory percentage assigned for IPsec processing (rounded to increments of 10)",
			},
			"fpga_decrypt": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable FPGA decryption offload; 'disable': Disable FPGA decryption offload;",
						},
					},
				},
			},
			"packet_round_robin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet round robin for IPsec packets",
			},
			"qat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HW assisted QAT SSL module",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemIpsecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpsecRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpsecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpsecRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpsecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpsecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpsecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpsec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemIpsecFpgaDecrypt1592(d []interface{}) edpt.SystemIpsecFpgaDecrypt1592 {

	count1 := len(d)
	var ret edpt.SystemIpsecFpgaDecrypt1592
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
	}
	return ret
}

func dataToEndpointSystemIpsec(d *schema.ResourceData) edpt.SystemIpsec {
	var ret edpt.SystemIpsec
	ret.Inst.CryptoCore = d.Get("crypto_core").(int)
	ret.Inst.CryptoMem = d.Get("crypto_mem").(int)
	ret.Inst.FpgaDecrypt = getObjectSystemIpsecFpgaDecrypt1592(d.Get("fpga_decrypt").([]interface{}))
	ret.Inst.PacketRoundRobin = d.Get("packet_round_robin").(int)
	ret.Inst.Qat = d.Get("qat").(int)
	//omit uuid
	return ret
}
