





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileWebServiceCertKey() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_web_service_cert_key`: Import web service's private key and certificate\n\n__PLACEHOLDER__",
        CreateContext: resourceFileWebServiceCertKeyCreate,
        UpdateContext: resourceFileWebServiceCertKeyUpdate,
        ReadContext: resourceFileWebServiceCertKeyRead,
        DeleteContext: resourceFileWebServiceCertKeyDelete,

        Schema: map[string]*schema.Schema{
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileWebServiceCertKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileWebServiceCertKeyCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileWebServiceCertKey(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileWebServiceCertKeyRead(ctx, d, meta)
    }
    return diags
}

func resourceFileWebServiceCertKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileWebServiceCertKeyUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileWebServiceCertKey(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileWebServiceCertKeyRead(ctx, d, meta)
    }
    return diags
}
func resourceFileWebServiceCertKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileWebServiceCertKeyDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileWebServiceCertKey(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileWebServiceCertKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileWebServiceCertKeyRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileWebServiceCertKey(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileWebServiceCertKey(d *schema.ResourceData) edpt.FileWebServiceCertKey {
    var ret edpt.FileWebServiceCertKey
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

