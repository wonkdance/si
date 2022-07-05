package siwebsocket_test

import (
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/go-wonk/si/siutils"
	"github.com/go-wonk/si/siwebsocket"
	"github.com/stretchr/testify/assert"
)

func TestHub(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	if !longtest {
		t.Skip("skipping long tests")
	}
	hub := siwebsocket.NewHub()
	go hub.Run()

	u := url.URL{Scheme: "ws", Host: ":48080", Path: "/push"}

	for i := 0; i < 50; i++ {
		log.Println("connect")
		conn, _, err := siwebsocket.DefaultConn(u, nil)
		siutils.AssertNilFail(t, err)
		c := siwebsocket.NewClientConfiguredWithHub(conn, 10*time.Second, 60*time.Second, 1024000, true, hub,
			siwebsocket.WithMessageHandler(&siwebsocket.DefaultMessageHandler{}))
		go c.Start()

		// c.SetID("9099909")
		err = hub.AddClient(c)
		if err != nil {
			log.Println(err)
			c.Stop()
			return
		}
	}

	err := hub.Broadcast([]byte("hey"))
	if err != nil {
		log.Println(err)
	}
	time.Sleep(6 * time.Second)
	log.Println("stopping...")
	hub.Stop()
	log.Println("stopped")
	hub.Wait()
	// time.Sleep(6 * time.Second)
}

func TestHub2(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	if !longtest {
		t.Skip("skipping long tests")
	}
	hub := siwebsocket.NewHub()
	go hub.Run()

	u := url.URL{Scheme: "ws", Host: ":48080", Path: "/push"}

	go func() {
		num := 0
		for {
			time.Sleep(80 * time.Millisecond)
			log.Println("connect")
			conn, _, err := siwebsocket.DefaultConn(u, nil)
			if err != nil {
				log.Println(err)
				return
			}
			c := siwebsocket.NewClientConfiguredWithHub(conn,
				10*time.Second, 60*time.Second, 1024000, true, hub,
				siwebsocket.WithMessageHandler(&siwebsocket.DefaultMessageHandler{}))
			go c.Start()

			// c.SetID("9099909")
			err = hub.AddClient(c)
			if err != nil {
				log.Println(err)
				c.Stop()
				return
			}

			num++
			if num > 300 {
				// hub.RemoveRandomClient()
				num--
				break
			}
		}
	}()

	go func() {
		for {
			time.Sleep(71 * time.Millisecond)
			err := hub.Broadcast([]byte("hey"))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()
	time.Sleep(20 * time.Second)
	log.Println("stopping...")
	hub.Stop()
	hub.Wait()
	// time.Sleep(12 * time.Second)
	log.Println("stopped", hub.LenClients())
}

func test() int {
	hub := siwebsocket.NewHub()
	go hub.Run()

	u := url.URL{Scheme: "ws", Host: ":48080", Path: "/push/randomclose"}

	go func() {
		num := 0
		for {
			time.Sleep(80 * time.Millisecond)
			log.Println("connect")
			conn, _, err := siwebsocket.DefaultConn(u, nil)
			if err != nil {
				log.Println(err)
				return
			}
			c := siwebsocket.NewClientConfiguredWithHub(conn,
				10*time.Second, 60*time.Second, 1024000, true, hub,
				siwebsocket.WithMessageHandler(&siwebsocket.DefaultMessageHandler{}))
			go c.Start()

			c.SetID("9099909")
			err = hub.AddClient(c)
			if err != nil {
				log.Println(err)
				c.Stop()
				return
			}

			num++
			if num > 200 {
				// hub.RemoveRandomClient()
				num--
			}
		}
	}()

	go func() {
		for {
			time.Sleep(71 * time.Millisecond)
			err := hub.Broadcast([]byte("hey"))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()
	time.Sleep(20 * time.Second)
	log.Println("stopping...")
	hub.Stop()
	hub.Wait()
	// time.Sleep(12 * time.Second)
	leftOver := hub.LenClients()
	log.Println("stopped", leftOver)

	return leftOver
}

func TestReconnects(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}
	if !longtest {
		t.Skip("skipping long tests")
	}
	for i := 0; i < 10; i++ {
		assert.EqualValues(t, 0, test())
	}
}
