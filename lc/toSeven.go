package main

func toSeven(num int) string {
	var t = "";
	for i := num; i > 0; i /= 7 {
		t += (string)num % 7;
	}

}

func main() {
	
}
