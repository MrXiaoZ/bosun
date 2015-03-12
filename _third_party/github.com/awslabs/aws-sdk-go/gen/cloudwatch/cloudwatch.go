// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatch provides a client for Amazon CloudWatch.
package cloudwatch

import (
	"net/http"
	"time"

	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/aws"
	"bosun.org/_third_party/github.com/awslabs/aws-sdk-go/gen/endpoints"
)

import (
	"encoding/xml"
	"io"
)

// CloudWatch is a client for Amazon CloudWatch.
type CloudWatch struct {
	client *aws.QueryClient
}

// New returns a new CloudWatch client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *CloudWatch {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("monitoring", region)

	return &CloudWatch{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-08-01",
		},
	}
}

// DeleteAlarms deletes all specified alarms. In the event of an error, no
// alarms are deleted.
func (c *CloudWatch) DeleteAlarms(req *DeleteAlarmsInput) (err error) {
	// NRE
	err = c.client.Do("DeleteAlarms", "POST", "/", req, nil)
	return
}

// DescribeAlarmHistory retrieves history for the specified alarm. Filter
// alarms by date range or item type. If an alarm name is not specified,
// Amazon CloudWatch returns histories for all of the owner's alarms.
func (c *CloudWatch) DescribeAlarmHistory(req *DescribeAlarmHistoryInput) (resp *DescribeAlarmHistoryResult, err error) {
	resp = &DescribeAlarmHistoryResult{}
	err = c.client.Do("DescribeAlarmHistory", "POST", "/", req, resp)
	return
}

// DescribeAlarms retrieves alarms with the specified names. If no name is
// specified, all alarms for the user are returned. Alarms can be retrieved
// by using only a prefix for the alarm name, the alarm state, or a prefix
// for any action.
func (c *CloudWatch) DescribeAlarms(req *DescribeAlarmsInput) (resp *DescribeAlarmsResult, err error) {
	resp = &DescribeAlarmsResult{}
	err = c.client.Do("DescribeAlarms", "POST", "/", req, resp)
	return
}

// DescribeAlarmsForMetric retrieves all alarms for a single metric.
// Specify a statistic, period, or unit to filter the set of alarms
// further.
func (c *CloudWatch) DescribeAlarmsForMetric(req *DescribeAlarmsForMetricInput) (resp *DescribeAlarmsForMetricResult, err error) {
	resp = &DescribeAlarmsForMetricResult{}
	err = c.client.Do("DescribeAlarmsForMetric", "POST", "/", req, resp)
	return
}

// DisableAlarmActions disables actions for the specified alarms. When an
// alarm's actions are disabled the alarm's state may change, but none of
// the alarm's actions will execute.
func (c *CloudWatch) DisableAlarmActions(req *DisableAlarmActionsInput) (err error) {
	// NRE
	err = c.client.Do("DisableAlarmActions", "POST", "/", req, nil)
	return
}

// EnableAlarmActions is undocumented.
func (c *CloudWatch) EnableAlarmActions(req *EnableAlarmActionsInput) (err error) {
	// NRE
	err = c.client.Do("EnableAlarmActions", "POST", "/", req, nil)
	return
}

// GetMetricStatistics gets statistics for the specified metric. The
// maximum number of data points returned from a single GetMetricStatistics
// request is 1,440, wereas the maximum number of data points that can be
// queried is 50,850. If you make a request that generates more than 1,440
// data points, Amazon CloudWatch returns an error. In such a case, you can
// alter the request by narrowing the specified time range or increasing
// the specified period. Alternatively, you can make multiple requests
// across adjacent time ranges. Amazon CloudWatch aggregates data points
// based on the length of the period that you specify. For example, if you
// request statistics with a one-minute granularity, Amazon CloudWatch
// aggregates data points with time stamps that fall within the same
// one-minute period. In such a case, the data points queried can greatly
// outnumber the data points returned. The following examples show various
// statistics allowed by the data point query maximum of 50,850 when you
// call GetMetricStatistics on Amazon EC2 instances with detailed
// (one-minute) monitoring enabled: Statistics for up to 400 instances for
// a span of one hour Statistics for up to 35 instances over a span of 24
// hours Statistics for up to 2 instances over a span of 2 weeks For
// information about the namespace, metric names, and dimensions that other
// Amazon Web Services products use to send metrics to Cloudwatch, go to
// Amazon CloudWatch Metrics, Namespaces, and Dimensions Reference in the
// Amazon CloudWatch Developer Guide .
func (c *CloudWatch) GetMetricStatistics(req *GetMetricStatisticsInput) (resp *GetMetricStatisticsResult, err error) {
	resp = &GetMetricStatisticsResult{}
	err = c.client.Do("GetMetricStatistics", "POST", "/", req, resp)
	return
}

