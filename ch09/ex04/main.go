package main

func main() {}

type Pipe struct {
	in          chan int
	out         chan int
	passedCount int
}

func (p *Pipe) a(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			p.passedCount++
			out <- n
		}
		close(out)
	}()
	return out
}

func (p *Pipe) build(n int) {
	p.in = make(chan int)
	var in chan int

	in = p.in
	for i := 0; i < n; i++ {
		in = p.a(in)
	}
	p.out = in
}

func (p *Pipe) passThrough(num int) int {
	p.in <- num
	return <-p.out
}
