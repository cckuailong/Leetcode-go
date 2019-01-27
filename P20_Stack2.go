package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	l := list.New()
	l.PushBack('0')
	for _, c:= range(s){
		if c=='(' || c=='[' || c=='{'{
			l.PushBack(c)
		} else if c== ')'{
			back := l.Back()
			if back.Value == '('{
				l.Remove(back)
			}else{
				return false
			}
		}else if c== ']'{
			back := l.Back()
			if back.Value == '['{
				l.Remove(back)
			}else{
				return false
			}
		}else if c== '}'{
			back := l.Back()
			if back.Value == '{'{
				l.Remove(back)
			}else{
				return false
			}
		}else{
			return false
		}
	}

	if l.Front().Next() == nil{
		return true
	}else{
		return false
	}
}

func main(){
	s := "}]"
	fmt.Println(isValid(s))
}
