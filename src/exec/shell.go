package main

import (
	"os/exec"
	"time"
)

func main() {

	go func() {
		// fmt.Println("rtmp://uplive.v0.upaiyun.com/live/" + strconv.Itoa(i))
		cmd := exec.Command("ffmpeg", "-re", "-i", "time.300kbps.flv", "-c",
			"copy", "-f", "flv", "rtmp://uplive.v0.upaiyun.com/live/110")
		_, err := cmd.Output()
		if err != nil {
			panic(err.Error())
		}
	}()

	go func() {
		// fmt.Println("rtmp://uplive.v0.upaiyun.com/live/" + strconv.Itoa(i))
		cmd := exec.Command("ffmpeg", "-re", "-i", "time.300kbps.flv", "-c",
			"copy", "-f", "flv", "rtmp://uplive.v0.upaiyun.com/live/111")
		_, err := cmd.Output()
		if err != nil {
			panic(err.Error())
		}
	}()

	// go func() {
	// 	// fmt.Println("rtmp://uplive.v0.upaiyun.com/live/" + strconv.Itoa(i))
	// 	cmd := exec.Command("ffmpeg", "-re", "-i", "time.300kbps.flv", "-c",
	// 		"copy", "-f", "flv", "rtmp://uplive.v0.upaiyun.com/live/13")
	// 	_, err := cmd.Output()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }()
	//
	// go func() {
	// 	// fmt.Println("rtmp://uplive.v0.upaiyun.com/live/" + strconv.Itoa(i))
	// 	cmd := exec.Command("ffmpeg", "-re", "-i", "time.300kbps.flv", "-c",
	// 		"copy", "-f", "flv", "rtmp://uplive.v0.upaiyun.com/live/14")
	// 	_, err := cmd.Output()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }()
	//
	// go func() {
	// 	// fmt.Println("rtmp://uplive.v0.upaiyun.com/live/" + strconv.Itoa(i))
	// 	cmd := exec.Command("ffmpeg", "-re", "-i", "time.300kbps.flv", "-c",
	// 		"copy", "-f", "flv", "rtmp://uplive.v0.upaiyun.com/live/15")
	// 	_, err := cmd.Output()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }()

	time.Sleep(30 * time.Second)
}
