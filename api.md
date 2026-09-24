# Dedalus Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

- [`Machines`](#machines)
  - [List machines](#list-machines)
  - [Create machine](#create-machine)
  - [Get machine](#get-machine)
  - [Update machine](#update-machine)
  - [Destroy machine](#destroy-machine)
  - [Sleep a running machine](#sleep-a-running-machine)
  - [Wake a sleeping machine](#wake-a-sleeping-machine)
  - [Reboot a machine with fresh memory](#reboot-a-machine-with-fresh-memory)
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
    - [`Machines Executions Logs`](#machines-executions-logs)
      - [Get execution log status](#get-execution-log-status)
      - [Reauthorize execution log publication](#reauthorize-execution-log-publication)
      - [Create execution log read token](#create-execution-log-read-token)
  - [`Machines Autoresizing`](#machines-autoresizing)
    - [Read this machine's RAM autoresizing settings](#read-this-machines-ram-autoresizing-settings)
    - [Set this machine's RAM autoresizing settings](#set-this-machines-ram-autoresizing-settings)
- [`Organization`](#organization)
  - [`Organization Autoresizing`](#organization-autoresizing)
    - [Read organization RAM autoresizing policy](#read-organization-ram-autoresizing-policy)
    - [Set organization RAM autoresizing policy](#set-organization-ram-autoresizing-policy)

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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID:    "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Sleep a running machine

| Direction | Type |
| --- | --- |
| Request | [`MachineSleepParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Sleep(context.Background(), sdk.MachineSleepParams{
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(machine.MachineID)
```

### Reboot a machine with fresh memory

Checkpoints files and replaces the runtime. The machine ID and filesystem are preserved. RAM, processes, and temporary mounts are cleared. Poll the machine until its phase is running. Retry the same Idempotency-Key after a lost response.

| Direction | Type |
| --- | --- |
| Request | [`MachineRebootParams`](./machine.go) |
| Response | [`Machine`](./machine.go) |

```go
machine, err := client.Machines.Reboot(context.Background(), sdk.MachineRebootParams{
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(machine)
```

### `Machines Ssh`

#### List SSH sessions

| Direction | Type |
| --- | --- |
| Request | [`MachineSSHListParams`](./machinessh.go) |
| Response | [`SSHSession`](./machinessh.go) |

```go
page, err := client.Machines.SSH.List(context.Background(), sdk.MachineSSHListParams{
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
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
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

#### `Machines Executions Logs`

##### Get execution log status

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionLogGetParams`](./machineexecutionlog.go) |
| Response | [`Status`](./machineexecutionlog.go) |

```go
log, err := client.Machines.Executions.Logs.Get(context.Background(), sdk.MachineExecutionLogGetParams{
	ExecutionID: "executionID",
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(log)
```

##### Reauthorize execution log publication

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionLogReauthorizeParams`](./machineexecutionlog.go) |
| Response | [`Status`](./machineexecutionlog.go) |

```go
log, err := client.Machines.Executions.Logs.Reauthorize(context.Background(), sdk.MachineExecutionLogReauthorizeParams{
	ExecutionID: "executionID",
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(log)
```

##### Create execution log read token

| Direction | Type |
| --- | --- |
| Request | [`MachineExecutionLogNewTokenParams`](./machineexecutionlog.go) |
| Response | [`ReadToken`](./machineexecutionlog.go) |

```go
log, err := client.Machines.Executions.Logs.NewToken(context.Background(), sdk.MachineExecutionLogNewTokenParams{
	ExecutionID: "executionID",
	MachineID:   "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(log)
```

### `Machines Autoresizing`

#### Read this machine's RAM autoresizing settings

| Direction | Type |
| --- | --- |
| Request | [`MachineAutoresizingGetParams`](./machineautoresizing.go) |
| Response | [`Settings`](./machineautoresizing.go) |

```go
autoresizing, err := client.Machines.Autoresizing.Get(context.Background(), sdk.MachineAutoresizingGetParams{
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
})
if err != nil {
	panic(err)
}

fmt.Println(autoresizing)
```

#### Set this machine's RAM autoresizing settings

| Direction | Type |
| --- | --- |
| Request | [`MachineAutoresizingUpdateParams`](./machineautoresizing.go) |
| Response | [`Settings`](./machineautoresizing.go) |

```go
autoresizing, err := client.Machines.Autoresizing.Update(context.Background(), sdk.MachineAutoresizingUpdateParams{
	MachineID: "017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
	Settings: sdk.SettingsParam{
		Enabled: sdk.F[bool](false),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(autoresizing)
```

## `Organization`

### `Organization Autoresizing`

#### Read organization RAM autoresizing policy

| Direction | Type |
| --- | --- |
| Response | [`Policy`](./organizationautoresizing.go) |

```go
autoresizing, err := client.Organization.Autoresizing.Get(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(autoresizing)
```

#### Set organization RAM autoresizing policy

| Direction | Type |
| --- | --- |
| Request | [`OrganizationAutoresizingUpdateParams`](./organizationautoresizing.go) |
| Response | [`Policy`](./organizationautoresizing.go) |

```go
autoresizing, err := client.Organization.Autoresizing.Update(context.Background(), sdk.OrganizationAutoresizingUpdateParams{
	Policy: sdk.PolicyParam{
		Enabled: sdk.F[bool](false),
	},
})
if err != nil {
	panic(err)
}

fmt.Println(autoresizing)
```
