/*

*   *        *       *       ******* **      * *****         *      
*  *        * *      *          *    * *     * *            * *     
* *        *   *     *          *    *  *    * *           *   *    
**        *******    *          *    *   *   * * *****    *******   
* *      *       *   *          *    *    *  * *     *   *       *  
*  *    *         *  *          *    *     * * *     *  *         * 
*   *  *           * ******* ******* *      **  ****   *           *


*/
package main
import "fmt"
func main(){
	x:=5
	a:=7
	b:=7
	c:=7
	d:=7
	n:=2
	line:=8
	for i:=1;i<=7;i++{
		for j:=1;j<=7;j++{
			if ((j==1)|| (j==x)){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i<4{
			x--
		}else{
			x++
		}
		for j:=1;j<=13;j++{
			if (j==a || j==b || i==4 && j<=b && j>a){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		a--
		b++
		if b==6{
			continue
		}
		for k:=1;k<=1;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=7;j++{
			if (j==1 || i==7){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for k:=1;k<=1;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=7;j++{
			if (i==1 ||i==7 ||j==4  ){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for k:=1;k<=1;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=9;j++{
			if (j==1 || j==9 || j==n){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		n++
		for k:=1;k<=1;k++{
			fmt.Print(" ")
		}
		for j:= 1; j <=line; j++{ 
            if ((j == 1 && i != 0 && i != line - 1) || 
              ((i == 1 || i == line - 1) && j > 1 && 
              j < line - 2) || (i == (line/ 2) &&  
              j > 2 && j < line-1) || (j == line - 1 &&  
              i != 0 && i >= ((line ) / 2) && i != line - 1)){ 
                fmt.Print("*") 
			  }else{
                fmt.Print(" ") 
  
        	} 
		} 
		for j:=1;j<=13;j++{
			if (j==c || j==d || i==4 && j<=d && j>c){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		c--
		d++
		if d==6{
			continue
		}
		fmt.Println()
	}
}
