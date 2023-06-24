print() => print anything you want without any extra spacies and dont go to next line at the end of print
println() => print anything you want with space between arguments and go to next line at the end of print

fmt.printf() => %d for integer numbers - %s for strings - %f for float numbers - %b for binary of a variable - %T for type of a variable

type "go doc fmt " => for see fmt documenetation

========================================
variable := "this is test"
fmt.Println(strings.Contains(variable,"test"))    => this return true or false to us
========================================
fmt.Println(strings.ConatainAny(variable,"test")) => if only one word matche return true   => true
fmt.Println(strings.ConatainAny(variable,"g")) => if any word dont match return false   => false
========================================
fmt.Println(strings.Count(variable,"test")) => count how many "test" exist inside of variable  
========================================
fmt.Println(strings.Cut(variable,"test")) => delete first "test" from variable 
=======================================
spilited_string := strings.Spilit(variable," ")
joined_string := strings.Join(spilited_string,",")
========================================
10TimeTestString := strings.Repeat("test",10)
========================================
strings.Replace(variable,"test","testnew",1) => replace first test with testnew in variable
strings.Replace(variable,"test","testnew",2) => replace first and sencond test with testnew in variable
strings.ReplaceAll(variable,"test","testnew") => replace all "test"s with testnew in variable
========================================
strings.Compare("test","testnew") => compare two string and then if two string is equal return 0 else return 1 or -1
strings.EqualFold("test","testnew") => compare two string and then if two string is equal retun true else return false. this method is not case sensetive
========================================
strings.HasPrefix("Iran","Ir") => check prefix - return true
strings.HasSuffix("Iran","IR") => check suffix
======================================
strings.Index("test","st") => return index number of sub string -in this example retuen 2 to us
======================================
strings.ToLowen("test")  
strings.ToUpper("test")  
strings.Title("test")  

=======================================
strings.Trim("  fsdfd  "," ")
