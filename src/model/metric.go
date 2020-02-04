package model

import (
	"fmt"
)

type MetricInfo struct {
	Id       int    `json:"id"`
	Endpoint string `json:"endpoint"`
	Metric   string `json:"metric"`
	Step     int64  `json:"step"`
	Type     string `json:"type"`
	Tags     string `json:"tags"`
	Name     string `json:"name"`
	Open     bool   `json:"open"`
}

type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
	Name      string      `json:"name"`
}

func (this *MetricValue) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Metric:%s, Type:%s, Tags:%s, Step:%d, Time:%d, Value:%v>",
		this.Endpoint,
		this.Metric,
		this.Type,
		this.Tags,
		this.Step,
		this.Timestamp,
		this.Value,
	)
}

// Same As `MetricValue`
type JsonMetaData struct {
	Metric      string      `json:"metric"`
	Endpoint    string      `json:"endpoint"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
	Value       interface{} `json:"value"`
	Name        string      `json:"name"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
}

func (t *JsonMetaData) String() string {
	return fmt.Sprintf("<JsonMetaData Endpoint:%s, Metric:%s, Tags:%s, DsType:%s, Step:%d, Value:%v, Timestamp:%d>",
		t.Endpoint, t.Metric, t.Tags, t.CounterType, t.Step, t.Value, t.Timestamp)
}

type MetaData struct {
	Metric      string            `json:"metric"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       interface{}       `json:"value"`
	CounterType string            `json:"counterType"`
	Tags        map[string]string `json:"tags"`
}

func (t *MetaData) String() string {
	return fmt.Sprintf("<MetaData Endpoint:%s, Metric:%s, Timestamp:%d, Step:%d, Value:%f, Tags:%v>",
		t.Endpoint, t.Metric, t.Timestamp, t.Step, t.Value, t.Tags)
}

type BuildMetric struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Host   string `json:"host"`
	Pid    int64  `json:"pid"`
	Type   int    `json:"type"x`
	Params string `json:"params"`
}
