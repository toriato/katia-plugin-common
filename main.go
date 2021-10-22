package main

import (
	"github.com/toriato/katia"
)

var Plugin = katia.Plugin{
	Name:        "katia-plugin-common",
	Description: "메모리, Redis 등을 통한 키-값 캐시를 제공합니다",
	Author:      "Sangha Lee <totoriato@gmail.com>",
	Version:     [3]int{0, 1, 0},

	OnEnable: func(bot *katia.Bot, plugin *katia.Plugin) error {
		return nil
	},
}
