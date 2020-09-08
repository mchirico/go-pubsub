package pubsub

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"testing"
	"time"
)

func Test_creds(t *testing.T) {
	g := NewG()
	if len(g.Credential.Project_id) > 3 {
		t.Logf("We have project: %s\n", g.Credential.Project_id)
	} else {
		t.Fatalf("Can't read credential file: ../credentials/credential.json")
	}
}

func TestFindFile(t *testing.T) {
	_, s := FindFile()
	if strings.Contains(s, ".json") {
		t.Logf("found: %s\n", s)
	} else {
		t.Fatalf("Cannot find .json")
	}
}

func TestG_Publish(t *testing.T) {
	g := NewG()
	var buf bytes.Buffer
	id, err := g.Publish(&buf, "test", "test")
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	fmt.Printf("id: %v\n", id)
}

func TestG_CreateTopic(t *testing.T) {
	g := NewG()
	_, err := g.CreateTopic("test")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func TestG_CreateSubForCloudFunctions(t *testing.T) {
	g := NewG()
	topic, _ := g.CreateTopic("gocloud")
	_, err := g.CreateSub("sub-gocloud", topic)
	if err != nil {
		if strings.Contains(err.Error(), "code = AlreadyExists desc = Resource ") {
			t.Logf("This is okay... it should exist")
		} else {
			t.Fatal("Sub error")
		}
	}
	var buf bytes.Buffer
	g.Publish(&buf, "gocloud", "test")
	topic.Stop()
}

func TestG_CreateSub(t *testing.T) {
	g := NewG()
	topic, _ := g.CreateTopic("test")
	_, err := g.CreateSub("sub-test", topic)
	if err != nil {
		if strings.Contains(err.Error(), "code = AlreadyExists desc = Resource ") {
			t.Logf("This is okay... it should exist")
		} else {
			t.Fatal("Sub error")
		}
	}
	topic.Stop()
}

func CreateMsg() {
	g := NewG()
	var buf bytes.Buffer
	id, err := g.Publish(&buf, "test", "test")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("id: %v\n", id)
}

func TestG_PullMsgs(t *testing.T) {

	CreateMsg()
	g := NewG()
	var buf bytes.Buffer
	msg, err := g.PullMsgs(&buf, "sub-test")
	if err != nil {
		t.Fatalf("No message")
	}
	fmt.Printf("msg: %s\n", msg)

}

// Block for N seconds
func TestG_PullMsgsTimeOut(t *testing.T) {

	go func() {
		time.Sleep(14 * time.Second)
		CreateMsg()
	}()
	g := NewG()
	var buf bytes.Buffer

	msg, n, err := g.PullMsgsTimeOut(&buf, "sub-test", 3)
	for n==0 {
		t.Log("Trying again")
		msg, n, err = g.PullMsgsTimeOut(&buf, "sub-test", 3)
	}
	if err != nil {
		t.Fatalf("No message")
	}
	fmt.Printf("msg: %s\n", msg)

}

func Test_ListSubscriptions(t *testing.T) {
	ctx := context.Background()
	g := NewG()
	topic, _ := g.CreateTopic("test")

	for subs := topic.Subscriptions(ctx); ; {
		sub, err := subs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		_ = sub // TODO: use the subscription.
		t.Logf("sub: %s\n", sub.String())
	}

}

// Real pull. Safe to uncomment and live pull
//func TestG_PullMsgsExpress(t *testing.T) {
//
//
//	g := NewG()
//	var buf bytes.Buffer
//	msg, err := g.PullMsgs(&buf, "sub-npubsub")
//	if err != nil {
//		t.Fatalf("No message")
//	}
//	fmt.Printf("msg: %s\n", msg)
//
//}
