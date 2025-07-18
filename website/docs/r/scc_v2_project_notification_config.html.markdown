---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenterv2/ProjectNotificationConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Security Command Center (SCC) v2 API"
description: |-
  This is a continuous export that exports findings to a Pub/Sub topic.
---

# google_scc_v2_project_notification_config

This is a continuous export that exports findings to a Pub/Sub topic.


To get more information about ProjectNotificationConfig, see:

* [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v2/projects.locations.notificationConfigs)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/security-command-center/docs)

## Example Usage - Scc V2 Project Notification Config Basic


```hcl
resource "google_pubsub_topic" "scc_v2_project_notification" {
  name = "my-topic"
}

resource "google_scc_v2_project_notification_config" "custom_notification_config" {
  config_id    = "my-config"
  project      = "my-project-name"
  location     = "global"
  description  = "My custom Cloud Security Command Center Finding Notification Configuration"
  pubsub_topic =  google_pubsub_topic.scc_v2_project_notification.id

  streaming_config {
    filter = "category = \"OPEN_FIREWALL\" AND state = \"ACTIVE\""
  }
}
```

## Argument Reference

The following arguments are supported:


* `streaming_config` -
  (Required)
  The config for triggering streaming-based notifications.
  Structure is [documented below](#nested_streaming_config).

* `config_id` -
  (Required)
  This must be unique within the project.


* `description` -
  (Optional)
  The description of the notification config (max of 1024 characters).

* `pubsub_topic` -
  (Optional)
  The Pub/Sub topic to send notifications to. Its format is
  "projects/[project_id]/topics/[topic]".

* `location` -
  (Optional)
  Location ID of the parent organization. Only global is supported at the moment.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_streaming_config"></a>The `streaming_config` block supports:

* `filter` -
  (Required)
  Expression that defines the filter to apply across create/update
  events of assets or findings as specified by the event type. The
  expression is a list of zero or more restrictions combined via
  logical operators AND and OR. Parentheses are supported, and OR
  has higher precedence than AND.
  Restrictions have the form <field> <operator> <value> and may have
  a - character in front of them to indicate negation. The fields
  map to those defined in the corresponding resource.
  The supported operators are:
  * = for all value types.
  * >, <, >=, <= for integer values.
  * :, meaning substring matching, for strings.
  The supported value types are:
  * string literals in quotes.
  * integer literals without quotes.
  * boolean literals true and false without quotes.
  See
  [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
  for information on how to write a filter.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{name}}`

* `name` -
  The resource name of this notification config, in the format
  `projects/{{projectId}}/locations/{{location}}/notificationConfigs/{{config_id}}`.

* `service_account` -
  The service account that needs "pubsub.topics.publish" permission to
  publish to the Pub/Sub topic.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ProjectNotificationConfig can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/notificationConfigs/{{config_id}}`
* `{{project}}/{{location}}/{{config_id}}`
* `{{location}}/{{config_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ProjectNotificationConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/notificationConfigs/{{config_id}}"
  to = google_scc_v2_project_notification_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ProjectNotificationConfig can be imported using one of the formats above. For example:

```
$ terraform import google_scc_v2_project_notification_config.default projects/{{project}}/locations/{{location}}/notificationConfigs/{{config_id}}
$ terraform import google_scc_v2_project_notification_config.default {{project}}/{{location}}/{{config_id}}
$ terraform import google_scc_v2_project_notification_config.default {{location}}/{{config_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
