type WordDictionary struct {
    IsEnd bool
	Nodes [26]*WordDictionary
}

func Constructor() WordDictionary {
    return WordDictionary{
		Nodes: [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    next := this
	for _,c := range word {
		char := c - 'a'
		if next.Nodes[char] == nil {
			next.Nodes[char] = &WordDictionary{
				Nodes: [26]*WordDictionary{},
			}
		}
		next = next.Nodes[char]
	}
	next.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    // Gọi hàm đệ quy bắt đầu từ vị trí 0 và nút gốc (this)
    return dfs(0, this, word)
}

func dfs(index int, node *WordDictionary, word string) bool {
    curr := node
    
    for i := index; i < len(word); i++ {
        c := word[i]
        
        if c == '.' {
            // Nếu là dấu '.', ta phải thử TẤT CẢ 26 nhánh con
            for _, child := range curr.Nodes {
                if child != nil {
                    // Nếu nhánh này dẫn đến kết quả đúng, trả về true luôn
                    if dfs(i+1, child, word) {
                        return true
                    }
                }
            }
            // Thử hết 26 nhánh mà không cái nào khớp thì từ này không tồn tại
            return false
        } else {
            // Nếu là ký tự bình thường (a-z)
            char := c - 'a'
            if curr.Nodes[char] == nil {
                return false
            }
            curr = curr.Nodes[char]
        }
    }
    
    // Sau khi đi hết từ, phải kiểm tra xem đây có phải là điểm kết thúc của 1 từ không
    return curr.IsEnd
}