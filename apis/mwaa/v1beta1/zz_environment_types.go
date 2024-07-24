// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DagProcessingLogsInitParameters struct {

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type DagProcessingLogsObservation struct {

	// Provides the ARN for the CloudWatch group where the logs will be published
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn,omitempty" tf:"cloud_watch_log_group_arn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type DagProcessingLogsParameters struct {

	// Enabling or disabling the collection of logs
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type EnvironmentInitParameters struct {
	AirflowConfigurationOptions map[string]*string `json:"airflowConfigurationOptionsSecretRef,omitempty" tf:"-"`

	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion *string `json:"airflowVersion,omitempty" tf:"airflow_version,omitempty"`

	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see Importing DAGs on Amazon MWAA.
	DagS3Path *string `json:"dagS3Path,omitempty" tf:"dag_s3_path,omitempty"`

	EndpointManagement *string `json:"endpointManagement,omitempty" tf:"endpoint_management,omitempty"`

	// Environment class for the cluster. Possible options are mw1.small, mw1.medium, mw1.large. Will be set by default to mw1.small. Please check the AWS Pricing for more information about the environment classes.
	EnvironmentClass *string `json:"environmentClass,omitempty" tf:"environment_class,omitempty"`

	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the official AWS documentation for the detailed role specification.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key aws/airflow by default. Please check the Official Documentation for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`

	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration *LoggingConfigurationInitParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The maximum number of workers that can be automatically scaled up. Value need to be between 1 and 25. Will be 10 by default.
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// The minimum number of workers that you want to run in your environment. Will be 1 by default.
	MinWorkers *float64 `json:"minWorkers,omitempty" tf:"min_workers,omitempty"`

	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration *NetworkConfigurationInitParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion,omitempty" tf:"plugins_s3_object_version,omitempty"`

	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	PluginsS3Path *string `json:"pluginsS3Path,omitempty" tf:"plugins_s3_path,omitempty"`

	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion,omitempty" tf:"requirements_s3_object_version,omitempty"`

	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	RequirementsS3Path *string `json:"requirementsS3Path,omitempty" tf:"requirements_s3_path,omitempty"`

	// The number of schedulers that you want to run in your environment. v2.0.2 and above accepts 2 - 5, default 2. v1.10.12 accepts 1.
	Schedulers *float64 `json:"schedulers,omitempty" tf:"schedulers,omitempty"`

	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SourceBucketArn *string `json:"sourceBucketArn,omitempty" tf:"source_bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate sourceBucketArn.
	// +kubebuilder:validation:Optional
	SourceBucketArnRef *v1.Reference `json:"sourceBucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate sourceBucketArn.
	// +kubebuilder:validation:Optional
	SourceBucketArnSelector *v1.Selector `json:"sourceBucketArnSelector,omitempty" tf:"-"`

	// The version of the startup shell script you want to use. You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	StartupScriptS3ObjectVersion *string `json:"startupScriptS3ObjectVersion,omitempty" tf:"startup_script_s3_object_version,omitempty"`

	// The relative path to the script hosted in your bucket. The script runs as your environment starts before starting the Apache Airflow process. Use this script to install dependencies, modify configuration options, and set environment variables. See Using a startup script. Supported for environment versions 2.x and later.
	StartupScriptS3Path *string `json:"startupScriptS3Path,omitempty" tf:"startup_script_s3_path,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: PRIVATE_ONLY (default) and PUBLIC_ONLY.
	WebserverAccessMode *string `json:"webserverAccessMode,omitempty" tf:"webserver_access_mode,omitempty"`

	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart,omitempty" tf:"weekly_maintenance_window_start,omitempty"`
}

type EnvironmentObservation struct {

	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion *string `json:"airflowVersion,omitempty" tf:"airflow_version,omitempty"`

	// The ARN of the MWAA Environment
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Created At date of the MWAA Environment
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see Importing DAGs on Amazon MWAA.
	DagS3Path *string `json:"dagS3Path,omitempty" tf:"dag_s3_path,omitempty"`

	// The VPC endpoint for the environment's Amazon RDS database
	DatabaseVPCEndpointService *string `json:"databaseVpcEndpointService,omitempty" tf:"database_vpc_endpoint_service,omitempty"`

	EndpointManagement *string `json:"endpointManagement,omitempty" tf:"endpoint_management,omitempty"`

	// Environment class for the cluster. Possible options are mw1.small, mw1.medium, mw1.large. Will be set by default to mw1.small. Please check the AWS Pricing for more information about the environment classes.
	EnvironmentClass *string `json:"environmentClass,omitempty" tf:"environment_class,omitempty"`

	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the official AWS documentation for the detailed role specification.
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key aws/airflow by default. Please check the Official Documentation for more information.
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	LastUpdated []LastUpdatedObservation `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`

	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration *LoggingConfigurationObservation `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The maximum number of workers that can be automatically scaled up. Value need to be between 1 and 25. Will be 10 by default.
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// The minimum number of workers that you want to run in your environment. Will be 1 by default.
	MinWorkers *float64 `json:"minWorkers,omitempty" tf:"min_workers,omitempty"`

	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration *NetworkConfigurationObservation `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion,omitempty" tf:"plugins_s3_object_version,omitempty"`

	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	PluginsS3Path *string `json:"pluginsS3Path,omitempty" tf:"plugins_s3_path,omitempty"`

	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion,omitempty" tf:"requirements_s3_object_version,omitempty"`

	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	RequirementsS3Path *string `json:"requirementsS3Path,omitempty" tf:"requirements_s3_path,omitempty"`

	// The number of schedulers that you want to run in your environment. v2.0.2 and above accepts 2 - 5, default 2. v1.10.12 accepts 1.
	Schedulers *float64 `json:"schedulers,omitempty" tf:"schedulers,omitempty"`

	// The Service Role ARN of the Amazon MWAA Environment
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn *string `json:"sourceBucketArn,omitempty" tf:"source_bucket_arn,omitempty"`

	// The version of the startup shell script you want to use. You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	StartupScriptS3ObjectVersion *string `json:"startupScriptS3ObjectVersion,omitempty" tf:"startup_script_s3_object_version,omitempty"`

	// The relative path to the script hosted in your bucket. The script runs as your environment starts before starting the Apache Airflow process. Use this script to install dependencies, modify configuration options, and set environment variables. See Using a startup script. Supported for environment versions 2.x and later.
	StartupScriptS3Path *string `json:"startupScriptS3Path,omitempty" tf:"startup_script_s3_path,omitempty"`

	// The status of the Amazon MWAA Environment
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: PRIVATE_ONLY (default) and PUBLIC_ONLY.
	WebserverAccessMode *string `json:"webserverAccessMode,omitempty" tf:"webserver_access_mode,omitempty"`

	// The webserver URL of the MWAA Environment
	WebserverURL *string `json:"webserverUrl,omitempty" tf:"webserver_url,omitempty"`

	// The VPC endpoint for the environment's web server
	WebserverVPCEndpointService *string `json:"webserverVpcEndpointService,omitempty" tf:"webserver_vpc_endpoint_service,omitempty"`

	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart,omitempty" tf:"weekly_maintenance_window_start,omitempty"`
}

type EnvironmentParameters struct {

	// The airflow_configuration_options parameter specifies airflow override options. Check the Official documentation for all possible configuration options.
	// +kubebuilder:validation:Optional
	AirflowConfigurationOptionsSecretRef *v1.SecretReference `json:"airflowConfigurationOptionsSecretRef,omitempty" tf:"-"`

	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	// +kubebuilder:validation:Optional
	AirflowVersion *string `json:"airflowVersion,omitempty" tf:"airflow_version,omitempty"`

	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see Importing DAGs on Amazon MWAA.
	// +kubebuilder:validation:Optional
	DagS3Path *string `json:"dagS3Path,omitempty" tf:"dag_s3_path,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointManagement *string `json:"endpointManagement,omitempty" tf:"endpoint_management,omitempty"`

	// Environment class for the cluster. Possible options are mw1.small, mw1.medium, mw1.large. Will be set by default to mw1.small. Please check the AWS Pricing for more information about the environment classes.
	// +kubebuilder:validation:Optional
	EnvironmentClass *string `json:"environmentClass,omitempty" tf:"environment_class,omitempty"`

	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the official AWS documentation for the detailed role specification.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key aws/airflow by default. Please check the Official Documentation for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`

	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	// +kubebuilder:validation:Optional
	LoggingConfiguration *LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The maximum number of workers that can be automatically scaled up. Value need to be between 1 and 25. Will be 10 by default.
	// +kubebuilder:validation:Optional
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// The minimum number of workers that you want to run in your environment. Will be 1 by default.
	// +kubebuilder:validation:Optional
	MinWorkers *float64 `json:"minWorkers,omitempty" tf:"min_workers,omitempty"`

	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	// +kubebuilder:validation:Optional
	NetworkConfiguration *NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// The plugins.zip file version you want to use.
	// +kubebuilder:validation:Optional
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion,omitempty" tf:"plugins_s3_object_version,omitempty"`

	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	// +kubebuilder:validation:Optional
	PluginsS3Path *string `json:"pluginsS3Path,omitempty" tf:"plugins_s3_path,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The requirements.txt file version you want to use.
	// +kubebuilder:validation:Optional
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion,omitempty" tf:"requirements_s3_object_version,omitempty"`

	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see Importing DAGs on Amazon MWAA.
	// +kubebuilder:validation:Optional
	RequirementsS3Path *string `json:"requirementsS3Path,omitempty" tf:"requirements_s3_path,omitempty"`

	// The number of schedulers that you want to run in your environment. v2.0.2 and above accepts 2 - 5, default 2. v1.10.12 accepts 1.
	// +kubebuilder:validation:Optional
	Schedulers *float64 `json:"schedulers,omitempty" tf:"schedulers,omitempty"`

	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SourceBucketArn *string `json:"sourceBucketArn,omitempty" tf:"source_bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate sourceBucketArn.
	// +kubebuilder:validation:Optional
	SourceBucketArnRef *v1.Reference `json:"sourceBucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate sourceBucketArn.
	// +kubebuilder:validation:Optional
	SourceBucketArnSelector *v1.Selector `json:"sourceBucketArnSelector,omitempty" tf:"-"`

	// The version of the startup shell script you want to use. You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	// +kubebuilder:validation:Optional
	StartupScriptS3ObjectVersion *string `json:"startupScriptS3ObjectVersion,omitempty" tf:"startup_script_s3_object_version,omitempty"`

	// The relative path to the script hosted in your bucket. The script runs as your environment starts before starting the Apache Airflow process. Use this script to install dependencies, modify configuration options, and set environment variables. See Using a startup script. Supported for environment versions 2.x and later.
	// +kubebuilder:validation:Optional
	StartupScriptS3Path *string `json:"startupScriptS3Path,omitempty" tf:"startup_script_s3_path,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: PRIVATE_ONLY (default) and PUBLIC_ONLY.
	// +kubebuilder:validation:Optional
	WebserverAccessMode *string `json:"webserverAccessMode,omitempty" tf:"webserver_access_mode,omitempty"`

	// Specifies the start date for the weekly maintenance window.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart,omitempty" tf:"weekly_maintenance_window_start,omitempty"`
}

type ErrorInitParameters struct {
}

type ErrorObservation struct {
	ErrorCode *string `json:"errorCode,omitempty" tf:"error_code,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`
}

type ErrorParameters struct {
}

type LastUpdatedInitParameters struct {
}

type LastUpdatedObservation struct {

	// The Created At date of the MWAA Environment
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Error []ErrorObservation `json:"error,omitempty" tf:"error,omitempty"`

	// The status of the Amazon MWAA Environment
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type LastUpdatedParameters struct {
}

type LoggingConfigurationInitParameters struct {

	// Log configuration options for processing DAGs. See Module logging configuration for more information. Disabled by default.
	DagProcessingLogs *DagProcessingLogsInitParameters `json:"dagProcessingLogs,omitempty" tf:"dag_processing_logs,omitempty"`

	// Log configuration options for the schedulers. See Module logging configuration for more information. Disabled by default.
	SchedulerLogs *SchedulerLogsInitParameters `json:"schedulerLogs,omitempty" tf:"scheduler_logs,omitempty"`

	// Log configuration options for DAG tasks. See Module logging configuration for more information. Enabled by default with INFO log level.
	TaskLogs *TaskLogsInitParameters `json:"taskLogs,omitempty" tf:"task_logs,omitempty"`

	// Log configuration options for the webservers. See Module logging configuration for more information. Disabled by default.
	WebserverLogs *WebserverLogsInitParameters `json:"webserverLogs,omitempty" tf:"webserver_logs,omitempty"`

	// Log configuration options for the workers. See Module logging configuration for more information. Disabled by default.
	WorkerLogs *WorkerLogsInitParameters `json:"workerLogs,omitempty" tf:"worker_logs,omitempty"`
}

type LoggingConfigurationObservation struct {

	// Log configuration options for processing DAGs. See Module logging configuration for more information. Disabled by default.
	DagProcessingLogs *DagProcessingLogsObservation `json:"dagProcessingLogs,omitempty" tf:"dag_processing_logs,omitempty"`

	// Log configuration options for the schedulers. See Module logging configuration for more information. Disabled by default.
	SchedulerLogs *SchedulerLogsObservation `json:"schedulerLogs,omitempty" tf:"scheduler_logs,omitempty"`

	// Log configuration options for DAG tasks. See Module logging configuration for more information. Enabled by default with INFO log level.
	TaskLogs *TaskLogsObservation `json:"taskLogs,omitempty" tf:"task_logs,omitempty"`

	// Log configuration options for the webservers. See Module logging configuration for more information. Disabled by default.
	WebserverLogs *WebserverLogsObservation `json:"webserverLogs,omitempty" tf:"webserver_logs,omitempty"`

	// Log configuration options for the workers. See Module logging configuration for more information. Disabled by default.
	WorkerLogs *WorkerLogsObservation `json:"workerLogs,omitempty" tf:"worker_logs,omitempty"`
}

type LoggingConfigurationParameters struct {

	// Log configuration options for processing DAGs. See Module logging configuration for more information. Disabled by default.
	// +kubebuilder:validation:Optional
	DagProcessingLogs *DagProcessingLogsParameters `json:"dagProcessingLogs,omitempty" tf:"dag_processing_logs,omitempty"`

	// Log configuration options for the schedulers. See Module logging configuration for more information. Disabled by default.
	// +kubebuilder:validation:Optional
	SchedulerLogs *SchedulerLogsParameters `json:"schedulerLogs,omitempty" tf:"scheduler_logs,omitempty"`

	// Log configuration options for DAG tasks. See Module logging configuration for more information. Enabled by default with INFO log level.
	// +kubebuilder:validation:Optional
	TaskLogs *TaskLogsParameters `json:"taskLogs,omitempty" tf:"task_logs,omitempty"`

	// Log configuration options for the webservers. See Module logging configuration for more information. Disabled by default.
	// +kubebuilder:validation:Optional
	WebserverLogs *WebserverLogsParameters `json:"webserverLogs,omitempty" tf:"webserver_logs,omitempty"`

	// Log configuration options for the workers. See Module logging configuration for more information. Disabled by default.
	// +kubebuilder:validation:Optional
	WorkerLogs *WorkerLogsParameters `json:"workerLogs,omitempty" tf:"worker_logs,omitempty"`
}

type NetworkConfigurationInitParameters struct {

	// Security groups IDs for the environment. At least one of the security group needs to allow MWAA resources to talk to each other, otherwise MWAA cannot be provisioned.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// The private subnet IDs in which the environment should be created. MWAA requires two subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type NetworkConfigurationObservation struct {

	// Security groups IDs for the environment. At least one of the security group needs to allow MWAA resources to talk to each other, otherwise MWAA cannot be provisioned.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The private subnet IDs in which the environment should be created. MWAA requires two subnets.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type NetworkConfigurationParameters struct {

	// Security groups IDs for the environment. At least one of the security group needs to allow MWAA resources to talk to each other, otherwise MWAA cannot be provisioned.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// The private subnet IDs in which the environment should be created. MWAA requires two subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type SchedulerLogsInitParameters struct {

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type SchedulerLogsObservation struct {

	// Provides the ARN for the CloudWatch group where the logs will be published
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn,omitempty" tf:"cloud_watch_log_group_arn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type SchedulerLogsParameters struct {

	// Enabling or disabling the collection of logs
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type TaskLogsInitParameters struct {

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type TaskLogsObservation struct {

	// Provides the ARN for the CloudWatch group where the logs will be published
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn,omitempty" tf:"cloud_watch_log_group_arn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type TaskLogsParameters struct {

	// Enabling or disabling the collection of logs
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WebserverLogsInitParameters struct {

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WebserverLogsObservation struct {

	// Provides the ARN for the CloudWatch group where the logs will be published
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn,omitempty" tf:"cloud_watch_log_group_arn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WebserverLogsParameters struct {

	// Enabling or disabling the collection of logs
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WorkerLogsInitParameters struct {

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WorkerLogsObservation struct {

	// Provides the ARN for the CloudWatch group where the logs will be published
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn,omitempty" tf:"cloud_watch_log_group_arn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type WorkerLogsParameters struct {

	// Enabling or disabling the collection of logs
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Logging level. Valid values: CRITICAL, ERROR, WARNING, INFO, DEBUG. Will be INFO by default.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Environment is the Schema for the Environments API. Creates a MWAA Environment
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dagS3Path) || (has(self.initProvider) && has(self.initProvider.dagS3Path))",message="spec.forProvider.dagS3Path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkConfiguration) || (has(self.initProvider) && has(self.initProvider.networkConfiguration))",message="spec.forProvider.networkConfiguration is a required parameter"
	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}