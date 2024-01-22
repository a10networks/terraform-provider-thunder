package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMergeModeAddSlb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_merge_mode_add_slb`: Control block-merge mode behavior for slb objects\n\n__PLACEHOLDER__",
		CreateContext: resourceMergeModeAddSlbCreate,
		UpdateContext: resourceMergeModeAddSlbUpdate,
		ReadContext:   resourceMergeModeAddSlbRead,
		DeleteContext: resourceMergeModeAddSlbDelete,

		Schema: map[string]*schema.Schema{
			"member": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Control block-merge behavior for slb service-group member",
			},
			"server_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Control block-merge behavior for slb server port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_server_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Control block-merge behavior for slb virtual-server port",
			},
		},
	}
}
func resourceMergeModeAddSlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMergeModeAddSlbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMergeModeAddSlb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMergeModeAddSlbRead(ctx, d, meta)
	}
	return diags
}

func resourceMergeModeAddSlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMergeModeAddSlbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMergeModeAddSlb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMergeModeAddSlbRead(ctx, d, meta)
	}
	return diags
}
func resourceMergeModeAddSlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMergeModeAddSlbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMergeModeAddSlb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceMergeModeAddSlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMergeModeAddSlbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMergeModeAddSlb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointMergeModeAddSlb(d *schema.ResourceData) edpt.MergeModeAddSlb {
	var ret edpt.MergeModeAddSlb
	ret.Inst.Member = d.Get("member").(int)
	ret.Inst.ServerPort = d.Get("server_port").(int)
	//omit uuid
	ret.Inst.VirtualServerPort = d.Get("virtual_server_port").(int)
	return ret
}
