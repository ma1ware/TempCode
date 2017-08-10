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
