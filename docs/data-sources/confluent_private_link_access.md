---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "confluent_private_link_access Data Source - terraform-provider-confluent"
subcategory: ""
description: |-
  
---

# confluent_private_link_access Data Source

`confluent_private_link_access` describes a Network data source.

## Example Usage

```terraform
data "confluent_private_link_access" "example_using_id" {
  id = "pla-abc123"
  environment {
    id = "env-xyz456"
  }
}

output "example_using_id" {
  value = data.confluent_private_link_access.example_using_id
}

data "confluent_private_link_access" "example_using_name" {
  display_name = "my_pla"
  environment {
    id = "env-xyz456"
  }
}

output "example_using_name" {
  value = data.confluent_private_link_access.example_using_name
}
```

<!-- schema generated by tfplugindocs -->
## Argument Reference

The following arguments are supported:

- `id` - (Optional String) The ID of the Private Link Access, for example, `pla-abc123`.
- `display_name` - (Optional String) A human-readable name for the Private Link Access.
- `environment` (Required Configuration Block) supports the following:
  - `id` - (Required String) The ID of the Environment that the Private Link Access belongs to, for example, `env-xyz456`.

-> **Note:** Exactly one from the `id` and `display_name` attributes must be specified.

## Attributes Reference

In addition to the preceding arguments, the following attributes are exported:

- `id` - (Required String) The ID of the Network, for example, `n-abc123`.
- `display_name` - (Optional String) The name of the Private Link Access.
- `environment` (Required Configuration Block) supports the following:
  - `id` - (Required String) The ID of the Environment that the Private Link Access belongs to, for example, `env-abc123`.
- `network` (Required Configuration Block) supports the following:
  - `id` - (Required String) The ID of the Network that the Private Link Access belongs to, for example, `n-abc123`.
- `aws` - (Optional Configuration Block) The AWS-specific Private Link Access details if available. It supports the following:
  - `account` - (Required String) The AWS account ID to enable for the Private Link Access. You can find your AWS account ID [here](https://console.aws.amazon.com/billing/home?#/account) under **My Account** in your AWS Management Console. Must be a **12 character string**.
- `azure` - (Optional Configuration Block) The Azure-specific Private Link Access details if available. It supports the following:
  - `subscription` - (Required String) The Azure subscription ID to enable for the Private Link Access. You can find your Azure subscription ID in the subscription section of your [Microsoft Azure Portal](https://portal.azure.com/#blade/Microsoft_Azure_Billing/SubscriptionsBlade). Must be a valid **32 character UUID string**.

-> **Note:** Use the `aws[0]` or `azure[0]` prefix for referencing these attributes, for example, `data.confluent_private_link_access.example_using_name.aws[0].account`.