// ListMetrics returns a list of valid metrics stored for the AWS account
// owner. Returned metrics can be used with GetMetricStatistics to obtain
// statistical data for a given metric.
func (c *CloudWatch) ListMetrics(req *ListMetricsInput) (resp *ListMetricsResult, err error) {
	resp = &ListMetricsResult{}
	err = c.client.Do("ListMetrics", "POST", "/", req, resp)
	return
}

// PutMetricAlarm creates or updates an alarm and associates it with the
// specified Amazon CloudWatch metric. Optionally, this operation can
// associate one or more Amazon Simple Notification Service resources with
// the alarm. When this operation creates an alarm, the alarm state is
// immediately set to . The alarm is evaluated and its StateValue is set
// appropriately. Any actions associated with the StateValue is then
// executed.
func (c *CloudWatch) PutMetricAlarm(req *PutMetricAlarmInput) (err error) {
	// NRE
	err = c.client.Do("PutMetricAlarm", "POST", "/", req, nil)
	return
}

// PutMetricData publishes metric data points to Amazon CloudWatch. Amazon
// Cloudwatch associates the data points with the specified metric. If the
// specified metric does not exist, Amazon CloudWatch creates the metric.
// It can take up to fifteen minutes for a new metric to appear in calls to
// the ListMetrics action. request is limited to 8 KB for GET requests and
// 40 KB for requests. Although the Value parameter accepts numbers of type
// Double , Amazon CloudWatch truncates values with very large exponents.
// Values with base-10 exponents greater than 126 (1 x 10^126) are
// truncated. Likewise, values with base-10 exponents less than -130 (1 x
// 10^-130) are also truncated. Data that is timestamped 24 hours or more
// in the past may take in excess of 48 hours to become available from
// submission time using GetMetricStatistics
func (c *CloudWatch) PutMetricData(req *PutMetricDataInput) (err error) {
	// NRE
	err = c.client.Do("PutMetricData", "POST", "/", req, nil)
	return
}

// SetAlarmState temporarily sets the state of an alarm. When the updated
// StateValue differs from the previous value, the action configured for
// the appropriate state is invoked. This is not a permanent change. The
// next periodic alarm check (in about a minute) will set the alarm to its
// actual state.
func (c *CloudWatch) SetAlarmState(req *SetAlarmStateInput) (err error) {
	// NRE
	err = c.client.Do("SetAlarmState", "POST", "/", req, nil)
	return
}

// AlarmHistoryItem is undocumented.
type AlarmHistoryItem struct {
	AlarmName       aws.StringValue `query:"AlarmName" xml:"AlarmName"`
	HistoryData     aws.StringValue `query:"HistoryData" xml:"HistoryData"`
	HistoryItemType aws.StringValue `query:"HistoryItemType" xml:"HistoryItemType"`
	HistorySummary  aws.StringValue `query:"HistorySummary" xml:"HistorySummary"`
	Timestamp       time.Time       `query:"Timestamp" xml:"Timestamp"`
}

// Possible values for CloudWatch.
const (
	ComparisonOperatorGreaterThanOrEqualToThreshold = "GreaterThanOrEqualToThreshold"
	ComparisonOperatorGreaterThanThreshold          = "GreaterThanThreshold"
	ComparisonOperatorLessThanOrEqualToThreshold    = "LessThanOrEqualToThreshold"
	ComparisonOperatorLessThanThreshold             = "LessThanThreshold"
)

