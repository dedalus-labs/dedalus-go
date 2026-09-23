# Dedalus Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

- [`Machines`](#machines)
  - [List machines](#list-machines)
  - [Create machine](#create-machine)
  - [Get machine](#get-machine)
  - [Update machine](#update-machine)
  - [Destroy machine](#destroy-machine)
  - [Watch machine lifecycle status](#watch-machine-lifecycle-status)
  - [Sleep a running machine](#sleep-a-running-machine)
  - [Wake a sleeping machine](#wake-a-sleeping-machine)
  - [`Machines Network`](#machines-network)
    - [Get machine network identity](#get-machine-network-identity)
  - [`Machines Artifacts`](#machines-artifacts)
    - [List artifacts](#list-artifacts)
    - [Get artifact](#get-artifact)
    - [Delete artifact](#delete-artifact)
  - [`Machines Ports`](#machines-ports)
    - [List ports](#list-ports)
    - [Create port](#create-port)
    - [Get port](#get-port)
    - [Delete port](#delete-port)
  - [`Machines Ssh`](#machines-ssh)
    - [List SSH sessions](#list-ssh-sessions)
    - [Create SSH session](#create-ssh-session)
    - [Get SSH session](#get-ssh-session)
    - [Delete SSH session](#delete-ssh-session)
  - [`Machines Executions`](#machines-executions)
    - [List executions](#list-executions)
    - [Create execution](#create-execution)
    - [Get execution](#get-execution)
    - [Delete execution](#delete-execution)
    - [Get execution output](#get-execution-output)
    - [List execution events](#list-execution-events)
  - [`Machines Terminals`](#machines-terminals)
    - [List terminals](#list-terminals)
    - [Create terminal](#create-terminal)
    - [Get terminal](#get-terminal)
    - [Delete terminal](#delete-terminal)
- [`Networks`](#networks)
  - [Get network details](#get-network-details)
- [`Usage`](#usage)
  - [Get usage summary](#get-usage-summary)
  - [List machine compute usage breakdown](#list-machine-compute-usage-breakdown)
  - [List machine storage usage breakdown](#list-machine-storage-usage-breakdown)

## Setup

```go
import (
	"context"
	"fmt"

	sdk "github.com/dedalus-labs/dedalus-go"
)

client := sdk.NewClient()
```

## `Machines`

### List machines

| Direction | Type |
| --- | --- |
| Request | [`MachineListParams`](./machine.go) |
| Response | [`MachineListItem`](./machine.go) |

```go
page, err := client.Machines.List(context.Background(), sdk.MachineListParams{})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

### Create machine

| Direction | Type |
| --- | --- |
| Request | [`MachineNewParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.New(context.Background(), sdk.MachineNewParams{
	CreateParams: sdk.CreateParams{
		Autosleep:  sdk.F[string]("300s"),
		MemoryMib:  sdk.F[int64](4096),
		StorageGib: sdk.F[int64](10),
		Vcpu:       sdk.F[float64](1),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Get machine

| Direction | Type |
| --- | --- |
| Request | [`MachineGetParams`](./machine.go) |
| Response | [`MachineGetResponse`](./machine.go) |

```go
machine, err := client.Machines.Get(context.Background(), sdk.MachineGetParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Update machine

| Direction | Type |
| --- | --- |
| Request | [`MachineUpdateParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Update(context.Background(), sdk.MachineUpdateParams{
	MachineID:    "machineID",
	UpdateParams: sdk.UpdateParams{},
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Destroy machine

| Direction | Type |
| --- | --- |
| Request | [`MachineDeleteParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Delete(context.Background(), sdk.MachineDeleteParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Watch machine lifecycle status

Streams machine lifecycle updates over Server-Sent Events. Each `status` event contains a full `LifecycleResponse` payload. The stream closes after the machine reaches its current desired state.

| Direction | Type |
| --- | --- |
| Request | [`MachineWatchParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
stream := client.Machines.WatchStreaming(context.Background(), sdk.MachineWatchParams{
	MachineID: "machineID",
})
defer stream.Close()

for stream.Next() {
	event := stream.Current()
	fmt.Println(event)
}
if err := stream.Err(); err != nil {
	panic(err)
}
```

### Sleep a running machine

| Direction | Type |
| --- | --- |
| Request | [`MachineSleepParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Sleep(context.Background(), sdk.MachineSleepParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Wake a sleeping machine

| Direction | Type |
| --- | --- |
| Request | [`MachineWakeParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Wake(context.Background(), sdk.MachineWakeParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### `Machines Network`

#### Get machine network identity

| Direction | Type |
| --- | --- |
| Request | [`MachineNetworkGetParams`](./machinenetwork.go) |
| Response | [`MachineNetwork`](./machinenetwork.go) |

```go
network, err := client.Machines.Network.Get(context.Background(), sdk.MachineNetworkGetParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(network)
```

### `Machines Artifacts`

#### List artifacts

| Direction | Type |
| --- | --- |
| Request | [`MachineArtifactListParams`](./machineartifact.go) |
| Response | [`Artifact`](./machineartifact.go) |

```go
page, err := client.Machines.Artifacts.List(context.Background(), sdk.MachineArtifactListParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### Get artifact

| Direction | Type |
| --- | --- |
| Request | [`MachineArtifactGetParams`](./machineartifact.go) |
| Response | [`Artifact`](./machineartifact.go) |

```go
artifact, err := client.Machines.Artifacts.Get(context.Background(), sdk.MachineArtifactGetParams{
	ArtifactID: "artifactID",
	MachineID:  "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(artifact.ArtifactID)
```

#### Delete artifact

| Direction | Type |
| --- | --- |
| Request | [`MachineArtifactDeleteParams`](./machineartifact.go) |
| Response | [`Artifact`](./machineartifact.go) |

```go
artifact, err := client.Machines.Artifacts.Delete(context.Background(), sdk.MachineArtifactDeleteParams{
	ArtifactID: "artifactID",
	MachineID:  "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(artifact.ArtifactID)
```

### `Machines Ports`

#### List ports

| Direction | Type |
| --- | --- |
| Request | [`MachinePortListParams`](./machineport.go) |
| Response | [`Port`](./machineport.go) |

```go
page, err := client.Machines.Ports.List(context.Background(), sdk.MachinePortListParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### Create port

| Direction | Type |
| --- | --- |
| Request | [`MachinePortNewParams`](./machineport.go) |
| Response | [`Port`](./machineport.go) |

```go
port, err := client.Machines.Ports.New(context.Background(), sdk.MachinePortNewParams{
	MachineID: "machineID",
	PortCreateParams: sdk.PortCreateParams{
		Port: sdk.F[int64](0),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(port.PortID)
```

#### Get port

| Direction | Type |
| --- | --- |
| Request | [`MachinePortGetParams`](./machineport.go) |
| Response | [`Port`](./machineport.go) |

```go
port, err := client.Machines.Ports.Get(context.Background(), sdk.MachinePortGetParams{
	MachineID: "machineID",
	PortID:    "portID",
})
if err != nil {
	panic(err)
}

fmt.Println(port.PortID)
```

#### Delete port

| Direction | Type |
| --- | --- |
| Request | [`MachinePortDeleteParams`](./machineport.go) |
| Response | [`Port`](./machineport.go) |

```go
port, err := client.Machines.Ports.Delete(context.Background(), sdk.MachinePortDeleteParams{
	MachineID: "machineID",
	PortID:    "portID",
})
if err != nil {
	panic(err)
}

fmt.Println(port.PortID)
```

### `Machines Ssh`

#### List SSH sessions

| Direction | Type |
| --- | --- |
| Request | [`MachineSSHListParams`](./machinessh.go) |
| Response | [`SSHSession`](./machinessh.go) |

```go
page, err := client.Machines.SSH.List(context.Background(), sdk.MachineSSHListParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### Create SSH session

| Direction | Type |
| --- | --- |
| Request | [`MachineSSHNewParams`](./machinessh.go) |
| Response | [`SSHSession`](./machinessh.go) |

```go
ssh, err := client.Machines.SSH.New(context.Background(), sdk.MachineSSHNewParams{
	MachineID: "machineID",
	SSHSessionCreateParams: sdk.SSHSessionCreateParams{
		PublicKey: sdk.F[string](""),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(ssh.SessionID)
```

#### Get SSH session

| Direction | Type |
| --- | --- |
| Request | [`MachineSSHGetParams`](./machinessh.go) |
| Response | [`SSHSession`](./machinessh.go) |

```go
ssh, err := client.Machines.SSH.Get(context.Background(), sdk.MachineSSHGetParams{
	MachineID: "machineID",
	SessionID: "sessionID",
})
if err != nil {
	panic(err)
}

fmt.Println(ssh.SessionID)
```

#### Delete SSH session

| Direction | Type |
| --- | --- |
| Request | [`MachineSSHDeleteParams`](./machinessh.go) |
| Response | [`SSHSession`](./machinessh.go) |

```go
ssh, err := client.Machines.SSH.Delete(context.Background(), sdk.MachineSSHDeleteParams{
	MachineID: "machineID",
	SessionID: "sessionID",
})
if err != nil {
	panic(err)
}

fmt.Println(ssh.SessionID)
```

### `Machines Executions`

#### List executions

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionListParams`](./machineexecution.go) |
| Response | [`Execution`](./machineexecution.go) |

```go
page, err := client.Machines.Executions.List(context.Background(), sdk.MachineExecutionListParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### Create execution

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionNewParams`](./machineexecution.go) |
| Response | [`Execution`](./machineexecution.go) |

```go
execution, err := client.Machines.Executions.New(context.Background(), sdk.MachineExecutionNewParams{
	MachineID: "machineID",
	ExecutionCreateParams: sdk.ExecutionCreateParams{
		Command: sdk.F[[]string]([]string{""}),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(execution.ExecutionID)
```

#### Get execution

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionGetParams`](./machineexecution.go) |
| Response | [`Execution`](./machineexecution.go) |

```go
execution, err := client.Machines.Executions.Get(context.Background(), sdk.MachineExecutionGetParams{
	ExecutionID: "executionID",
	MachineID:   "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(execution.ExecutionID)
```

#### Delete execution

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionDeleteParams`](./machineexecution.go) |
| Response | [`Execution`](./machineexecution.go) |

```go
execution, err := client.Machines.Executions.Delete(context.Background(), sdk.MachineExecutionDeleteParams{
	ExecutionID: "executionID",
	MachineID:   "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(execution.ExecutionID)
```

#### Get execution output

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionOutputParams`](./machineexecution.go) |
| Response | [`ExecutionOutput`](./machineexecution.go) |

```go
execution, err := client.Machines.Executions.Output(context.Background(), sdk.MachineExecutionOutputParams{
	ExecutionID: "executionID",
	MachineID:   "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(execution.ExecutionID)
```

#### List execution events

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionEventsParams`](./machineexecution.go) |
| Response | [`ExecutionEvent`](./machineexecution.go) |

```go
page, err := client.Machines.Executions.Events(context.Background(), sdk.MachineExecutionEventsParams{
	ExecutionID: "executionID",
	MachineID:   "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

### `Machines Terminals`

#### List terminals

| Direction | Type |
| --- | --- |
| Request | [`MachineTerminalListParams`](./machineterminal.go) |
| Response | [`Terminal`](./machineterminal.go) |

```go
page, err := client.Machines.Terminals.List(context.Background(), sdk.MachineTerminalListParams{
	MachineID: "machineID",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### Create terminal

| Direction | Type |
| --- | --- |
| Request | [`MachineTerminalNewParams`](./machineterminal.go) |
| Response | [`Terminal`](./machineterminal.go) |

```go
terminal, err := client.Machines.Terminals.New(context.Background(), sdk.MachineTerminalNewParams{
	MachineID: "machineID",
	TerminalCreateParams: sdk.TerminalCreateParams{
		Height: sdk.F[int64](0),
		Width:  sdk.F[int64](0),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(terminal.TerminalID)
```

#### Get terminal

| Direction | Type |
| --- | --- |
| Request | [`MachineTerminalGetParams`](./machineterminal.go) |
| Response | [`Terminal`](./machineterminal.go) |

```go
terminal, err := client.Machines.Terminals.Get(context.Background(), sdk.MachineTerminalGetParams{
	MachineID:  "machineID",
	TerminalID: "terminalID",
})
if err != nil {
	panic(err)
}

fmt.Println(terminal.TerminalID)
```

#### Delete terminal

| Direction | Type |
| --- | --- |
| Request | [`MachineTerminalDeleteParams`](./machineterminal.go) |
| Response | [`Terminal`](./machineterminal.go) |

```go
terminal, err := client.Machines.Terminals.Delete(context.Background(), sdk.MachineTerminalDeleteParams{
	MachineID:  "machineID",
	TerminalID: "terminalID",
})
if err != nil {
	panic(err)
}

fmt.Println(terminal.TerminalID)
```

## `Networks`

### Get network details

| Direction | Type |
| --- | --- |
| Request | [`NetworkGetParams`](./network.go) |
| Response | [`Network`](./network.go) |

```go
network, err := client.Networks.Get(context.Background(), sdk.NetworkGetParams{
	NetworkID: "networkID",
})
if err != nil {
	panic(err)
}

fmt.Println(network)
```

## `Usage`

### Get usage summary

| Direction | Type |
| --- | --- |
| Request | [`UsageGetParams`](./usage.go) |
| Response | [`OrgUsage`](./usage.go) |

```go
usage, err := client.Usage.Get(context.Background(), sdk.UsageGetParams{})
if err != nil {
	panic(err)
}

fmt.Println(usage)
```

### List machine compute usage breakdown

| Direction | Type |
| --- | --- |
| Request | [`UsageMachineComputeParams`](./usage.go) |
| Response | [`MachineComputeUsage`](./usage.go) |

```go
usage, err := client.Usage.MachineCompute(context.Background(), sdk.UsageMachineComputeParams{})
if err != nil {
	panic(err)
}

fmt.Println(usage)
```

### List machine storage usage breakdown

| Direction | Type |
| --- | --- |
| Request | [`UsageMachineStorageParams`](./usage.go) |
| Response | [`MachineStorageUsage`](./usage.go) |

```go
usage, err := client.Usage.MachineStorage(context.Background(), sdk.UsageMachineStorageParams{})
if err != nil {
	panic(err)
}

fmt.Println(usage)
```
