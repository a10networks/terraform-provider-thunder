package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingMgmtSvcAcl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_mgmt_svc_acl`: Poll MGMT Service ACL counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingMgmtSvcAclCreate,
		UpdateContext: resourceSflowPollingMgmtSvcAclUpdate,
		ReadContext:   resourceSflowPollingMgmtSvcAclRead,
		DeleteContext: resourceSflowPollingMgmtSvcAclDelete,

		Schema: map[string]*schema.Schema{
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling MGMT Service ACL counters; 'disable': Disable polling MGMT Service ACL counters;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingMgmtSvcAclCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingMgmtSvcAclCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingMgmtSvcAcl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingMgmtSvcAclRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingMgmtSvcAclUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingMgmtSvcAclUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingMgmtSvcAcl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingMgmtSvcAclRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingMgmtSvcAclDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingMgmtSvcAclDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingMgmtSvcAcl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingMgmtSvcAclRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingMgmtSvcAclRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingMgmtSvcAcl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingMgmtSvcAcl(d *schema.ResourceData) edpt.SflowPollingMgmtSvcAcl {
	var ret edpt.SflowPollingMgmtSvcAcl
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
