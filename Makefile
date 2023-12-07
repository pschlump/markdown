
all:
	go build
	( cd examples ; ./build.sh )
	( cd cli/mdtohtml ; go build )
	( cd cli/mdtomd ; go build )

