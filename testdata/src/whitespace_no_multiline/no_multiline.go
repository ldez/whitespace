package whitespace

import "fmt"

// MultiFunc (FuncDecl), MultiIf and MultiFunc(FuncLit) want to remove newline.
func fn1(
	arg1 int,
	arg2 int,
) { // want "unnecessary leading newline"

	fmt.Println("Hello, World")

	if true &&
		false { // want "unnecessary leading newline"

		fmt.Println("Hello, World")
	}

	_ = func(
		arg1 int,
		arg2 int,
	) { // want "unnecessary leading newline"

		fmt.Println("Hello, World")
	}

	_ = func(
		arg1 int,
		arg2 int,
	) {
		_ = func(
			arg1 int,
			arg2 int,
		) { // want "unnecessary leading newline"

			fmt.Println("Hello, World")
		}
	}
}

// MultiFunc (FuncDecl) with comment.
func fn2(
	arg1 int,
	arg2 int,
) { // want "unnecessary leading newline"

	// A comment.
	fmt.Println("Hello, World")
}

// MultiFunc (FuncDecl) that's not `gofmt`:ed.
func fn3(
	arg1 int,
	arg2 int,
) { // want "unnecessary leading newline"
  
          
  
    // A comment.
	fmt.Println("Hello, World")

	if true { // want "unnecessary leading newline"



        fmt.Println("No comments")


    } // want "unnecessary trailing newline"
       // Also at end


  
} // want "unnecessary trailing newline"

// Regular func (FuncDecl) that's not `gofmt`:ed.
func fn4() { // want "unnecessary leading newline"
  
          
	  

	fmt.Println("Hello, World")

	if true { // want "unnecessary leading newline"



        fmt.Println("No comments")


    } // want "unnecessary trailing newline"



        
} // want "unnecessary trailing newline"

// Regular func (FuncDecl) that's not `gofmt`:ed. with comments
func fn5() {
	// A comment that should still exist after this
	// This one should also still exist



	fmt.Println("Hello, World")

	if true {
		// A comment that should still exist after this
		// This one should also still exist



		fmt.Println("No comments")


	} // want "unnecessary trailing newline"




} // want "unnecessary trailing newline"

// Regular func (FuncDecl) that's not `gofmt`:ed. with comment blocks
func fn6() {
	/*
		Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Curabitur ornare dolor at nulla ultrices cursus.
		Mauris pharetra metus ac condimentum sodales.
		Fusce viverra libero vitae tellus dictum, sed congue risus sodales.
	*/



	fmt.Println("Hello, World")

	if true {
		/*
			Lorem ipsum dolor sit amet, consectetur adipiscing elit.
			Curabitur ornare dolor at nulla ultrices cursus.
			Mauris pharetra metus ac condimentum sodales.
			Fusce viverra libero vitae tellus dictum, sed congue risus sodales.
		*/


		fmt.Println("No comments")


	} // want "unnecessary trailing newline"




} // want "unnecessary trailing newline"

func fn7() { /* Multiline spaning 1 line */ // want "unnecessary leading newline"

	fmt.Println("No comments")
}

func fn8() { /* Multiline spaning
	over several lines */

	fmt.Println("No comments")
}

func fn9() { /* Multiline spaning

	over several lines */

	fmt.Println("No comments")
}
