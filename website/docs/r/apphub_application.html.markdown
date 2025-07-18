---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apphub/Application.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "App Hub"
description: |-
  Application is a functional grouping of Services and Workloads that helps achieve a desired end-to-end business functionality.
---

# google_apphub_application

Application is a functional grouping of Services and Workloads that helps achieve a desired end-to-end business functionality. Services and Workloads are owned by the Application.



<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=apphub_application_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Apphub Application Basic


```hcl
resource "google_apphub_application" "example" {
  location = "us-east1"
  application_id = "example-application"
  scope {
    type = "REGIONAL"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=apphub_application_global_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Apphub Application Global Basic


```hcl
resource "google_apphub_application" "example" {
  location = "global"
  application_id = "example-application"
  scope {
    type = "GLOBAL"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=apphub_application_full&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Apphub Application Full


```hcl
resource "google_apphub_application" "example2" {
  location = "us-east1"
  application_id = "example-application"
  display_name = "Application Full"
  scope {
    type = "REGIONAL"
  }
  description = "Application for testing"
  attributes {
    environment {
      type = "STAGING"
		}
		criticality {  
      type = "MISSION_CRITICAL"
		}
		business_owners {
		  display_name =  "Alice"
		  email        =  "alice@google.com"
		}
		developer_owners {
		  display_name =  "Bob"
		  email        =  "bob@google.com"
		}
		operator_owners {
		  display_name =  "Charlie"
		  email        =  "charlie@google.com"
		}
  }
}
```

## Argument Reference

The following arguments are supported:


* `scope` -
  (Required)
  Scope of an application.
  Structure is [documented below](#nested_scope).

* `location` -
  (Required)
  Part of `parent`. See documentation of `projectsId`.

* `application_id` -
  (Required)
  Required. The Application identifier.


* `display_name` -
  (Optional)
  Optional. User-defined name for the Application.

* `description` -
  (Optional)
  Optional. User-defined description of an Application.

* `attributes` -
  (Optional)
  Consumer provided attributes.
  Structure is [documented below](#nested_attributes).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_scope"></a>The `scope` block supports:

* `type` -
  (Required)
  Required. Scope Type. 
   Possible values:
  REGIONAL
  GLOBAL
  Possible values are: `REGIONAL`, `GLOBAL`.

<a name="nested_attributes"></a>The `attributes` block supports:

* `criticality` -
  (Optional)
  Criticality of the Application, Service, or Workload
  Structure is [documented below](#nested_attributes_criticality).

* `environment` -
  (Optional)
  Environment of the Application, Service, or Workload
  Structure is [documented below](#nested_attributes_environment).

* `developer_owners` -
  (Optional)
  Optional. Developer team that owns development and coding.
  Structure is [documented below](#nested_attributes_developer_owners).

* `operator_owners` -
  (Optional)
  Optional. Operator team that ensures runtime and operations.
  Structure is [documented below](#nested_attributes_operator_owners).

* `business_owners` -
  (Optional)
  Optional. Business team that ensures user needs are met and value is delivered
  Structure is [documented below](#nested_attributes_business_owners).


<a name="nested_attributes_criticality"></a>The `criticality` block supports:

* `type` -
  (Required)
  Criticality type.
  Possible values are: `MISSION_CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.

<a name="nested_attributes_environment"></a>The `environment` block supports:

* `type` -
  (Required)
  Environment type.
  Possible values are: `PRODUCTION`, `STAGING`, `TEST`, `DEVELOPMENT`.

<a name="nested_attributes_developer_owners"></a>The `developer_owners` block supports:

* `display_name` -
  (Optional)
  Optional. Contact's name.

* `email` -
  (Required)
  Required. Email address of the contacts.

<a name="nested_attributes_operator_owners"></a>The `operator_owners` block supports:

* `display_name` -
  (Optional)
  Optional. Contact's name.

* `email` -
  (Required)
  Required. Email address of the contacts.

<a name="nested_attributes_business_owners"></a>The `business_owners` block supports:

* `display_name` -
  (Optional)
  Optional. Contact's name.

* `email` -
  (Required)
  Required. Email address of the contacts.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/applications/{{application_id}}`

* `name` -
  Identifier. The resource name of an Application. Format:
  "projects/{host-project-id}/locations/{location}/applications/{application-id}"

* `create_time` -
  Output only. Create time.

* `update_time` -
  Output only. Update time.

* `uid` -
  Output only. A universally unique identifier (in UUID4 format) for the `Application`.

* `state` -
  Output only. Application state. 
   Possible values:
   STATE_UNSPECIFIED
  CREATING
  ACTIVE
  DELETING


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Application can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/applications/{{application_id}}`
* `{{project}}/{{location}}/{{application_id}}`
* `{{location}}/{{application_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Application using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/applications/{{application_id}}"
  to = google_apphub_application.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Application can be imported using one of the formats above. For example:

```
$ terraform import google_apphub_application.default projects/{{project}}/locations/{{location}}/applications/{{application_id}}
$ terraform import google_apphub_application.default {{project}}/{{location}}/{{application_id}}
$ terraform import google_apphub_application.default {{location}}/{{application_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
