# Azd Modules Example

This is a repo with examples of breaking azd into multiple modules.

## Modules

- **cli** - `azd` CLI
  - Uses: core, ext, bicep & terraform
- **core** - `azd` core module
  - Uses: ext
- **ext** - Extensibility SDK module
  - Uses: None
- **bicep** - Bicep implementation of infra provider
  - Uses: ext
- **terraform** - Terraform implementation of infra provider
  - Uses: ext
- **ai** - AI command extension for custom AI commands
  - Uses: ext, core