nextRunTime := time.Now()
if !s.trigger.StartTime.IsZero() {
	nextRunTime = s.trigger.StartTime
}
for {
    duration := nextRunTime.Sub(time.Now())
    duration = time.Duration(float32(duration) * s.getTriggerScale())
    <- time.After(duration)
    nextRunTime = time.Now()
    if s.trigger.IsFinished {
        return
    }
    s.strategy.Tick(s.triggerTicker, nextRunTime)
    nextRunTime = s.trigger.NextTime(nextRunTime)
    if s.trigger.ListenMode {
        break
    }
}



//nextRunTime := time.Now()
		//若为监听模式则只触发一次
		/*
			if s.trigger.ListenMode {
				if !s.trigger.StartTime.IsZero() && time.Now().Before(s.trigger.StartTime) {
					nextRunTime = s.trigger.StartTime
				}
				if s.trigger.IsFinished() {
					return
				}
				s.strategy.Tick(s.triggerTicker, nextRunTime)
				return
			}*/
		/*
			if !s.trigger.StartTime.IsZero() && time.Now().After(s.trigger.StartTime) {
				for {
					fmt.Println("start-------------------->", s.trigger.StartTime)
					nextRunTime = s.trigger.NextTime(nextRunTime)
					if s.trigger.CronExpr == "" {
						return
					}
					fmt.Println("trigger------------------>", s.trigger.CronExpr)
					fmt.Println("next-------------------->", nextRunTime)
					duration := nextRunTime.Sub(time.Now())
					duration = time.Duration(float32(duration) * s.getTriggerScale())
					<-time.After(duration)
					nextRunTime = time.Now()
					if s.trigger.IsFinished() {
						return
					}
					s.strategy.Tick(s.triggerTicker, nextRunTime)
					time.Sleep(time.Microsecond * 1)
				}

			}
			if !s.trigger.StartTime.IsZero() && time.Now().Before(s.trigger.StartTime) {

				nextRunTime = s.trigger.StartTime
				for {
					duration := nextRunTime.Sub(time.Now())

					duration = time.Duration(float32(duration) * s.getTriggerScale())

					<-time.After(duration)
					nextRunTime = time.Now()
					nextRunTime = s.trigger.NextTime(nextRunTime)
					if s.trigger.IsFinished() {
						return
					}
					s.strategy.Tick(s.triggerTicker, nextRunTime)
					time.Sleep(time.Microsecond * 1)
				}
			}*/
