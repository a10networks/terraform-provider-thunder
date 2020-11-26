package thunder

//Thunder resource SnmpServerUser

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerUserCreate,
		Update: resourceSnmpServerUserUpdate,
		Read:   resourceSnmpServerUserRead,
		Delete: resourceSnmpServerUserDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_val": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encpasswd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passwd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_pw_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"v3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pw_encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerUserCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerUser (Inside resourceSnmpServerUserCreate) ")
		name1 := d.Get("username").(string)
		data := dataToSnmpServerUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerUser --")
		d.SetId(name1)
		go_thunder.PostSnmpServerUser(client.Token, data, client.Host)

		return resourceSnmpServerUserRead(d, meta)

	}
	return nil
}

func resourceSnmpServerUserRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerUser (Inside resourceSnmpServerUserRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerUser(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerUserUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerUser   (Inside resourceSnmpServerUserUpdate) ")
		data := dataToSnmpServerUser(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerUser ")
		go_thunder.PutSnmpServerUser(client.Token, name1, data, client.Host)

		return resourceSnmpServerUserRead(d, meta)

	}
	return nil
}

func resourceSnmpServerUserDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerUserDelete) " + name1)
		err := go_thunder.DeleteSnmpServerUser(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerUser(d *schema.ResourceData) go_thunder.SnmpServerUser {
	var vc go_thunder.SnmpServerUser
	var c go_thunder.SnmpServerUserInstance
	c.Username = d.Get("username").(string)
	c.AuthVal = d.Get("auth_val").(string)
	c.Group = d.Get("group").(string)
	c.Encpasswd = d.Get("encpasswd").(string)
	c.Passwd = d.Get("passwd").(string)
	c.PrivPwEncrypted = d.Get("priv_pw_encrypted").(string)
	c.V3 = d.Get("v3").(string)
	c.PwEncrypted = d.Get("pw_encrypted").(string)
	c.Priv = d.Get("priv").(string)

	vc.Username = c
	return vc
}
