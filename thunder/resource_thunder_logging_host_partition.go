package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoggingHostPartition() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceLoggingHostPartitionCreate,
        UpdateContext: resourceLoggingHostPartitionUpdate,
        ReadContext: resourceLoggingHostPartitionRead,
        DeleteContext: resourceLoggingHostPartitionDelete,
        Schema: map[string]*schema.Schema{
            "shared": {
                Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
                ValidateFunc: validation.IntBetween(0, 1),
            },
            "partition_name": {
                Type: schema.TypeString, Optional: true, Description: "Select partition name for logging",
                ValidateFunc: validation.StringLenBetween(1, 14),
            },
            "uuid": {
                Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
            },
        },
    }
}

func resourceLoggingHostPartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostPartitionCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostPartition(d)
        d.SetId(obj.GetId())
        err := PostEndpointLoggingHostPartition(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostPartitionRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostPartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostPartitionRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointLoggingHostPartition(client.Token, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceLoggingHostPartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostPartitionUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostPartition(d)
        err := PutEndpointLoggingHostPartition(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostPartitionRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostPartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostPartitionDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointLoggingHostPartition(client.Token, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointLoggingHostPartition(d *schema.ResourceData) EndpointLoggingHostPartition{
    var ret EndpointLoggingHostPartition
    ret.Inst.Shared = d.Get("shared").(int)
    ret.Inst.PartitionName = d.Get("partition_name").(string)
    //omit uuid
    return ret
}
