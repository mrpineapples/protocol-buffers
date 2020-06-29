package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/mrpineapples/go-protobuf/addressbookpb"
	"github.com/mrpineapples/go-protobuf/complexpb"
	"github.com/mrpineapples/go-protobuf/enumpb"
	"github.com/mrpineapples/go-protobuf/simplepb"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()
	doComplex()

	fmt.Println("===== Address Book Demo Below =====")
	addressBookDemo("addressbook.bin")
}

func addressBookDemo(fname string) {
	book := &addressbookpb.AddressBook{
		People: []*addressbookpb.Person{
			{
				Id:    1,
				Name:  "Michael Miranda",
				Email: "mmiranda@example.com",
				Phones: []*addressbookpb.Person_PhoneNumber{
					{
						Number: "718-555-4321",
						Type:   addressbookpb.Person_HOME,
					},
				},
			},
			{
				Id:    2,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*addressbookpb.Person_PhoneNumber{
					{Number: "212-555-4321", Type: addressbookpb.Person_HOME},
				},
			},
		},
	}

	writeToFile(fname, book)
	book2 := &addressbookpb.AddressBook{}
	readFromFile(fname, book2)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println("Complex message:", cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println("smAsString", smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't unmarshal JSON into pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("sm2:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't put bytes into the protocol buffers struct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)
	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
