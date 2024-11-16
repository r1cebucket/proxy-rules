package generator

import "proxy-rules/internal/generator/mode"

const (
	MODE_SING_BOX     = "sing-box"
	MODE_CLASH        = "clash"
	MODE_QUAN_X       = "quan-x"
	MODE_MATSURI      = "matsuri"
	MODE_SUREG        = "surge"
	MODE_SHADOWROCKET = "shadowrocket"
)

var MODES_ALLOWED = map[string]bool{
	MODE_SING_BOX:     true,
	MODE_CLASH:        true,
	MODE_QUAN_X:       true,
	MODE_MATSURI:      true,
	MODE_SUREG:        true,
	MODE_SHADOWROCKET: true,
}

var ModeGenerator = map[string]Generator{
	MODE_CLASH:        mode.Clash{},
	MODE_SHADOWROCKET: mode.ShadowRocket{},
	MODE_SING_BOX:     mode.SingBox{},
}
