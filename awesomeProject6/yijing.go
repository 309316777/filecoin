package main

import (
	"fmt"
)

func main() {



	yijing3()
}
 func yijing1() {
	 var a = 56
	 var b = a % 9
	 if (b == 1) {
		 fmt.Printf("坤\n")
	 }
	 if (b == 2) {
		 fmt.Printf("巽\n")
	 }
	 if (b == 3) {
		 fmt.Printf("离\n")
	 }
	 if (b == 4) {
		 fmt.Printf("兑\n")
	 }
	 if (b == 5) {
		 fmt.Printf("艮\n")
	 }
	 if (b == 6) {
		 fmt.Printf("艮\n")
	 }
	 if (b == 7) {
		 fmt.Printf("坎\n")
	 }
	 if (b == 8) {
		 fmt.Printf("震\n")
	 }
	 if (b == 9) {
		 fmt.Printf("乾\n")
	 }
	 if (b == 0) {
		 fmt.Printf("乾\n")
	 }
 }
func yijing2() {
	var a = 67
	var b = a % 9
	if (b == 1) {
		fmt.Printf("坤\n")
	}else if (b == 2) {
		fmt.Printf("巽\n")
	}else if (b == 3) {
		fmt.Printf("离\n")
	}else if (b == 4) {
		fmt.Printf("兑\n")
	}else if (b == 5) {
		fmt.Printf("艮\n")
	}else if (b == 6) {
		fmt.Printf("艮\n")
	}else if (b == 7) {
		fmt.Printf("坎\n")
	} else if (b == 8) {
		fmt.Printf("震\n")
	}else if (b == 9) {
		fmt.Printf("乾\n")
	}else if (b == 0) {
		fmt.Printf("乾\n")
	}
}





func yijing3() {

	var a = 89
	var b = a % 12
	if (b == 1) {
		fmt.Printf("地分：子\n")
	} else if (b == 2) {
		fmt.Printf("地分：丑\n")
	} else if (b == 3) {
		fmt.Printf("地分：寅\n")
	} else if (b == 4) {
		fmt.Printf("地分：卯\n")
	} else if (b == 5) {
		fmt.Printf("地分：辰\n")
	} else if (b == 6) {
		fmt.Printf("地分：巳\n")
	} else if (b == 7) {
		fmt.Printf("地分：午\n")
	} else if (b == 8) {
		fmt.Printf("地分：未\n")
	} else if (b == 9) {
		fmt.Printf("地分：申\n")
	} else if (b == 10) {
		fmt.Printf("地分：酉\n")
	} else if (b == 11) {
		fmt.Printf("地分：戌\n")
	} else if (b == 12) {
		fmt.Printf("地分：亥\n")
	}
}