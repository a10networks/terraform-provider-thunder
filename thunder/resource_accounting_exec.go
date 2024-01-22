package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccountingExec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_accounting_exec`: Configuration for EXEC <shell> accounting\n\n__PLACEHOLDER__",
		CreateContext: resourceAccountingExecCreate,
		UpdateContext: resourceAccountingExecUpdate,
		ReadContext:   resourceAccountingExecRead,
		DeleteContext: resourceAccountingExecDelete,

		Schema: map[string]*schema.Schema{
			"accounting_exec_method": {
				Type: schema.TypeString, Optional: true, Description: "'tacplus': Use TACACS+ servers for accounting; 'radius': Use radius servers for accounting;",
			},
			"accounting_exec_type": {
				Type: schema.TypeString, Optional: true, Description: "'start-stop': Record start and stop without waiting; 'stop-only': Record stop when service terminates;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAccountingExecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingExecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingExec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingExecRead(ctx, d, meta)
	}
	return diags
}

func resourceAccountingExecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingExecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingExec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingExecRead(ctx, d, meta)
	}
	return diags
}
func resourceAccountingExecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingExecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingExec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAccountingExecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingExecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingExec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAccountingExec(d *schema.ResourceData) edpt.AccountingExec {
	var ret edpt.AccountingExec
	ret.Inst.AccountingExecMethod = d.Get("accounting_exec_method").(string)
	ret.Inst.AccountingExecType = d.Get("accounting_exec_type").(string)
	//omit uuid
	return ret
}
