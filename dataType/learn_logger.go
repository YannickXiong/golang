/**
 * @Author : Yannick
 * @File   : learn_array.go
 * @Date   : 2017-11-10
 * @Desc   : This is a demo which I could learn golang logger through it.
 */

package main

import (
	"log"
	"os"
	"strconv"
)

// Logger : log.Logger
type Logger log.Logger

// Logger has the same underlying type to log.Logger, and log.New() returns *log.Logger,
// however, code below will lead to compile err: cannot use log.New(os.Stdout, "Logger_prefix",
// log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds) (type *log.Logger) as type Logger in return argument
// This shows that Logger and log.Logger can't automatically conversions type, though their underlying types(ptr) are the same.

// func LoggerInit() *Logger {
// 	logger := log.New(os.Stdout, "[Logger_prefix] ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds)
// 	// *log.Logger
// 	fmt.Printf("type => %s\n", reflect.TypeOf(logger))
// 	// reflect.ValueOf().Kind() returns the underlying type, also called basic type.
//  // and reflect.Type() returns the static type.
// 	// ptr
// 	fmt.Printf("kind => %s\n", reflect.ValueOf(logger).Kind())
// 	myLogger := &Logger{}
// 	// *main.Logger
// 	fmt.Printf("type => %s\n", reflect.TypeOf(myLogger))
// 	// ptr
// 	fmt.Printf("kind => %s\n", reflect.ValueOf(myLogger).Kind())

// 	return logger
// }

// LoggerInit :
func LoggerInit() *log.Logger {
	return log.New(os.Stdout, "[Logger_prefix] ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds)
}

// LearnLogger :
func LearnLogger() {

	logger := LoggerInit()
	for i := 0; i < 10; i++ {
		logger.Printf("from logger write log ...%d + 1000 = %d, order is %s", i, i+1000, strconv.Itoa(i))
	}

}
