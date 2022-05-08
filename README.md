## RunAsUser Plugin

This plugin borrow logic from [Kubernetes Admission Controller for RunAsUser](https://github.com/ElisaOyj/runasuser-admission-controller) as running as vcluster sidecar container.

Idea is that when this plugin is enabled users can use vcluster [isolated mode](https://www.vcluster.com/docs/operator/security#isolated-mode) with **Restricted** [Pod Security Standard](https://kubernetes.io/docs/concepts/security/pod-security-standards/) and still use `kubectl create deployment ...` normal way.


## Using the Plugin

To use the plugin, create a new vcluster with the `plugin.yaml`:

```
# Use public plugin.yaml
vcluster create my-vcluster -n my-vcluster -f https://raw.githubusercontent.com/k8s4u/vcluster-runasuser-plugin/main/plugin.yaml
```

After that, wait for vcluster to start up and create deployment:

```
vcluster connect my-vcluster --namespace my-vcluster -- kubectl create deployment test --image=busybox -- sleep infinity

# Check if pod has started:
vcluster connect my-vcluster --namespace my-vcluster -- kubectl get pods

# Check if pod run with non-root user
vcluster connect my-vcluster --namespace my-vcluster -- kubectl exec -it <pod name> -- whoami

```

## Building the Plugin
To just build the plugin image and push it to the registry, run:
```
# Build
docker build . -t k8s4u/vcluster-runasuser-plugin:dev

# Push
docker push k8s4u/vcluster-runasuser-plugin:dev
```

Then exchange the image in the `plugin.yaml`.

## Development

General vcluster plugin project structure:
```
.
├── go.mod              # Go module definition
├── go.sum
├── devspace.yaml       # Development environment definition
├── devspace_start.sh   # Development entrypoint script
├── Dockerfile          # Production Dockerfile 
├── Dockerfile.dev      # Development Dockerfile
├── main.go             # Go Entrypoint
├── plugin.yaml         # Plugin Helm Values
├── syncers/            # Plugin Syncers
└── manifests/          # Additional plugin resources
```

Before starting to develop, make sure you have installed the following tools on your computer:
- [docker](https://docs.docker.com/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/) with a valid kube context configured
- [helm](https://helm.sh/docs/intro/install/), which is used to deploy vcluster and the plugin
- [vcluster CLI](https://www.vcluster.com/docs/getting-started/setup) v0.6.0 or higher
- [DevSpace](https://devspace.sh/cli/docs/quickstart), which is used to spin up a development environment

If you want to develop within a remote Kubernetes cluster (as opposed to docker-desktop or minikube), make sure to exchange `PLUGIN_IMAGE` in the `devspace.yaml` with a valid registry path you can push to.

After successfully setting up the tools, start the development environment with:
```
devspace dev -n vcluster
```

After a while a terminal should show up with additional instructions. Enter the following command to start the plugin:
```
go run -mod vendor ./cmd/main.go
```

The output should look something like this:
```
I0124 11:20:14.702799    4185 logr.go:249] plugin: Try creating context...
I0124 11:20:14.730044    4185 logr.go:249] plugin: Waiting for vcluster to become leader...
I0124 11:20:14.731097    4185 logr.go:249] plugin: Starting syncers...
[...]
I0124 11:20:15.957331    4185 logr.go:249] plugin: Successfully started plugin.
```

You can now change a file locally in your IDE and then restart the command in the terminal to apply the changes to the plugin.

Delete the development environment with:
```
devspace purge -n vcluster
```
