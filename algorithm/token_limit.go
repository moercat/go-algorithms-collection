package algorithm

import "time"

type TokenB struct {
	tokenCap      int
	tokeNum       int
	tokenMs       int64
	tokenLastTime int64
}

func New(tn int, tm int64) *TokenB {
	return &TokenB{
		tokenCap:      tn,
		tokeNum:       tn,
		tokenMs:       tm,
		tokenLastTime: time.Now().Unix(),
	}
}

func (t *TokenB) GetToken() bool {

	if time.Now().Unix() > t.tokenLastTime+t.tokenMs {
		t.tokenCap = t.tokeNum - 1
		t.tokenLastTime = time.Now().Unix()
		return true
	}
	if t.tokenCap <= 0 {
		return false
	}
	t.tokenCap -= 1
	t.tokenLastTime = time.Now().Unix()
	return true
}
