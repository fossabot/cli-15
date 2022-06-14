package k8s

// https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
const (
	LabelVersion   = "app.kubernetes.io/version"
	LabelPartOf    = "app.kubernetes.io/part-of"
	LabelManagedBy = "app.kubernetes.io/managed-by"
	LabelCreatedBy = "app.kubernetes.io/created-by"
)

const (
	LabelProjectID      = "calyptia_project_id"
	LabelAggregatorID   = "aggregatorID"
	LabelAggregatorName = "calyptia_aggregator_name"
	LabelPipelineID     = "calyptia_pipeline_id"
	LabelPipelineName   = "calyptia_pipeline_name"
)
