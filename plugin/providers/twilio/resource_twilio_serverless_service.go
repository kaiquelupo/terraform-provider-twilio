package twilio

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	log "github.com/sirupsen/logrus"
)

func resourceTwilioServerlessService() *schema.Resource {
	return &schema.Resource{
		Create: resourceTwilioServerlessServiceCreate,
		Read:   resourceTwilioServerlessServiceRead,
		Update: resourceTwilioServerlessServiceUpdate,
		Delete: resourceTwilioServerlessServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"include_credentials": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ui_editable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func flattenServerlessServiceForCreate(d *schema.ResourceData) url.Values {
	v := make(url.Values)

	v.Add("FriendlyName", d.Get("friendly_name").(string))
	v.Add("UniqueName", d.Get("unique_name").(string))
	v.Add("IncludeCredentials", d.Get("include_credentials").(string))
	v.Add("UiEditable", d.Get("ui_editable").(string))

	
	return v
}

func resourceTwilioServerlessServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioServerlessServiceCreate")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	createParams := flattenServerlessServiceForCreate(d)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.Serverless.Service")

	service, err := client.Serverless.Create(context, createParams)
	if err != nil {
		log.WithFields(
			log.Fields{
				"account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.Serverless.Service failed")

		return err
	}
	d.SetId(service.Sid)
	d.Set("friendly_name", service.FriendlyName)
	d.Set("unique_name", service.UniqueName)

	return nil
}

func resourceTwilioServerlessServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioServerlessServiceRead")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.Serverless.Service.Get")

	service, err := client.Serverless.Get(context, sid)
	if err != nil {
		log.WithFields(
			log.Fields{
				"parent_account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.Serverless.Service.Get failed")

		return err
	}
	d.Set("friendly_name", service.FriendlyName)
	d.Set("unique_name", service.UniqueName)
	return nil
}

func resourceTwilioServerlessServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioServerlessServiceUpdate")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()
	createParams := flattenServerlessServiceForCreate(d)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.Serverless.Update")

	service, err := client.Serverless.Update(context, sid, createParams)
	if err != nil {
		log.WithFields(
			log.Fields{
				"account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.Serverless.Update failed")

		return err
	}
	d.SetId(service.Sid)
	d.Set("friendly_name", service.FriendlyName)
	d.Set("unique_name", service.UniqueName)
	
	return nil
}

func resourceTwilioServerlessServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioServerlessServiceDelete")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
			"queue_sid":   sid,
		},
	).Debug("START client.Serverless.Delete")

	err := client.Serverless.Delete(context, sid)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
			"queue_sid":   sid,
		},
	).Debug("END client.Serverless.Delete")
	
	if err != nil {
		return fmt.Errorf("Failed to delete serverless service: %s", err.Error())
	}
	return nil
}
