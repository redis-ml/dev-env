package main

import (
	"container/ring"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	. "kunlun_ai/storage_server/protobuf_idl"
	"kunlun_ai/testing/fake"
	"log"
	"net"
	"time"
)

type Pool struct {
	c Config
}

type Config struct {
	name string
}

func NewPool(c Config) *Pool {
	ret := &Pool{
		c: c,
	}
	return ret
}

func (p *Pool) GetConf() Config {
	return p.c
}

type Tag TagIdIndexData

func (t Tag) Print() {
	fmt.Printf("print %v, addr: %p", t, &t)
}
func (t *Tag) Marshal() (data []byte, err error) {
	fmt.Printf("Unmarshal, pointer of t: %p\n", &t)
	tmp := TagIdIndexData(*t)
	fmt.Printf("Unmarshal, pointer of tmp: %p\n", &tmp)
	return proto.Marshal(&tmp)
}

func (t *Tag) Unmarshal(data []byte) (err error) {
	fmt.Printf("Unmarshal, pointer of t: %p\n", t)
	tmp := TagIdIndexData(*t)
	fmt.Printf("Unmarshal, pointer of tmp: %p\n", &tmp)
	err = proto.Unmarshal(data, &tmp)
	if err == nil {
		*t = Tag(tmp)
	}
	return
}

type T struct {
	a string
}

type P T

func (p *P) ttt() {
	fmt.Println("in:", p.a)
}
func testRing() {
	r := ring.New(10)
	r1 := ring.New(1)
	r1.Value = "hthe"
	r = r.Link(r1)
	fmt.Println("ring:", r.Value, ", len:", r.Len())
	r.Value = "r1"
	fmt.Println("ring:", r.Value, ", len:", r.Len())
	r = r.Next()
	r.Value = 10
	fmt.Println("ring:", r.Value, ", len:", r.Len())
	num := 0
	r.Do(func(d interface{}) {
		if d == "r1" {
			fmt.Println("found 1")
		} else if d != nil {
			num++
		}
	})
	fmt.Println("total value:", num)
}
func otherTest() {
	cl := make([]chan string, 10)
	for i, ch := range cl {
		ch = make(chan string, 1)
		cl[i] = ch
		ch <- "hehe"
	}
	ch := make(chan string, 10)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("item:", i)
	}
	for i := 0; i < 20; i++ {
		select {
		case b := <-ch:
			fmt.Println(i, ". Got one from channel, ", b)
			break
		default:
			fmt.Println(i, ". Fall to default")
		}
	}
	/*
		c := Config{name: "hehe"}
		pc := &c

		c3 := Config{name: "nmo"}
		*pc = c3
		fmt.Println("c3:", c3, ", c:", c)
		p := NewPool(c)
		c.name = "haha"
		fmt.Print("now pool is :", p, "\n")
		c2 := p.GetConf()
		fmt.Print("now getConf.name:", c2.name, "\n")
		c2.name = "c2"
		fmt.Print("now pool is :", p, "\n")

		/*
			tag := &Tag{}
			fmt.Printf("pointer of tag: %p\n", tag)
			tag.IdList = []int64{1, 2, 3}
			data, err := tag.Marshal()
			fmt.Printf("data: %v, err: %v, addr of data: %p\n", data, err, &tag)
			tag.Print()
			fmt.Printf("data: %v, err: %v, addr of data: %p\n", data, err, &tag)
			tag1 := &Tag{}
			err = tag1.Unmarshal(data)
			fmt.Print("err:", err, ", unmarshaled:", tag1, ", IdList:", tag1.IdList)
			/*
				t := &T{}
				p := P(*t)
				p.ttt()

				fmt.Println("vim-go")
	*/
}
func testChanLen() {
	ch := make(chan string, 10)
	fmt.Println("len:", len(ch))
	for i := 0; i < 11; i++ {
		ch <- "hehe"
		fmt.Println("len:", len(ch))
	}
}

func testSelect() {
	ctx := context.Background()
	ch := make(chan string, 1)
	go func(c context.Context) {
		ctx, cancel := context.WithTimeout(c, 1*time.Millisecond)
		fmt.Println("canceler: start to wait.")
		select {
		case <-ctx.Done():
			fmt.Println(" ctx done.", ctx.Err())
		}
		fmt.Println("canceler: to cancel")
		cancel()
	}(ctx)
	go func(ch chan string) {
		fmt.Println("generator: sleep 1 sec")
		select {
		case <-time.After(1 * time.Second):
			break
		}
		fmt.Println("generator: generate dat")
		ch <- "here we go"
		fmt.Println("generator: exiting...")
	}(ch)

	fmt.Println("main: sleep 3 sec")
	select {
	case <-time.After(3 * time.Second):
		break
	}
	fmt.Println("main: select from data")
	select {
	case <-ctx.Done():
		fmt.Println("main: context cancelled.")
		break
	case d := <-ch:
		fmt.Println("main: OK. got from chanal ", d)
		break
	}
	fmt.Println("main: exiting...")
}

func testServer() {
	server := fake.NewTcpServer()
	fmt.Println("created server, to start.")
	server.Start()

	serverAddr := server.GetServerAddr()
	log.Printf("server Addr: %s, type: %T", serverAddr, serverAddr)
	conn, err := net.DialTimeout("tcp", server.GetServerAddr(), 1*time.Second)
	log.Println("conn errr:", err, ", conn:", conn)
	conn, err = net.DialTimeout("tcp", server.GetServerAddr(), 1*time.Second)
	log.Println("conn errr:", err, ", conn:", conn)

	fmt.Println("Started server, to stop.")
	ctx := server.Stop()
	fmt.Println("Stopped server, to check.")
	select {
	case <-ctx.Done():
		log.Println("Successfully stopped")
		break
	case <-time.After(4 * time.Second):
		if server.IsRunning() {
			log.Fatal("Server still running after stopped 4 sec..")
		} else {
			log.Println("Successfully stopped")
		}
	}
}
func main() {
	testServer()
	//testSelect()
	//testRing()
	//testChanLen()
	var ch net.Conn

	if ch == nil {
		fmt.Print("bingle")
	}
}
