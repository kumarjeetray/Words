/*

*****  **      *  *******  *******  *****  *******  *     *  *******  *******
  *    * *     *  *           *       *       *     *     *     *     *      
  *    *  *    *  *           *       *       *     *     *     *     *      
  *    *   *   *  *******     *       *       *     *     *     *     *******
  *    *    *  *        *     *       *       *     *     *     *     *      
  *    *     * *        *     *       *       *     *     *     *     *      
*****  *      **  *******     *     *****     *      *****      *     *******

*/

package main
import "fmt"
func main(){
	g:=2
	e:=1
	f:=5
	for i:=1;i<=7;i++{
		for j:=1;j<=5;j++{
			if (i==1 ||i==7 ||j==3  ){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}
		for j:=1;j<=9;j++{
			if (j==1 || j==9 || j==g){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		g++
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}

		for j:=1;j<=7;j++{
			if (i==1||i==7||i==4||j==1 && i==e||j==7 && i==f){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i==f{
			f++
		}
		if (i<3){
			e++
		}else{
			e=0
		}
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}

		for j:=1;j<=7;j++{
			if (i==1 || j==4){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}

		for j:=1;j<=5;j++{
			if (i==1 ||i==7 ||j==3  ){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}
		for j:=1;j<=7;j++{
			if (i==1 || j==4){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for i:=1;i<=2;i++{
			fmt.Print(" ")
		}
			for j:=1;j<=7;j++{
				if (j==1 && i<7|| j==7 && i<7|| i==7 && j>1 && j<7){
					fmt.Print("*")
				}else{
					fmt.Print(" ")
				}
			}
			for i:=1;i<=2;i++{
				fmt.Print(" ")
			}
			for j:=1;j<=7;j++{
				if (i==1 || j==4){
					fmt.Print("*")
				}else{
					fmt.Print(" ")
				}
			}
			for i:=1;i<=2;i++{
				fmt.Print(" ")
			}
			for j:=1;j<=7;j++{
				if (i==1 || i==7 || i==4 ||j==1){
					fmt.Print("*")
					}else{
					fmt.Print(" ")
				}
			}
			fmt.Println()
	}
}
