/* 

***** **      * *****   *     * ******* ******* ******* *****     **       *    
  *   * *     * *    *  *     * *          *    *     *   *      *  *      *    
  *   *  *    * *    *  *     * *          *    *     *   *     ******     *    
  *   *   *   * *    *  *     * *******    *    *******   *    *      *    *    
  *   *    *  * *    *  *     *       *    *    *  *      *   *        *   *    
  *   *     * * *    *  *     *       *    *    *   *     *  *          *  *    
***** *      ** *****    *****  *******    *    *    *  *****            * *****

package main
import "fmt"
func main(){
	t:=2
	s:=8
	u:=4
	h:=1
	g:=5
	a1:=6
	b1:=7
	fmt.Println()
	for i:=1;i<=7;i++{
		for j:=1;j<=5;j++{
			if (i==1 ||i==7 ||j==3  ){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print(" ")
		for j:=1;j<=9;j++{
			if (j==1 || j==9 || j==t){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		t++
		fmt.Print(" ")
		for j:= 1; j <= s; j++{ 
            if (j == 1 || ((i == 1 || i == s-1) && 
               (j > 1 && j < s-2)) || (j == s-2 && 
                i != 1 && i != s-1)){
                fmt.Print("*")  
                }else{
                fmt.Print(" ") 
                }
		}
		for j:=1;j<=7;j++{
			if (j==1 && i<7|| j==7 && i<7|| i==7 && j>1 && j<7){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print(" ")
		for j:=1;j<=7;j++{
			if (i==1||i==7||i==4||j==1 && i==h||j==7 && i==g){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i==g{
			g++
		}
		if (i<3){
			h++
		}else{
			h=0
		}
		fmt.Print(" ")
		for j:=1;j<=7;j++{
			if (i==1 || j==4){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print(" ")
		for j:=1;j<=7;j++{
			if (i==1 ||i==4 ||j==1||j==7 && i<=4){
				fmt.Print("*")
				}else if (i>4 && j==u){
					fmt.Print("*")
					}else{
				fmt.Print(" ")
			}
		}
		if i>4{
			u++
		} 
		fmt.Print(" ")
		for j:=1;j<=5;j++{
			if (i==1 ||i==7 ||j==3  ){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for j:=1;j<=13;j++{
			if (j==a1 || j==b1 || i==3 && j<=b1 && j>a1){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		a1--
		b1++
		fmt.Print(" ")
		for j:=1;j<=5;j++{
			if (j==1 || i==7){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}

		fmt.Println();
	}
}
