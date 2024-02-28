package models

import "time"

type Swipe struct {
	ID        string    `json:"id"`
	SwiperID  string    `json:"swiper_id"`
	SwipedID  string    `json:"swiped_id"`
	SwipeType string    `json:"swipe_type"`
	SwipeDate time.Time `json:"swipe_date"`
}
