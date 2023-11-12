package response

import "time"

type Deployment struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
	} `json:"metadata"`
	Items []struct {
		Kind       string `json:"kind"`
		APIVersion string `json:"apiVersion"`
		Metadata   struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Run string `json:"run"`
			} `json:"labels"`
			Annotations struct {
				DeploymentKubernetesIoRevision    string `json:"deployment.kubernetes.io/revision"`
				ReplicatingperfectionNetPushImage string `json:"replicatingperfection.net/push-image"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Selector struct {
				MatchLabels struct {
					Run string `json:"run"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						AutoPushedImagePwittrockAPIDocs string `json:"auto-pushed-image-pwittrock/api-docs"`
						Run                             string `json:"run"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					Containers []struct {
						Name      string `json:"name"`
						Image     string `json:"image"`
						Resources struct {
						} `json:"resources"`
						TerminationMessagePath string `json:"terminationMessagePath"`
						ImagePullPolicy        string `json:"imagePullPolicy"`
					} `json:"containers"`
					RestartPolicy                 string `json:"restartPolicy"`
					TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
					DNSPolicy                     string `json:"dnsPolicy"`
					SecurityContext               struct {
					} `json:"securityContext"`
				} `json:"spec"`
			} `json:"template"`
			Strategy struct {
				Type          string `json:"type"`
				RollingUpdate struct {
					MaxUnavailable int `json:"maxUnavailable"`
					MaxSurge       int `json:"maxSurge"`
				} `json:"rollingUpdate"`
			} `json:"strategy"`
		} `json:"spec"`
		Status struct {
			ObservedGeneration int `json:"observedGeneration"`
			Replicas           int `json:"replicas"`
			UpdatedReplicas    int `json:"updatedReplicas"`
			AvailableReplicas  int `json:"availableReplicas"`
		} `json:"status"`
	} `json:"items"`
}
