---
title: "meshctl debug snapshot"
weight: 5
---
## meshctl debug snapshot

Input and Output snapshots for the discovery and networking pod. Requires jq to be installed if the --json flag is not being used.

```
meshctl debug snapshot [flags]
```

### Options

```
  -f, --file string   file to write output to
  -h, --help          help for snapshot
      --json          display the entire json snapshot. The output can be piped into a command like jq (https://stedolan.github.io/jq/tutorial/). For example:
                       meshctl debug snapshot discovery input | jq '.'
      --zip string    zip file output
```

### Options inherited from parent commands

```
  -v, --verbose   Enable verbose logging
```

### SEE ALSO

* [meshctl debug](../meshctl_debug)	 - Debug Gloo Mesh resources
* [meshctl debug snapshot discovery](../meshctl_debug_snapshot_discovery)	 - Input and output snapshots for the discovery pod
* [meshctl debug snapshot networking](../meshctl_debug_snapshot_networking)	 - Input and output snapshots for the networking pod

