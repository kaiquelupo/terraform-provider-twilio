package twilio

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	log "github.com/sirupsen/logrus"
)

func resourceTwilioWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceTwilioWorkspaceCreate,
		Read:   resourceTwilioWorkspaceRead,
		Update: resourceTwilioWorkspaceUpdate,
		Delete: resourceTwilioWorkspaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func flattenWorkspaceForCreate(d *schema.ResourceData) url.Values {
	v := make(url.Values)

	v.Add("FriendlyName", d.Get("friendly_name").(string))
	// TODO: pass through multi task enabled
	// https://www.twilio.com/docs/taskrouter/api/workspace#create-a-workspace-resource
	return v
}

func resourceTwilioWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioWorkspaceCreate")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	createParams := flattenWorkspaceForCreate(d)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.WorkspaceCreator.Create")

	workspace, err := client.WorkspaceCreator.Create(context, createParams)
	if err != nil {
		log.WithFields(
			log.Fields{
				"account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.WorkspaceCreator.Create failed")

		return err
	}
	d.SetId(workspace.Sid)
	d.Set("friendly_name", workspace.FriendlyName)
	d.Set("date_created", workspace.DateCreated)
	d.Set("date_updated", workspace.DateUpdated)
	d.Set("multi_task_enabled", workspace.MultiTaskEnabled)
	d.Set("sid", workspace.Sid)
	return nil
}

func resourceTwilioWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioWorkspaceRead")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.WorkspaceCreator.Get")

	workspace, err := client.WorkspaceCreator.Get(context, sid)
	if err != nil {
		log.WithFields(
			log.Fields{
				"account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.WorkspaceCreator.Get failed")

		return err
	}
	d.Set("friendly_name", workspace.FriendlyName)
	d.Set("date_created", workspace.DateCreated)
	d.Set("date_updated", workspace.DateUpdated)
	d.Set("multi_task_enabled", workspace.MultiTaskEnabled)
	d.Set("sid", workspace.Sid)
	return nil
}

func resourceTwilioWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {

	log.Debug("ENTER resourceTwilioWorkspaceUpdate")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()
	createParams := flattenWorkspaceForCreate(d)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.WorkspaceCreator.Update")

	workspace, err := client.WorkspaceCreator.Update(context, sid, createParams)

	if err != nil {
		log.WithFields(
			log.Fields{
				"account_sid": config.AccountSID,
			},
		).WithError(err).Error("client.WorkspaceCreator.Update failed")

		return err
	}
	d.Set("friendly_name", workspace.FriendlyName)
	d.Set("date_created", workspace.DateCreated)
	d.Set("date_updated", workspace.DateUpdated)
	d.Set("multi_task_enabled", workspace.MultiTaskEnabled)
	d.Set("sid", workspace.Sid)
	return nil
}

func resourceTwilioWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Debug("ENTER resourceTwilioWorkspaceDelete")

	client := meta.(*TerraformTwilioContext).client
	config := meta.(*TerraformTwilioContext).configuration
	context := context.TODO()

	sid := d.Id()

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("START client.WorkspaceCreator.Delete")

	err := client.WorkspaceCreator.Delete(context, sid)

	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("END client.WorkspaceCreator.Delete")
	if err != nil {
		return fmt.Errorf("Failed to delete Workspace: %s", err.Error())
	}
	return nil
}
