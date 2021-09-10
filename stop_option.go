package xcho

import (
	`time`
)

type (
	stopOption interface {
		applyStop(options *stopOptions)
	}

	stopOptions struct {
		*options

		// 退出超时时间
		timeout time.Duration
	}
)

func defaultStopOptions() *stopOptions {
	return &stopOptions{
		options: defaultOptions,

		timeout: 30 * time.Second,
	}
}
