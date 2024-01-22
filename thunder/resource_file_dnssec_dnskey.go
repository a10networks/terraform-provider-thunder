





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileDnssecDnskey() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_dnssec_dnskey`: dnssec dnskey  file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileDnssecDnskeyCreate,
        UpdateContext: resourceFileDnssecDnskeyUpdate,
        ReadContext: resourceFileDnssecDnskeyRead,
        DeleteContext: resourceFileDnssecDnskeyDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "dnssec dnskey local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileDnssecDnskeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDnskeyCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDnskey(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDnssecDnskeyRead(ctx, d, meta)
    }
    return diags
}

func resourceFileDnssecDnskeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDnskeyUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDnskey(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDnssecDnskeyRead(ctx, d, meta)
    }
    return diags
}
func resourceFileDnssecDnskeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDnskeyDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDnskey(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileDnssecDnskeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDnskeyRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDnskey(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileDnssecDnskey(d *schema.ResourceData) edpt.FileDnssecDnskey {
    var ret edpt.FileDnssecDnskey
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

