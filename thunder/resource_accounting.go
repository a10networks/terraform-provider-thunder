package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccounting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_accounting`: Configuration for accounting\n\n__PLACEHOLDER__",
		CreateContext: resourceAccountingCreate,
		UpdateContext: resourceAccountingUpdate,
		ReadContext:   resourceAccountingRead,
		DeleteContext: resourceAccountingDelete,

		Schema: map[string]*schema.Schema{
			"commands": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable level for commands accounting",
			},
			"debug": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the debug level for accounting (Debug level for command accounting. bitwise OR of the following: 1(common), 2(packet),4(packet detail), 8(md5))",
			},
			"exec": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accounting_exec_type": {
							Type: schema.TypeString, Optional: true, Description: "'start-stop': Record start and stop without waiting; 'stop-only': Record stop when service terminates;",
						},
						"accounting_exec_method": {
							Type: schema.TypeString, Optional: true, Description: "'tacplus': Use TACACS+ servers for accounting; 'radius': Use radius servers for accounting;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"stop_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Record stop when service terminates",
			},
			"tacplus": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use TACACS+ servers for accounting",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAccountingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccounting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingRead(ctx, d, meta)
	}
	return diags
}

func resourceAccountingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccounting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingRead(ctx, d, meta)
	}
	return diags
}
func resourceAccountingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccounting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAccountingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccounting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAccountingExec40(d []interface{}) edpt.AccountingExec40 {

	count1 := len(d)
	var ret edpt.AccountingExec40
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AccountingExecType = in["accounting_exec_type"].(string)
		ret.AccountingExecMethod = in["accounting_exec_method"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointAccounting(d *schema.ResourceData) edpt.Accounting {
	var ret edpt.Accounting
	ret.Inst.Commands = d.Get("commands").(int)
	ret.Inst.Debug = d.Get("debug").(int)
	ret.Inst.Exec = getObjectAccountingExec40(d.Get("exec").([]interface{}))
	ret.Inst.StopOnly = d.Get("stop_only").(int)
	ret.Inst.Tacplus = d.Get("tacplus").(int)
	//omit uuid
	return ret
}
