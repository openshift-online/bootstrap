package clusters

import (
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"github.com/openshift/hive/apis/hive/v1/aws"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewMachinePool(config *ClusterDeploymentConfig) hivev1.MachinePool {
	return hivev1.MachinePool{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "hive.openshift.io/v1",
			Kind:       "MachinePool",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      config.Name + "-worker",
			Namespace: config.Name,
		},
		Spec: hivev1.MachinePoolSpec{
			ClusterDeploymentRef: corev1.LocalObjectReference{
				Name: config.Name,
			},
			Name:     "worker",
			Replicas: &config.WorkerReplicas,
			Platform: hivev1.MachinePoolPlatform{
				AWS: &aws.MachinePoolPlatform{
					InstanceType: config.WorkerInstanceType,
					Zones: []string{
						config.Region,
					},
					EC2RootVolume: aws.EC2RootVolume{
						IOPS: 2000,
						Size: 100,
						Type: "io1",
					},
				},
			},
		},
	}
}
