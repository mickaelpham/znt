# Zuora Notifications Terraform

Manage Zuora notifications as code

## Usage

```
Zuora Notifications Terraform (ZNT)

A manager for Zuora notification definitions
built in Go. Complete documentation is available at
https://github.com/mickaelpham/znt

Usage:
  znt [command]

Available Commands:
  help        Help about any command
  verify      Verify notifications exist

Flags:
  -c, --config string     config file (default is $HOME/.znt.yaml)
  -h, --help              help for znt
  -t, --template string   template file

Use "znt [command] --help" for more information about a command.
```

### Verify

Running the `verify` subcommand, given a `template.json` file, will verify the
presence of the rendered triggers and notification definitions against the
targeted Zuora environment.

```
2020/10/01 21:25:38 Using config file: /Users/mickael/.znt.yaml
2020/10/01 21:25:38 GET /events/event-triggers
2020/10/01 21:25:38 GET /events/event-triggers?start=10&limit=10

--- Trigger Diff

These triggers will be created:
  * {Account on "changeType == 'INSERT'"}
  * {Account on "changeType == 'UPDATE'"}


--- Communication Profiles
  * (profile-id-123) Profile A
  * (profile-id-789) Profile B

2020/10/01 21:25:40 GET /notifications/notification-definitions

--- Notification Diff

These notifications will be created:
  * (profile-id-789) znt-Account-onInsert
  * (profile-id-123) znt-Account-onInsert
  * (profile-id-789) znt-Account-onUpdate
  * (profile-id-123) znt-Account-onUpdate
```

## Roadmap

- [x] Verify an event trigger exists and is active
- [x] Ensure a notification definition is created for each event trigger \*
      communication profile
- [x] Use the shared `callout` (in the template) and add the params from the
      notifications
- [x] Prefix the `eventTypeName` with something like `znt-` to uniquely identify
      triggers managed by this tool
- [x] Similarly, prefix all notification definition name by `znt-` and construct
      the name like: `znt-on<Object><ConditionKey>`
- Add an `apply` and a `destroy` command (self-explanatory)
- Update the notification instead of destroying/adding them back
