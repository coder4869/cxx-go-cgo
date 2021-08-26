## GoWeb

## Introduction
	This WEB is an go web demo, now including WEB API used and mongodb DataBase.
	
Project Directory:
	|--go-web
		|--	docs
		|--	src
			|-- WebMain (entry of the Project)
			|-- conf (config operations for DataBase)
			|-- controllers (business operation)
			|-- libs (depends libs)
			|-- log (business logs)
			|-- models
			|-- routers (web api register)
		|-- tests (test cases)	
		
Function Call:
	WebMain --> routers --> controllers --> models & conf & libs
	
Developers:
	coder4869@gmail.com
	
Depends:
	"github.com/coder4869/golibs"
	"gopkg.in/mgo.v2"

	
# Run & Deployment
	Refer to "script/Run&Deploy.md", this document provides the modify points 
	for run and deploy the project. 
	
	
# More
 Other README.md or docs

 
