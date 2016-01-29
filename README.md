# BufioReaderCrash

Following steps will help you to simulate the problem on the same pc
1. Download the zip 
2. Build and Run the main.go location BufioCrashSolution/main.go
3. Open another terminal and enter the command "nc localhost 8090" press enter
4. then type some text atleast 33 characters and press enter
 example "dfbvjdfvbjkdsabvkjdffvjkdskbvkjdsbvsadb vmsdn vmnsd v jksdnvjkdsnvkjd vkjsd vkds vs"
5. Now go to the console window in which your main.go was running
  then see the program crash
##########################################
  panic: runtime error: index out of range

goroutine 7 [running]:
bufio.(*Reader).ReadByte(0xc20803e060, 0x2, 0x0, 0x0)
    /home/XYZ/go/src/bufio/bufio.go:209 +0x596
main.readln(0xc20803e060, 0x0, 0x0)
    /home/XYX/BufioCrashSolution/main.go:41 +0x8f
created by main.func·001
    /home/XYZ/BufioCrashSolution/main.go:26 +0xc8


 
