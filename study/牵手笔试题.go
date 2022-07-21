package study

/*
变形词是对一个单词的字母进行重新排列得到的新词。例如，你可以重新排列 terraced，得到 retraced。你也可以再次重排，得到cratered。除此外，你还能通过重排得到dactrere redatrec，尽管这两个单词并不是有实际意义的英语单词。

输入
输入为每行一个单词。每个单词由大、小写字母（a-z）组成，也可能包含各种字符。输入以文件末尾结束。

输出
对于每个输入单词，输出这个单词可以得到的变形词个数。就此问题而言，大、小写字母是有区别的。

样本输入 1

at
ordeals
abcdefghijklmnopqrstuvwxyz
abcdefghijklmabcdefghijklm
abcdABCDabcd

样本输出 1
2
5040
403291461126605635584000000
49229914688306352000000
29937600
、兆、京、垓、秭、穰、沟、涧、正、载、极、恒河沙、阿僧祇、那由他、不可思议、无量大数
*/
func CalcPossibilities(str string) int {
	var (
		distinctMap = make(map[string]struct{})
		l           = len(str)
	)
	distinctMap[str] = struct{}{}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			for s := range distinctMap {
				t := []byte(s)
				t[i], t[j] = t[j], t[i]
				distinctMap[string(t)] = struct{}{}
			}
		}
	}
	return len(distinctMap)
}
