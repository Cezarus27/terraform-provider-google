---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/storage/ObjectAccessControl.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Storage"
description: |-
  The ObjectAccessControls resources represent the Access Control Lists
  (ACLs) for objects within Google Cloud Storage.
---

# google_storage_object_access_control

The ObjectAccessControls resources represent the Access Control Lists
(ACLs) for objects within Google Cloud Storage. ACLs let you specify
who has access to your data and to what extent.

There are two roles that can be assigned to an entity:

READERs can get an object, though the acl property will not be revealed.
OWNERs are READERs, and they can get the acl property, update an object,
and call all objectAccessControls methods on the object. The owner of an
object is always an OWNER.
For more information, see Access Control, with the caveat that this API
uses READER and OWNER instead of READ and FULL_CONTROL.


To get more information about ObjectAccessControl, see:

* [API documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=storage_object_access_control_public_object&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Storage Object Access Control Public Object


```hcl
resource "google_storage_object_access_control" "public_rule" {
  object = google_storage_bucket_object.object.output_name
  bucket = google_storage_bucket.bucket.name
  role   = "READER"
  entity = "allUsers"
}

resource "google_storage_bucket" "bucket" {
  name     = "static-content-bucket"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "public-object"
  bucket = google_storage_bucket.bucket.name
  source = "../static/img/header-logo.png"
}
```

## Argument Reference

The following arguments are supported:


* `bucket` -
  (Required)
  The name of the bucket.

* `entity` -
  (Required)
  The entity holding the permission, in one of the following forms:
    * user-{{userId}}
    * user-{{email}} (such as "user-liz@example.com")
    * group-{{groupId}}
    * group-{{email}} (such as "group-example@googlegroups.com")
    * domain-{{domain}} (such as "domain-example.com")
    * project-team-{{projectId}}
    * allUsers
    * allAuthenticatedUsers

* `object` -
  (Required)
  The name of the object to apply the access control to.

* `role` -
  (Required)
  The access permission for the entity.
  Possible values are: `OWNER`, `READER`.




## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{bucket}}/{{object}}/{{entity}}`

* `domain` -
  The domain associated with the entity.

* `email` -
  The email address associated with the entity.

* `entity_id` -
  The ID for the entity

* `generation` -
  The content generation of the object, if applied to an object.

* `project_team` -
  The project team associated with the entity
  Structure is [documented below](#nested_project_team).


<a name="nested_project_team"></a>The `project_team` block contains:

* `project_number` -
  (Optional)
  The project team associated with the entity

* `team` -
  (Optional)
  The team.
  Possible values are: `editors`, `owners`, `viewers`.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ObjectAccessControl can be imported using any of these accepted formats:

* `{{bucket}}/{{object}}/{{entity}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ObjectAccessControl using one of the formats above. For example:

```tf
import {
  id = "{{bucket}}/{{object}}/{{entity}}"
  to = google_storage_object_access_control.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ObjectAccessControl can be imported using one of the formats above. For example:

```
$ terraform import google_storage_object_access_control.default {{bucket}}/{{object}}/{{entity}}
```
