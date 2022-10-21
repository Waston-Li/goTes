package pack

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IaoTest() {
	//user输入接收
	//整行读入
	fmt.Println("Please enter rules: ")
	var input string
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n') //回车为结束符；包含空格符；
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("Your input is %s", input)
	fmt.Println("字符串的长度:", len(input))
	//单个字符串读入; 不计入回车符；包含空格
	var input2, input3 string
	fmt.Println("Please enter rules: ")
	fmt.Scanln(&input2) //阻塞函数，若有未接受完的参数 ，只接受以空格分割，遇回车结束 (字符串无法包含回车)
	fmt.Printf("Your input is %s", input2)
	fmt.Println("字符串的长度:", len(input2))
	//
	fmt.Println("Please enter rules: ")
	fmt.Scanf("%s", &input3) //阻塞函数，若有未接受完的参数 ，接受空格/回车分割，直到参数接收完
	fmt.Printf("Your input is %s ", input3)
	fmt.Println("字符串的长度:", len(input3))

}

func IaoArr() {
	//读取数据按空格拆分存入切片
	fmt.Println("Please enter rules: ")
	var input string
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	for _, v := range strings.Fields(input) {
		fmt.Println(v)
	}
	//fmt.Printf("Your input is %s", input)

}

func FileInput() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string

	inputFile, inputError := os.Open("input.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	//inputReader := bufio.NewReader(inputFile) //读入器

	// for {  //逐行输出
	// 	inputString, readerError := inputReader.ReadString('\n')
	// 	fmt.Printf("The input was: %s", inputString)
	// 	if readerError == io.EOF {
	// 		return // error or EOF
	// 	}
	// }

	input_str := "input.txt"
	outputFile := "output.sh"
	buf, err := ioutil.ReadFile(input_str) //将整个文件读到一个字符串里
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	// 将[]byte的值写入文件
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

func FileOutput() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile("output_test0", os.O_WRONLY|os.O_CREATE, 0666)
	// 	`os.O_RDONLY`：只读  ; 此参数 用 | 连接
	//  `os.O_WRONLY`：只写
	//  `os.O_CREATE`：创建：如果指定文件不存在，就创建该文件
	//  `os.O_TRUNC`：截断：如果指定文件已存在，就将该文件的长度截为 0 。
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	outputWriter.WriteString(outputString)

	outputWriter.Flush()
}

func CmdFlag() {
	var Rules = flag.String("r", "rm,shutdown", "rule1,rule2...") // of type *string
	var RuleConfigPath = flag.String("f", "/r/r", ".../RuleConfigPath")
	var tof bool
	flag.BoolVar(&tof, "h", false, "help message") //TypeVar形式,功能为将flag绑定到一个变量上

	flag.PrintDefaults() //打印 参数 的使用帮助信息
	flag.Parse()         // Scans the arg list and sets up flags

	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		// flag.NArg()   //命令行参数后的其他参数
		s += flag.Arg(i) //解析后的参数 与 os.Args() 不一样
		s += " "
	}

	os.Stdout.WriteString(s)
	fmt.Printf("\n%s\n%s\n%v\n", *RuleConfigPath, *Rules, *&tof)

	var rule_slice []string
	rule_slice = strings.Split(*Rules, ",")
	fmt.Println(rule_slice)
}
