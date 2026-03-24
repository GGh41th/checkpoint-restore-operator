# Container Checkpoint CRD:


**ContainerCheckPoint** is a namespaced resource that represents the intention to checkpoint a container.


## ContainerCheckpoint Spec
```go
type ContainerCheckpointSpec struct {

	PodName string `json:"podname"`

	ContainerName string `json:"containerName,omitempty"`

	NameSpace string `json:"namespace,omitempty"`

	// Compression field is yet to be discussed
	// Compression string `json:"compression,omitempty"`


	Policy CheckpointPolicy
}
```
***Fields Description***

**PodName**: The pod of the container to be checkpointed\
**ContainerName**: The container to be checkpointed\
**Namespace**: The pod's namespace\
**Compression**: If we can do it efficiently, we might compress the checkpoints.\
**Policy**: A map of mutually execlusive fields that specify the trigger for the checkpoint process.


```go
type CheckpointPolicy struct {

	Schedule string `json:"schedule,omitempty"`

	Resources corev1.ResourceList `json:"resources,omitempty"`

	NodeConditions []Condition `json:"nodeConditions,omitempty"`

	OnDrain bool `json:"onDrain,omitempty"`
}

type Condition struct {

	Type   string

	Status metav1.ConditionStatus
}

type ResourceThreshold struct {
	cpu    string
	memory string
}
```


## ContainerCheckpoint Status

```go
type ContainerCheckpointStatus struct {

	Phase ContainerCheckpointPhase `json:"phase,omitempty"`

	CheckpointPath string `json:"checkpointPath,omitempty"`

	NodeName string `json:"nodeName,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type ContainerCheckpointPhase string

const (
	ContainerCheckpointPending    ContainerCheckpointPhase = "Pending"
	ContainerCheckPointInProgress ContainerCheckpointPhase = "InProgress"
	ContainerCheckpointReady      ContainerCheckpointPhase = "Ready"
	ContainerCheckpointFailed     ContainerCheckpointPhase = "Failed"
)

```
***Fields Description***

**Phase**: The checkpointing phase, possible values are:

- Pending: Checkpointing hasn't started yet
- InProgress: Checkpointing in progress.
- Ready: Checkpointing has finished succesfully.
- Failed: Checkpointing has failed with a fatal error.

**CheckpointPath**: The checkpoint path on disk (the item field in the json response)

**NodeName**: The pod's node name.


### Alternative Implementations:
Another possible implementation is to allow users to specify higher level constructs (Deployments,DaemonSets,StatefulSets) where all replicas will be checkpointed. This remains open to discuss with mentors.
