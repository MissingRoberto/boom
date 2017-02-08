# boom

Boom is a tool for manipulating bosh manifests. This tool targets not only to help with the quick manifest modification, but also to with the maintenance of Bosh configuration.

## Installation

* Mac OS

```

$ wget https://github.com/jszroberto/boom/releases/download/0.1/boom ; chmod +x boom; mv boom /usr/local/bin
```

* Linux:

```

$ wget https://github.com/jszroberto/boom/releases/download/0.1/boom-linux ; chmod +x boom-linux; mv boom-linux /usr/local/bin/boom
```


## Usage

```
NAME:
   boom - a simple and quick tool for bosh manifest maintenance

USAGE:
   boom [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     mask, si             creates a mask for a given list and a given key
     set-instances, si    Set the number of instances
     scale-instances, sc  Scale number of instances
     help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Examples of use

### List the jobs

```
$ boom mask manifest.yml jobs
---
jobs:
- name: database
- name: brain
- name: cell
- name: cc_bridge
- name: route_emitter
- name: access
- name: acceptance_tests
```

If you want to know the value of other keys, you only need to run something like this:

```
$ boom mask manifest.yml jobs
---
jobs:
- instances: 1
  name: database
- instances: 1
  name: brain
- instances: 4
  name: cell
```

### List the releases installed

```
$ boom mask manifest.yml releases
---
releases:
- name: diego
- name: cf
- name: garden-runc

```
If you want to get the versions:

```
boom mask lon-fabric.yml releases version
---
releases:
- name: diego
  version: 1.0.1
- name: cf
  version: 251
- name: garden-runc
  version: 1.1.1
```
