package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/younghyun-ahn/go-etc/protobuf/simple/simplepb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	sp := doSimple()

	//readAndWriteDemo(sp)
	jsonDemo(sp)
}

func jsonDemo(sp *simplepb.SimpleMessage) {
	spAsString := toJSON(sp)
	fmt.Println(spAsString)

	sp2 := &simplepb.SimpleMessage{}
	fromJSON(spAsString, sp2)
	fmt.Println("Successfully created proto struct:", sp2)
}

func toJSON(pb *simplepb.SimpleMessage) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb *simplepb.SimpleMessage) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into ", err)
	}
}

func readAndWriteDemo(sp *simplepb.SimpleMessage) {
	writeToFile("simple.bin", sp)
	sp2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sp2)
	fmt.Println("Read the content:", sp2)
}

func writeToFile(fname string, sp proto.Message) error {
	out, err := proto.Marshal(sp)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, sp proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, sp)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers sturct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sp := simplepb.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "My Simple Message",
		SimpleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(sp)

	sp.Name = "I renamed you"
	fmt.Println(sp)

	fmt.Println("The ID is: ", sp.GetId())

	return &sp
}