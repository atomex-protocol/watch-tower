// Code generated by "stringer --output ../../internal/atomex/stringer.go -type WebsocketType ../../internal/atomex/"; DO NOT EDIT.

package atomex

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WebsocketTypeMarketData-1]
	_ = x[WebsocketTypeExchange-2]
}

const _WebsocketType_name = "WebsocketTypeMarketDataWebsocketTypeExchange"

var _WebsocketType_index = [...]uint8{0, 23, 44}

func (i WebsocketType) String() string {
	i -= 1
	if i < 0 || i >= WebsocketType(len(_WebsocketType_index)-1) {
		return "WebsocketType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _WebsocketType_name[_WebsocketType_index[i]:_WebsocketType_index[i+1]]
}
