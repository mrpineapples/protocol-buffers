protoc -I simplepb/ --go_out=simplepb/ simplepb/simple.proto
protoc -I enumpb/ --go_out=enumpb/ enumpb/enum.proto
protoc -I complexpb/ --go_out=complexpb/ complexpb/complex.proto
protoc -I addressbookpb/ --go_out=addressbookpb/ addressbookpb/addressbook.proto