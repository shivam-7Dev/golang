/*
standdard log package doesnt have bilt in  support for logging lever like
debug ifro warn error.
however we can create custom fuction to handle different level
*/
package intermediate

import "log"

func LogIntro() {
	log.Println("this is from log file")
	LogInfo("fetching data from api")
}

func LogInfo(info string) {
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Default().Flags() | log.LUTC | log.Lshortfile | log.Llongfile)
	log.Println(info)
}
