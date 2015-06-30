package transaction

import (
	"testing"

	"appengine"
	"appengine/aetest"
	"appengine/datastore"
)

func balance(c appengine.Context, t *testing.T) int {
	b := BankAccount{}
	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
	if err := datastore.Get(c, key, &b); err != nil {
		t.Fatal(err)
	}
	return b.Balance
}

func TestWithdrawNoAcc(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	sc := make(chan string)

	err = withdraw(c, sc, "myid", 42, 0)
	if err != datastore.ErrNoSuchEntity {
		t.Errorf("Want ErrNoSuchEntity; got %v", err)
	}
}

func TestWithdrawOkay(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	sc := make(chan string)
	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
	if _, err := datastore.Put(c, key, &BankAccount{100}); err != nil {
		t.Fatal(err)
	}

	donec := make(chan bool)
	go func() {
		if msg, want := <-sc, "myid: balance is $0000100  (attempt number 0)\n"; msg != want {
			t.Errorf("Message %q, want %q", msg, want)
		}
		donec <- true
	}()

	err = withdraw(c, sc, "myid", 42, 0)
	if err != nil {
		t.Fatalf("Error: %v; want no error", err)
	}

	if bal, want := balance(c, t), 58; bal != want {
		t.Errorf("Balance %d, want %d", bal, want)
	}

	<-donec
}

func TestWithdrawLowBal(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	sc := make(chan string, 1)
	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
	if _, err := datastore.Put(c, key, &BankAccount{100}); err != nil {
		t.Fatal(err)
	}

	err = withdraw(c, sc, "myid", 128, 0)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("Error: %v; want insufficient funds error", err)
	}

	if bal, want := balance(c, t), 100; bal != want {
		t.Errorf("Balance %d, want %d", bal, want)
	}
}
