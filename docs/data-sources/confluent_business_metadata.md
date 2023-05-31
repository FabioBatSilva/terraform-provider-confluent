---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "confluent_business_metadata Data Source - terraform-provider-confluent"
subcategory: ""
description: |-
   
---

# confluent_business_metadata Data Source

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)

-> **Note:** `confluent_business_metadata` data source is available in **Preview** for early adopters. Preview features are introduced to gather customer feedback. This feature should be used only for evaluation and non-production testing purposes or to provide feedback to Confluent, particularly as it becomes more widely available in follow-on editions.  
**Preview** features are intended for evaluation use in development and testing environments only, and not for production use. The warranty, SLA, and Support Services provisions of your agreement with Confluent do not apply to Preview features. Preview features are considered to be a Proof of Concept as defined in the Confluent Cloud Terms of Service. Confluent may discontinue providing preview releases of the Preview features at any time in Confluent’s sole discretion.

`confluent_business_metadata` describes a Business Metadata data source.

## Example Usage

### Option #1: Manage multiple Schema Registry clusters in the same Terraform workspace

```terraform
provider "confluent" {
  cloud_api_key    = var.confluent_cloud_api_key    # optionally use CONFLUENT_CLOUD_API_KEY env var
  cloud_api_secret = var.confluent_cloud_api_secret # optionally use CONFLUENT_CLOUD_API_SECRET env var
}

data "confluent_business_metadata" "pii" {
  schema_registry_cluster {
    id = confluent_schema_registry_cluster.essentials.id
  }
  rest_endpoint = confluent_schema_registry_cluster.essentials.rest_endpoint
  credentials {
    key    = "<Schema Registry API Key for confluent_schema_registry_cluster.essentials>"
    secret = "<Schema Registry API Secret for confluent_schema_registry_cluster.essentials>"
  }
  
  name = "PII"
}
```

### Option #2: Manage a single Schema Registry cluster in the same Terraform workspace

```terraform
provider "confluent" {
  schema_registry_id            = var.schema_registry_id            # optionally use SCHEMA_REGISTRY_ID env var
  schema_registry_rest_endpoint = var.schema_registry_rest_endpoint # optionally use SCHEMA_REGISTRY_REST_ENDPOINT env var
  schema_registry_api_key       = var.schema_registry_api_key       # optionally use SCHEMA_REGISTRY_API_KEY env var
  schema_registry_api_secret    = var.schema_registry_api_secret    # optionally use SCHEMA_REGISTRY_API_SECRET env var
}

data "confluent_business_metadata" "pii" {
  name = "PII"
}
```

<!-- schema generated by tfplugindocs -->
## Argument Reference

The following arguments are supported:

- `schema_registry_cluster` - (Optional Configuration Block) supports the following:
    - `id` - (Required String) The ID of the Schema Registry cluster, for example, `lsrc-abc123`.
- `rest_endpoint` - (Optional String) The REST endpoint of the Schema Registry cluster, for example, `https://psrc-00000.us-central1.gcp.confluent.cloud:443`).
- `credentials` (Optional Configuration Block) supports the following:
    - `key` - (Required String) The Schema Registry API Key.
    - `secret` - (Required String, Sensitive) The Schema Registry API Secret.
- `name` - (Required String) The name of the Business Metadata, for example, `PII`. The name must not be empty and consist of a letter followed by a sequence of letter, number, space, or _ characters.

-> **Note:** A Schema Registry API key consists of a key and a secret. Schema Registry API keys are required to interact with Schema Registry clusters in Confluent Cloud. Each Schema Registry API key is valid for one specific Schema Registry cluster.

## Attributes Reference

In addition to the preceding arguments, the following attributes are exported:

- `id` - (Required String) The ID of the Business Metadata, in the format `<Schema Registry cluster ID>/<Business Metadata name>`, for example, `lsrc-8wrx70/PII`.
- `description` - (Optional String) The description of the Business Metadata.
- `version` - (Required Integer) The version of the Business Metadata, for example, `1`.
- `attribute_definition` - (Optional List) The list of attribute definitions (see [Business Metadata](https://docs.confluent.io/cloud/current/stream-governance/stream-catalog.html#business-metadata-for-schemas) for more details):
  - `name` - (Required String) The name of the attribute.
  - `type` - (Required String) The type of the attribute, it always returns `string`.
  - `is_optional` - (Optional Boolean) An optional flag to control whether the attribute should be optional or required.
  - `default_value` - (Optional String) The default value of this attribute.
  - `description` - (Optional String) The description of this attribute.
  - `options` - (Optional Map) Block for the attribute options:
      - `applicableEntityTypes` - (Optional String) The entity types that the attribute is applicable, it always returns `[\"cf_entity\"]`.
      - `maxStrLength` - (Optional String) The maximum length of the string value, it always returns `5000`.