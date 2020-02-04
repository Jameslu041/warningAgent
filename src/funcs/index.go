// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
	"warningAgent/src/model"
	"warningAgent/src/util"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func initDataHistory(interval int) {
	for {
		UpdateCpuStat()
		UpdateDiskStats()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func BuildLopper(interval int) {
	go initDataHistory(interval)

	Mappers = []FuncsAndInterval{
		{
			Fs: []func() []*model.MetricValue{
				CpuMetrics,
				NetMetrics,
				LoadAvgMetrics,
				MemMetrics,
				DiskIOMetrics,
				IOStatsMetrics,
				DeviceMetrics,
			},
			Interval: interval,
		},
	}
}

func StartCollect() {
	fmt.Println("server metric collect [OPEN]")
	for _, v := range Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*model.MetricValue) {

	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()
	for {
		<-t.C

		hostname := "hostname"

		mvs := []*model.MetricValue{}

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			mvs = append(mvs, items...)
		}

		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ {
			Warning(mvs[j])
			mvs[j].Step = sec
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
		}

		//g.SendToTransfer(mvs)
		fmt.Println(len(mvs), " -- ", mvs)
	}
}

func Warning(mv *model.MetricValue) {
	switch mv.Metric {
	case "cpu.busy":
		i := viper.GetFloat64("alarm_index.cpu.busy")
		if mv.Value.(float64) > i {
			util.PostErrorf("慈善捐赠平台:\n'%s': 当前数值 %.2f 大于阈值%.2f (百分率), 请注意!\n", "cpu使用率", mv.Value.(float64), i)

		}
	case "mem.memused.percent":
		i := viper.GetFloat64("alarm_index.mem.memused")
		if mv.Value.(float64) > i {
			util.PostErrorf("慈善捐赠平台:\n'%s': 当前数值 %.2f 大于阈值%.2f (百分率), 请注意!\n", "内存使用率", mv.Value.(float64), i)

		}
	case "disk.io.util":
		i := viper.GetFloat64("alarm_index.disk.io_util")
		if mv.Value.(float64) > i {
			util.PostErrorf("慈善捐赠平台:\n'%s': 当前数值 %.2f 大于阈值%.2f (百分率), 请注意!\n", "IO使用率", mv.Value.(float64), i)

		}
	case "df.statistics.used.percent":
		i := viper.GetFloat64("alarm_index.df.statistics_used")
		if mv.Value.(float64) > i {
			util.PostErrorf("慈善捐赠平台:\n'%s': 当前数值 %.2f 大于阈值%.2f (百分率), 请注意!\n", "数据硬盘使用率", mv.Value.(float64), i)

		}
	default:
		return
	}
}
