package main

import (
	"testing"
)

func Test_CreateAttTask(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"https://dereking.myds.me:1443/ielts/2022%20IELTS%20Actual%20Test/test4/Listening%20Test%204%20Section%203.mp3"},
	}
	//24.bc11eabccbcb1574f875dae1b914ed2b.2592000.1683040789.282335-31887214
	//2023/04/02 22:49:02 {"log_id":16804469424149814,"task_status":"Created","task_id":"642995deea05a000018dea25"}
	c := NewBaiduClient("31887214", "BE2fHZo8WUISCwvenLd4pMiB----", "tfLaDQBrHqGkdzOGTLW0uGDWTZOAZ4kf")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.CreateAttTask(tt.name)
		})
	}
}

func Test_GetTaskResult(t *testing.T) {
	//2023/04/02 22:49:02 {"log_id":16804469424149814,"task_status":"Created","task_id":"642995deea05a000018dea25"}
	c := NewBaiduClient("31887214", "BE2fHZo8WUISCwvenLd4pMiB", "tfLaDQBrHqGkdzOGTLW0uGDWTZOAZ4kf")
	c.GetTaskResult([]string{"642995deea05a000018dea25"})
}
