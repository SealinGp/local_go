package main;
import(
	"fmt"
	"os"
)


func main() {
	args := os.Args;
	var fun1 func();
	for i := 0; i < len(args); i++ {
		fun1 = args[i];
		fun1();
	}

	// code1();

	// code2();	
}

func code1() {
	messages := make(chan string);
	signals  := make(chan bool);

	select {
	case msg := <-messages:
		fmt.Println("received message",msg);
	default:
		fmt.Println("no message received");
	}

	msg := "hi";
	select  {
	case messages <- msg:
		fmt.Println("sent message",msg);
	default:
		fmt.Println("no message sent");
	}

	select {
	case msg := <-messages:
		fmt.Println("received message",msg);
	case sig := <-signals:
		fmt.Println("received signal",sig);
	default:
		fmt.Println("no activity");
	}
}

func code2() {
	messages := make(chan string,1);
	signals  := make(chan bool);

	select {
	case msg := <-messages:
		fmt.Println("received message",msg);
	default:
		fmt.Println("no message received");
	}

	msg := "hi world";
	select {
	case messages <- msg:
		fmt.Println("sent message",msg);
	default:
		fmt.Println("no message sent");
	}

	select {
	case msg := <-messages:
		fmt.Println("received message1",msg);
	case sig := <-signals:
		fmt.Println("received signal",sig);
	default:
		fmt.Println("no activity");
	}
}