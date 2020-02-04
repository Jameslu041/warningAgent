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

func LoadAvgMetrics() []*model.MetricValue {
	load, err := nux.LoadAvg()
	if err != nil {
		log.Println(err)
		return nil
	}

	return []*model.MetricValue{
		GaugeValueN("load.1min", "1min机器负载", load.Avg1min),
		GaugeValueN("load.5min", "5min机器负载", load.Avg5min),
		GaugeValueN("load.15min", "15min机器负载", load.Avg15min),
	}

}
