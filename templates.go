package main

const (
	templateName1 = "aws_i3en_2xlarge"
	templateName2 = "aws_i4i_2xlarge"
	templateName3 = "aws_i5i_2xlarge"
	templateName4 = "aws_i6i_2xlarge"
)

type spec map[string]float64

var templates = map[string]spec{
	templateName1: {
		storage_pct_metric_name: 50,
	},
	templateName2: {
		storage_pct_metric_name: 40,
	},
	templateName3: {
		storage_pct_metric_name: 50,
	},
	templateName4: {
		storage_pct_metric_name: 0,
	},
}
