package thunder

//Thunder resource SnmpServerGroup

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerGroupCreate,
		UpdateContext: resourceSnmpServerGroupUpdate,
		ReadContext:   resourceSnmpServerGroupRead,
		DeleteContext: resourceSnmpServerGroupDelete,
		Schema: map[string]*schema.Schema{
			"read": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"groupname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"v3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerGroup (Inside resourceSnmpServerGroupCreate) ")
		name1 := d.Get("groupname").(string)
		data := dataToSnmpServerGroup(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerGroup --")
		d.SetId(name1)
		err := go_thunder.PostSnmpServerGroup(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerGroup (Inside resourceSnmpServerGroupRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerGroup(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceSnmpServerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerGroup   (Inside resourceSnmpServerGroupUpdate) ")
		data := dataToSnmpServerGroup(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerGroup ")
		err := go_thunder.PutSnmpServerGroup(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerGroupDelete) " + name1)
		err := go_thunder.DeleteSnmpServerGroup(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSnmpServerGroup(d *schema.ResourceData) go_thunder.SnmpServerGroup {
	var vc go_thunder.SnmpServerGroup
	var c go_thunder.SnmpServerGroupInstance
	c.Read = d.Get("read").(string)
	c.Groupname = d.Get("groupname").(string)
	c.V3 = d.Get("v3").(string)

	vc.Read = c
	return vc
}
