package study

func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i <= len(arr); i++ {
		var flag bool
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func QuicklySort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuicklySort(arr, start, j)
		}
		if end > i {
			QuicklySort(arr, i, end)
		}
	}
}

/*
4
100 100 120 130
40 30 60 50
*/
//func main() {
//	var (
//		a int
//	)
//
//	for {
//		n, _ := fmt.Scan(&a)
//		if n == 0 {
//			break
//		} else {
//			b := make([]int,a)
//			nn, _ := fmt.Scan(&b)
//			fmt.Println("-----", nn,b)
//			if nn == 0 {
//				break
//			} else {
//				var (
//					ll []int
//					cc []int
//				)
//				ll = append(ll, b)
//				cc = append(cc, c)
//				fmt.Printf("%d\n", F(ll, cc))
//			}
//
//		}
//	}
//}
//
//type P []pe
//
//type pe struct {
//	L int // 身高
//	C int // 体重
//	I int // 序号
//}
//
//func (p P) Less(i, j int) bool {
//	if p[i].L < p[j].L {
//		return true
//	} else if p[i].L == p[j].L {
//		if p[i].C < p[j].C {
//			return true
//		} else if p[i].C == p[j].C {
//			if p[i].I < p[j].I {
//				return true
//			} else {
//				return false
//			}
//		} else {
//			return false
//		}
//	} else {
//		return false
//	}
//}
//
//func (p P) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//
//func (p P) Len() int {
//	return len(p)
//}
//
//func SIn(p P) sort.Interface {
//	return &p
//}
//
//func F(l, c []int) []int {
//	var (
//		peo []pe
//		res []int
//	)
//	for i := 0; i < len(l); i++ {
//		peo = append(peo, pe{
//			L: l[i],
//			C: c[i],
//			I: i + 1,
//		})
//	}
//	in := SIn(peo)
//	sort.Sort(in)
//	for _, val := range peo {
//		res = append(res, val.I)
//	}
//	return res
//}