package api

type PositionService struct {
	client *Client
}

const (
	urlPathGetPosition   = "/private/get_position"
	urlPathGetPositions  = "/private/get_positions"
	urlPathClosePosition = "/private/close_position"
	urlPathGetMargins    = "/private/get_margins"
)
