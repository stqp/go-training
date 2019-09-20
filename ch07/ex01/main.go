package main

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordLineCounter struct{ w, l int }

func (c *WordLineCounter) Split(s []byte, seps string) (res []string) {
	var tmp string
	var on bool
	for _, v := range s {
		var isSep bool
		for _, sep := range []byte(seps) {
			if v == sep {
				isSep = true
			}
		}
		if !isSep {
			on = true
			tmp += string(v)
		} else {
			if on {
				on = false
				res = append(res, tmp)
				tmp = ""
			}
		}
	}
	if on {
		on = false
		res = append(res, tmp)
		tmp = ""
	}
	return
}

func (c *WordLineCounter) Write(p []byte) (int, error) {
	c.l++
	for _, v := range p {
		if v == '\n' {
			c.l++
		}
	}
	c.w = len(c.Split(p, " \n"))
	return len(p), nil
}

func main() {}
