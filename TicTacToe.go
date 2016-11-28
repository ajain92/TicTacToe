
package main

import (
"fmt"
)

var player int = 0;
var winner int;
var choice string;
var row int;
var column int;
var i int;
var j int;
var P1 string = "X";
var P2 string = "O";
var Counter int;
var returnValue string;
var Player1Counter int = 0;
var Player2Counter int = 0;

var boardTest = [3][3]string{}
var isWin bool = false;
    
func main() {

    boardTest[0][0] = "1";
    boardTest[0][1] = "2";
    boardTest[0][2] = "3";
    boardTest[1][0] = "4";
    boardTest[1][1] = "5";
    boardTest[1][2] = "6";
    boardTest[2][0] = "7";
    boardTest[2][1] = "8";
    boardTest[2][2] = "9";


	for i = 0; i < 3; i++ {
            fmt.Print("\n | ");
            for j = 0; j < 3; j++ {
                fmt.Print(boardTest[i][j] + " | ");
            }
	}

	fmt.Print("\n\n Player 1 : X \n Player 2 : O");

	for Counter = 0; Counter < 9; Counter++ {
		if(Counter%2 == 0){
			fmt.Print("\n\n Player 1 Enter Choice: ");
			Player1Counter++;
		} else {
			fmt.Print("\n\n Player 2 Enter Choice: ");
			Player2Counter++;
		}

		fmt.Scanln(&choice);
		NumChoice(choice);

		if(Counter > 3){
			CheckWinning(0, Counter)
        }

        if(isWin){
        	break
        }
        
		if(Player1Counter == 5 && Player2Counter == 4){
			fmt.Print("\n\n Game Draw !! ");
			isWin =  true;
		}
	}
}

func CheckWinning(i int, counterNew int){
	if(boardTest[i][i] == boardTest[i][i+1]){
		if(boardTest[i][i+2] == boardTest[i][i]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i+1][i] == boardTest[i+1][i+1]){
		if(boardTest[i+1][i+2] == boardTest[i+1][i]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i+2][i] == boardTest[i+2][i+1]){
		if(boardTest[i+2][i+1] == boardTest[i+2][i+2]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i][i] == boardTest[i+1][i]){
		if(boardTest[i+1][i] == boardTest[i+2][i]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i][i+1] == boardTest[i+1][i+1]){
		if(boardTest[i+1][i+1] == boardTest[i+2][i+1]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i][i+2] == boardTest[i+1][i+2]){
		if(boardTest[i+1][i+2] == boardTest[i+2][i+2]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i][i] == boardTest[i+1][i+1]){
		if(boardTest[i+1][i+1] == boardTest[i+2][i+2]){
			ValidateWin(counterNew);
		}
	}
	if(boardTest[i][i+2] == boardTest[i+1][i+1]){
		if(boardTest[i+1][i+1] == boardTest[i+2][i]){
			ValidateWin(counterNew);
		}
	}
}

func ValidateWin(counterValidateWin int){
	if(counterValidateWin%2==0){
	    fmt.Print("\n\n Player 1 Wins !! ");
	    isWin = true;
	} else {
		fmt.Print("\n\n Player 2 Wins !! ");
		isWin =  true;
	}
}

func NumChoice(choice1 string) {
		
	switch choice1 {
	case "1":
		if (Counter%2 == 0){
		  boardTest[0][0] = "X"
		} else {
			boardTest[0][0] = "O"
		}
		CreateMatrix(0, 0);	
	case "2":
		if (Counter%2 == 0){
		  boardTest[0][1] = "X"
		} else {
			boardTest[0][1] = "O"
		}
		CreateMatrix(0, 1);
	case "3":
		if (Counter%2 == 0){
		  boardTest[0][2] = "X"
		} else {
			boardTest[0][2] = "O"
		}
		CreateMatrix(0, 2);
	case "4":
		if (Counter%2 == 0){
		  boardTest[1][0] = "X"
		} else {
			boardTest[1][0] = "O"
		}
		CreateMatrix(1, 0);
	case "5":
		if (Counter%2 == 0){
		  boardTest[1][1] = "X"
		} else {
			boardTest[1][1] = "O"
		}
		CreateMatrix(1, 1);
	case "6":
		if (Counter%2 == 0){
		  boardTest[1][2] = "X"
		} else {
			boardTest[1][2] = "O"
		}
		CreateMatrix(1, 2);
	case "7":
		if (Counter%2 == 0){
		  boardTest[2][0] = "X"
		} else {
			boardTest[2][0] = "O"
		}
			CreateMatrix(2, 0);
		
	case "8":
		if (Counter%2 == 0){
		  boardTest[2][1] = "X"
		} else {
			boardTest[2][1] = "O"
		}
		CreateMatrix(2, 1);	
		
	case "9":
		if (Counter%2 == 0){
		  boardTest[2][2] = "X"
		} else {
			boardTest[2][2] = "O"
		}
		CreateMatrix(2, 2);
	}
}

func CreateMatrix(valueI int, valueJ int){
	for i = 0; i < 3; i++ {
            fmt.Print("\n | ");
            for j = 0; j < 3; j++ {
                fmt.Print(boardTest[i][j] + " | ");
            }
		}
}
