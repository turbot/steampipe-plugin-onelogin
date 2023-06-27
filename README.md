![image](https://hub.steampipe.io/images/plugins/turbot/onelogin-social-graphic.png)

# OneLogin Plugin for Steampipe

Use SQL to query audiences, automation workflows, campaigns, and more from OneLogin

- **[Get started →](https://hub.steampipe.io/plugins/turbot/onelogin)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/onelogin/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-onelogin/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install onelogin
```

Run a query:

```sql
select
  id,
  title,
  content_type,
  create_time,
  emails_sent,
  send_time,
  status,
  type
from
  onelogin_campaign;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-onelogin.git
cd steampipe-plugin-onelogin
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/onelogin.spc
```

Try it!

```
steampipe query
> .inspect onelogin
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-onelogin/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [OneLogin Plugin](https://github.com/turbot/steampipe-plugin-onelogin/labels/help%20wanted)
