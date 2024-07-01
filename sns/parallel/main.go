package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	post := fetchPost()

	// channelの初期化
	// 2個のバッファを持ったchannelを作成
	resChan := make(chan any, 2)

	// カウントが0になるまでメインゴルーチンの終了を待たせる
	var wg sync.WaitGroup
	wg.Add(2)

	// 以下の２つは同時進行する。終了時にwgの数値を1減らす
	// 参照渡しをしないとコピーが作成されるため、上記のwgから数値が減ることはない。
	go fetchPostLikes(post, resChan, &wg)
	go fetchPostComments(post, resChan, &wg)

	// wgの数値が0になるまでメインゴルーチンの終了を待つ
	wg.Wait()

	// resChan channelへの送信を終了しchannelを閉じる
	close(resChan)

	// channelが閉じられるまでループする
	for res := range resChan {
		fmt.Println("res:", res)
	}

	fmt.Println("took:", time.Since(start))
}

func fetchPost() string {
	time.Sleep(time.Millisecond * 50)
	return "What programming languages do you prefer?"
}

func fetchPostLikes(_ string, resChan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	resChan <- 10
	// wgの数値を1減らす
	wg.Done()
}

func fetchPostComments(_ string, resChan chan any, wg  *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	resChan <- []string{"Golang", "Java", "Rust"}
	// wgの数値を1減らす
	wg.Done()
}