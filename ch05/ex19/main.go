package main

func main() {}

type num struct{}

func a() (ret int) {
	defer func() {
		switch p := recover(); p {
		case num{}:
			ret = 10
		}
	}()
	panic(num{})
}
