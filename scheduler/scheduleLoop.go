package scheduler

import (
	"time"
)

func sch() {
	l2Lock := false
	l3Index := 0

	var l2TC <-chan time.Time

	// 任务调度大循环
	for {
		// 1级自检任务
		// 1级任务采取检查表的形式，每次逐一检查所有项，满足条件则执行相应的处置程序
		// 1级任务所需的各种状态信息，依赖于采集服务写入redis中
		L1_SelfInspection()

		// 2级手动任务
		// 2级任务采用channel的方式接收任务，同时加定时器
		// 一旦接收过2级任务，短时间大循环体锁定2级任务状态，不再执行任何3级任务。
		// 超时未再接收到2级任务，解除锁定状态；2级任务中有一个结束2级任务，收到此任务后立即解除2级任务锁定

		//missionChannel chan
		//lockTimeout int
		if l2Lock == false {
			l2TC = time.After(1 * time.Second)
		}

		select {
		case <-l2TC:
			l2Lock = false
		case l2m := <-missionChannel:
			l2Lock = true
			l2TC = time.After(lockTimeout * time.Second)
			switch l2m.Type {

			}
		}

		// 3级周期任务
		// 周期逐一执行任务书中的所有任务，大循环体每循环一次，执行一条3级任务
		// 3级任务可包含子任务，例如：升降杆的升降、拍摄、测温等等
		// 每执行一条3级任务，同时执行其中的所有子任务（即执行子任务过程中不受打断）
		// 3级任务的任务书更新是通过一个2级任务来完成
		l3Index = L3_PeriodicTask(l3Index)

	}
}
