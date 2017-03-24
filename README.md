# Demo of Whispir templates

This demo wraps the [API of template resource of Whispir](https://whispir.github.io/api/?go#whispir-platform-api) into CLI.

## Build

This demo is written in Go. So Golang should be installed before build.

Then just build as following:

```
make
```

## Run

API key, user name and password are needed before you can run this demo. 
Or you can run this demo with Docker without Golang. See [here](#Docker) for more details.

### Get help

Get details of all available commands as following:

```
$ ./demo -h
Usage:
  demo [command]

Available Commands:
  create      Create template from specified file
  delete      Delete specified template
  get         Get specified template
  help        Help about any command
  list        List all templates
  send        Send message(s) with specified template
  update      update template from specified file

Flags:
  -k, --apikey string      The API key to query Whispir API
  -p, --password string    The password of Whispir
  -u, --user string        The user name of Whispir
  -w, --workspace string   The workspace

Use "demo [command] --help" for more information about a command.
```

Get details of a command as following:

```
$ ./demo create -h
Create template from specified file

Usage:
  demo create [flags]

Flags:
      --file string   Path of template resource file

Global Flags:
  -k, --apikey string      The API key to query Whispir API
  -p, --password string    The password of Whispir
  -u, --user string        The user name of Whispir
  -w, --workspace string   The workspace
```

### Select workspace

You may want to list workspaces at first because templates are the resources can be workspaced. Type the following command:

```
$./demo list -k=<your API key> -u=<your user name> -p=<your password> --resource=workspaces
workspaces: [
  Demo for Habor: {
    id: 3986158A11F9166B
    status: A
    link: https://api.whispir.com/workspaces/3986158A11F9166B?apikey=bvtaqrzveu86gpz8khd66gm7
  }
  ......
]

```

Then select the 'id' to be used as '-w' command flag to query templates in the specified workspace, or you can ignore this
flag to query in the default workspace.

### Create a template

resource-demos/template-demo.json can be used to create a template:

```
$ ./demo create -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B --file=resource-demos/template-demo.json
Create successfully
```

### List templates

```
$ ./demo list -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B
templates: [
  Sample SMS Template: {
    id: 58DBAE17028295AB
    description: Template to provide an example on whispir.io
    link: https://api.whispir.com/workspaces/3986158A11F9166B/templates/58DBAE17028295AB?apikey=bvtaqrzveu86gpz8khd66gm7&limit=20&offset=0
  }
  habor template: {
    id: A00A0E9DD994E4B9
    description:
    link: https://api.whispir.com/workspaces/3986158A11F9166B/templates/A00A0E9DD994E4B9?apikey=bvtaqrzveu86gpz8khd66gm7&limit=20&offset=0
  }
]
```

### Get a template

The 'id' in the list can be used to query a template details:

```
$ ./demo get -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B 58DBAE17028295AB
template 'Sample SMS Template': {
  description: Template to provide an example on whispir.io
  subject: Test SMS Message
  SMS: This is the body of my test SMS message
  ------------------- email message -------------------
    body:
    footer:
  ------------------- email message -------------------
  ------------------- web message -------------------
    body:
  ------------------- web message -------------------
  ------------------- voice message -------------------
    header:
    body:
  ------------------- voice message -------------------
  link: [
    deleteTemplate: (DELETE)https://api.whispir.com/workspaces/3986158A11F9166B/templates/58DBAE17028295AB?apikey=bvtaqrzveu86gpz8khd66gm7
    updateTemplate: (PUT)https://api.whispir.com/workspaces/3986158A11F9166B/templates/58DBAE17028295AB?apikey=bvtaqrzveu86gpz8khd66gm7
  ]
}
```

### Update a template

The usage of 'update' command is similar with 'create' but an additional argument of template id is needed.

Modify the resource-demos/template-demo.json e.g. change messageTemplateName to "Updated Sample SMS Template". Then type:

```
$ ./demo update -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B --file=resource-demos/template-demo.json 58DBAE17028295AB
Update successfully

```

List templates again to check the result:

```
$ ./demo list -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B
templates: [
  Updated Sample SMS Template: {
    id: 58DBAE17028295AB
    description: Template to provide an example on whispir.io
    link: https://api.whispir.com/workspaces/3986158A11F9166B/templates/58DBAE17028295AB?apikey=bvtaqrzveu86gpz8khd66gm7&limit=20&offset=0
  }
  habor template: {
    id: A00A0E9DD994E4B9
    description:
    link: https://api.whispir.com/workspaces/3986158A11F9166B/templates/A00A0E9DD994E4B9?apikey=bvtaqrzveu86gpz8khd66gm7&limit=20&offset=0
  }
]
```

### Delete a template

Delete a template with an argument of template id:

```
$ ./demo delete -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B 58DBAE17028295AB
Delete successfully

```

List templates again to check the result:

```
$ ./demo list -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B
templates: [
  habor template: {
    id: A00A0E9DD994E4B9
    description: 
    link: https://api.whispir.com/workspaces/3986158A11F9166B/templates/A00A0E9DD994E4B9?apikey=bvtaqrzveu86gpz8khd66gm7&limit=20&offset=0
  }
]
```

### Send message with template

You can send messages with a template. Just specify desitination and the template id:

```
$ ./demo send -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B "+8613811112222" A00A0E9DD994E4B9
Message info: https://api.whispir.com/workspaces/3986158A11F9166B/messages/86CF7CCE99C5D1086E4B9A8EAA95685D?apikey=bvtaqrzveu86gpz8khd66gm7
```

## Docker

With the help of Docker, we can wrap the build and running environments into images without installing them. 
If the Docker has been already installed in your environment, you can build an image of this demo:

```
make docker-build
```

A local image named "demo" is available after that. Then you can run this demo as following:

```
$  docker run demo create -k=<your API key> -u=<your user name> -p=<your password> -w=3986158A11F9166B --file=resource-demos/template-demo.json
Create successfully
```