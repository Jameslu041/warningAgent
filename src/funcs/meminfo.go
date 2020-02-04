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
	"github.com/toolkits/nux"
	"log"
	"warningAgent/src/model"
)

func MemMetrics() []*model.MetricValue {
	m, err := nux.MemInfo()
	if err != nil {
		log.Println(err)
		return nil
	}

	memFree := m.MemFree + m.Buffers + m.Cached
	memUsed := m.MemTotal - memFree

	pmemFree := 0.0
	pmemUsed := 0.0
	if m.MemTotal != 0 {
		pmemFree = float64(memFree) * 100.0 / float64(m.MemTotal)
		pmemUsed = float64(memUsed) * 100.0 / float64(m.MemTotal)
	}

	pswapFree := 0.0
	pswapUsed := 0.0
	if m.SwapTotal != 0 {
		pswapFree = float64(m.SwapFree) * 100.0 / float64(m.SwapTotal)
		pswapUsed = float64(m.SwapUsed) * 100.0 / float64(m.SwapTotal)
	}

	return []*model.MetricValue{
		GaugeValueN("mem.memtotal", "内存总量", m.MemTotal),
		GaugeValueN("mem.memused", "内存已使用", memUsed),
		GaugeValueN("mem.memfree", "内存空闲", memFree),
		GaugeValueN("mem.swaptotal", "swap总大小", m.SwapTotal),
		GaugeValueN("mem.swapused", "swap已使用", m.SwapUsed),
		GaugeValueN("mem.swapfree", "swap空闲", m.SwapFree),
		GaugeValueN("mem.memfree.percent", "内存空闲百分比", pmemFree),
		GaugeValueN("mem.memused.percent", "内存已使用百分比", pmemUsed),
		GaugeValueN("mem.swapfree.percent", "swap空闲百分比", pswapFree),
		GaugeValueN("mem.swapused.percent", "swap已使用百分比", pswapUsed),
	}

}
