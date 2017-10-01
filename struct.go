package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool // 小文字は閉じたスコープ
}

// コンストラクタを定義出来ないので自前でメソッドを用意
func newTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

// stringに変換するメソッド
func (task Task) String() string {
	str := fmt.Sprintf("%d:%s", task.ID, task.Detail)
	return str
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "goの練習",
		done:   true,
	}
	fmt.Println(task)
}