// Datapoint is undocumented.
type Datapoint struct {
	Average     aws.DoubleValue `query:"Average" xml:"Average"`
	Maximum     aws.DoubleValue `query:"Maximum" xml:"Maximum"`
	Minimum     aws.DoubleValue `query:"Minimum" xml:"Minimum"`
	SampleCount aws.DoubleValue `query:"SampleCount" xml:"SampleCount"`
	Sum         aws.DoubleValue `query:"Sum" xml:"Sum"`
	Timestamp   time.Time       `query:"Timestamp" xml:"Timestamp"`
	Unit        aws.StringValue `query:"Unit" xml:"Unit"`
}

// DeleteAlarmsInput is undocumented.
type DeleteAlarmsInput struct {
	AlarmNames []string `query:"AlarmNames.member" xml:"AlarmNames>member"`
}

// DescribeAlarmHistoryInput is undocumented.
type DescribeAlarmHistoryInput struct {
	AlarmName       aws.StringValue  `query:"AlarmName" xml:"AlarmName"`
	EndDate         time.Time        `query:"EndDate" xml:"EndDate"`
	HistoryItemType aws.StringValue  `query:"HistoryItemType" xml:"HistoryItemType"`
	MaxRecords      aws.IntegerValue `query:"MaxRecords" xml:"MaxRecords"`
	NextToken       aws.StringValue  `query:"NextToken" xml:"NextToken"`
	StartDate       time.Time        `query:"StartDate" xml:"StartDate"`
}

// DescribeAlarmHistoryOutput is undocumented.
type DescribeAlarmHistoryOutput struct {
	AlarmHistoryItems []AlarmHistoryItem `query:"AlarmHistoryItems.member" xml:"DescribeAlarmHistoryResult>AlarmHistoryItems>member"`
	NextToken         aws.StringValue    `query:"NextToken" xml:"DescribeAlarmHistoryResult>NextToken"`
}

// DescribeAlarmsForMetricInput is undocumented.
type DescribeAlarmsForMetricInput struct {
	Dimensions []Dimension      `query:"Dimensions.member" xml:"Dimensions>member"`
	MetricName aws.StringValue  `query:"MetricName" xml:"MetricName"`
	Namespace  aws.StringValue  `query:"Namespace" xml:"Namespace"`
	Period     aws.IntegerValue `query:"Period" xml:"Period"`
	Statistic  aws.StringValue  `query:"Statistic" xml:"Statistic"`
	Unit       aws.StringValue  `query:"Unit" xml:"Unit"`
}

// DescribeAlarmsForMetricOutput is undocumented.
type DescribeAlarmsForMetricOutput struct {
	MetricAlarms []MetricAlarm `query:"MetricAlarms.member" xml:"DescribeAlarmsForMetricResult>MetricAlarms>member"`
}

// DescribeAlarmsInput is undocumented.
type DescribeAlarmsInput struct {
	ActionPrefix    aws.StringValue  `query:"ActionPrefix" xml:"ActionPrefix"`
	AlarmNamePrefix aws.StringValue  `query:"AlarmNamePrefix" xml:"AlarmNamePrefix"`
	AlarmNames      []string         `query:"AlarmNames.member" xml:"AlarmNames>member"`
	MaxRecords      aws.IntegerValue `query:"MaxRecords" xml:"MaxRecords"`
	NextToken       aws.StringValue  `query:"NextToken" xml:"NextToken"`
	StateValue      aws.StringValue  `query:"StateValue" xml:"StateValue"`
}

// DescribeAlarmsOutput is undocumented.
type DescribeAlarmsOutput struct {
	MetricAlarms []MetricAlarm   `query:"MetricAlarms.member" xml:"DescribeAlarmsResult>MetricAlarms>member"`
	NextToken    aws.StringValue `query:"NextToken" xml:"DescribeAlarmsResult>NextToken"`
}

// Dimension is undocumented.
type Dimension struct {
	Name  aws.StringValue `query:"Name" xml:"Name"`
	Value aws.StringValue `query:"Value" xml:"Value"`
}

