





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceDeleteIps() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_delete_ips`: Delete ips related file\n\n__PLACEHOLDER__",
        CreateContext: resourceDeleteIpsCreate,
        UpdateContext: resourceDeleteIpsUpdate,
        ReadContext: resourceDeleteIpsRead,
        DeleteContext: resourceDeleteIpsDelete,

        Schema: map[string]*schema.Schema{
        "profile": {
        Type: schema.TypeString, Optional: true, Description: "Specify profile to be deleted",
        },
        "signature": {
        Type: schema.TypeInt, Optional: true, Description: "Specify custom signature sid to be deleted",
        },
        },
    }
}
func resourceDeleteIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDeleteIpsCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDeleteIps(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDeleteIpsRead(ctx, d, meta)
    }
    return diags
}

func resourceDeleteIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDeleteIpsUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDeleteIps(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDeleteIpsRead(ctx, d, meta)
    }
    return diags
}
func resourceDeleteIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDeleteIpsDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDeleteIps(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceDeleteIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDeleteIpsRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDeleteIps(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointDeleteIps(d *schema.ResourceData) edpt.DeleteIps {
    var ret edpt.DeleteIps
	ret.Inst.Profile = d.Get("profile").(string)
	ret.Inst.Signature = d.Get("signature").(int)
    return ret
}

