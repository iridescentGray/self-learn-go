func doSomething() (err error) {
	defer func() {
		switch e := recover().(type) {
		case nil:
		case error:
			err = e
		default:
			panic(e) // 重新抛出此恐慌
		}
	}()

	doStep1()
	doStep2()
	doStep3()
	doStep4()
	doStep5()

	return
}

// 在现实中，各个doStepN函数的原型可能不同。
// 这里，每个doStepN函数的行为如下：
// * 如果已经成功，则调用panic(nil)来制造一个恐慌
//   以示不需继续；
// * 如果本步失败，则调用panic(err)来制造一个恐慌
//   以示不需继续；
// * 不制造任何恐慌表示继续下一步。
func doStepN() {
	...
	if err != nil {
		panic(err)
	}
	...
	if done {
		panic(nil)
	}
}