// DimensionFilter is undocumented.
type DimensionFilter struct {
	Name  aws.StringValue `query:"Name" xml:"Name"`
	Value aws.StringValue `query:"Value" xml:"Value"`
}

// DisableAlarmActionsInput is undocumented.
type DisableAlarmActionsInput struct {
	AlarmNames []string `query:"AlarmNames.member" xml:"AlarmNames>member"`
}

// EnableAlarmActionsInput is undocumented.
type EnableAlarmActionsInput struct {
	AlarmNames []string `query:"AlarmNames.member" xml:"AlarmNames>member"`
}

// GetMetricStatisticsInput is undocumented.
type GetMetricStatisticsInput struct {
	Dimensions []Dimension      `query:"Dimensions.member" xml:"Dimensions>member"`
	EndTime    time.Time        `query:"EndTime" xml:"EndTime"`
	MetricName aws.StringValue  `query:"MetricName" xml:"MetricName"`
	Namespace  aws.StringValue  `query:"Namespace" xml:"Namespace"`
	Period     aws.IntegerValue `query:"Period" xml:"Period"`
	StartTime  time.Time        `query:"StartTime" xml:"StartTime"`
	Statistics []string         `query:"Statistics.member" xml:"Statistics>member"`
	Unit       aws.StringValue  `query:"Unit" xml:"Unit"`
}

// GetMetricStatisticsOutput is undocumented.
type GetMetricStatisticsOutput struct {
	Datapoints []Datapoint     `query:"Datapoints.member" xml:"GetMetricStatisticsResult>Datapoints>member"`
	Label      aws.StringValue `query:"Label" xml:"GetMetricStatisticsResult>Label"`
}

// Possible values for CloudWatch.
const (
	HistoryItemTypeAction              = "Action"
	HistoryItemTypeConfigurationUpdate = "ConfigurationUpdate"
	HistoryItemTypeStateUpdate         = "StateUpdate"
)

// ListMetricsInput is undocumented.
type ListMetricsInput struct {
	Dimensions []DimensionFilter `query:"Dimensions.member" xml:"Dimensions>member"`
	MetricName aws.StringValue   `query:"MetricName" xml:"MetricName"`
	Namespace  aws.StringValue   `query:"Namespace" xml:"Namespace"`
	NextToken  aws.StringValue   `query:"NextToken" xml:"NextToken"`
}

// ListMetricsOutput is undocumented.
type ListMetricsOutput struct {
	Metrics   []Metric        `query:"Metrics.member" xml:"ListMetricsResult>Metrics>member"`
	NextToken aws.StringValue `query:"NextToken" xml:"ListMetricsResult>NextToken"`
}

// Metric is undocumented.
type Metric struct {
	Dimensions []Dimension     `query:"Dimensions.member" xml:"Dimensions>member"`
	MetricName aws.StringValue `query:"MetricName" xml:"MetricName"`
	Namespace  aws.StringValue `query:"Namespace" xml:"Namespace"`
}

// MetricAlarm is undocumented.
type MetricAlarm struct {
	ActionsEnabled                     aws.BooleanValue `query:"ActionsEnabled" xml:"ActionsEnabled"`
	AlarmActions                       []string         `query:"AlarmActions.member" xml:"AlarmActions>member"`
	AlarmARN                           aws.StringValue  `query:"AlarmArn" xml:"AlarmArn"`
	AlarmConfigurationUpdatedTimestamp time.Time        `query:"AlarmConfigurationUpdatedTimestamp" xml:"AlarmConfigurationUpdatedTimestamp"`
	AlarmDescription                   aws.StringValue  `query:"AlarmDescription" xml:"AlarmDescription"`
	AlarmName                          aws.StringValue  `query:"AlarmName" xml:"AlarmName"`
	ComparisonOperator                 aws.StringValue  `query:"ComparisonOperator" xml:"ComparisonOperator"`
	Dimensions                         []Dimension      `query:"Dimensions.member" xml:"Dimensions>member"`
	EvaluationPeriods                  aws.IntegerValue `query:"EvaluationPeriods" xml:"EvaluationPeriods"`
	InsufficientDataActions            []string         `query:"InsufficientDataActions.member" xml:"InsufficientDataActions>member"`
	MetricName                         aws.StringValue  `query:"MetricName" xml:"MetricName"`
	Namespace                          aws.StringValue  `query:"Namespace" xml:"Namespace"`
	OKActions                          []string         `query:"OKActions.member" xml:"OKActions>member"`
	Period                             aws.IntegerValue `query:"Period" xml:"Period"`
	StateReason                        aws.StringValue  `query:"StateReason" xml:"StateReason"`
	StateReasonData                    aws.StringValue  `query:"StateReasonData" xml:"StateReasonData"`
	StateUpdatedTimestamp              time.Time        `query:"StateUpdatedTimestamp" xml:"StateUpdatedTimestamp"`
	StateValue                         aws.StringValue  `query:"StateValue" xml:"StateValue"`
	Statistic                          aws.StringValue  `query:"Statistic" xml:"Statistic"`
	Threshold                          aws.DoubleValue  `query:"Threshold" xml:"Threshold"`
	Unit                               aws.StringValue  `query:"Unit" xml:"Unit"`
}

