package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseActionL2Mac() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action_l2_mac`: Mac Address\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionL2MacCreate,
		UpdateContext: resourceSysUtStateNextStateCaseActionL2MacUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseActionL2MacRead,
		DeleteContext: resourceSysUtStateNextStateCaseActionL2MacDelete,

		Schema: map[string]*schema.Schema{
			"address_type": {
				Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Nat pool",
			},
			"src_dst": {
				Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Description: "Mac Address",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
			},
			"virtual_server": {
				Type: schema.TypeString, Optional: true, Description: "vip",
			},
			"case_number": {
				Type: schema.TypeString, Required: true, Description: "CaseNumber",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseActionL2MacCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL2MacCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL2Mac(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionL2MacRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionL2MacUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL2MacUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL2Mac(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionL2MacRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionL2MacDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL2MacDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL2Mac(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionL2MacRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL2MacRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL2Mac(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtStateNextStateCaseActionL2Mac(d *schema.ResourceData) edpt.SysUtStateNextStateCaseActionL2Mac {
	var ret edpt.SysUtStateNextStateCaseActionL2Mac
	ret.Inst.AddressType = d.Get("address_type").(string)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcDst = d.Get("src_dst").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VirtualServer = d.Get("virtual_server").(string)
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
