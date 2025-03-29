package config

import (
	"time"
)

var (
	Server = map[string]string{
		"name":      "meshery-meshsync",
		"port":      "11000",
		"version":   "latest",
		"startedat": time.Now().String(),
	}

	DefaultPublishingSubject = "meshery.meshsync.core"
	K8sPublishingSubject = "meshery.meshsync.k8s"

	Pipelines = map[string]PipelineConfigs{
		GlobalResourceKey: []PipelineConfig{
			// Core Resources
			{
				Name:      "namespaces.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "configmaps.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "nodes.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "secrets.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "persistentvolumes.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "persistentvolumeclaims.v1.",
				PublishTo: K8sPublishingSubject,
			},
		},
		LocalResourceKey: []PipelineConfig{
			// Core Resources
			{
				Name:      "replicasets.v1.apps",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "pods.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "services.v1.",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "deployments.v1.apps",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "statefulsets.v1.apps",
				PublishTo: K8sPublishingSubject,
			},
			{
				Name:      "daemonsets.v1.apps",
				PublishTo: K8sPublishingSubject,
			},
			//Added Ingress support
			{
				Name:      "ingresses.v1.networking.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added 
			{
				Name:      "ingressclass.v1.networking.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			// Added endpoint support
			{
				Name:      "endpoints.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added endpointslice support
			{
				Name:      "endpointslices.v1.discovery.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added container support
			{
				Name:	  "container.v1.core",
				PublishTo: K8sPublishingSubject,
			},
			//Added job  support
			{
				Name:	  "job.v1.batch",
				PublishTo: K8sPublishingSubject,
			},
			//Added service APIs support
			{
				Name:	  "service.apis",
				PublishTo: K8sPublishingSubject,
			},
			//Added  csidriver support
			{
				Name:	  "csidriver.v1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added csinode  support
			{
				Name:	  "csinode.v1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added  csistoragecapacity support
			{
				Name:	  "csistoragecapacity.v1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added volume support
			{
				Name:	  "volume.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added volumeattributesclass support
			{
				Name:	  "volumeattributesclass.v1alpha1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added clustertrustbundle support
			{
				Name:	  "clustertrustbundle.v1alpha1.certificates.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added  controllerversion support
			{
				Name:	  "controllerrevision.v1.apps",
				PublishTo: K8sPublishingSubject,
			},
			//Added customresourcedefinition support
			{
				Name:	  "customresourcedefinition.v1.apiextensions.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added  event support
			{
				Name:	  "event.v1.events.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added limitrange support
			{
				Name:	  "limitrange.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added horizontalpodautoscaler support
			{
				Name:	  "horizontalpodautoscaler.v2.autoscaling",
				PublishTo: K8sPublishingSubject,
			},
			//Added mutatingwebhookconfiguration support
			{
				Name:	  "mutatingwebhookconfiguration.v1.admissionregistration.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added podschedulingcontext support
			{
				Name:	  "podschedulingcontext.v1alpha2.resource.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added podtemplate support
			{
				Name:	  "podtemplate.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added poddistruptionbudget support
			{
				Name:	  "poddisruptionbudget.v1.policy",
				PublishTo: K8sPublishingSubject,
			},
			//Added priorityclass support
			{
				Name:	  "priorityclass.v1.scheduling.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added resourceclaim support
			{
				Name:	  "resourceclaim.v1alpha2.resource.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added resourceclaimtemplate support
			{
				Name:	  "resourceclaimtemplate.v1alpha2.resource.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added resourceclass support
			{
				Name:	  "resourceclass.v1alpha2.resource.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added ValidatingWebhookConfiguration support
			{
				Name:	  "validatingwebhookconfiguration.v1.admissionregistration.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added ValidatingAdmissionPolicy support
			{
				Name:	  "validatingadmissionpolicy.v1beta1.admissionregistration.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added ValidatingAdmissionPolicyBinding support
			{
				Name:	  "validatingadmissionpolicybinding.v1beta1.admissionregistration.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added  binding support
			{
				Name:	  "binding.v1.",
				PublishTo: K8sPublishingSubject,
			},			
			//Added certificatesigningrequest support
			{
				Name:	  "certificatesigningrequest.v1.certificates.k8s.io",
				PublishTo: K8sPublishingSubject,
			},		
				
			//Added clusterrolebinding support
			{
				Name:	  "clusterrolebinding.v1.rbac.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added componentstatus support
			{
				Name:	  "componentstatus.v1.",
				PublishTo: K8sPublishingSubject,
			},			
			//Added flowschema support
			{
				Name:	  "flowschema.v1.flowcontrol.apiserver.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added IPAddress support
			{
				Name:	  "ipaddress.v1alpha1.networking.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added Lease support
			{
				Name:	  "lease.v1.coordination.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added LocalSubjectAccessReview support
			{
				Name:	  "localsubjectaccessreview.v1.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added Node support
			{
				Name:	  "node.v1.",
				PublishTo: K8sPublishingSubject,
			},			
			//Added NetworkPolicy support
			{
				Name:	  "networkpolicy.v1.networking.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added PriorityLevelConfiguration support
			{
				Name:	  "prioritylevelconfiguration.v1.flowcontrol.apiserver.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added resourcequota support
			{
				Name:	  "resourcequota.v1.",
				PublishTo: K8sPublishingSubject,
			},			
			//Added Role support
			{
				Name:	  "role.v1.rbac.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added RoleBinding support
			{
				Name:	  "rolebinding.v1.rbac.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added runtimeclass support
			{
				Name:	  "runtimeclass.v1.node.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added SelfSubjectAccessReview support
			{
				Name:	  "selfsubjectaccessreview.v1.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added SelfSubjectReview support
			{
				Name:	  "selfsubjectreview.v1.authentication.k8s.io",
				PublishTo: K8sPublishingSubject,
			},			
			//Added selfsubjectrulesreview support
			{
				Name:	  "selfsubjectrulesreview.v1.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added ServiceAccount support
			{
				Name:	  "serviceaccount.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added ServiceCIDR support
			{
				Name:	  "servicecidr.v1alpha1.networking.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added StorageVersion support
			{
				Name:	  "storageversion.v1alpha1.internal.apiserver.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added subjectaccessreview support
			{
				Name:	  "subjectaccessreview.v1.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added TokenRequest support
			{
				Name:	  "tokenrequest.v1.authentication.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added TokenReview support
			{
				Name:	  "tokenreview.v1.authentication.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			// Added cronJob support
			{
				Name:      "cronjobs.v1.batch",
				PublishTo: K8sPublishingSubject,
			},
			//Added ReplicationController support
			{
				Name:      "replicationcontrollers.v1.",
				PublishTo: K8sPublishingSubject,
			},
			//Added storageClass support
			{
				Name:      "storageclasses.v1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added ClusterRole support
			{
				Name:      "clusterroles.v1.rbac.authorization.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added VolumeAttachment support
			{
				Name:      "volumeattachments.v1.storage.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
			//Added apiservice support
			{
				Name:      "apiservices.v1.apiregistration.k8s.io",
				PublishTo: K8sPublishingSubject,
			},
		},
	}

	Listeners = map[string]ListenerConfig{
		LogStream: {
			Name:           LogStream,
			ConnectionName: "meshsync-logstream",
			PublishTo:      "meshery.meshsync.logs",
		},
		ExecShell: {
			Name:           ExecShell,
			ConnectionName: "meshsync-exec",
			PublishTo:      "meshery.meshsync.exec",
		},
		RequestStream: {
			Name:           RequestStream,
			ConnectionName: "meshsync-request-stream",
			SubscribeTo:    "meshery.meshsync.request",
		},
	}

	DefaultEvents = []string{"ADD", "UPDATE", "DELETE"}
)