// MetricDatum is undocumented.
type MetricDatum struct {
	Dimensions      []Dimension     `query:"Dimensions.member" xml:"Dimensions>member"`
	MetricName      aws.StringValue `query:"MetricName" xml:"MetricName"`
	StatisticValues *StatisticSet   `query:"StatisticValues" xml:"StatisticValues"`
	Timestamp       time.Time       `query:"Timestamp" xml:"Timestamp"`
	Unit            aws.StringValue `query:"Unit" xml:"Unit"`
	Value           aws.DoubleValue `query:"Value" xml:"Value"`
}

// PutMetricAlarmInput is undocumented.
type PutMetricAlarmInput struct {
	ActionsEnabled          aws.BooleanValue `query:"ActionsEnabled" xml:"ActionsEnabled"`
	AlarmActions            []string         `query:"AlarmActions.member" xml:"AlarmActions>member"`
	AlarmDescription        aws.StringValue  `query:"AlarmDescription" xml:"AlarmDescription"`
	AlarmName               aws.StringValue  `query:"AlarmName" xml:"AlarmName"`
	ComparisonOperator      aws.StringValue  `query:"ComparisonOperator" xml:"ComparisonOperator"`
	Dimensions              []Dimension      `query:"Dimensions.member" xml:"Dimensions>member"`
	EvaluationPeriods       aws.IntegerValue `query:"EvaluationPeriods" xml:"EvaluationPeriods"`
	InsufficientDataActions []string         `query:"InsufficientDataActions.member" xml:"InsufficientDataActions>member"`
	MetricName              aws.StringValue  `query:"MetricName" xml:"MetricName"`
	Namespace               aws.StringValue  `query:"Namespace" xml:"Namespace"`
	OKActions               []string         `query:"OKActions.member" xml:"OKActions>member"`
	Period                  aws.IntegerValue `query:"Period" xml:"Period"`
	Statistic               aws.StringValue  `query:"Statistic" xml:"Statistic"`
	Threshold               aws.DoubleValue  `query:"Threshold" xml:"Threshold"`
	Unit                    aws.StringValue  `query:"Unit" xml:"Unit"`
}

// PutMetricDataInput is undocumented.
type PutMetricDataInput struct {
	MetricData []MetricDatum   `query:"MetricData.member" xml:"MetricData>member"`
	Namespace  aws.StringValue `query:"Namespace" xml:"Namespace"`
}

// SetAlarmStateInput is undocumented.
type SetAlarmStateInput struct {
	AlarmName       aws.StringValue `query:"AlarmName" xml:"AlarmName"`
	StateReason     aws.StringValue `query:"StateReason" xml:"StateReason"`
	StateReasonData aws.StringValue `query:"StateReasonData" xml:"StateReasonData"`
	StateValue      aws.StringValue `query:"StateValue" xml:"StateValue"`
}

