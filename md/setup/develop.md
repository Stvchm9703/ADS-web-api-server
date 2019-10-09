# System setup for Development
*if you are looking for the command list

this part is mainly for developing setup.

## Working Dependecy
the following are necessary for developing
- Nodejs 
  - version (^ 12.8.0)
- Golang 
  - version (^ 1.12.5)

## File Structure

Currently, it only supports relative paths for the templates and static files.

```
--[ filepath for this project ]
  |
  |--[ client-render ] ( web client develop)
  |  |
  |  |--[ assets ]
  |  | ... (resource file, e.g.: css/png file)
  |  |
  |  |--[ components ]
  |  | ... (components template vue file)
  |  |
  |  |--[ dist ]
  |  | ... (distributed files after client build)
  |  |
  |  |--[ layouts ]
  |  | ... (common page layouts template vue file)
  |  |
  |  |--[ middleware ]
  |  | ... (middleware function, js file)
  |  |
  |  |--[ node_modules ]
  |  | ... (nodejs module libs, auto generated)
  |  |
  |  |--[ pages ]
  |  | ... (routing pages template vue file, related to nuxt-routing)
  |  |
  |  |--[ plugins ]
  |  | ... (plugin function for nuxt plugin, js files)
  |  |
  |  |--[ static ]
  |  | ... (static resourse file, that direct render to page)
  |  |
  |  |--[ store ]
  |  | ... (Vuex Module related fuction, js files)
  |  |
  |  |-- nuxt.config.js (js files for config nuxt related configuration)
  |  |
  |  |-- package.json (json files for nodejs module package management (NPM) )
  |  |
  |  |-- package-lock.json (json files for nodejs module package management)
  |  
  |
  |
  |--[ server ]
  |  |
  |  |--[ cmd ]
  |  | ... (go file for command setting)
  |  | 
  |  |--[ common ]
  |  | ... (go file for system common running variables)
  |  |
  |  |--[ controller ]
  |  | ... (go file for mvc controller functions part)
  |  |
  |  |--[ model ]
  |  | ... (go file for mvc model functions part)
  |  |
  |  |-- main.go ( go file for this program server function)
  |  |
  |  |-- router.go ( go file for server routing path, bindings to controller functions)
  |
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
  |  | ... (md files for documentation)
  |
  |--[ vendor ]
  |  | ... (go file for go libs, auto generate in importing)
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
## Installation for Project
*please check the version of Dependecy software*

After you download the project, you should install the go module libs. Please run:
```bash
go mod vendor
```
then, the related libs will be installed into the `vendor` (`{project_directory}/vendor/`) this folder.

If you also want to develop with UI components, you should also install the node_modules. Please run
```bash
cd ./client_render
npm i -D
```
then, the related libs will be installed into the `node_modules` (`{project_directory}/client_render/node_modules/`) this folder.

---
# API server
## Initial Setup for API server

*check the mongodb connection is available first*

This step is required for first time to use, it only does the mongodb create collection and view-collection.

### common command :
    go run main.go init -c=config.yaml

### Usage:
    go run main.go init [flags]

### Flags:
| shorthand | long             | foramt | description                                                                        |
| --------- | ---------------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf`         | string | start server with specific config file (default "{working directory}/config.toml") |
| `-x`      | `--exportSchema` | -      | export the createCollecction validator.$jsonSchema build json                      |
| `-h`      | `--help`         | -      | help message for init                                                              |
| `-m`      | `--schema`       | string | start server with specific config file (default "{working directory}" )            |

---

## Start Server for API server

*check the mongodb connection is available first*

This step is required for starting server.

### common command :
    go run main.go start -c=config.yaml

### Usage:
    go run main.go start [flags]

### Flags:
| shorthand | long     | foramt | description                                                                        |
| --------- | -------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf` | string | start server with specific config file (default "{working directory}/config.toml") |
| `-h`      | `--help` | -      | help message for init                                                              |
| `-m`      | `--mode` | string | server running mode [prod / dev / test] (default "prod")                           |

---

## Import json files for API server

*check the mongodb connection is available first*

This step is used for importing the json data from json files. It is option and please note that the data maybe conflict with the existing in mongodb.

### common command :
    go run main.go import -c=config.yaml

### Usage:
    go run main.go import [flags]

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
    go run main.go export -c=config.yaml

### Usage:
    go run main.go export [flags]

### Flags:
| shorthand | long       | foramt | description                                                                        |
| --------- | ---------- | ------ | ---------------------------------------------------------------------------------- |
| `-c`      | `--conf`   | string | start server with specific config file (default "{working directory}/config.toml") |
| `-h`      | `--help`   | -      | help message for init                                                              |
| `-x`      | `--export` | string | export json file for mongodb (default directory "{working directory}/data")        |

---
## Startup for webpage server in develop mode

This step is used to start webpage server for development.

>*make sure you have installed related dependencies*

```bash
npm run dev
```
or you have installed npx in global
```bash
npx nuxt
```
then the server should run in `localhost:3000`

---
## Startup for webpage server in production mode

This step is used to start webpage server for production.

>*make sure you have installed related dependencies*

```bash
npm run start
```
or you have installed npx in global
```bash
npx nuxt start
```
then the server should run in `localhost:3000`

---
## Build webpage for go server

This step is used to build webpage (in SPA mode) for production.

>*make sure you have installed related dependencies*

```bash
npm run generate
```
or you have installed npx in global
```bash
npx nuxt generate
```

then, you can push the files in `/client_render/dist` to :
- html or file in `/client_render/dist` to `template`
- static files in `/client_render/dist/static` to `static`

after this, you can run the go server with webpage

---
## Build webserver from Golang to Executable

This step is used to build an executable for production.

>*make sure you have installed related dependencies*

```bash
go build
```
then, the executable (`webserver`) is built under the project directory 
