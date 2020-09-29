# Zuora Notifications Terraform

Manage Zuora notifications as code

## Roadmap

- Verify an event trigger exists and is active
- Ensure a notification definition is created for each event trigger \*
  communication profile
- Use the shared `callout` (in the template) and add the params from the
  notifications
- Prefix the `eventTypeName` with something like `znt-` to uniquely identify
  triggers managed by this tool
- Similarly, prefix all notification definition name by `znt-` and construct the
  name like: `znt-on<Object><ConditionKey>`
- Add an `apply` and a `destroy` command (self-explanatory)
