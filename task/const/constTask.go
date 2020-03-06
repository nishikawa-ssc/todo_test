package constant

// DbName DB名
const DbName = "task"

const (
	// Text 内容
	Text = "text"
	// Stat 状態
	Stat = "stat"
)

// State 状態用の型
type State int

const (
	_ State = iota
	// YET 未着手
	YET
	// STARTED 着手
	STARTED
	// DONE 完了
	DONE
)

func (s State) String() string {
	switch s {
	case YET:
		return "未着手"
	case STARTED:
		return "着手"
	case DONE:
		return "完了"
	default:
		return "　"
	}
}
