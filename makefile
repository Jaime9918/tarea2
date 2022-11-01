combine: 
	go run Combine/main.go

rebeldes: 
	go run Rebeldes/main.go

datanode1: 
	go run DataNode1/main.go

datanode2: 
	go run DataNode2/main.go

datanode3: 
	go run DataNode3/main.go

datanode4:
	go run DataNode4/main.go

namenode:
	go run NameNode/main.go

#clean:
#	if [ -f SOLICITUDES.txt ]; then rm SOLICITUDES.txt -R; fi