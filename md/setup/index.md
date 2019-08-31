# System setup for Deployment
*if you are looking for the command list

this part is mainly for production release setup.

## File Structure

Currently, it only supports relative paths for the templates and static files.

```
--[ filepath for this server ]
  |
  |-- webserver (executable)
  |
  |--[ data ]
  |  | ... (json file)
  |
  |--[ log ]
  |  | (nothing but keep the directory as it is)
  |
  |--[ static ]
  |  | ... (js / css files)
  |
  |--[ template ]
  |  | ... (html files)
  |
  |--[ md ]
  |  | ... (md files)
  |
  |-- config.yaml
```

## Before Initial Setup
Before running the initial setup script, please make sure the `config.yaml` file is 
- set server ip and port
- set the mongodb connection.

*currently, this version is not supported the configuration of template and static files, but please keep the settings into file*
```yaml
server:
  ip: "0.0.0.0"
  port: 8080
  root_path: A:\Gitrepo\ADS-web-api-server\
  main_path: ./
  static_filepath: A:\Gitrepo\ADS-web-api-server\static
  static_outpath: ./static
  template_filepath: A:\Gitrepo\ADS-web-api-server\template
  template_outpath: ./template

database:
  connector: mongodb
  host: 127.0.0.1
  port: 27017
  username: ""
  password: ""
  database: scope_ADS1
  filepath: ""
  
db_fallback:
  schema: db
  filepath: A:\Gitrepo\Creator-Selling-Platform-API\db\Wildbase.fall.db
```
---

## Initial Setup

*check the mongodb connection is available first*

This step is required for first time to use, it only does the mongodb create collection and view-collection.

### common command :
    webserver init -c=config.yaml

### Usage:
    webserver init [flags]

### Flags:
| shorthand | long             | foramt | description                                                                        |
| --------- | ---------------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf`         | string | start server with specific config file (default "{working directory}/config.toml") |
| `-x`      | `--exportSchema` | -      | export the createCollecction validator.$jsonSchema build json                      |
| `-h`      | `--help`         | -      | help message for init                                                              |
| `-m`      | `--schema`       | string | start server with specific config file (default "{working directory}" )            |

---

## Start Server


*check the mongodb connection is available first*

This step is required for starting server.


### common command :
    webserver start -c=config.yaml

### Usage:
    webserver start [flags]

### Flags:
| shorthand | long     | foramt | description                                                                        |
| --------- | -------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf` | string | start server with specific config file (default "{working directory}/config.toml") |
| `-h`      | `--help` | -      | help message for init                                                              |
| `-m`      | `--mode` | string | server running mode [prod / dev / test] (default "prod")                           |

---

## Import json files

*check the mongodb connection is available first*

This step is used for importing the json data from json files. It is option and please note that the data maybe conflict with the existing in mongodb.

### common command :
    webserver import -c=config.yaml

### Usage:
    webserver import [flags]

### Flags:
| shorthand | long       | foramt | description                                                                        |
| --------- | ---------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf`   | string | start server with specific config file (default "{working directory}/config.toml") |
| `-h`      | `--help`   | -      | help message for init                                                              |
| `-i`      | `--import` | string | import json file for mongodb (default directory "{working directory}/data")        |

---

## Export json files

*check the mongodb connection is available first*

This step is used for importing the json data from json files. It is option and please note that the data maybe conflict with the existing in mongodb.

### common command :
    webserver export -c=config.yaml

### Usage:
    webserver export [flags]

### Flags:
| shorthand | long       | foramt | description                                                                        |
| --------- | ---------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf`   | string | start server with specific config file (default "{working directory}/config.toml") |
| `-h`      | `--help`   | -      | help message for init                                                              |
| `-x`      | `--export` | string | export json file for mongodb (default directory "{working directory}/data")        |

---