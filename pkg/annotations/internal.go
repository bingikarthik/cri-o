package annotations

const (
	// Annotations carries the received Kubelet annotations.
	Annotations = "io.kubernetes.cri-o.Annotations"

	// ContainerID is the container ID annotation.
	ContainerID = "io.kubernetes.cri-o.ContainerID"

	// ContainerName is the container name annotation.
	ContainerName = "io.kubernetes.cri-o.ContainerName"

	// ContainerType is the container type (sandbox or container) annotation.
	ContainerType = "io.kubernetes.cri-o.ContainerType"

	// Created is the container creation time annotation.
	Created = "io.kubernetes.cri-o.Created"

	// HostName is the container host name annotation.
	HostName = "io.kubernetes.cri-o.HostName"

	// CgroupParent is the sandbox cgroup parent.
	CgroupParent = "io.kubernetes.cri-o.CgroupParent"

	// IP is the container ipv4 or ipv6 address.
	IP = "io.kubernetes.cri-o.IP"

	// NamespaceOptions store the options for namespaces.
	NamespaceOptions = "io.kubernetes.cri-o.NamespaceOptions"

	// SeccompProfilePath is the node seccomp profile path.
	SeccompProfilePath = "io.kubernetes.cri-o.SeccompProfilePath"

	// UserRequestedImage is an annotation containing the image specified in the container spec
	// and used to look up the image when creating the container.
	// It might evaluate to a different image (or to a different kind of reference!) at any future time.
	UserRequestedImage = "io.kubernetes.cri-o.Image"

	// SomeNameOfTheImage is an annotation containing _some_ name of the image used to create the container.
	// May be "" if the image was referenced by ID and has no names.
	// It also has NO RELATIONSHIP to user input used to find the image.
	SomeNameOfTheImage = "io.kubernetes.cri-o.ImageName"

	// ImageRef is the container image ref annotation.
	ImageRef = "io.kubernetes.cri-o.ImageRef"

	// KubeName is the kubernetes name annotation.
	KubeName = "io.kubernetes.cri-o.KubeName"

	// PortMappings holds the port mappings for the sandbox.
	PortMappings = "io.kubernetes.cri-o.PortMappings"

	// Labels are the kubernetes labels annotation.
	Labels = "io.kubernetes.cri-o.Labels"

	// LogPath is the container logging path annotation.
	LogPath = "io.kubernetes.cri-o.LogPath"

	// Metadata is the container metadata annotation.
	Metadata = "io.kubernetes.cri-o.Metadata"

	// Name is the pod name annotation.
	Name = "io.kubernetes.cri-o.Name"

	// Namespace is the pod namespace annotation.
	Namespace = "io.kubernetes.cri-o.Namespace"

	// PrivilegedRuntime is the annotation for the privileged runtime path.
	PrivilegedRuntime = "io.kubernetes.cri-o.PrivilegedRuntime"

	// ResolvPath is the resolver configuration path annotation.
	ResolvPath = "io.kubernetes.cri-o.ResolvPath"

	// HostnamePath is the path to /etc/hostname to bind mount annotation.
	HostnamePath = "io.kubernetes.cri-o.HostnamePath"

	// SandboxID is the sandbox ID annotation.
	SandboxID = "io.kubernetes.cri-o.SandboxID"

	// SandboxName is the sandbox name annotation.
	SandboxName = "io.kubernetes.cri-o.SandboxName"

	// ShmPath is the shared memory path annotation.
	ShmPath = "io.kubernetes.cri-o.ShmPath"

	// MountPoint is the mount point of the container rootfs.
	MountPoint = "io.kubernetes.cri-o.MountPoint"

	// RuntimeHandler is the annotation for runtime handler.
	RuntimeHandler = "io.kubernetes.cri-o.RuntimeHandler"

	// TTY is the terminal path annotation.
	TTY = "io.kubernetes.cri-o.TTY"

	// Stdin is the stdin annotation.
	Stdin = "io.kubernetes.cri-o.Stdin"

	// StdinOnce is the stdin_once annotation.
	StdinOnce = "io.kubernetes.cri-o.StdinOnce"

	// Volumes is the volumes annotation.
	Volumes = "io.kubernetes.cri-o.Volumes"

	// HostNetwork indicates whether the host network namespace is used or not.
	HostNetwork = "io.kubernetes.cri-o.HostNetwork"

	// CNIResult is the JSON string representation of the Result from CNI.
	CNIResult = "io.kubernetes.cri-o.CNIResult"

	// ContainerManager is the annotation key for indicating the creator and
	// manager of the container.
	ContainerManager = "io.container.manager"
)

// ContainerType values.
const (
	// ContainerTypeSandbox represents a pod sandbox container.
	ContainerTypeSandbox = "sandbox"

	// ContainerTypeContainer represents a container running within a pod.
	ContainerTypeContainer = "container"
)

// ContainerManagerLibpod indicates that libpod created and manages the
// container.
const ContainerManagerLibpod = "libpod"
