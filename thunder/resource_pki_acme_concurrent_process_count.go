package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiAcmeConcurrentProcessCount() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_acme_concurrent_process_count`: Configure the number of cuncurrent ACME processes\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiAcmeConcurrentProcessCountCreate,
		UpdateContext: resourcePkiAcmeConcurrentProcessCountUpdate,
		ReadContext:   resourcePkiAcmeConcurrentProcessCountRead,
		DeleteContext: resourcePkiAcmeConcurrentProcessCountDelete,

		Schema: map[string]*schema.Schema{
			"process_num": {
				Type: schema.TypeInt, Optional: true, Description: "Set the concurrent processes number, the default and the maximum number are dependent on the number of control CPU",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourcePkiAcmeConcurrentProcessCountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeConcurrentProcessCountCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeConcurrentProcessCount(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiAcmeConcurrentProcessCountRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiAcmeConcurrentProcessCountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeConcurrentProcessCountUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeConcurrentProcessCount(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiAcmeConcurrentProcessCountRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiAcmeConcurrentProcessCountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeConcurrentProcessCountDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeConcurrentProcessCount(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiAcmeConcurrentProcessCountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeConcurrentProcessCountRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeConcurrentProcessCount(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiAcmeConcurrentProcessCount(d *schema.ResourceData) edpt.PkiAcmeConcurrentProcessCount {
	var ret edpt.PkiAcmeConcurrentProcessCount
	ret.Inst.ProcessNum = d.Get("process_num").(int)
	//omit uuid
	return ret
}
