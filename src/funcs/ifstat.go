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

func NetMetrics() []*model.MetricValue {
	prefixs := []string{}
	return CoreNetMetrics(prefixs)
}

func CoreNetMetrics(ifacePrefix []string) []*model.MetricValue {

	netIfs, err := nux.NetIfs(ifacePrefix)
	if err != nil {
		log.Println(err)
		return []*model.MetricValue{}
	}
	ret := []*model.MetricValue{}
	out := int64(0)
	in := int64(0)
	all := int64(0)
	speed := int64(0)
	for _, netIf := range netIfs {
		// iface := "iface=" + netIf.Iface
		in += netIf.InBytes
		out += netIf.OutBytes
		all += netIf.TotalBytes
		speed += netIf.SpeedBits

		// ret[idx*23+0] = CounterValue("net.if.in.bytes", netIf.InBytes, iface)
		// ret[idx*23+1] = CounterValue("net.if.in.packets", netIf.InPackages, iface)
		// ret[idx*23+2] = CounterValue("net.if.in.errors", netIf.InErrors, iface)
		// ret[idx*23+3] = CounterValue("net.if.in.dropped", netIf.InDropped, iface)
		// ret[idx*23+4] = CounterValue("net.if.in.fifo.errs", netIf.InFifoErrs, iface)
		// ret[idx*23+5] = CounterValue("net.if.in.frame.errs", netIf.InFrameErrs, iface)
		// ret[idx*23+6] = CounterValue("net.if.in.compressed", netIf.InCompressed, iface)
		// ret[idx*23+7] = CounterValue("net.if.in.multicast", netIf.InMulticast, iface)
		// ret[idx*23+8] = CounterValue("net.if.out.bytes", netIf.OutBytes, iface)
		// ret[idx*23+9] = CounterValue("net.if.out.packets", netIf.OutPackages, iface)
		// ret[idx*23+10] = CounterValue("net.if.out.errors", netIf.OutErrors, iface)
		// ret[idx*23+11] = CounterValue("net.if.out.dropped", netIf.OutDropped, iface)
		// ret[idx*23+12] = CounterValue("net.if.out.fifo.errs", netIf.OutFifoErrs, iface)
		// ret[idx*23+13] = CounterValue("net.if.out.collisions", netIf.OutCollisions, iface)
		// ret[idx*23+14] = CounterValue("net.if.out.carrier.errs", netIf.OutCarrierErrs, iface)
		// ret[idx*23+15] = CounterValue("net.if.out.compressed", netIf.OutCompressed, iface)
		// ret[idx*23+16] = CounterValue("net.if.total.bytes", netIf.TotalBytes, iface)
		// ret[idx*23+17] = CounterValue("net.if.total.packets", netIf.TotalPackages, iface)
		// ret[idx*23+18] = CounterValue("net.if.total.errors", netIf.TotalErrors, iface)
		// ret[idx*23+19] = CounterValue("net.if.total.dropped", netIf.TotalDropped, iface)
		// ret[idx*23+20] = GaugeValue("net.if.speed.bits", netIf.SpeedBits, iface)
		// ret[idx*23+21] = CounterValue("net.if.in.percent", netIf.InPercent, iface)
		// ret[idx*23+22] = CounterValue("net.if.out.percent", netIf.OutPercent, iface)
	}

	ret = append(ret, CounterValueN("net.if.in.total.bytes", "网络总流入", in, ""))
	ret = append(ret, CounterValueN("net.if.out.total.bytes", "网络总流出", out, ""))
	ret = append(ret, CounterValueN("net.if.total.bytes.all", "网络总流量", all, ""))
	ret = append(ret, CounterValueN("net.if.total.speed", "网络带宽", speed, ""))

	return ret
}
