package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSvmSourceNat() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_svm_source_nat`: Configure SVM Source-NAT Pool\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSvmSourceNatCreate,
		UpdateContext: resourceSlbSvmSourceNatUpdate,
		ReadContext:   resourceSlbSvmSourceNatRead,
		DeleteContext: resourceSlbSvmSourceNatDelete,

		Schema: map[string]*schema.Schema{
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbSvmSourceNatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSvmSourceNatCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSvmSourceNat(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSvmSourceNatRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSvmSourceNatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSvmSourceNatUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSvmSourceNat(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSvmSourceNatRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSvmSourceNatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSvmSourceNatDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSvmSourceNat(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSvmSourceNatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSvmSourceNatRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSvmSourceNat(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbSvmSourceNat(d *schema.ResourceData) edpt.SlbSvmSourceNat {
	var ret edpt.SlbSvmSourceNat
	ret.Inst.Pool = d.Get("pool").(string)
	//omit uuid
	return ret
}
