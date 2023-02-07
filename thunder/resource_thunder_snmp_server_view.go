package thunder

//Thunder resource SnmpServerView

import (
	"context"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerView() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerViewCreate,
		UpdateContext: resourceSnmpServerViewUpdate,
		ReadContext:   resourceSnmpServerViewRead,
		DeleteContext: resourceSnmpServerViewDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"viewname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerViewCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerView (Inside resourceSnmpServerViewCreate) ")
		name2 := d.Get("oid").(string)
		name1 := d.Get("viewname").(string)
		data := dataToSnmpServerView(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerView --")
		d.SetId(name1 + "," + name2)
		err := go_thunder.PostSnmpServerView(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerViewRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerViewRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerView (Inside resourceSnmpServerViewRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerView(client.Token, name1, name2, client.Host)
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

func resourceSnmpServerViewUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerView   (Inside resourceSnmpServerViewUpdate) ")
		data := dataToSnmpServerView(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerView ")
		err := go_thunder.PutSnmpServerView(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerViewRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerViewDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerViewDelete) " + name1)
		err := go_thunder.DeleteSnmpServerView(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSnmpServerView(d *schema.ResourceData) go_thunder.SnmpServerView {
	var vc go_thunder.SnmpServerView
	var c go_thunder.SnmpServerViewInstance
	c.Type = d.Get("type").(string)
	c.Oid = d.Get("oid").(string)
	c.Mask = d.Get("mask").(string)
	c.Viewname = d.Get("viewname").(string)

	vc.Type = c
	return vc
}