// Possible values for CloudWatch.
const (
	StandardUnitBits            = "Bits"
	StandardUnitBitsSecond      = "Bits/Second"
	StandardUnitBytes           = "Bytes"
	StandardUnitBytesSecond     = "Bytes/Second"
	StandardUnitCount           = "Count"
	StandardUnitCountSecond     = "Count/Second"
	StandardUnitGigabits        = "Gigabits"
	StandardUnitGigabitsSecond  = "Gigabits/Second"
	StandardUnitGigabytes       = "Gigabytes"
	StandardUnitGigabytesSecond = "Gigabytes/Second"
	StandardUnitKilobits        = "Kilobits"
	StandardUnitKilobitsSecond  = "Kilobits/Second"
	StandardUnitKilobytes       = "Kilobytes"
	StandardUnitKilobytesSecond = "Kilobytes/Second"
	StandardUnitMegabits        = "Megabits"
	StandardUnitMegabitsSecond  = "Megabits/Second"
	StandardUnitMegabytes       = "Megabytes"
	StandardUnitMegabytesSecond = "Megabytes/Second"
	StandardUnitMicroseconds    = "Microseconds"
	StandardUnitMilliseconds    = "Milliseconds"
	StandardUnitNone            = "None"
	StandardUnitPercent         = "Percent"
	StandardUnitSeconds         = "Seconds"
	StandardUnitTerabits        = "Terabits"
	StandardUnitTerabitsSecond  = "Terabits/Second"
	StandardUnitTerabytes       = "Terabytes"
	StandardUnitTerabytesSecond = "Terabytes/Second"
)

// Possible values for CloudWatch.
const (
	StateValueAlarm            = "ALARM"
	StateValueInsufficientData = "INSUFFICIENT_DATA"
	StateValueOK               = "OK"
)

// Possible values for CloudWatch.
const (
	StatisticAverage     = "Average"
	StatisticMaximum     = "Maximum"
	StatisticMinimum     = "Minimum"
	StatisticSampleCount = "SampleCount"
	StatisticSum         = "Sum"
)

// StatisticSet is undocumented.
type StatisticSet struct {
	Maximum     aws.DoubleValue `query:"Maximum" xml:"Maximum"`
	Minimum     aws.DoubleValue `query:"Minimum" xml:"Minimum"`
	SampleCount aws.DoubleValue `query:"SampleCount" xml:"SampleCount"`
	Sum         aws.DoubleValue `query:"Sum" xml:"Sum"`
}

// DescribeAlarmHistoryResult is a wrapper for DescribeAlarmHistoryOutput.
type DescribeAlarmHistoryResult struct {
	AlarmHistoryItems []AlarmHistoryItem `query:"AlarmHistoryItems.member" xml:"DescribeAlarmHistoryResult>AlarmHistoryItems>member"`
	NextToken         aws.StringValue    `query:"NextToken" xml:"DescribeAlarmHistoryResult>NextToken"`
}

// DescribeAlarmsForMetricResult is a wrapper for DescribeAlarmsForMetricOutput.
type DescribeAlarmsForMetricResult struct {
	MetricAlarms []MetricAlarm `query:"MetricAlarms.member" xml:"DescribeAlarmsForMetricResult>MetricAlarms>member"`
}

// DescribeAlarmsResult is a wrapper for DescribeAlarmsOutput.
type DescribeAlarmsResult struct {
	MetricAlarms []MetricAlarm   `query:"MetricAlarms.member" xml:"DescribeAlarmsResult>MetricAlarms>member"`
	NextToken    aws.StringValue `query:"NextToken" xml:"DescribeAlarmsResult>NextToken"`
}

// GetMetricStatisticsResult is a wrapper for GetMetricStatisticsOutput.
type GetMetricStatisticsResult struct {
	Datapoints []Datapoint     `query:"Datapoints.member" xml:"GetMetricStatisticsResult>Datapoints>member"`
	Label      aws.StringValue `query:"Label" xml:"GetMetricStatisticsResult>Label"`
}

// ListMetricsResult is a wrapper for ListMetricsOutput.
type ListMetricsResult struct {
	Metrics   []Metric        `query:"Metrics.member" xml:"ListMetricsResult>Metrics>member"`
	NextToken aws.StringValue `query:"NextToken" xml:"ListMetricsResult>NextToken"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ xml.Decoder
var _ = io.EOF
